package repo

import (
	"context"
	"database/sql"
	"errors"

	"github.com/harsh-jagtap-josh/RozgarLink/internal/pkg/apperrors"
	"github.com/jmoiron/sqlx"
)

type jobStore struct {
	BaseRepository
}

func NewJobRepo(db *sqlx.DB) JobStorer {
	return &jobStore{
		BaseRepository: BaseRepository{
			DB: db,
		},
	}
}

type JobStorer interface {
	CreateJob(ctx context.Context, jobData Job) (Job, error)
	UpdateJobById(ctx context.Context, jobData Job) (Job, error)
	FetchJobById(ctx context.Context, jobId int) (Job, error)
	DeleteJobById(ctx context.Context, jobId int) (int, error)
	FindJobById(ctx context.Context, jobId int) bool
}

// PostgreSQL Queries
const (
	createJobQuery             = `INSERT INTO jobs (employer_id, title, required_gender, location, description, duration_in_hours, skills_required, sectors, wage, vacancy, created_at, updated_at) VALUES (:employer_id, :title, :required_gender, :location, :description, :duration_in_hours, :skills_required, :sectors, :wage, :vacancy, NOW(), NOW()) RETURNING *;`
	updateJobByIdQuery         = `UPDATE jobs SET title=:title, required_gender=:required_gender, description=:description, duration_in_hours=:duration_in_hours, skills_required=:skills_required, sectors=:sectors, wage=:wage, vacancy=:vacancy, updated_at=NOW() where id=:id RETURNING *;`
	fetchJobByIdQuery          = `SELECT jobs.*, address.details, address.street, address.city, address.state, address.pincode from jobs inner join address on jobs.location = address.id where jobs.id = $1;`
	deleteJobByIdQuery         = `DELETE FROM jobs WHERE id=$1 RETURNING location;`
	findJobByIdQuery           = `SELECT id FROM jobs WHERE id = $1;`
	fetchJobsByIdEmployerQuery = `SELECT jobs.*, address.details, address.street, address.city, address.state, address.pincode from jobs inner join address on jobs.location = address.id where jobs.employer_id = $1;`
)

// Create New Job
func (jobS *jobStore) CreateJob(ctx context.Context, jobData Job) (Job, error) {

	var createdJob Job

	addressData := Address{
		Details: jobData.Details,
		Street:  jobData.Street,
		City:    jobData.City,
		State:   jobData.State,
	}

	address, err := CreateAddress(ctx, jobS.DB, addressData)
	if err != nil {
		return Job{}, err
	}

	jobData.Location = address.ID

	rows, err := jobS.DB.NamedQuery(createJobQuery, jobData)
	if err != nil {
		return Job{}, err
	}

	defer rows.Close()

	if rows.Next() {
		err = rows.StructScan(&createdJob)
		if err != nil {
			return Job{}, err
		}
	}
	createdJob = MapAddressToJob(createdJob, address)
	return createdJob, nil
}

// Update Job
func (jobS *jobStore) UpdateJobById(ctx context.Context, jobData Job) (Job, error) {
	var updatedJob Job
	var updatedAddress Address

	address, err := GetAddressById(ctx, jobS.DB, jobData.Location)
	if err != nil {
		return Job{}, nil
	}

	isAddressChanged := !MatchAddressJob(address, jobData)
	if isAddressChanged {
		updatedAddress, err = UpdateAddress(ctx, jobS.DB, Address{
			ID:      address.ID,
			Details: jobData.Details,
			Street:  jobData.Street,
			City:    jobData.City,
			State:   jobData.State,
			Pincode: jobData.Pincode,
		})
		if err != nil {
			return Job{}, err
		}
	}

	rows, err := jobS.DB.NamedQuery(updateJobByIdQuery, jobData)
	if err != nil {
		return Job{}, err
	}

	defer rows.Close()
	if rows.Next() {
		err = rows.StructScan(&updatedJob)
		if err != nil {
			return Job{}, err
		}
	}

	if isAddressChanged {
		updatedJob = MapAddressToJob(updatedJob, updatedAddress)
	} else {
		updatedJob = MapAddressToJob(updatedJob, address)
	}

	return updatedJob, nil
}

// Fetch Job Data by ID
func (jobS *jobStore) FetchJobById(ctx context.Context, jobId int) (Job, error) {

	var job Job
	err := jobS.DB.Get(&job, fetchJobByIdQuery, jobId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return Job{}, apperrors.ErrNoJobExists
		}

		return Job{}, err
	}

	return job, nil
}

// Delete Job by ID
func (jobS *jobStore) DeleteJobById(ctx context.Context, jobId int) (int, error) {
	var addressId int

	err := jobS.DB.Get(&addressId, deleteJobByIdQuery, jobId)

	if err != nil {
		return -1, err
	}

	err = DeleteAddress(ctx, jobS.DB, addressId)
	if err != nil {
		return -1, err
	}

	return jobId, nil
}

// Find Job By ID
func (jobS *jobStore) FindJobById(ctx context.Context, jobId int) bool {
	var ID int
	err := jobS.BaseRepository.DB.QueryRow(findJobByIdQuery, jobId).Scan(&ID)

	return err == nil
}

// Find Jobs By Employer ID
func (jobS *jobStore) FindJobByEmployerId(ctx context.Context, employerId int) ([]Job, error) {
	var jobs []Job
	err := jobS.DB.Select(&jobs, fetchJobsByIdEmployerQuery, employerId)
	if err != nil {
		return []Job{}, err
	}
	return jobs, nil
}
