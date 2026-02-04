package httputil

import (
	"log"
	"net/http"
)

// HandlerFunc is a handler that returns an error
type HandlerFunc func(w http.ResponseWriter, r *http.Request) error

// ErrorHandler wraps a HandlerFunc to handle errors consistently
func ErrorHandler(f HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			log.Printf("handling %q: %v", r.RequestURI, err)
			WriteError(w, http.StatusInternalServerError, err.Error())
		}
	}
}
