package repo

import (
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"
)

type jobStore struct {
	BaseRepository
}

func NewJobRepo(db *sql.DB) JobStorer {
	return &jobStore{
		BaseRepository: BaseRepository{
			DB: db,
		},
	}
}

type JobStorer interface {
	CreateJob(ctx context.Context, jobData JobRepoStruct) (JobRepoStruct, error)
	UpdateJobById(ctx context.Context, jobData JobRepoStruct) (JobRepoStruct, error)
}

// PostgreSQL Queries
const ()

// Create New Job
func (jobS *jobStore) CreateJob(ctx context.Context, jobData JobRepoStruct) (JobRepoStruct, error) {
	db := sqlx.NewDb(jobS.DB, "postgres")

	var createdJob JobRepoStruct

	query := `INSERT INTO jobs (employer_id, title, required_gender, location, description, duration_in_hours, skills_required, sectors, wage, vacancy, created_at, updated_at) 
		VALUES (:employer_id, :title, :required_gender, :location, :description, :duration_in_hours, :skills_required, :sectors, :wage, :vacancy, NOW(), NOW()) RETURNING *;
	`

	addressData := Address{
		Details: jobData.Details,
		Street:  jobData.Street,
		City:    jobData.City,
		State:   jobData.State,
	}

	address, err := CreateAddress(ctx, db, addressData)
	if err != nil {
		return JobRepoStruct{}, err
	}

	jobData.Location = address.ID

	rows, err := db.NamedQuery(query, jobData)
	if err != nil {
		return JobRepoStruct{}, err
	}

	defer rows.Close()

	if rows.Next() {
		err = rows.StructScan(&createdJob)
		if err != nil {
			return JobRepoStruct{}, err
		}
	}
	createdJob = MapAddressToJob(createdJob, address)
	return createdJob, nil
}

func (jobS *jobStore) UpdateJobById(ctx context.Context, jobData JobRepoStruct) (JobRepoStruct, error) {
	db := sqlx.NewDb(jobS.DB, "postgres")

	var updatedJob JobRepoStruct
	var updatedAddress Address

	address, err := GetAddressById(ctx, db, jobData.Location)
	if err != nil {
		return JobRepoStruct{}, nil
	}

	isAddressChanged := !MatchAddressJob(address, jobData)
	if isAddressChanged {
		updatedAddress, err = UpdateAddress(ctx, db, Address{
			ID:      address.ID,
			Details: jobData.Details,
			Street:  jobData.Street,
			City:    jobData.City,
			State:   jobData.State,
			Pincode: jobData.Pincode,
		})
		if err != nil {
			return JobRepoStruct{}, err
		}
	}

	const updateJobByIdQuery = `UPDATE jobs SET title=:title, required_gender=:required_gender, description=:description, duration_in_hours=:duration_in_hours, skills_required=:skills_required, sectors=:sectors, wage=:wage, vacancy=:vacancy, updated_at=NOW() where id=:id RETURNING *;`
	rows, err := db.NamedQuery(updateJobByIdQuery, jobData)
	if err != nil {
		return JobRepoStruct{}, err
	}

	defer rows.Close()
	if rows.Next() {
		err = rows.StructScan(&updatedJob)
		if err != nil {
			return JobRepoStruct{}, err
		}
	}

	if isAddressChanged {
		updatedJob = MapAddressToJob(updatedJob, updatedAddress)
	} else {
		updatedJob = MapAddressToJob(updatedJob, address)
	}

	return updatedJob, nil
}
