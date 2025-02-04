package worker

import (
	"context"

	"github.com/harsh-jagtap-josh/RozgarLink/internal/repository"
)

type service struct {
	workerRepo repository.WorkerStorer
}

type Service interface {
	GetWorkerByID(ctx context.Context, tx repository.Transaction, productID int64) (repository.Worker, error)
}

func NewService(workerRepo repository.WorkerStorer) Service {
	return &service{
		workerRepo: workerRepo,
	}
}

func (ws *service) GetWorkerByID(ctx context.Context, tx repository.Transaction, productID int64) (repository.Worker, error) {
	workerInfoDB, err := ws.workerRepo.GetWorkerByID(ctx, tx, productID)
	if err != nil {
		return repository.Worker{}, nil
	}

	if workerInfoDB.ID == 0 {
		return repository.Worker{}, nil
	}

	return workerInfoDB, nil
}
