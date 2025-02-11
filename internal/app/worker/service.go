package worker

import (
	"context"

	"github.com/harsh-jagtap-josh/RozgarLink/internal/repo"
)

type service struct {
	workerRepo repo.WorkerStorer
}

type Service interface {
	FetchWorkerByID(ctx context.Context, workerId int) (Worker, error)
}

func NewService(workerRepo repo.WorkerStorer) Service {
	return &service{
		workerRepo: workerRepo,
	}
}

func (ws *service) FetchWorkerByID(ctx context.Context, workerId int) (Worker, error) {

	workerInfoDB, err := ws.workerRepo.GetWorkerByID(ctx, workerId)
	if err != nil {
		return Worker{}, err
	}

	newWorker := Mapper(workerInfoDB)
	return newWorker, nil
}
