package job

import (
	"context"

	"github.com/harsh-jagtap-josh/RozgarLink/internal/pkg/apperrors"
	"github.com/harsh-jagtap-josh/RozgarLink/internal/repo"
)

type jobService struct {
	jobRepo repo.JobStorer
}

type Service interface {
	CreateJob(ctx context.Context, jobData Job) (Job, error)
	UpdateJobByID(ctx context.Context, jobData Job) (Job, error)
	FetchJobByID(ctx context.Context, jobId int) (Job, error)
	DeleteJobByID(ctx context.Context, jobId int) (int, error)
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

func (js jobService) FetchJobByID(ctx context.Context, jobId int) (Job, error) {
	job, err := js.jobRepo.FetchJobById(ctx, jobId)
	if err != nil {
		return Job{}, err
	}

	fetchedJob := MapJobRepoStructToService(job)
	return fetchedJob, nil
}

func (js jobService) DeleteJobByID(ctx context.Context, jobId int) (int, error) {
	exists := js.jobRepo.FindJobById(ctx, jobId)
	if !exists {
		return -1, apperrors.ErrNoJobExists
	}

	id, err := js.jobRepo.DeleteJobById(ctx, jobId)
	if err != nil {
		return -1, err
	}
	return id, nil
}
