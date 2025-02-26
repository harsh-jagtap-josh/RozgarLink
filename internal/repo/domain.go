package repo

import "time"

type Gender string

const (
	Male    Gender = "male"
	Female  Gender = "female"
	Unknown Gender = "unknown"
)

type Worker struct {
	ID              int       `db:"id"`
	Name            string    `db:"name"`
	ContactNumber   string    `db:"contact_number"`
	Email           string    `db:"email"`
	Gender          Gender    `db:"gender"`
	Password        string    `db:"password"`
	Sectors         string    `db:"sectors"`
	Skills          string    `db:"skills"`
	Location        *int      `db:"location"`
	IsAvailable     bool      `db:"is_available"`
	Rating          float64   `db:"rating"`
	TotalJobsWorked int       `db:"total_jobs_worked"`
	CreatedAt       time.Time `db:"created_at"`
	UpdatedAt       time.Time `db:"updated_at"`
	Language        string    `db:"language"`
}

type EmployerType string

const (
	Organization EmployerType = "organization"
	EmployerP    EmployerType = "employer"
)

type Employer struct {
	ID           int
	Name         string
	ContactNo    string
	Email        string
	Type         EmployerType
	Password     string
	Sectors      string
	Location     int
	IsVerified   bool
	Rating       float64
	WorkersHired int
	CreatedAt    time.Time
	UpdatedAt    time.Time
	Language     string
}
