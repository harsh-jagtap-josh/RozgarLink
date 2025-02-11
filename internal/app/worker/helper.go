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
		Location:        *repoWorker.Location, // the use of pointers helps in storing NULL values as its type is Integer a null is db can't be stored in struct as null it can only be int hence we use pointers, here pointers can be null..
		IsAvailable:     repoWorker.IsAvailable,
		Rating:          repoWorker.Rating,
		TotalJobsWorked: repoWorker.TotalJobsWorked,
		CreatedAt:       repoWorker.CreatedAt,
		UpdatedAt:       repoWorker.UpdatedAt,
	}
}
