package dto

import "time"

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
