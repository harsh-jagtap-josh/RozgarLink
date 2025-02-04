package dto

import "time"

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
