package repo

import (
	"context"
	"database/sql"
	"errors"

	"github.com/harsh-jagtap-josh/RozgarLink/internal/pkg/apperrors"
	"github.com/jmoiron/sqlx"
)

type WorkerStorer interface {
	FetchWorkerByID(ctx context.Context, workerID int) (WorkerResponse, error)
	CreateWorker(ctx context.Context, workerData WorkerResponse) (WorkerResponse, error)
	UpdateWorkerByID(ctx context.Context, workerData WorkerResponse) (WorkerResponse, error)
	DeleteWorkerByID(ctx context.Context, workerId int) (int, error)
	FindWorkerByEmail(ctx context.Context, email string) bool
	FindWorkerById(ctx context.Context, id int) bool
}

type workerStore struct {
	BaseRepository
}

func NewWorkerRepo(db *sql.DB) WorkerStorer {
	return &workerStore{
		BaseRepository: BaseRepository{db},
	}
}

// PostgreSQL Queries
const (
	fetchWorkerByIDQuery  = `SELECT workers.id, name, contact_number, email, gender, sectors, skills, location, is_available, rating, total_jobs_worked, created_at, updated_at, language from workers inner join address on workers.location = address.id where workers.id = $1;`
	createWorkerQuery     = `INSERT INTO Workers (name, contact_number, email, gender, password, sectors, skills, location, is_available, rating, total_jobs_worked, created_at, updated_at, language) VALUES (:name, :contact_number, :email, :gender, :password, :sectors, :skills, :location, :is_available, :rating, :total_jobs_worked, NOW(), NOW(), :language) RETURNING *;`
	updateWorkerByIDQuery = `UPDATE Workers SET name=:name, contact_number=:contact_number, email=:email, gender=:gender, sectors=:sectors, skills=:skills, is_available=:is_available, rating=:rating, total_jobs_worked=:total_jobs_worked, updated_at=NOW(), language=:language WHERE id=:id RETURNING *;`
	deleteWorkerByIdQuery = `DELETE FROM workers WHERE id=$1 RETURNING location;`
	findEmailExistsQuery  = "SELECT id FROM workers WHERE email = $1;"
	findIdExistsQuery     = "SELECT id FROM workers WHERE id = $1;"
)

// Create a New Worker
func (ws *workerStore) CreateWorker(ctx context.Context, workerData WorkerResponse) (WorkerResponse, error) {

	db := sqlx.NewDb(ws.DB, "postgres")

	var worker WorkerResponse
	addressData := Address{
		Details: workerData.Details,
		Street:  workerData.Street,
		City:    workerData.City,
		State:   workerData.State,
	}

	address, err := CreateAddress(ctx, db, addressData)
	if err != nil {
		return WorkerResponse{}, err
	}

	workerData.Location = address.ID

	rows, err := db.NamedQuery(createWorkerQuery, workerData)
	if err != nil {
		return WorkerResponse{}, err
	}

	defer rows.Close()

	if rows.Next() {
		err = rows.StructScan(&worker)
		if err != nil {
			return WorkerResponse{}, err
		}
	}
	worker = MapAddressToWorker(worker, address)
	return worker, nil
}

// Fetch Worker Details by Worker ID
func (ws *workerStore) FetchWorkerByID(ctx context.Context, workerID int) (WorkerResponse, error) {

	db := sqlx.NewDb(ws.DB, "postgres")

	var worker WorkerResponse

	err := db.Get(&worker, fetchWorkerByIDQuery, workerID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return WorkerResponse{}, apperrors.ErrNoWorkerExists
		}
		return WorkerResponse{}, err
	}

	address, err := GetAddressById(ctx, db, worker.Location)
	if err != nil {
		return WorkerResponse{}, err
	}

	worker = MapAddressToWorker(worker, address)
	return worker, nil
}

// Update Worker Details By ID
func (ws *workerStore) UpdateWorkerByID(ctx context.Context, workerData WorkerResponse) (WorkerResponse, error) {
	db := sqlx.NewDb(ws.DB, "postgres")

	updatedworker := workerData

	address, err := GetAddressByWorkerId(ctx, db, workerData.ID)
	if err != nil {
		return WorkerResponse{}, err
	}
	isAddressChanged := !MatchAddress(address, workerData)
	var updAddress Address
	if isAddressChanged {
		updAddress, err = UpdateAddress(ctx, db, Address{
			ID:      address.ID,
			Details: workerData.Details,
			Street:  workerData.Street,
			City:    workerData.City,
			State:   workerData.State,
			Pincode: workerData.Pincode,
		})
		if err != nil {
			return WorkerResponse{}, err
		}
	}

	rows, err := db.NamedQuery(updateWorkerByIDQuery, workerData)
	if err != nil {
		return WorkerResponse{}, err
	}

	defer rows.Close()

	if rows.Next() {
		err = rows.StructScan(&updatedworker)
		if err != nil {
			return WorkerResponse{}, err
		}
	}

	if isAddressChanged {
		updatedworker = MapAddressToWorker(updatedworker, updAddress)
	}

	return updatedworker, nil
}

// Delete Worker data and address
func (ws *workerStore) DeleteWorkerByID(ctx context.Context, workerId int) (int, error) {

	db := sqlx.NewDb(ws.DB, "postgres")
	var addressId int

	err := db.Get(&addressId, deleteWorkerByIdQuery, workerId)

	if err != nil {
		return -1, nil
	}

	err = DeleteAddress(ctx, db, addressId)
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
