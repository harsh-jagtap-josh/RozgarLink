package repo

import (
	"context"
	"database/sql"
	"errors"

	"github.com/harsh-jagtap-josh/RozgarLink/internal/pkg/apperrors"
	"github.com/jmoiron/sqlx"
)

type EmployerStorer interface {
	RegisterEmployer(ctx context.Context, employerData EmployerResponse) (EmployerResponse, error)
	FetchEmployerByID(ctx context.Context, employerId int) (EmployerResponse, error)
	UpdateEmployerById(ctx context.Context, employerData EmployerResponse) (EmployerResponse, error)
	DeleteEmployerByID(ctx context.Context, employerId int) (int, error)
	FindEmployerByEmail(ctx context.Context, employerEmail string) bool
	FindEmployerById(ctx context.Context, employerId int) bool
}

// PostgreSQL Queries
const (
	registerWorkerQuery      = `INSERT INTO employers (name, contact_number, email, type, password, sectors, location, is_verified, rating, workers_hired, created_at, updated_at, language) VALUES (:name, :contact_number, :email, :type, :password, :sectors, :location, :is_verified, :rating, :workers_hired, NOW(), NOW(), :language) RETURNING *;`
	fetchEmployerByIDQuery   = `SELECT employers.*, address.details, address.street, address.city, address.state, address.pincode from employers inner join address on employers.location = address.id where employers.id = $1;`
	updateEmployerByIdQuery  = `UPDATE employers SET name=:name, contact_number=:contact_number, email=:email, type=:type, sectors=:sectors, rating=:rating, workers_hired=:workers_hired, is_verified=:is_verified, updated_at=NOW(), language=:language WHERE id=:id RETURNING *;`
	deleteEmployerByIdQuery  = `DELETE from employers where id=$1 RETURNING location;`
	findEmployerByEmailQuery = `SELECT id from employers where email=$1;`
	findEmployerByIDQuery    = `SELECT id from employers where id=$1;`
)

type employerStore struct {
	BaseRepository
}

func NewEmployerRepo(db *sql.DB) EmployerStorer {
	return &employerStore{
		BaseRepository: BaseRepository{db},
	}
}

func (es *employerStore) RegisterEmployer(ctx context.Context, employerData EmployerResponse) (EmployerResponse, error) {
	db := sqlx.NewDb(es.DB, "postgres")

	var newEmployer EmployerResponse

	addressData := Address{
		Details: employerData.Details,
		Street:  employerData.Street,
		City:    employerData.City,
		State:   employerData.State,
	}

	address, err := CreateAddress(ctx, db, addressData)
	if err != nil {
		return EmployerResponse{}, err
	}

	employerData.Location = address.ID

	rows, err := db.NamedQuery(registerWorkerQuery, employerData)
	if err != nil {
		return EmployerResponse{}, err
	}

	defer rows.Close()

	if rows.Next() {
		err = rows.StructScan(&newEmployer)
		if err != nil {
			return EmployerResponse{}, err
		}
	}
	newEmployer = MapAddressToEmployer(newEmployer, address)
	return newEmployer, nil
}

func (es *employerStore) FetchEmployerByID(ctx context.Context, employerId int) (EmployerResponse, error) {
	db := sqlx.NewDb(es.DB, "postgres")

	var employer EmployerResponse

	err := db.Get(&employer, fetchEmployerByIDQuery, employerId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return EmployerResponse{}, apperrors.ErrNoEmployerExists
		}
		return EmployerResponse{}, err
	}

	return employer, nil
}

func (es *employerStore) UpdateEmployerById(ctx context.Context, employerData EmployerResponse) (EmployerResponse, error) {
	db := sqlx.NewDb(es.DB, "postgres")
	var updatedAddress Address
	var employerUpdated EmployerResponse

	address, err := GetAddressById(ctx, db, employerData.Location)
	if err != nil {
		return EmployerResponse{}, nil
	}

	isAddressChanged := !MatchAddressEmployer(address, employerData)
	if isAddressChanged {
		updatedAddress, err = UpdateAddress(ctx, db, Address{
			ID:      address.ID,
			Details: employerData.Details,
			Street:  employerData.Street,
			City:    employerData.City,
			State:   employerData.State,
			Pincode: employerData.Pincode,
		})
		if err != nil {
			return EmployerResponse{}, err
		}
	}

	rows, err := db.NamedQuery(updateEmployerByIdQuery, employerData)
	if err != nil {
		return EmployerResponse{}, err
	}
	defer rows.Close()
	if rows.Next() {
		err = rows.StructScan(&employerUpdated)
		if err != nil {
			return EmployerResponse{}, err
		}
	}

	if isAddressChanged {
		employerUpdated = MapAddressToEmployer(employerUpdated, updatedAddress)
	} else {
		employerUpdated = MapAddressToEmployer(employerUpdated, address)
	}

	return employerUpdated, nil
}

func (es *employerStore) DeleteEmployerByID(ctx context.Context, employerId int) (int, error) {

	db := sqlx.NewDb(es.DB, "postgres")
	var addressId int

	err := db.Get(&addressId, deleteEmployerByIdQuery, employerId)

	if err != nil {
		return -1, err
	}

	err = DeleteAddress(ctx, db, addressId)
	if err != nil {
		return -1, err
	}
	return employerId, nil
}

func (es *employerStore) FindEmployerByEmail(ctx context.Context, employerEmail string) bool {
	var ID int

	err := es.BaseRepository.DB.QueryRow(findEmployerByEmailQuery, employerEmail).Scan(&ID)
	return err == nil
}

func (es *employerStore) FindEmployerById(ctx context.Context, employerId int) bool {
	var ID int
	err := es.BaseRepository.DB.QueryRow(findEmployerByIDQuery, employerId).Scan(&ID)
	return err == nil
}
