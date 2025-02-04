package dto

import "time"

type Gender string

const (
	Male    Gender = "Male"
	Female  Gender = "Female"
	Unknown Gender = "Unknown"
)

type EmployerType string

const (
	Organization EmployerType = "Organization"
	EmployerP    EmployerType = "Employer"
)

type RequiredGender string

const (
	ReqMale   RequiredGender = "Male"
	ReqFemale RequiredGender = "Female"
	ReqAny    RequiredGender = "ANY"
)

type ApplicationStatus string

const (
	Pending     ApplicationStatus = "Pending"
	Shortlisted ApplicationStatus = "Shortlisted"
	Confirmed   ApplicationStatus = "Confirmed"
)

type ModeOfArrival string

const (
	Personal ModeOfArrival = "Personal"
	PickUp   ModeOfArrival = "Pick-Up"
)

// Structs
type Address struct {
	ID      int    `json:"id"`
	Details string `json:"details"`
	Street  string `json:"street,omitempty"`
	City    string `json:"city"`
	State   string `json:"state"`
}

type Worker struct {
	ID              int       `json:"id"`
	Name            string    `json:"name"`
	ContactNumber   string    `json:"contact_number"`
	Email           string    `json:"email,omitempty"`
	Gender          Gender    `json:"gender"`
	Password        string    `json:"password"`
	Sectors         string    `json:"sectors"`
	Skills          string    `json:"skills"`
	Location        *int      `json:"location,omitempty"`
	IsAvailable     *bool     `json:"is_available,omitempty"`
	Rating          *float64  `json:"rating,omitempty"`
	TotalJobsWorked *int      `json:"total_jobs_worked,omitempty"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
	Language        string    `json:"language"`
}

type Employer struct {
	ID           int          `json:"id"`
	Name         string       `json:"name"`
	ContactNo    string       `json:"contact_no"`
	Email        string       `json:"email,omitempty"`
	Type         EmployerType `json:"type"`
	Password     string       `json:"password"`
	Sectors      string       `json:"sectors"`
	Location     *int         `json:"location,omitempty"`
	IsVerified   bool         `json:"is_verified"`
	Rating       *float64     `json:"rating,omitempty"`
	WorkersHired *int         `json:"workers_hired,omitempty"`
	CreatedAt    time.Time    `json:"created_at"`
	UpdatedAt    time.Time    `json:"updated_at"`
	Language     string       `json:"language"`
}

type Job struct {
	ID             int            `json:"id"`
	EmployerID     int            `json:"employer_id"`
	Title          string         `json:"title"`
	RequiredGender RequiredGender `json:"required_gender"`
	Description    string         `json:"description"`
	DurationHours  int            `json:"duration_in_hours"`
	SkillsRequired string         `json:"skills_required"`
	Sectors        string         `json:"sectors"`
	Wage           int            `json:"wage"`
	Vacancy        int            `json:"vacancy"`
	Location       *int           `json:"location,omitempty"`
	CreatedAt      time.Time      `json:"created_at"`
	UpdatedAt      time.Time      `json:"updated_at"`
	Deadline       *time.Time     `json:"deadline,omitempty"`
}

type Application struct {
	ID             int               `json:"id"`
	JobID          int               `json:"job_id"`
	WorkerID       int               `json:"worker_id"`
	Status         ApplicationStatus `json:"status"`
	ExpectedWage   *float64          `json:"expected_wage,omitempty"`
	ModeOfArrival  *ModeOfArrival    `json:"mode_of_arrival,omitempty"`
	PickUpLocation *int              `json:"pick_up_location,omitempty"`
	WorkerComments *string           `json:"worker_comments,omitempty"`
	AppliedAt      time.Time         `json:"applied_at"`
	UpdatedAt      time.Time         `json:"updated_at"`
}

type Sector struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
