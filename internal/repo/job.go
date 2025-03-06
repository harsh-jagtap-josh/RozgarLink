package repo

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

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
	FetchApplicationsByJobId(ctx context.Context, jobId int) ([]ApplicationCompleteEmp, error)
	FetchAllJobs(ctx context.Context, filters JobFilters) ([]Job, error)
}

// PostgreSQL Queries
const (
	createJobQuery                = `INSERT INTO jobs (employer_id, title, required_gender, location, description, duration_in_hours, skills_required, sectors, wage, vacancy, date, start_hour, end_hour, created_at, updated_at) VALUES (:employer_id, :title, :required_gender, :location, :description, :duration_in_hours, :skills_required, :sectors, :wage, :vacancy, :date, :start_hour, :end_hour, NOW(), NOW()) RETURNING *;`
	updateJobByIdQuery            = `UPDATE jobs SET title=:title, required_gender=:required_gender, description=:description, duration_in_hours=:duration_in_hours, skills_required=:skills_required, sectors=:sectors, wage=:wage, vacancy=:vacancy, date=:date, start_hour=:start_hour, end_hour=:end_hour, updated_at=NOW() where id=:id RETURNING *;`
	fetchJobByIdQuery             = `SELECT jobs.*, address.details, address.street, address.city, address.state, address.pincode from jobs inner join address on jobs.location = address.id where jobs.id = $1;`
	deleteJobByIdQuery            = `DELETE FROM jobs WHERE id=$1 RETURNING location;`
	findJobByIdQuery              = `SELECT id FROM jobs WHERE id = $1;`
	fetchApplicationsByJobIdQuery = `select applications.*, address.details, address.street, address.state, address.city, address.pincode, jobs.title, jobs.description, jobs.skills_required, jobs.sectors, jobs.wage, jobs.vacancy, jobs.date, workers.name, workers.contact_number, workers.email, workers.gender from applications inner join address on applications.pick_up_location = address.id inner join jobs on applications.job_id = jobs.id inner join workers on applications.worker_id = workers.id where applications.job_id = $1;`
)

// Create New Job
func (jobS *jobStore) CreateJob(ctx context.Context, jobData Job) (Job, error) {
	var createdJob Job

	addressData := Address{
		Details: jobData.Details,
		Street:  jobData.Street,
		City:    jobData.City,
		State:   jobData.State,
		Pincode: jobData.Pincode,
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
		return Job{}, err
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

// fetch job applications
func (jobS *jobStore) FetchApplicationsByJobId(ctx context.Context, jobId int) ([]ApplicationCompleteEmp, error) {

	var applications []ApplicationCompleteEmp

	err := jobS.DB.Select(&applications, fetchApplicationsByJobIdQuery, jobId)
	if err != nil {
		return []ApplicationCompleteEmp{}, err
	}

	return applications, nil
}

func (jobS *jobStore) FetchAllJobs(ctx context.Context, filters JobFilters) ([]Job, error) {
	var jobs []Job
	query := `SELECT jobs.*, address.details, address.street, address.city, address.state, address.pincode FROM jobs INNER JOIN address ON jobs.location = address.id WHERE 1=1`
	args := []interface{}{}
	argIndex := 1

	// Apply filters dynamically
	if len(filters.Title) > 0 {
		query += fmt.Sprintf(" AND jobs.title ILIKE $%d", argIndex)
		args = append(args, "%"+filters.Title+"%")
		argIndex++
	}
	if len(filters.Sector) > 0 {
		query += fmt.Sprintf(" AND jobs.sectors ILIKE $%d", argIndex)
		args = append(args, "%"+filters.Sector+"%")
		argIndex++
	}
	if filters.WageMin > 0 {
		query += fmt.Sprintf(" AND jobs.wage >= $%d", argIndex)
		args = append(args, filters.WageMin)
		argIndex++
	}
	if filters.WageMax > 0 {
		query += fmt.Sprintf(" AND jobs.wage <= $%d", argIndex)
		args = append(args, filters.WageMax)
		argIndex++
	}
	if !filters.StartDate.IsZero() {
		query += fmt.Sprintf(" AND jobs.date >= $%d", argIndex)
		args = append(args, filters.StartDate)
		argIndex++
	}
	if !filters.EndDate.IsZero() {
		query += fmt.Sprintf(" AND jobs.date <= $%d", argIndex)
		args = append(args, filters.EndDate)
		argIndex++
	}
	if len(filters.City) > 0 {
		query += fmt.Sprintf(" AND address.city ILIKE $%d", argIndex)
		args = append(args, "%"+filters.City+"%")
		argIndex++
	}
	if len(filters.Gender) > 0 {
		query += fmt.Sprintf(" AND jobs.required_gender = $%d", argIndex)
		args = append(args, filters.Gender)
		argIndex++
	}

	err := jobS.DB.Select(&jobs, query, args...)
	if err != nil {
		return []Job{}, err
	}
	return jobs, nil
}
