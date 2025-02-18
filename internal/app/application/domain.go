package application

import (
	"time"

	"github.com/harsh-jagtap-josh/RozgarLink/internal/app/worker"
)

type Status string
type ModeOfArrival string

const (
	Pending     Status        = "pending"
	Shortlisted Status        = "shortlisted"
	Confirmed   Status        = "confirmed"
	Personal    ModeOfArrival = "personal"
	PickUp      ModeOfArrival = "pickup"
)

type Application struct {
	ID             int            `json:"id"`
	JobID          int            `json:"job_id"`
	WorkerID       int            `json:"worker_id"`
	Status         Status         `json:"status"`
	ExpectedWage   int            `json:"expected_wage"`
	ModeOfArrival  ModeOfArrival  `json:"mode_of_arrival"`
	PickUpLocation worker.Address `json:"pick_up_location"`
	WorkerComment  string         `json:"worker_comments"`
	AppliedAt      time.Time      `json:"applied_at"`
	UpdatedAt      time.Time      `json:"updated_at"`
}
