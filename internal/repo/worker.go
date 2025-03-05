package repo

import (
	"context"
	"database/sql"
	"errors"

	"github.com/harsh-jagtap-josh/RozgarLink/internal/pkg/apperrors"
	"github.com/jmoiron/sqlx"
)

type WorkerStorer interface {
	FetchWorkerByID(ctx context.Context, workerID int) (Worker, error)
	CreateWorker(ctx context.Context, workerData Worker) (Worker, error)
	UpdateWorkerByID(ctx context.Context, workerData Worker) (Worker, error)
	DeleteWorkerByID(ctx context.Context, workerId int) (int, error)
	FindWorkerByEmail(ctx context.Context, email string) bool
	FindWorkerById(ctx context.Context, id int) bool
	FetchApplicationsByWorkerId(ctx context.Context, workerId int) ([]ApplicationComplete, error)
	FetchAllWorkers(ctx context.Context) ([]Worker, error)
}

type workerStore struct {
	BaseRepository
}

func NewWorkerRepo(db *sqlx.DB) WorkerStorer {
	return &workerStore{
		BaseRepository: BaseRepository{db},
	}
}

// PostgreSQL Queries
const (
	fetchWorkerByIDQuery             = `SELECT workers.id, name, contact_number, email, gender, sectors, skills, location, is_available, rating, total_jobs_worked, created_at, updated_at, language from workers inner join address on workers.location = address.id where workers.id = $1;`
	createWorkerQuery                = `INSERT INTO Workers (name, contact_number, email, gender, password, sectors, skills, location, is_available, rating, total_jobs_worked, created_at, updated_at, language) VALUES (:name, :contact_number, :email, :gender, :password, :sectors, :skills, :location, :is_available, :rating, :total_jobs_worked, NOW(), NOW(), :language) RETURNING *;`
	updateWorkerByIDQuery            = `UPDATE Workers SET name=:name, contact_number=:contact_number, email=:email, gender=:gender, sectors=:sectors, skills=:skills, is_available=:is_available, rating=:rating, total_jobs_worked=:total_jobs_worked, updated_at=NOW(), language=:language WHERE id=:id RETURNING *;`
	deleteWorkerByIdQuery            = `DELETE FROM workers WHERE id=$1 RETURNING location;`
	findEmailExistsQuery             = "SELECT id FROM workers WHERE email = $1;"
	findIdExistsQuery                = "SELECT id FROM workers WHERE id = $1;"
	fetchApplicationsByWorkerIdQuery = `select applications.*, address.details, address.street, address.state, address.city, address.pincode, jobs.title, jobs.description, jobs.skills_required, jobs.sectors, jobs.wage, jobs.vacancy, jobs.date, employers.name, employers.contact_number, employers.email, employers.type from applications inner join address on applications.pick_up_location = address.id inner join jobs on applications.job_id = jobs.id inner join employers on jobs.employer_id = employers.id WHERE applications.worker_id = $1`
	fetchAllWorkersQuery             = `SELECT workers.*, address.details, address.street, address.city, address.state, address.pincode FROM workers inner join address on workers.location = address.id;`
)

// Create a New Worker
func (ws *workerStore) CreateWorker(ctx context.Context, workerData Worker) (Worker, error) {

	var worker Worker
	addressData := Address{
		Details: workerData.Details,
		Street:  workerData.Street,
		City:    workerData.City,
		State:   workerData.State,
		Pincode: worker.Pincode,
	}

	address, err := CreateAddress(ctx, ws.DB, addressData)
	if err != nil {
		return Worker{}, err
	}

	workerData.Location = address.ID

	rows, err := ws.DB.NamedQuery(createWorkerQuery, workerData)
	if err != nil {
		return Worker{}, err
	}

	defer rows.Close()

	if rows.Next() {
		err = rows.StructScan(&worker)
		if err != nil {
			return Worker{}, err
		}
	}
	worker = MapAddressToWorker(worker, address)
	return worker, nil
}

// Fetch Worker Details by Worker ID
func (ws *workerStore) FetchWorkerByID(ctx context.Context, workerID int) (Worker, error) {

	var worker Worker

	err := ws.DB.Get(&worker, fetchWorkerByIDQuery, workerID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return Worker{}, apperrors.ErrNoWorkerExists
		}
		return Worker{}, err
	}

	address, err := GetAddressById(ctx, ws.DB, worker.Location)
	if err != nil {
		return Worker{}, err
	}

	worker = MapAddressToWorker(worker, address)
	return worker, nil
}

// Update Worker Details By ID
func (ws *workerStore) UpdateWorkerByID(ctx context.Context, workerData Worker) (Worker, error) {

	updatedworker := workerData

	address, err := GetAddressByWorkerId(ctx, ws.DB, workerData.ID)
	if err != nil {
		return Worker{}, err
	}
	isAddressChanged := !MatchAddressWorker(address, workerData)
	var updAddress Address
	if isAddressChanged {
		updAddress, err = UpdateAddress(ctx, ws.DB, Address{
			ID:      address.ID,
			Details: workerData.Details,
			Street:  workerData.Street,
			City:    workerData.City,
			State:   workerData.State,
			Pincode: workerData.Pincode,
		})
		if err != nil {
			return Worker{}, err
		}
	}

	rows, err := ws.DB.NamedQuery(updateWorkerByIDQuery, workerData)
	if err != nil {
		return Worker{}, err
	}

	defer rows.Close()

	if rows.Next() {
		err = rows.StructScan(&updatedworker)
		if err != nil {
			return Worker{}, err
		}
	}

	if isAddressChanged {
		updatedworker = MapAddressToWorker(updatedworker, updAddress)
	}

	return updatedworker, nil
}

// Delete Worker data and address
func (ws *workerStore) DeleteWorkerByID(ctx context.Context, workerId int) (int, error) {

	var addressId int

	err := ws.DB.Get(&addressId, deleteWorkerByIdQuery, workerId)

	if err != nil {
		return -1, err
	}

	err = DeleteAddress(ctx, ws.DB, addressId)
	if err != nil {
		return -1, err
	}

	return workerId, nil
}

// Find Worker By Email Exists
func (ws *workerStore) FindWorkerByEmail(ctx context.Context, email string) bool {
	var ID int
	err := ws.BaseRepository.DB.QueryRow(findEmailExistsQuery, email).Scan(&ID)

	return err == nil
}

// Find Worker By ID Exists
func (ws *workerStore) FindWorkerById(ctx context.Context, id int) bool {
	var ID int
	err := ws.BaseRepository.DB.QueryRow(findIdExistsQuery, id).Scan(&ID)

	return err == nil
}

func (ws *workerStore) FetchApplicationsByWorkerId(ctx context.Context, workerId int) ([]ApplicationComplete, error) {

	var applications []ApplicationComplete

	err := ws.DB.Select(&applications, fetchApplicationsByWorkerIdQuery, workerId)
	if err != nil {
		return []ApplicationComplete{}, err
	}
	return applications, nil
}

func (ws *workerStore) FetchAllWorkers(ctx context.Context) ([]Worker, error) {
	workers := make([]Worker, 0)
	err := ws.DB.Select(&workers, fetchAllWorkersQuery)
	if err != nil {
		return []Worker{}, err
	}
	return workers, nil

}
