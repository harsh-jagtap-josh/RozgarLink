package employer

import (
	"context"
	"fmt"

	"github.com/harsh-jagtap-josh/RozgarLink/internal/repo"
)

type service struct {
	employerRepo repo.EmployerStorer
}

type Service interface {
	FetchEmployerByID(ctx context.Context, employerId int) (Employer, error)
}

func NewService(employerRepo repo.EmployerStorer) Service {
	return &service{
		employerRepo: employerRepo,
	}
}

func (empS *service) FetchEmployerByID(ctx context.Context, employerId int) (Employer, error) {

	response, err := empS.employerRepo.FetchEmployerByID(ctx, employerId)
	if err != nil {
		fmt.Println(err.Error())
		return Employer{}, err
	}

	employer := MapRepoToServiceDomain(response)

	return employer, nil
}
