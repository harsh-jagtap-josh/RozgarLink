package worker

import (
	"context"

	"github.com/harsh-jagtap-josh/RozgarLink/internal/repo"
)

type service struct {
	workerRepo repo.WorkerStorer
}

type Service interface {
	GetWorkerByID(ctx context.Context, productID int) (repo.Worker, error)
}

func NewService(workerRepo repo.WorkerStorer) Service {
	return &service{
		workerRepo: workerRepo,
	}
}

func (ws *service) GetWorkerByID(ctx context.Context, workerId int) (repo.Worker, error) {
	workerInfoDB, err := ws.workerRepo.GetWorkerByID(ctx, workerId)
	if err != nil {
		return repo.Worker{}, err
	}

	return workerInfoDB, err
}
