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
	Date            string         `json:"date"`
	StartHour       string         `json:"start_hour"`
	EndHour         string         `json:"end_hour"`
	CreatedAt       time.Time      `json:"created_at"`
	UpdatedAt       time.Time      `json:"updated_at"`
}

type JobFilters struct {
	Title     string
	Sector    string
	WageMin   int
	WageMax   int
	StartDate time.Time
	EndDate   time.Time
	City      string
	Gender    string
}
