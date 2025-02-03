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

func (ps *service) GetWorkerByID(ctx context.Context, tx repository.Transaction, productID int64) (repository.Worker, error) {
	productInfoDB, err := ps.workerRepo.GetWorkerByID(ctx, tx, productID)
	if err != nil {
		return repository.Worker{}, nil
	}

	if productInfoDB.ID == 0 {
		return repository.Worker{}, nil
	}

	return productInfoDB, nil
}
