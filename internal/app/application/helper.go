package application

import (
	"github.com/harsh-jagtap-josh/RozgarLink/internal/repo"
)

func MapServiceApplicationToRepo(application Application) repo.Application {
	return repo.Application{
		ID:             application.ID,
		JobID:          application.JobID,
		WorkerID:       application.WorkerID,
		Status:         repo.Status(application.Status),
		ExpectedWage:   application.ExpectedWage,
		ModeOfArrival:  repo.ModeOfArrival(application.ModeOfArrival),
		PickUpLocation: application.PickUpLocation.ID,
		WorkerComment:  application.WorkerComment,
		AppliedAt:      application.AppliedAt,
		UpdatedAt:      application.UpdatedAt,
		Details:        application.PickUpLocation.Details,
		Street:         application.PickUpLocation.Street,
		City:           application.PickUpLocation.City,
		State:          application.PickUpLocation.State,
		Pincode:        application.PickUpLocation.Pincode,
	}
}

func MapRepoApplicationToService(application repo.Application) Application {
	return Application{
		ID:            application.ID,
		JobID:         application.JobID,
		WorkerID:      application.WorkerID,
		Status:        Status(application.Status),
		ExpectedWage:  application.ExpectedWage,
		ModeOfArrival: ModeOfArrival(application.ModeOfArrival),
		PickUpLocation: Address{
			ID:      application.PickUpLocation,
			Details: application.Details,
			Street:  application.Street,
			City:    application.City,
			State:   application.State,
			Pincode: application.Pincode,
		},
		WorkerComment: application.WorkerComment,
		AppliedAt:     application.AppliedAt,
		UpdatedAt:     application.UpdatedAt,
	}
}
