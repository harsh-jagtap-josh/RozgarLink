package repo

import "time"

type Gender string

const (
	Male    Gender = "male"
	Female  Gender = "female"
	Unknown Gender = "unknown"
)

type Worker struct {
	ID              int
	Name            string
	ContactNumber   string
	Email           string
	Gender          Gender
	Password        string
	Sectors         string
	Skills          string
	Location        *int
	IsAvailable     *bool
	Rating          *float64
	TotalJobsWorked *int
	CreatedAt       time.Time
	UpdatedAt       time.Time
	Language        string
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
	Location     *int
	IsVerified   bool
	Rating       *float64
	WorkersHired *int
	CreatedAt    time.Time
	UpdatedAt    time.Time
	Language     string
}
