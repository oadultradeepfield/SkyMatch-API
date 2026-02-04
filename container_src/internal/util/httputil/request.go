package httputil

import (
	"net/http"
	"strings"
)

// QueryParam returns a lowercase, trimmed query parameter
func QueryParam(r *http.Request, key string) string {
	return strings.ToLower(strings.TrimSpace(r.URL.Query().Get(key)))
}
