package repository

import (
	"context"

	"github.com/harsh-jagtap-josh/RozgarLink/internal/pkg/dto"
)

type WorkerStorer interface {
	GetWorkerByID(ctx context.Context, productID int64) (dto.Worker, error)
}
