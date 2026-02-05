package controller

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"

	apperrors "server/internal/errors"
	"server/internal/model"
	"server/internal/service"
	"server/internal/util/httputil"
	"server/internal/view"
)

type SolveController struct {
	service service.SolveService
}

func NewSolveController(svc service.SolveService) *SolveController {
	return &SolveController{service: svc}
}

func (c *SolveController) SubmitImage(w http.ResponseWriter, r *http.Request) error {
	if err := r.ParseMultipartForm(32 << 20); err != nil {
		return apperrors.NewValidationError("invalid form")
	}
	file, header, err := r.FormFile("image")
	if err != nil {
		return apperrors.NewValidationError("missing image")
	}
	defer func() {
		if err := file.Close(); err != nil {
			log.Printf("failed to close file: %v", err)
		}
	}()

	subID, err := c.service.Submit(r.Context(), file, header.Filename)
	if err != nil {
		return fmt.Errorf("submit: %w", err)
	}
	httputil.WriteJSON(w, http.StatusAccepted, view.SubmitResponse{JobID: fmt.Sprintf("%d", subID)})
	return nil
}

func (c *SolveController) GetSolveStatus(w http.ResponseWriter, r *http.Request) error {
	jobID, err := strconv.Atoi(chi.URLParam(r, "jobId"))
	if err != nil {
		return apperrors.NewValidationError("invalid jobId")
	}
	result, err := c.service.GetStatus(r.Context(), jobID, r.URL.Query().Get("fetch") == "true")
	if err != nil {
		return err
	}
	httputil.WriteJSON(w, http.StatusOK, view.NewSolveStatusResponse(result))
	return nil
}

func (c *SolveController) CancelSolve(w http.ResponseWriter, r *http.Request) error {
	jobID := chi.URLParam(r, "jobId")
	if _, err := strconv.Atoi(jobID); err != nil {
		return apperrors.NewValidationError("invalid jobId")
	}
	httputil.WriteJSON(w, http.StatusOK, view.CancelResponse{JobID: jobID, Status: string(model.StatusCancelled)})
	return nil
}
