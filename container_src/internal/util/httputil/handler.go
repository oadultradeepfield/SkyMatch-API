package httputil

import (
	"context"
	"errors"
	"log"
	"net/http"

	apperrors "server/internal/errors"
)

type HandlerFunc func(w http.ResponseWriter, r *http.Request) error

func ErrorHandler(f HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			if errors.Is(err, context.Canceled) {
				return
			}

			if errors.Is(err, context.DeadlineExceeded) {
				WriteError(w, http.StatusGatewayTimeout, "request timeout")
				return
			}

			var apiErr *apperrors.APIError
			if errors.As(err, &apiErr) {
				log.Printf("handling %q: %v", r.RequestURI, err)
				WriteError(w, apiErr.Code, apiErr.Message)
				return
			}

			log.Printf("handling %q: %v", r.RequestURI, err)
			WriteError(w, http.StatusInternalServerError, "internal server error")
		}
	}
}
