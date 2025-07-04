package application

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/harsh-jagtap-josh/RozgarLink/internal/pkg/apperrors"
	"github.com/harsh-jagtap-josh/RozgarLink/internal/repo"
)

type applicationService struct {
	applicationRepo repo.ApplicationStorer
}

type Service interface {
	CreateNewApplication(ctx context.Context, applicationData Application) (Application, error)
	UpdateApplicationById(ctx context.Context, applicationData Application) (Application, error)
	FetchApplicationById(ctx context.Context, applicationId int) (Application, error)
	DeleteApplicationById(ctx context.Context, applicationId int) (int, error)
	FetchAllApplications(ctx context.Context) ([]ApplicationComplete, error)
}

func NewService(applicationRepo repo.ApplicationStorer) Service {
	return &applicationService{
		applicationRepo: applicationRepo,
	}
}

func (appS *applicationService) CreateNewApplication(ctx context.Context, applicationData Application) (Application, error) {
	var createApplication Application

	repoAppObj := MapServiceApplicationToRepo(applicationData)

	application, err := appS.applicationRepo.CreateNewApplication(ctx, repoAppObj)
	if err != nil {
		return Application{}, err
	}

	createApplication = MapRepoApplicationToService(application)

	return createApplication, nil
}

func (appS *applicationService) UpdateApplicationById(ctx context.Context, applicationData Application) (Application, error) {
	applRepoObj := MapServiceApplicationToRepo(applicationData)

	application, err := appS.applicationRepo.UpdateApplicationByID(ctx, applRepoObj)
	if err != nil {
		return Application{}, err
	}

	updatedApplication := MapRepoApplicationToService(application)

	return updatedApplication, nil
}

func (appS *applicationService) FetchApplicationById(ctx context.Context, applicationId int) (Application, error) {
	application, err := appS.applicationRepo.FetchApplicationByID(ctx, applicationId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return Application{}, fmt.Errorf("%w, %w", apperrors.ErrNoApplicationExists, err)
		}
		return Application{}, err
	}

	fetchedApplication := MapRepoApplicationToService(application)
	return fetchedApplication, nil
}

func (appS *applicationService) DeleteApplicationById(ctx context.Context, applicationId int) (int, error) {
	exists := appS.applicationRepo.FindApplicationById(ctx, applicationId)
	if !exists {
		return -1, apperrors.ErrNoApplicationExists
	}

	id, err := appS.applicationRepo.DeleteApplicationByID(ctx, applicationId)
	if err != nil {
		return -1, err
	}
	return id, nil
}

func (appS *applicationService) FetchAllApplications(ctx context.Context) ([]ApplicationComplete, error) {
	applications, err := appS.applicationRepo.FetchAllApplications(ctx)
	if err != nil {
		return []ApplicationComplete{}, err
	}

	fetchedApplications := make([]ApplicationComplete, 0)
	for _, appl := range applications {
		fetchedApplications = append(fetchedApplications, MapRepoApplCompToService(appl))
	}

	return fetchedApplications, nil
}
