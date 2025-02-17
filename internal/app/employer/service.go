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
