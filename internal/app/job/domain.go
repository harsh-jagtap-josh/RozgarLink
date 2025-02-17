package job

import (
	"time"

	"github.com/harsh-jagtap-josh/RozgarLink/internal/app/worker"
)

type Job struct {
	ID              int            `json:"id"`
	EmployerID      int            `json:"employer_id"`
	Title           string         `json:"title"`
	RequiredGender  string         `json:"required_gender,omitempty"`
	Description     string         `json:"description,omitempty"`
	DurationInHours int            `json:"duration_in_hours"`
	SkillsRequired  string         `json:"skills_required,omitempty"`
	Sectors         string         `json:"sectors,omitempty"`
	Wage            int            `json:"wage"`
	Vacancy         int            `json:"vacancy"`
	Location        worker.Address `json:"location,omitempty"`
	CreatedAt       time.Time      `json:"created_at"`
	UpdatedAt       time.Time      `json:"updated_at"`
}

type JobServiceStruct struct {
	ID              int       `json:"id"`
	EmployerID      int       `json:"employer_id"`
	Title           string    `json:"title"`
	RequiredGender  string    `json:"required_gender,omitempty"`
	Description     string    `json:"description,omitempty"`
	DurationInHours int       `json:"duration_in_hours"`
	SkillsRequired  string    `json:"skills_required,omitempty"`
	Sectors         string    `json:"sectors,omitempty"`
	Wage            int       `json:"wage"`
	Vacancy         int       `json:"vacancy"`
	Location        int       `json:"location,omitempty"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
	Details         string    `json:"details"`
	Street          string    `json:"street"`
	City            string    `json:"city"`
	State           string    `json:"state"`
	Pincode         int       `json:"pincode"`
}
