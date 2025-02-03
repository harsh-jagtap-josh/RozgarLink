package postgres

import (
	"database/sql"
	"errors"

	models "github.com/harsh-jagtap-josh/RozgarLink/internal/repository/domain"
)

// Create a new sector
func CreateSector(db *sql.DB, sector models.Sector) (int, error) {
	query := `INSERT INTO Sectors (name, description) VALUES ($1, $2) RETURNING id;`
	var id int
	err := db.QueryRow(query, sector.Name, sector.Description).Scan(&id)
	return id, err
}

// Get sector by ID
func GetSectorByID(db *sql.DB, id int) (*models.Sector, error) {
	query := `SELECT * FROM Sectors WHERE id = $1;`
	row := db.QueryRow(query, id)

	var sector models.Sector
	err := row.Scan(&sector.ID, &sector.Name, &sector.Description)

	if err != nil {
		return nil, err
	}
	return &sector, nil
}

// Get all sectors
func GetAllSectors(db *sql.DB) ([]models.Sector, error) {
	query := `SELECT * FROM Sectors;`
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var sectors []models.Sector
	for rows.Next() {
		var sector models.Sector
		err := rows.Scan(&sector.ID, &sector.Name, &sector.Description)

		if err != nil {
			return nil, err
		}
		sectors = append(sectors, sector)
	}
	return sectors, nil
}

// Update a sector
func UpdateSector(db *sql.DB, sector models.Sector) (models.Sector, error) {
	query := `UPDATE Sectors SET name=$1, description=$2 WHERE id=$3;`
	_, err := db.Exec(query, sector.Name, sector.Description, sector.ID)
	if errors.Is(err, sql.ErrNoRows) {
		return sector, sql.ErrNoRows
	}
	return sector, err
}

// Delete a sector
func DeleteSector(db *sql.DB, id int) (int64, error) {
	query := `DELETE FROM Sectors WHERE id=$1;`
	rows, err := db.Exec(query, id)
	if err != nil {
		return 0, err
	}

	affectedRows, err := rows.RowsAffected()
	if err != nil {
		return 0, err
	}
	return affectedRows, err
}
