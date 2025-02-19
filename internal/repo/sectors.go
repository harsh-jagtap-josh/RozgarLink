package repo

import (
	"context"
	"database/sql"
	"errors"

	"github.com/harsh-jagtap-josh/RozgarLink/internal/pkg/apperrors"
	"github.com/jmoiron/sqlx"
)

type sectorStore struct {
	BaseRepository
}

type SectoreStorer interface {
	CreateNewSector(ctx context.Context, sectorData Sector) (Sector, error)
	FetchSectorById(ctx context.Context, sectorId int) (Sector, error)
	UpdateSectorById(ctx context.Context, sectorData Sector) (Sector, error)
	DeleteSectorById(ctx context.Context, sectorId int) (int, error)
	FetchAllSectors(ctx context.Context) ([]Sector, error)
}

func NewSectorRepo(db *sql.DB) SectoreStorer {
	return &sectorStore{
		BaseRepository: BaseRepository{DB: db},
	}
}

// PostgreSQL Queries
const (
	createNewSectorQuery  = `INSERT INTO sectors (name, description) VALUES (:name, :description) RETURNING *;`
	fetchSectorByIdQuery  = `SELECT * FROM sectors where id=$1;`
	updateSectorByIdQuery = `UPDATE sectors SET name=:name, description=:description where id=:id RETURNING *;`
	deleteSectorByIdQuery = `DELETE FROM sectors WHERE id=$1 RETURNING id;`
	fetchAllSectorQuery   = `SELECT * FROM sectors;`
)

func (sectorS *sectorStore) CreateNewSector(ctx context.Context, sectorData Sector) (Sector, error) {
	db := sqlx.NewDb(sectorS.DB, "postgres")
	var createdSector Sector

	rows, err := db.NamedQuery(createNewSectorQuery, sectorData)
	if err != nil {
		return Sector{}, err
	}

	defer rows.Close()

	if rows.Next() {
		err = rows.StructScan(&createdSector)
		if err != nil {
			return Sector{}, err
		}
	}
	return createdSector, nil
}

func (sectorS *sectorStore) FetchSectorById(ctx context.Context, sectorId int) (Sector, error) {
	db := sqlx.NewDb(sectorS.DB, "postgres")
	var sector Sector

	err := db.Get(&sector, fetchSectorByIdQuery, sectorId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return Sector{}, apperrors.ErrNoSectorExists
		}
		return Sector{}, err
	}

	return sector, nil
}

func (sectorS *sectorStore) UpdateSectorById(ctx context.Context, sectorData Sector) (Sector, error) {

	db := sqlx.NewDb(sectorS.DB, "postgres")

	var sector Sector

	rows, err := db.NamedQuery(updateSectorByIdQuery, sectorData)
	if err != nil {
		return Sector{}, err
	}

	defer rows.Close()

	if rows.Next() {
		err = rows.StructScan(&sector)
		if err != nil {
			return Sector{}, err
		}
	}

	return sector, nil
}

func (sectorS *sectorStore) DeleteSectorById(ctx context.Context, sectorId int) (int, error) {

	db := sqlx.NewDb(sectorS.DB, "postgres")

	var id int

	err := db.Get(&id, deleteSectorByIdQuery, sectorId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return -1, apperrors.ErrNoSectorExists
		}
		return -1, err
	}

	return id, nil
}

func (sectorS *sectorStore) FetchAllSectors(ctx context.Context) ([]Sector, error) {
	var sectors []Sector
	db := sqlx.NewDb(sectorS.DB, "postgres")

	err := db.Select(&sectors, fetchAllSectorQuery)
	if err != nil {
		return []Sector{}, err
	}
	return sectors, nil
}
