package worker

import (
	"time"
)

type Gender string

const (
	Male    Gender = "male"
	Female  Gender = "female"
	Unknown Gender = "unknown"
)

type Address struct {
	ID      int    `json:"id,omitempty"`
	Details string `json:"details"`
	Street  string `json:"street"`
	City    string `json:"city"`
	State   string `json:"state"`
	Pincode int    `json:"pincode"`
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
	Location        int       `json:"location,omitempty"`
	IsAvailable     bool      `json:"is_available,omitempty"`
	Rating          float64   `json:"rating,omitempty"`
	TotalJobsWorked int       `json:"total_jobs_worked,omitempty"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
	Language        string    `json:"language"`
}

type WorkerRequest struct {
	ID              int       `json:"id"`
	Name            string    `json:"name"`
	ContactNumber   string    `json:"contact_number"`
	Email           string    `json:"email,omitempty"`
	Gender          Gender    `json:"gender"`
	Password        string    `json:"password,omitempty"`
	Sectors         string    `json:"sectors"`
	Skills          string    `json:"skills"`
	Location        Address   `json:"location,omitempty"`
	IsAvailable     bool      `json:"is_available,omitempty"`
	Rating          float64   `json:"rating,omitempty"`
	TotalJobsWorked int       `json:"total_jobs_worked,omitempty"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
	Language        string    `json:"language"`
}
