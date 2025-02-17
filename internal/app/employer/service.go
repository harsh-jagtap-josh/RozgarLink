package employer

import (
	"context"

	"github.com/harsh-jagtap-josh/RozgarLink/internal/pkg/apperrors"
	"github.com/harsh-jagtap-josh/RozgarLink/internal/pkg/utils"
	"github.com/harsh-jagtap-josh/RozgarLink/internal/repo"
)

type service struct {
	employerRepo repo.EmployerStorer
}

type Service interface {
	FetchEmployerByID(ctx context.Context, employerId int) (Employer, error)
	UpdateEmployerById(ctx context.Context, employerData Employer) (Employer, error, error)
	RegisterEmployer(ctx context.Context, employerData Employer) (Employer, error, error)
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

func (empS *service) UpdateEmployerById(ctx context.Context, employerData Employer) (Employer, error, error) {

	var updatedEmployer Employer

	errType, err := utils.ValidateUpdateUser(employerData.Name, employerData.ContactNo, employerData.Email)
	if err != nil {
		return Employer{}, err, errType
	}

	repoEmployerData := MapServiceToRepoDomain(employerData)

	updEmpRepo, err := empS.employerRepo.UpdateEmployerById(ctx, repoEmployerData)
	if err != nil {
		return Employer{}, err, apperrors.ErrUpdateEmployer
	}

	updatedEmployer = MapRepoToServiceDomain(updEmpRepo)

	return updatedEmployer, nil, nil
}

func (empS *service) RegisterEmployer(ctx context.Context, employerData Employer) (Employer, error, error) {

	// validate fields of user
	errType, err := utils.ValidateUser(employerData.Name, employerData.ContactNo, employerData.Email, employerData.Password)
	if err != nil {
		return Employer{}, err, errType
	}

	// check if employer with email already exists
	alreadyExists := empS.employerRepo.FindEmployerByEmail(ctx, employerData.Email)
	if alreadyExists {
		return Employer{}, apperrors.ErrCreateEmployer, apperrors.ErrEmployerAlreadyExists
	}

	// create an encrypted password using bcrypt utility function
	hashed_password, err := utils.HashPassword(employerData.Password)
	if err != nil {
		return Employer{}, err, apperrors.ErrEncrPassword
	}
	employerData.Password = hashed_password

	// Map from Service domain to repo domain struct
	repoEmployerStruct := MapServiceToRepoDomain(employerData)
	employer, err := empS.employerRepo.RegisterEmployer(ctx, repoEmployerStruct)
	if err != nil {
		return Employer{}, nil, nil
	}

	newEmployer := MapRepoToServiceDomain(employer)
	return newEmployer, nil, nil
}
