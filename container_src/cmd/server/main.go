package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/httprate"

	"server/internal/client"
	"server/internal/client/kv"
	"server/internal/client/nova"
	"server/internal/client/simbad"
	"server/internal/config"
	"server/internal/controller"
	"server/internal/service/solve"
	"server/internal/util/httputil"
)

func main() {
	cfg := config.Load()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	novaClient := nova.NewClient(cfg.Nova)

	var simbadClient client.SimbadClient = simbad.NewClient(cfg.Simbad)
	if cfg.KV.Enabled {
		kvClient := kv.NewClient(cfg.KV)
		simbadClient = simbad.NewCachedClient(simbadClient, kvClient, 30*24*3600)
	}

	solveService := solve.NewService(novaClient, simbadClient, cfg.Nova.APIKey)

	solveController := controller.NewSolveController(solveService)

	router := chi.NewRouter()
	router.Use(httprate.LimitByIP(100, time.Second))
	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte("Hello from SkyMatch API"))
		if err != nil {
			return
		}
	})

	router.Get("/api/constellations", httputil.ErrorHandler(controller.SearchConstellations))

	router.Post("/api/solve", httputil.ErrorHandler(solveController.SubmitImage))
	router.Get("/api/solve/{jobId}", httputil.ErrorHandler(solveController.GetSolveStatus))
	router.Delete("/api/solve/{jobId}", httputil.ErrorHandler(solveController.CancelSolve))

	server := &http.Server{
		Addr:         ":" + cfg.Server.Port,
		Handler:      router,
		ReadTimeout:  cfg.Server.ReadTimeout,
		WriteTimeout: cfg.Server.WriteTimeout,
	}

	go func() {
		log.Printf("Server listening on %s\n", server.Addr)
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatal(err)
		}
	}()

	sig := <-stop
	log.Printf("Received signal (%s), shutting down server...", sig)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatal(err)
	}

	log.Println("Server shutdown successfully")
}
