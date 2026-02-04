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

	"server/internal/controller"
	"server/internal/util/httputil"
)

func main() {
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	router := chi.NewRouter()
	router.Use(httprate.LimitByIP(100, time.Second))
	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte("Hello from SkyMatch API"))
		if err != nil {
			return
		}
	})

	router.Get("/api/constellations", httputil.ErrorHandler(controller.SearchConstellations))

	router.Post("/api/solve", httputil.ErrorHandler(controller.SubmitImage))
	router.Get("/api/solve/{jobId}", httputil.ErrorHandler(controller.GetSolveStatus))
	router.Delete("/api/solve/{jobId}", httputil.ErrorHandler(controller.CancelSolve))

	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
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
