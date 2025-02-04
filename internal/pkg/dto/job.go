package dto

import "time"

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
