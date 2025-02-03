package postgres

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"strings"

	models "github.com/harsh-jagtap-josh/RozgarLink/internal/repository/domain"
)

// Create a new job
func CreateJob(db *sql.DB, job models.Job) (int, error) {
	// query to insert into table, that also return id of newly created if there's no error
	query := `
		INSERT INTO Jobs (employer_id, title, required_gender, description, duration_in_hours, skills_required,
			sectors, wage, vacancy, location, created_at, updated_at, deadline) 
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, NOW(), NOW(), $11) RETURNING id;
	`
	var id int
	err := db.QueryRow(query, job.EmployerID, job.Title, job.RequiredGender, job.Description,
		job.DurationHours, job.SkillsRequired, job.Sectors, job.Wage, job.Vacancy,
		job.Location, job.Deadline).Scan(&id)
	return id, err
}

// Get job by ID
func GetJobByID(db *sql.DB, id int) (*models.Job, error) {
	query := `SELECT * FROM Jobs WHERE id = $1;`
	row := db.QueryRow(query, id)

	var job models.Job
	err := row.Scan(&job.ID, &job.EmployerID, &job.Title, &job.RequiredGender, &job.Description,
		&job.DurationHours, &job.SkillsRequired, &job.Sectors, &job.Wage, &job.Vacancy,
		&job.Location, &job.CreatedAt, &job.UpdatedAt, &job.Deadline)

	if err != nil {
		return nil, err
	}
	return &job, nil
}

// Get all jobs
func GetAllJobs(db *sql.DB) ([]models.Job, error) {
	query := `SELECT * FROM Jobs;`
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var jobs []models.Job
	for rows.Next() {
		var job models.Job
		err := rows.Scan(&job.ID, &job.EmployerID, &job.Title, &job.RequiredGender, &job.Description,
			&job.DurationHours, &job.SkillsRequired, &job.Sectors, &job.Wage, &job.Vacancy,
			&job.Location, &job.CreatedAt, &job.UpdatedAt, &job.Deadline)

		if err != nil {
			return nil, err
		}
		jobs = append(jobs, job)
	}
	return jobs, nil
}

// Update a job
func UpdateJob(db *sql.DB, job models.Job) (models.Job, error) {
	query := `
		UPDATE Jobs SET title=$1, required_gender=$2, description=$3, duration_in_hours=$4, 
		skills_required=$5, sectors=$6, wage=$7, vacancy=$8, location=$9, updated_at=NOW(), deadline=$10 
		WHERE id=$11;
	`
	_, err := db.Exec(query, job.Title, job.RequiredGender, job.Description, job.DurationHours,
		job.SkillsRequired, job.Sectors, job.Wage, job.Vacancy, job.Location, job.Deadline, job.ID)
	if errors.Is(err, sql.ErrNoRows) {
		return job, sql.ErrNoRows
	}
	return job, err
}

// Delete a job
func DeleteJob(db *sql.DB, id int) (int64, error) {
	query := `DELETE FROM Jobs WHERE id=$1;`
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

func GetFilteredJobs(db *sql.DB, title *string, sectors *string, wage *string, rating *string) ([]models.Job, error) {
	var jobs []models.Job
	var conditions []string
	var args []interface{}
	argIndex := 1

	if len(*title) > 0 {
		conditions = append(conditions, fmt.Sprintf("title ILIKE $%d", argIndex))
		args = append(args, "%"+*title+"%")
		argIndex++
	}
	if len(*sectors) > 0 {
		conditions = append(conditions, fmt.Sprintf("sectors ILIKE $%d", argIndex))
		args = append(args, "%"+*sectors+"%")
		argIndex++
	}
	if len(*rating) > 0 {
		conditions = append(conditions, fmt.Sprintf("rating >= $%d", argIndex))
		args = append(args, *rating)
		argIndex++
	}
	if len(*wage) > 0 {
		conditions = append(conditions, fmt.Sprintf("wage >= $%d", argIndex))
		args = append(args, *wage)
		argIndex++
	}
	// also can add required gender and duration in args
	query := "SELECT * FROM Jobs"
	if len(conditions) > 0 {
		query += " WHERE " + strings.Join(conditions, " AND ")
	}

	rows, err := db.Query(query, args...)
	if err != nil {
		log.Printf("Error fetching filtered jobs: %v", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var job models.Job
		if err := rows.Scan(&job.ID, &job.EmployerID, &job.Title, &job.RequiredGender, &job.Description, &job.DurationHours, &job.SkillsRequired, &job.Sectors, &job.Wage, &job.Vacancy, &job.Location, &job.CreatedAt, &job.UpdatedAt, &job.Deadline); err != nil {
			log.Printf("Error scanning job: %v", err)
			return nil, err
		}
		jobs = append(jobs, job)
	}

	return jobs, nil
}

// Get Applications by Worker ID
func GetApplicationsByJobID(db *sql.DB, id int) (*[]models.Application, error) {
	var applications []models.Application

	query := `SELECT * FROM Applications WHERE job_id = $1;`

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
