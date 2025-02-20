package worker

import (
	"context"

	"github.com/harsh-jagtap-josh/RozgarLink/internal/app/application"
	"github.com/harsh-jagtap-josh/RozgarLink/internal/pkg/apperrors"
	"github.com/harsh-jagtap-josh/RozgarLink/internal/pkg/utils"
	"github.com/harsh-jagtap-josh/RozgarLink/internal/repo"
)

type service struct {
	workerRepo repo.WorkerStorer
}

type Service interface {
	FetchWorkerByID(ctx context.Context, workerId int) (WorkerRequest, error)
	CreateWorker(ctx context.Context, workerData WorkerRequest) (WorkerRequest, error, error)
	UpdateWorkerByID(ctx context.Context, workerData WorkerRequest) (WorkerRequest, error, error)
	DeleteWorkerByID(ctx context.Context, workerId int) (int, error)
	FetchApplicationsByWorkerId(ctx context.Context, workerId int) ([]application.Application, error)
}

func NewService(workerRepo repo.WorkerStorer) Service {
	return &service{
		workerRepo: workerRepo,
	}
}

func (ws *service) FetchWorkerByID(ctx context.Context, workerId int) (WorkerRequest, error) {

	workerData, err := ws.workerRepo.FetchWorkerByID(ctx, workerId)
	if err != nil {
		return WorkerRequest{}, err
	}

	newWorker := MapRepoDomainToService(workerData)
	return newWorker, nil
}

func (ws *service) CreateWorker(ctx context.Context, workerData WorkerRequest) (WorkerRequest, error, error) {

	// validate fields of user - can also be done in front-end itself
	errType, err := utils.ValidateUser(workerData.Name, workerData.ContactNumber, workerData.Email, workerData.Password)
	if err != nil {
		return WorkerRequest{}, err, errType
	}

	alreadyExists := ws.workerRepo.FindWorkerByEmail(ctx, workerData.Email)
	if alreadyExists {
		return WorkerRequest{}, apperrors.ErrWorkerAlreadyExists, nil
	}

	hashed_password, err := utils.HashPassword(workerData.Password)
	if err != nil {
		return WorkerRequest{}, err, apperrors.ErrEncrPassword
	}

	workerData.Password = hashed_password
	repoWorkerObj := MapServiceDomainToRepo(workerData)

	newWorkerData, err := ws.workerRepo.CreateWorker(ctx, repoWorkerObj)
	if err != nil {
		return WorkerRequest{}, err, apperrors.ErrCreateWorker
	}

	mappedWorkerData := MapRepoDomainToService(newWorkerData)
	return mappedWorkerData, nil, nil
}

func (ws *service) UpdateWorkerByID(ctx context.Context, workerData WorkerRequest) (WorkerRequest, error, error) {

	// validate all user fields in update same as register- can also be done in frontend itself
	errType, err := utils.ValidateUpdateUser(workerData.Name, workerData.ContactNumber, workerData.Email)
	if err != nil {
		return WorkerRequest{}, err, errType
	}

	repoWorkerObj := MapServiceDomainToRepo(workerData)

	newWorkerData, err := ws.workerRepo.UpdateWorkerByID(ctx, repoWorkerObj)
	if err != nil {
		return WorkerRequest{}, err, apperrors.ErrUpdateWorker
	}

	mappedWorkerData := MapRepoDomainToService(newWorkerData)
	return mappedWorkerData, nil, nil
}

func (ws *service) DeleteWorkerByID(ctx context.Context, workerId int) (int, error) {
	workerExists := ws.workerRepo.FindWorkerById(ctx, workerId)
	if !workerExists {
		return -1, apperrors.ErrNoWorkerExists
	}

	id, err := ws.workerRepo.DeleteWorkerByID(ctx, workerId)
	if err != nil {
		return -1, err
	}
	return id, nil
}

func (ws *service) FetchApplicationsByWorkerId(ctx context.Context, workerId int) ([]application.Application, error) {
	workerExists := ws.workerRepo.FindWorkerById(ctx, workerId)
	if !workerExists {
		return []application.Application{}, apperrors.ErrNoWorkerExists
	}

	applications, err := ws.workerRepo.FetchApplicationsByWorkerId(ctx, workerId)
	if err != nil {
		return []application.Application{}, err
	}

	var fetchedApplications []application.Application
	for _, appl := range applications {
		fetchedApplications = append(fetchedApplications, application.MapRepoApplicationToService(appl))
	}

	return fetchedApplications, nil
}
