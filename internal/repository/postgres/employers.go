package postgres

import (
	"database/sql"
	"errors"
	"log"

	models "github.com/harsh-jagtap-josh/RozgarLink/internal/repository/domain"
)

// // Create an employer
// func CreateEmployer(db *sql.DB, employer models.Employer) (int, error) {
// 	query := `
// 		INSERT INTO Employers (name, contact_no, email, type, password, sectors, location,
// 			is_verified, rating, workers_hired, created_at, updated_at, language)
// 		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, NOW(), NOW(), $11) RETURNING id;
// 	`
// 	var id int
// 	err := db.QueryRow(query, employer.Name, employer.ContactNo, employer.Email, employer.Type, employer.Password,
// 		employer.Sectors, employer.Location, employer.IsVerified, employer.Rating,
// 		employer.WorkersHired, employer.Language).Scan(&id)
// 	return id, err
// }

// Get employer by ID
func GetEmployerByID(db *sql.DB, id int) (*models.Employer, error) {
	query := `SELECT * FROM Employers WHERE id = $1;`
	row := db.QueryRow(query, id)

	var employer models.Employer
	err := row.Scan(&employer.ID, &employer.Name, &employer.ContactNo, &employer.Email, &employer.Type,
		&employer.Password, &employer.Sectors, &employer.Location, &employer.IsVerified, &employer.Rating,
		&employer.WorkersHired, &employer.CreatedAt, &employer.UpdatedAt, &employer.Language)

	if err != nil {
		return nil, err
	}
	return &employer, nil
}

// Update employer
func UpdateEmployer(db *sql.DB, employer models.Employer) (models.Employer, error) {
	query := `
		UPDATE Employers SET name=$1, contact_no=$2, email=$3, type=$4, sectors=$5, location=$6, 
		is_verified=$7, rating=$8, workers_hired=$9, updated_at=NOW(), language=$10 WHERE id=$11;
	`
	_, err := db.Exec(query, employer.Name, employer.ContactNo, employer.Email, employer.Type,
		employer.Sectors, employer.Location, employer.IsVerified, employer.Rating,
		employer.WorkersHired, employer.Language, employer.ID)

	if errors.Is(err, sql.ErrNoRows) {
		return employer, sql.ErrNoRows
	}

	return employer, err
}

// Delete employer
func DeleteEmployer(db *sql.DB, id int) (int64, error) {
	query := `DELETE FROM Employers WHERE id=$1;`
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

// Get Jobs by Employer ID
func GetJobsByEmployerId(db *sql.DB, id int) (*[]models.Job, error) {
	var jobs []models.Job

	query := `SELECT * FROM Jobs WHERE employer_id = $1;`

	// Execute query
	rows, err := db.Query(query, id)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, sql.ErrNoRows
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Parse results
	for rows.Next() {
		var job models.Job
		if err := rows.Scan(&job.ID, &job.EmployerID, &job.Title, &job.RequiredGender, &job.Description, &job.DurationHours, &job.SkillsRequired, &job.Sectors, &job.Wage, &job.Vacancy, &job.Location, &job.CreatedAt, &job.UpdatedAt, &job.Deadline); err != nil {
			log.Printf("Error scanning application: %v", err)
			return nil, err
		}
		jobs = append(jobs, job)
	}

	return &jobs, nil
}
