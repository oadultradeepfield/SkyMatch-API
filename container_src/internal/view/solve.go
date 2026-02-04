package view

import "server/internal/model"

type SubmitResponse struct {
	JobID string `json:"jobId"`
}

type SolveStatusResponse struct {
	JobID             string             `json:"jobId"`
	Status            string             `json:"status"`
	AnnotatedImageURL string             `json:"annotatedImageUrl,omitempty"`
	IdentifiedObjects []IdentifiedObject `json:"identifiedObjects,omitempty"`
}

type IdentifiedObject struct {
	Type           string          `json:"type"`
	Identifier     string          `json:"identifier"`
	Name           string          `json:"name,omitempty"`
	Constellation  *Constellation  `json:"constellation,omitempty"`
	XCoordinate    float64         `json:"xCoordinate"`
	YCoordinate    float64         `json:"yCoordinate"`
	StarDetails    *StarDetails    `json:"starDetails,omitempty"`
	DeepSkyDetails *DeepSkyDetails `json:"deepSkyDetails,omitempty"`
}

type Constellation struct {
	LatinName   string `json:"latinName"`
	EnglishName string `json:"englishName"`
}

type StarDetails struct {
	VisualMagnitude *float64 `json:"visualMagnitude,omitempty"`
	SpectralClass   string   `json:"spectralClass,omitempty"`
	DistanceParsecs *float64 `json:"distanceParsecs,omitempty"`
}

type DeepSkyDetails struct {
	ObjectType string `json:"objectType"`
}

type CancelResponse struct {
	JobID  string `json:"jobId"`
	Status string `json:"status"`
}

func NewSolveStatusResponse(r *model.SolveResult) SolveStatusResponse {
	resp := SolveStatusResponse{
		JobID:             r.JobID,
		Status:            string(r.Status),
		AnnotatedImageURL: r.AnnotatedImageURL,
	}
	for _, obj := range r.Objects {
		resp.IdentifiedObjects = append(resp.IdentifiedObjects, toIdentifiedObject(obj))
	}
	return resp
}

func toIdentifiedObject(obj model.IdentifiedObject) IdentifiedObject {
	v := IdentifiedObject{
		Type:        string(obj.Type),
		Identifier:  obj.Identifier,
		Name:        obj.Name,
		XCoordinate: obj.XCoordinate,
		YCoordinate: obj.YCoordinate,
	}
	if obj.Constellation != nil {
		v.Constellation = &Constellation{
			LatinName:   obj.Constellation.LatinName,
			EnglishName: obj.Constellation.EnglishName,
		}
	}
	if obj.Type == model.ObjectTypeStar {
		v.StarDetails = &StarDetails{
			VisualMagnitude: obj.VMagnitude,
			SpectralClass:   string(obj.SpectralClass),
			DistanceParsecs: obj.DistanceParsecs,
		}
	} else {
		v.DeepSkyDetails = &DeepSkyDetails{ObjectType: string(obj.DSOType)}
	}
	return v
}
