package job

import (
	"github.com/harsh-jagtap-josh/RozgarLink/internal/app/worker"
	"github.com/harsh-jagtap-josh/RozgarLink/internal/repo"
)

func MapJobRepoStructToService(job repo.Job) Job {
	return Job{
		ID:              job.ID,
		EmployerID:      job.EmployerID,
		Title:           job.Title,
		RequiredGender:  job.RequiredGender,
		Description:     job.Description,
		DurationInHours: job.DurationInHours,
		SkillsRequired:  job.SkillsRequired,
		Sectors:         job.Sectors,
		Wage:            job.Wage,
		Vacancy:         job.Vacancy,
		Location: worker.Address{
			ID:      job.Location,
			Details: job.Details,
			Street:  job.Street,
			City:    job.City,
			State:   job.State,
			Pincode: job.Pincode,
		},
		CreatedAt: job.CreatedAt,
		UpdatedAt: job.UpdatedAt,
	}
}

func MapJobServiceStructToRepo(job Job) repo.Job {
	return repo.Job{
		ID:              job.ID,
		EmployerID:      job.EmployerID,
		Title:           job.Title,
		RequiredGender:  job.RequiredGender,
		Description:     job.Description,
		DurationInHours: job.DurationInHours,
		SkillsRequired:  job.SkillsRequired,
		Sectors:         job.Sectors,
		Wage:            job.Wage,
		Vacancy:         job.Vacancy,
		Location:        job.Location.ID,
		CreatedAt:       job.CreatedAt,
		UpdatedAt:       job.UpdatedAt,
		Details:         job.Location.Details,
		Street:          job.Location.Street,
		City:            job.Location.City,
		State:           job.Location.State,
		Pincode:         job.Location.Pincode,
	}
}
