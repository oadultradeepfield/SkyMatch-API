package service

import (
	"context"
	"io"

	"server/internal/model"
)

type SolveService interface {
	Submit(ctx context.Context, file io.Reader, filename string) (int, error)
	GetStatus(ctx context.Context, subID int, fetch bool) (*model.SolveResult, error)
}
