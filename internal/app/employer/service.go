package employer

import (
	"context"
	"fmt"

	"github.com/harsh-jagtap-josh/RozgarLink/internal/app/job"
	"github.com/harsh-jagtap-josh/RozgarLink/internal/pkg/apperrors"
	"github.com/harsh-jagtap-josh/RozgarLink/internal/pkg/utils"
	"github.com/harsh-jagtap-josh/RozgarLink/internal/repo"
)

type service struct {
	employerRepo repo.EmployerStorer
}

type Service interface {
	FetchEmployerByID(ctx context.Context, employerId int) (Employer, error)
	UpdateEmployerById(ctx context.Context, employerData Employer) (Employer, error)
	RegisterEmployer(ctx context.Context, employerData Employer) (Employer, error)
	DeleteEmployerById(ctx context.Context, employerId int) (int, error)
	FetchJobsByEmployerId(ctx context.Context, employerId int) ([]job.Job, error)
	FetchAllEmployers(ctx context.Context) ([]Employer, error)
}

func NewService(employerRepo repo.EmployerStorer) Service {
	return &service{
		employerRepo: employerRepo,
	}
}

func (empS *service) FetchEmployerByID(ctx context.Context, employerId int) (Employer, error) {

	response, err := empS.employerRepo.FetchEmployerByID(ctx, employerId)
	if err != nil {
		return Employer{}, err
	}

	employer := MapRepoToServiceDomain(response)

	return employer, nil
}

func (empS *service) UpdateEmployerById(ctx context.Context, employerData Employer) (Employer, error) {

	var updatedEmployer Employer

	err := utils.ValidateUpdateUser(employerData.Name, employerData.ContactNo, employerData.Email)
	if err != nil {
		return Employer{}, fmt.Errorf("%w: %w", apperrors.ErrInvalidUserDetails, err)
	}

	repoEmployerData := MapServiceToRepoDomain(employerData)

	updEmpRepo, err := empS.employerRepo.UpdateEmployerById(ctx, repoEmployerData)
	if err != nil {
		return Employer{}, err
	}

	updatedEmployer = MapRepoToServiceDomain(updEmpRepo)

	return updatedEmployer, nil
}

func (empS *service) RegisterEmployer(ctx context.Context, employerData Employer) (Employer, error) {

	// validate fields of user
	err := utils.ValidateUser(employerData.Name, employerData.ContactNo, employerData.Email, employerData.Password)
	if err != nil {
		return Employer{}, fmt.Errorf("%w: %w", apperrors.ErrInvalidUserDetails, err)
	}

	// check if employer with email already exists
	alreadyExists := empS.employerRepo.FindEmployerByEmail(ctx, employerData.Email)
	if alreadyExists {
		return Employer{}, apperrors.ErrEmployerAlreadyExists
	}

	// create an encrypted password using bcrypt utility function
	hashed_password, err := utils.HashPassword(employerData.Password)
	if err != nil {
		return Employer{}, fmt.Errorf("%w: %w", apperrors.ErrEncrPassword, err)
	}
	employerData.Password = hashed_password

	// Map from Service domain to repo domain struct
	repoEmployerStruct := MapServiceToRepoDomain(employerData)
	employer, err := empS.employerRepo.RegisterEmployer(ctx, repoEmployerStruct)
	if err != nil {
		return Employer{}, err
	}

	newEmployer := MapRepoToServiceDomain(employer)
	return newEmployer, nil
}

func (empS *service) DeleteEmployerById(ctx context.Context, employerId int) (int, error) {
	exists := empS.employerRepo.FindEmployerById(ctx, employerId)
	if !exists {
		return -1, apperrors.ErrNoEmployerExists
	}

	id, err := empS.employerRepo.DeleteEmployerByID(ctx, employerId)
	if err != nil {
		return -1, err
	}

	return id, nil
}

func (es *service) FetchJobsByEmployerId(ctx context.Context, employerId int) ([]job.Job, error) {
	exists := es.employerRepo.FindEmployerById(ctx, employerId)
	if !exists {
		return []job.Job{}, apperrors.ErrNoEmployerExists
	}
	jobs, err := es.employerRepo.FindJobByEmployerId(ctx, employerId)
	if err != nil {
		return []job.Job{}, err
	}
	var mappedJobs []job.Job
	for _, newJob := range jobs {
		mappedJobs = append(mappedJobs, job.MapJobRepoStructToService(newJob))
	}
	return mappedJobs, nil
}

func (es *service) FetchAllEmployers(ctx context.Context) ([]Employer, error) {
	employers, err := es.employerRepo.FetchAllEmployers(ctx)
	if err != nil {
		return []Employer{}, err
	}

	fetchedEmployers := make([]Employer, 0)
	for _, emp := range employers {
		fetchedEmployers = append(fetchedEmployers, MapRepoToServiceDomain(emp))
	}

	return fetchedEmployers, nil
}
