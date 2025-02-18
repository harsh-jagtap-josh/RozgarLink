package repo

import (
	"context"
	"database/sql"
	"errors"

	"github.com/harsh-jagtap-josh/RozgarLink/internal/pkg/apperrors"
	"github.com/jmoiron/sqlx"
)

type applicationStore struct {
	BaseRepository
}

type ApplicationStorer interface {
	CreateNewApplication(ctx context.Context, applicationData Application) (Application, error)
	UpdateApplicationByID(ctx context.Context, applicationData Application) (Application, error)
	FetchApplicationByID(ctx context.Context, applicationId int) (Application, error)
	DeleteApplicationByID(ctx context.Context, applicationId int) (int, error)
	FindApplicationById(ctx context.Context, applicationId int) bool
}

func NewApplicationRepo(db *sql.DB) ApplicationStorer {
	return &applicationStore{
		BaseRepository: BaseRepository{DB: db},
	}
}

// PostgreSQL Queries
const (
	createApplicationQuery     = `INSERT INTO applications (job_id, worker_id, status, expected_wage, mode_of_arrival, pick_up_location, worker_comments, applied_at, updated_at) VALUES (:job_id, :worker_id, :status, :expected_wage, :mode_of_arrival, :pick_up_location, :worker_comments, NOW(), NOW()) RETURNING *;`
	updateApplicationByIdQuery = `UPDATE applications SET status=:status, expected_wage=:expected_wage, mode_of_arrival=:mode_of_arrival, pick_up_location=:pick_up_location, worker_comments=:worker_comments, updated_at=NOW() where id=:id RETURNING *;`
	fethcApplicationByIdQuery  = `SELECT applications.*, address.details, address.street, address.city, address.state, address.pincode from applications inner join address on applications.pick_up_location = address.id where applications.id = $1;`
	deleteApplicationByIdQuery = `DELETE FROM applications WHERE id=$1 RETURNING pick_up_location;`
	findApplicationByIdQuery   = `SELECT id FROM applications WHERE id = $1;`
)

func (appS *applicationStore) CreateNewApplication(ctx context.Context, applicationData Application) (Application, error) {
	db := sqlx.NewDb(appS.DB, "postgres")

	var createdApplication Application

	addressData := Address{
		Details: applicationData.Details,
		Street:  applicationData.Street,
		City:    applicationData.City,
		State:   applicationData.State,
	}

	address, err := CreateAddress(ctx, db, addressData)
	if err != nil {
		return Application{}, err
	}

	applicationData.PickUpLocation = address.ID

	rows, err := db.NamedQuery(createApplicationQuery, applicationData)
	if err != nil {
		return Application{}, err
	}

	defer rows.Close()

	if rows.Next() {
		err = rows.StructScan(&createdApplication)
		if err != nil {
			return Application{}, err
		}
	}

	return createdApplication, nil
}

func (appS *applicationStore) UpdateApplicationByID(ctx context.Context, applicationData Application) (Application, error) {
	db := sqlx.NewDb(appS.DB, "postgres")

	var updatedApplication Application
	var updatedAddress Address

	address, err := GetAddressById(ctx, db, applicationData.PickUpLocation)
	if err != nil {
		return Application{}, nil
	}

	isAddressChanged := !MatchAddressApplication(address, applicationData)
	if isAddressChanged {
		updatedAddress, err = UpdateAddress(ctx, db, Address{
			ID:      address.ID,
			Details: applicationData.Details,
			Street:  applicationData.Street,
			City:    applicationData.City,
			State:   applicationData.State,
			Pincode: applicationData.Pincode,
		})
		if err != nil {
			return Application{}, err
		}
	}

	rows, err := db.NamedQuery(updateApplicationByIdQuery, applicationData)
	if err != nil {
		return Application{}, err
	}

	defer rows.Close()
	if rows.Next() {
		err = rows.StructScan(&updatedApplication)
		if err != nil {
			return Application{}, err
		}
	}

	if isAddressChanged {
		updatedApplication = MapAddressToApplication(updatedApplication, updatedAddress)
	} else {
		updatedApplication = MapAddressToApplication(updatedApplication, address)
	}

	return updatedApplication, nil
}

func (appS *applicationStore) FetchApplicationByID(ctx context.Context, applicationId int) (Application, error) {

	db := sqlx.NewDb(appS.DB, "postgres")

	var application Application

	err := db.Get(&application, fethcApplicationByIdQuery, applicationId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return Application{}, apperrors.ErrNoApplicationExists
		}
		return Application{}, err
	}

	return application, nil
}

func (appS *applicationStore) DeleteApplicationByID(ctx context.Context, applicationId int) (int, error) {
	db := sqlx.NewDb(appS.DB, "postgres")
	var addressId int

	err := db.Get(&applicationId, deleteApplicationByIdQuery, addressId)

	if err != nil {
		return -1, err
	}

	err = DeleteAddress(ctx, db, addressId)
	if err != nil {
		return -1, err
	}

	return -1, nil
}

func (appS *applicationStore) FindApplicationById(ctx context.Context, applicationId int) bool {
	var ID int
	err := appS.BaseRepository.DB.QueryRow(findApplicationByIdQuery, applicationId).Scan(&ID)
	return err == nil
}
