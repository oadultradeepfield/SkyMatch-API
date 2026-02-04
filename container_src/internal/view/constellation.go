package view

import "server/internal/model"

// ConstellationView is the JSON response representation of a constellation
type ConstellationView struct {
	LatinName   string `json:"latinName"`
	EnglishName string `json:"englishName"`
	ImageURL    string `json:"imageUrl"`
}

// GetViewFromModel converts a model.Constellation to ConstellationView
func GetViewFromModel(c model.Constellation) ConstellationView {
	return ConstellationView{
		LatinName:   c.LatinName,
		EnglishName: c.EnglishName,
		ImageURL:    c.ImageURL(),
	}
}

// ConstellationsResponse is the JSON response for the search endpoint
type ConstellationsResponse struct {
	Constellations []ConstellationView `json:"constellations"`
}

// GetViewFromModels converts a slice of model.Constellation to ConstellationsResponse
func GetViewFromModels(constellations []model.Constellation) ConstellationsResponse {
	views := make([]ConstellationView, len(constellations))
	for i, c := range constellations {
		views[i] = GetViewFromModel(c)
	}
	return ConstellationsResponse{Constellations: views}
}
