package repo

import (
	"time"
)

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
	Location        int       `db:"location"`
	IsAvailable     bool      `db:"is_available"`
	Rating          float64   `db:"rating"`
	TotalJobsWorked int       `db:"total_jobs_worked"`
	CreatedAt       time.Time `db:"created_at"`
	UpdatedAt       time.Time `db:"updated_at"`
	Language        string    `db:"language"`
}

type WorkerResponse struct {
	ID              int       `db:"id"`
	Name            string    `db:"name"`
	ContactNumber   string    `db:"contact_number"`
	Email           string    `db:"email"`
	Gender          Gender    `db:"gender"`
	Password        string    `db:"password"`
	Sectors         string    `db:"sectors"`
	Skills          string    `db:"skills"`
	Location        int       `db:"location"`
	IsAvailable     bool      `db:"is_available"`
	Rating          float64   `db:"rating"`
	TotalJobsWorked int       `db:"total_jobs_worked"`
	CreatedAt       time.Time `db:"created_at"`
	UpdatedAt       time.Time `db:"updated_at"`
	Language        string    `db:"language"`
	Details         string    `db:"details"`
	Street          string    `db:"street"`
	City            string    `db:"city"`
	State           string    `db:"state"`
	Pincode         int       `db:"pincode"`
}

type EmployerType string

const (
	Organization EmployerType = "organization"
	EmployerP    EmployerType = "employer"
)

type Employer struct {
	ID           int          `db:"id"`
	Name         string       `db:"name"`
	ContactNo    string       `db:"contact_number"`
	Email        string       `db:"email"`
	Type         EmployerType `db:"type"`
	Password     string       `db:"password"`
	Sectors      string       `db:"sectors"`
	Location     int          `db:"location"`
	IsVerified   bool         `db:"is_verified"`
	Rating       float64      `db:"rating"`
	WorkersHired int          `db:"workers_hired"`
	CreatedAt    time.Time    `db:"created_at"`
	UpdatedAt    time.Time    `db:"updated_at"`
	Language     string       `db:"language"`
}

type EmployerResponse struct {
	ID           int          `db:"id"`
	Name         string       `db:"name"`
	ContactNo    string       `db:"contact_no"`
	Email        string       `db:"email"`
	Type         EmployerType `db:"type"`
	Password     string       `db:"password"`
	Sectors      string       `db:"sectors"`
	Location     int          `db:"location"`
	IsVerified   bool         `db:"is_verified"`
	Rating       float64      `db:"rating"`
	WorkersHired int          `db:"workers_hired"`
	CreatedAt    time.Time    `db:"created_at"`
	UpdatedAt    time.Time    `db:"updated_at"`
	Language     string       `db:"language"`
	Details      string       `db:"details"`
	Street       string       `db:"street"`
	City         string       `db:"city"`
	State        string       `db:"state"`
	Pincode      int          `db:"pincode"`
}

type Address struct {
	ID      int    `db:"id"`
	Details string `db:"details"`
	Street  string `db:"street"`
	City    string `db:"city"`
	State   string `db:"state"`
	Pincode int    `db:"pincode"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginUserData struct {
	ID       int    `db:"id"`
	Name     string `db:"name"`
	Email    string `db:"email"`
	Password string `db:"password"`
	Role     string `db:"role"`
}

type LoginResponse struct {
	Token string
	User  LoginUserData
}
