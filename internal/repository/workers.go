package repository

import (
	"context"
	"time"
)

type WorkerStorer interface {
	GetWorkerByID(ctx context.Context, tx Transaction, productID int64) (Worker, error)
}

type Gender string

const (
	Male    Gender = "Male"
	Female  Gender = "Female"
	Unknown Gender = "Unknown"
)

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
