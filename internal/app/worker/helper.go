package worker

import "github.com/harsh-jagtap-josh/RozgarLink/internal/repo"

func Mapper(repoWorker repo.Worker) Worker {
	return Worker{
		ID:              repoWorker.ID,
		Name:            repoWorker.Name,
		ContactNumber:   repoWorker.ContactNumber,
		Email:           repoWorker.Email,
		Gender:          Gender(repoWorker.Gender),
		Sectors:         repoWorker.Sectors,
		Skills:          repoWorker.Skills,
		Location:        *repoWorker.Location,
		IsAvailable:     repoWorker.IsAvailable,
		Rating:          repoWorker.Rating,
		TotalJobsWorked: repoWorker.TotalJobsWorked,
		CreatedAt:       repoWorker.CreatedAt,
		UpdatedAt:       repoWorker.UpdatedAt,
	}
}
