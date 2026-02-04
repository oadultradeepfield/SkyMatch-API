package controller

import (
	"fmt"
	"mime/multipart"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"

	"server/internal/model"
	"server/internal/service/solve"
	"server/internal/util/httputil"
	"server/internal/view"
)

type SolveController struct {
	service *solve.Service
}

func NewSolveController(service *solve.Service) *SolveController {
	return &SolveController{service: service}
}

func (c *SolveController) SubmitImage(w http.ResponseWriter, r *http.Request) error {
	if err := r.ParseMultipartForm(32 << 20); err != nil {
		httputil.WriteError(w, http.StatusBadRequest, "invalid form")
		return nil
	}
	file, header, err := r.FormFile("image")
	if err != nil {
		httputil.WriteError(w, http.StatusBadRequest, "missing image")
		return nil
	}
	defer func(file multipart.File) {
		closeErr := file.Close()
		if closeErr != nil {
			return
		}
	}(file)

	subID, err := c.service.Submit(file, header.Filename)
	if err != nil {
		return fmt.Errorf("submit: %w", err)
	}
	httputil.WriteJSON(w, http.StatusAccepted, view.SubmitResponse{JobID: fmt.Sprintf("%d", subID)})
	return nil
}

func (c *SolveController) GetSolveStatus(w http.ResponseWriter, r *http.Request) error {
	jobID, err := strconv.Atoi(chi.URLParam(r, "jobId"))
	if err != nil {
		httputil.WriteError(w, http.StatusBadRequest, "invalid jobId")
		return nil
	}
	result, err := c.service.GetStatus(jobID, r.URL.Query().Get("fetch") == "true")
	if err != nil {
		return err
	}
	httputil.WriteJSON(w, http.StatusOK, view.NewSolveStatusResponse(result))
	return nil
}

func (c *SolveController) CancelSolve(w http.ResponseWriter, r *http.Request) error {
	jobID := chi.URLParam(r, "jobId")
	if _, err := strconv.Atoi(jobID); err != nil {
		httputil.WriteError(w, http.StatusBadRequest, "invalid jobId")
		return nil
	}
	httputil.WriteJSON(w, http.StatusOK, view.CancelResponse{JobID: jobID, Status: string(model.StatusCancelled)})
	return nil
}
