package worker

import "time"

type Gender string

const (
	Male    Gender = "Male"
	Female  Gender = "Female"
	Unknown Gender = "Unknown"
)

type Worker struct {
	ID              int       `db:"id"`
	Name            string    `db:"name"`
	ContactNumber   string    `db:"contact_number"`
	Email           string    `db:"email,omitempty"`
	Gender          Gender    `db:"gender"`
	Password        string    `db:"password"`
	Sectors         string    `db:"sectors"`
	Skills          string    `db:"skills"`
	Location        *int      `db:"location,omitempty"`
	IsAvailable     *bool     `db:"is_available,omitempty"`
	Rating          *float64  `db:"rating,omitempty"`
	TotalJobsWorked *int      `db:"total_jobs_worked,omitempty"`
	CreatedAt       time.Time `db:"created_at"`
	UpdatedAt       time.Time `db:"updated_at"`
	Language        string    `db:"language"`
}
