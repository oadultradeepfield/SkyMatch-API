package client

import (
	"context"
	"io"

	"server/internal/client/nova"
	"server/internal/client/simbad"
)

// NovaClient defines the contract for Nova API operations.
type NovaClient interface {
	Login(ctx context.Context, apiKey string) (string, error)
	Upload(ctx context.Context, session string, file io.Reader, filename string) (int, error)
	GetSubmission(ctx context.Context, subID int) (*nova.Submission, error)
	GetJobStatus(ctx context.Context, jobID int) (string, error)
	GetJobInfo(ctx context.Context, jobID int) (*nova.JobInfo, error)
	GetAnnotations(ctx context.Context, jobID int) ([]nova.Annotation, error)
	AnnotatedImageURL(jobID int) string
}

// SimbadClient defines the contract for SIMBAD queries.
type SimbadClient interface {
	QueryObject(ctx context.Context, identifier string) (*simbad.ObjectInfo, error)
}
