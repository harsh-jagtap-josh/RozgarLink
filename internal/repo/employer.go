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
	FindEmployerByEmail(ctx context.Context, employerEmail string) bool
}

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

	query := `INSERT INTO employers (name, contact_number, email, type, password, sectors, location, is_verified, rating, workers_hired, created_at, updated_at, language) VALUES (:name, :contact_number, :email, :type, :password, :sectors, :location, :is_verified, :rating, :workers_hired, NOW(), NOW(), :language) RETURNING *;`

	rows, err := db.NamedQuery(query, employerData)
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
	query := `SELECT employers.*, address.details, address.street, address.city, address.state, address.pincode from employers inner join address on employers.location = address.id where employers.id = $1;`

	err := db.Get(&employer, query, employerId)
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

	query := `UPDATE employers SET name=:name, contact_number=:contact_number, email=:email, type=:type, sectors=:sectors, rating=:rating, workers_hired=:workers_hired, is_verified=:is_verified, updated_at=NOW(), language=:language WHERE id=:id RETURNING *;`
	rows, err := db.NamedQuery(query, employerData)
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

func (es *employerStore) FindEmployerByEmail(ctx context.Context, employerEmail string) bool {
	var ID int
	query := `SELECT id from employers where email=$1;`
	err := es.BaseRepository.DB.QueryRow(query, employerEmail).Scan(&ID)
	return err == nil
}
