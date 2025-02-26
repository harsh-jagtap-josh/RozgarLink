package worker

import "time"

type Gender string

const (
	Male    Gender = "male"
	Female  Gender = "female"
	Unknown Gender = "unknown"
)

type Worker struct {
	ID              int       `json:"id"`
	Name            string    `json:"name"`
	ContactNumber   string    `json:"contact_number"`
	Email           string    `json:"email,omitempty"`
	Gender          Gender    `json:"gender"`
	Password        string    `json:"-"`
	Sectors         string    `json:"sectors"`
	Skills          string    `json:"skills"`
	Location        int       `json:"location,omitempty"`
	IsAvailable     bool      `json:"is_available,omitempty"`
	Rating          float64   `json:"rating,omitempty"`
	TotalJobsWorked int       `json:"total_jobs_worked,omitempty"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
	Language        string    `json:"language"`
}
