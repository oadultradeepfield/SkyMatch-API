package client

import (
	"io"

	"server/internal/client/nova"
	"server/internal/client/simbad"
)

// NovaClient defines the contract for Nova API operations.
type NovaClient interface {
	Login(apiKey string) (string, error)
	Upload(session string, file io.Reader, filename string) (int, error)
	GetSubmission(subID int) (*nova.Submission, error)
	GetJobStatus(jobID int) (string, error)
	GetJobInfo(jobID int) (*nova.JobInfo, error)
	GetAnnotations(jobID int) ([]nova.Annotation, error)
	AnnotatedImageURL(jobID int) string
}

// SimbadClient defines the contract for SIMBAD queries.
type SimbadClient interface {
	QueryObject(identifier string) (*simbad.ObjectInfo, error)
}
