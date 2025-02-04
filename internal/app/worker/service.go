package worker

import (
	"context"

	"github.com/harsh-jagtap-josh/RozgarLink/internal/pkg/dto"
	"github.com/harsh-jagtap-josh/RozgarLink/internal/repository"
)

type service struct {
	workerRepo repository.WorkerStorer
}

type Service interface {
	GetWorkerByID(ctx context.Context, productID int64) (dto.Worker, error)
}

func NewService(workerRepo repository.WorkerStorer) Service {
	return &service{
		workerRepo: workerRepo,
	}
}

func (ws *service) GetWorkerByID(ctx context.Context, productID int64) (dto.Worker, error) {
	workerInfoDB, err := ws.workerRepo.GetWorkerByID(ctx, productID)
	if err != nil {
		return dto.Worker{}, err
	}

	if workerInfoDB.ID == 0 {
		return dto.Worker{}, nil
	}

	return workerInfoDB, nil
}
