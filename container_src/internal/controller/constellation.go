package controller

import (
	"net/http"

	"server/internal/model"
	"server/internal/util/httputil"
	"server/internal/view"
)

// SearchConstellations handles GET /api/constellations
func SearchConstellations(w http.ResponseWriter, r *http.Request) error {
	query := httputil.QueryParam(r, "query")
	results := model.SearchConstellations(query)
	response := view.GetViewFromModels(results)
	httputil.WriteJSON(w, http.StatusOK, response)
	return nil
}
