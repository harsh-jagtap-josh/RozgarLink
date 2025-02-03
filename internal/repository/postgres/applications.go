package postgres

import (
	"database/sql"
	"errors"

	models "github.com/harsh-jagtap-josh/RozgarLink/internal/repository/domain"
)

// Create a new application
func CreateApplication(db *sql.DB, app models.Application) (int, error) {
	query := `
		INSERT INTO Applications (job_id, worker_id, status, expected_wage, mode_of_arrival, 
			pick_up_location, worker_comments, applied_at, updated_at) 
		VALUES ($1, $2, $3, $4, $5, $6, $7, NOW(), NOW()) RETURNING id;
	`
	var id int
	err := db.QueryRow(query, app.JobID, app.WorkerID, app.Status, app.ExpectedWage,
		app.ModeOfArrival, app.PickUpLocation, app.WorkerComments).Scan(&id)
	return id, err
}

// Get application by ID
func GetApplicationByID(db *sql.DB, id int) (*models.Application, error) {
	query := `SELECT * FROM Applications WHERE id = $1;`
	row := db.QueryRow(query, id)

	var app models.Application
	err := row.Scan(&app.ID, &app.JobID, &app.WorkerID, &app.Status, &app.ExpectedWage,
		&app.ModeOfArrival, &app.PickUpLocation, &app.WorkerComments, &app.AppliedAt, &app.UpdatedAt)

	if err != nil {
		return nil, err
	}
	return &app, nil
}

// Get all applications
func GetAllApplications(db *sql.DB) ([]models.Application, error) {
	query := `SELECT * FROM Applications;`
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var applications []models.Application
	for rows.Next() {
		var app models.Application
		err := rows.Scan(&app.ID, &app.JobID, &app.WorkerID, &app.Status, &app.ExpectedWage,
			&app.ModeOfArrival, &app.PickUpLocation, &app.WorkerComments, &app.AppliedAt, &app.UpdatedAt)

		if err != nil {
			return nil, err
		}
		applications = append(applications, app)
	}
	return applications, nil
}

// Update an application
func UpdateApplication(db *sql.DB, app models.Application) (models.Application, error) {
	query := `
		UPDATE Applications SET status=$1, expected_wage=$2, mode_of_arrival=$3, pick_up_location=$4, 
		worker_comments=$5, updated_at=NOW() WHERE id=$6;
	`
	_, err := db.Exec(query, app.Status, app.ExpectedWage, app.ModeOfArrival, app.PickUpLocation,
		app.WorkerComments, app.ID)

	if errors.Is(err, sql.ErrNoRows) {
		return app, sql.ErrNoRows
	}

	return app, err
}

// Delete an application
func DeleteApplication(db *sql.DB, id int) (int64, error) {
	query := `DELETE FROM Applications WHERE id=$1;`
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
