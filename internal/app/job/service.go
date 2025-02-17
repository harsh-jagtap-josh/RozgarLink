package job

import (
	"context"

	"github.com/harsh-jagtap-josh/RozgarLink/internal/repo"
)

type jobService struct {
	jobRepo repo.JobStorer
}

type Service interface {
	CreateJob(ctx context.Context, jobData Job) (Job, error)
	UpdateJobByID(ctx context.Context, jobData Job) (Job, error)
}

func NewService(jobRepo repo.JobStorer) Service {
	return jobService{
		jobRepo: jobRepo,
	}
}

func (js jobService) CreateJob(ctx context.Context, jobData Job) (Job, error) {

	jobRepoObj := MapJobServiceStructToRepo(jobData)

	job, err := js.jobRepo.CreateJob(ctx, jobRepoObj)
	if err != nil {
		return Job{}, err
	}

	createdJob := MapJobRepoStructToService(job)

	return createdJob, nil
}

func (js jobService) UpdateJobByID(ctx context.Context, jobData Job) (Job, error) {

	jobRepoObj := MapJobServiceStructToRepo(jobData)

	job, err := js.jobRepo.UpdateJobById(ctx, jobRepoObj)
	if err != nil {
		return Job{}, err
	}

	updatedJob := MapJobRepoStructToService(job)

	return updatedJob, nil
}
