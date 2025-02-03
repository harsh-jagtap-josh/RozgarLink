package auth

import (
	"database/sql"

	"github.com/harsh-jagtap-josh/RozgarLink/internal/repository/domain"
)

type User struct {
	ID       int
	Email    string
	Password string
}

func CreateWorker(db *sql.DB, worker domain.Worker) (int, error) {
	var workerID int
	err := db.QueryRow(
		"INSERT INTO Workers (name, contact_number, email, gender, password, sectors, skills, location, is_available, rating, total_jobs_worked, created_at, updated_at, language) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, NOW(), NOW(), $12) RETURNING id",
		worker.Name, worker.ContactNumber, worker.Email, worker.Gender, worker.Password, worker.Sectors, worker.Skills, worker.Location, worker.IsAvailable, worker.Rating, worker.TotalJobsWorked, worker.Language,
	).Scan(&workerID)
	return workerID, err
}

func CreateEmployer(db *sql.DB, employer domain.Employer) (int, error) {
	var employerID int
	err := db.QueryRow(
		"INSERT INTO Employers (name, contact_no, email, type, password, sectors, location, is_verified, rating, workers_hired, created_at, updated_at, language) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, NOW(), NOW(), $11) RETURNING id",
		employer.Name, employer.ContactNo, employer.Email, employer.Type, employer.Password, employer.Sectors, employer.Location, employer.IsVerified, employer.Rating, employer.WorkersHired, employer.Language,
	).Scan(&employerID)
	return employerID, err
}

func GetWorkerByEmail(db *sql.DB, email string) (User, error) {
	var user User
	err := db.QueryRow("SELECT id, email, password FROM Workers WHERE email = $1", email).Scan(&user.ID, &user.Email, &user.Password)
	return user, err
}

func GetEmployerByEmail(db *sql.DB, email string) (User, error) {
	var user User
	err := db.QueryRow("SELECT id, email, password FROM Employers WHERE email = $1", email).Scan(&user.ID, &user.Email, &user.Password)
	return user, err
}
