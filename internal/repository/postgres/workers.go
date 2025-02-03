package postgres

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"strings"

	models "github.com/harsh-jagtap-josh/RozgarLink/internal/repository/domain"
)

func GetFilteredWorkers(db *sql.DB, isAvailable *string, name *string, sector *string, rating *string, jobsWorked *string) ([]models.Worker, error) {
	var workers []models.Worker
	var conditions []string
	var args []interface{}
	argIndex := 1

	if len(*isAvailable) > 0 {
		conditions = append(conditions, fmt.Sprintf("is_available = $%d", argIndex))
		args = append(args, *isAvailable)
		argIndex++
	}
	if len(*name) > 0 {
		conditions = append(conditions, fmt.Sprintf("name = $%d", argIndex))
		args = append(args, *name)
		argIndex++
	}
	if len(*sector) > 0 {
		conditions = append(conditions, fmt.Sprintf("sectors ILIKE $%d", argIndex))
		args = append(args, "%"+*sector+"%")
		argIndex++
	}
	if len(*rating) > 0 {
		conditions = append(conditions, fmt.Sprintf("rating >= $%d", argIndex))
		args = append(args, *rating)
		argIndex++
	}
	if len(*jobsWorked) > 0 {
		conditions = append(conditions, fmt.Sprintf("rating >= $%d", argIndex))
		args = append(args, *rating)
		argIndex++
	}

	query := "SELECT id, name, contact_number, email, gender, password, sectors, skills, location, is_available, rating, total_jobs_worked, language FROM Workers"
	if len(conditions) > 0 {
		query += " WHERE " + strings.Join(conditions, " AND ")
	}

	rows, err := db.Query(query, args...)
	if err != nil {
		log.Printf("Error fetching filtered workers: %v", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var worker models.Worker
		if err := rows.Scan(&worker.ID, &worker.Name, &worker.ContactNumber, &worker.Email, &worker.Gender, &worker.Password, &worker.Sectors, &worker.Skills, &worker.Location, &worker.IsAvailable, &worker.Rating, &worker.TotalJobsWorked, &worker.Language); err != nil {
			log.Printf("Error scanning worker: %v", err)
			return nil, err
		}
		workers = append(workers, worker)
	}

	return workers, nil
}

// // Create a new worker
// func CreateWorker(db *sql.DB, worker models.Worker) (int, error) {
// 	query := `
// 		INSERT INTO Workers (name, contact_number, email, gender, password, sectors, skills, location,
// 			is_available, rating, total_jobs_worked, created_at, updated_at, language)
// 		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, NOW(), NOW(), $12) RETURNING id;
// 	`
// 	var id int
// 	err := db.QueryRow(query, worker.Name, worker.ContactNumber, worker.Email, worker.Gender, worker.Password,
// 		worker.Sectors, worker.Skills, worker.Location, worker.IsAvailable, worker.Rating,
// 		worker.TotalJobsWorked, worker.Language).Scan(&id)
// 	return id, err
// }

// Get worker by ID
func GetWorkerByID(db *sql.DB, id int) (*models.Worker, error) {
	query := `SELECT * FROM Workers WHERE id = $1;`
	row := db.QueryRow(query, id)

	var worker models.Worker
	err := row.Scan(&worker.ID, &worker.Name, &worker.ContactNumber, &worker.Email, &worker.Gender,
		&worker.Password, &worker.Sectors, &worker.Skills, &worker.Location, &worker.IsAvailable,
		&worker.Rating, &worker.TotalJobsWorked, &worker.CreatedAt, &worker.UpdatedAt, &worker.Language)

	if errors.Is(err, sql.ErrNoRows) {
		return nil, sql.ErrNoRows
	}
	if err != nil {
		return nil, err
	}
	return &worker, nil
}

// Update worker
func UpdateWorker(db *sql.DB, worker models.Worker) (models.Worker, error) {
	query := `
		UPDATE Workers SET name=$1, contact_number=$2, email=$3, gender=$4, sectors=$5, skills=$6, location=$7,
		is_available=$8, rating=$9, total_jobs_worked=$10, updated_at=NOW(), language=$11 WHERE id=$12;
	`
	_, err := db.Exec(query, worker.Name, worker.ContactNumber, worker.Email, worker.Gender, worker.Sectors,
		worker.Skills, worker.Location, worker.IsAvailable, worker.Rating, worker.TotalJobsWorked, worker.Language, worker.ID)

	if errors.Is(err, sql.ErrNoRows) {
		return worker, sql.ErrNoRows
	}

	return worker, err
}

// Function to delete worker that returns the Number of rows affected and Error
func DeleteWorker(db *sql.DB, id int) (int64, error) {
	query := `DELETE FROM Workers WHERE id=$1;`
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

// Get Applications by Worker ID
func GetApplicationsByWorkerID(db *sql.DB, id int) (*[]models.Application, error) {
	var applications []models.Application

	query := `SELECT * FROM Applications WHERE worker_id = $1;`

	// Execute query
	rows, err := db.Query(query, id)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, sql.ErrNoRows
	}
	if err != nil {
		log.Printf("internal error: %v", err)
		return nil, err
	}
	defer rows.Close()

	// Parse results
	for rows.Next() {
		var application models.Application
		if err := rows.Scan(&application.ID, &application.JobID, &application.WorkerID, &application.Status, &application.ExpectedWage, &application.ModeOfArrival, &application.PickUpLocation, &application.WorkerComments, &application.AppliedAt, &application.UpdatedAt); err != nil {
			log.Printf("Error scanning application: %v", err)
			return nil, err
		}
		applications = append(applications, application)
	}

	return &applications, nil
}
