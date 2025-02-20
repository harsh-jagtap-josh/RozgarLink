package worker

import "github.com/harsh-jagtap-josh/RozgarLink/internal/repo"

func MapRepoDomainToService(repoWorker repo.Worker) WorkerRequest {
	return WorkerRequest{
		ID:            repoWorker.ID,
		Name:          repoWorker.Name,
		ContactNumber: repoWorker.ContactNumber,
		Email:         repoWorker.Email,
		Gender:        Gender(repoWorker.Gender),
		Sectors:       repoWorker.Sectors,
		Skills:        repoWorker.Skills,
		Location: Address{
			ID:      repoWorker.Location,
			Details: repoWorker.Details,
			Street:  repoWorker.Street,
			City:    repoWorker.City,
			State:   repoWorker.State,
			Pincode: repoWorker.Pincode,
		},
		IsAvailable:     repoWorker.IsAvailable,
		Rating:          repoWorker.Rating,
		TotalJobsWorked: repoWorker.TotalJobsWorked,
		CreatedAt:       repoWorker.CreatedAt,
		UpdatedAt:       repoWorker.UpdatedAt,
	}
}

func MapServiceDomainToRepo(Worker WorkerRequest) repo.Worker {
	return repo.Worker{
		ID:              Worker.ID,
		Name:            Worker.Name,
		ContactNumber:   Worker.ContactNumber,
		Email:           Worker.Email,
		Gender:          repo.Gender(Worker.Gender),
		Password:        Worker.Password,
		Sectors:         Worker.Sectors,
		Skills:          Worker.Skills,
		Location:        Worker.Location.ID,
		IsAvailable:     Worker.IsAvailable,
		Rating:          Worker.Rating,
		TotalJobsWorked: Worker.TotalJobsWorked,
		CreatedAt:       Worker.CreatedAt,
		UpdatedAt:       Worker.UpdatedAt,
		Details:         Worker.Location.Details,
		Street:          Worker.Location.Street,
		City:            Worker.Location.City,
		State:           Worker.Location.State,
		Pincode:         Worker.Location.Pincode,
	}
}
