package employer

import "time"

type EmployerType string

const (
	Organization EmployerType = "organization"
	EmployerP    EmployerType = "employer"
)

type Address struct {
	ID      int    `json:"id,omitempty"`
	Details string `json:"details"`
	Street  string `json:"street"`
	City    string `json:"city"`
	State   string `json:"state"`
	Pincode int    `json:"pincode"`
}

type Employer struct {
	ID           int          `json:"id"`
	Name         string       `json:"name"`
	ContactNo    string       `json:"contact_number"`
	Email        string       `json:"email"`
	Type         EmployerType `json:"type"`
	Password     string       `json:"password,omitempty"`
	Sectors      string       `json:"sectors"`
	Location     Address      `json:"location"`
	IsVerified   bool         `json:"is_verified"`
	Rating       float64      `json:"rating"`
	WorkersHired int          `json:"workers_hired"`
	CreatedAt    time.Time    `json:"created_at"`
	UpdatedAt    time.Time    `json:"updated_at"`
	Language     string       `json:"language"`
}
