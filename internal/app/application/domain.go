package application

import (
	"time"
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

type Address struct {
	ID      int    `json:"id,omitempty"`
	Details string `json:"details"`
	Street  string `json:"street"`
	City    string `json:"city"`
	State   string `json:"state"`
	Pincode int    `json:"pincode"`
}

type Application struct {
	ID             int           `json:"id"`
	JobID          int           `json:"job_id"`
	WorkerID       int           `json:"worker_id"`
	Status         Status        `json:"status"`
	ExpectedWage   int           `json:"expected_wage"`
	ModeOfArrival  ModeOfArrival `json:"mode_of_arrival"`
	PickUpLocation Address       `json:"pick_up_location"`
	WorkerComment  string        `json:"worker_comments"`
	AppliedAt      time.Time     `json:"applied_at"`
	UpdatedAt      time.Time     `json:"updated_at"`
}

type ApplicationComplete struct {
	ID             int           `json:"id"`
	JobID          int           `json:"job_id"`
	WorkerID       int           `json:"worker_id"`
	Status         Status        `json:"status"`
	ExpectedWage   int           `json:"expected_wage"`
	ModeOfArrival  ModeOfArrival `json:"mode_of_arrival"`
	PickUpLocation int           `json:"pick_up_location"`
	WorkerComment  string        `json:"worker_comments"`
	AppliedAt      time.Time     `json:"applied_at"`
	UpdatedAt      time.Time     `json:"updated_at"`
	Details        string        `json:"details"`
	Street         string        `json:"street"`
	City           string        `json:"city"`
	State          string        `json:"state"`
	Pincode        int           `json:"pincode"`
	JobTitle       string        `json:"title"`
	Description    string        `json:"description"`
	SkillsRequired string        `json:"skills_required"`
	JobSectors     string        `json:"sectors"`
	JobWage        int           `json:"wage"`
	Vacancy        int           `json:"vacancy"`
	JobDate        string        `json:"date"`
	EmployerName   string        `json:"name"`
	ContactNumber  string        `json:"contact_number"`
	EmployerEmail  string        `json:"email"`
	EmployerType   string        `json:"type"`
}

type ApplicationCompleteEmp struct {
	ID             int           `json:"id"`
	JobID          int           `json:"job_id"`
	WorkerID       int           `json:"worker_id"`
	Status         Status        `json:"status"`
	ExpectedWage   int           `json:"expected_wage"`
	ModeOfArrival  ModeOfArrival `json:"mode_of_arrival"`
	PickUpLocation int           `json:"pick_up_location"`
	WorkerComment  string        `json:"worker_comments"`
	AppliedAt      time.Time     `json:"applied_at"`
	UpdatedAt      time.Time     `json:"updated_at"`
	Details        string        `json:"details"`
	Street         string        `json:"street"`
	City           string        `json:"city"`
	State          string        `json:"state"`
	Pincode        int           `json:"pincode"`
	JobTitle       string        `json:"title"`
	Description    string        `json:"description"`
	SkillsRequired string        `json:"skills_required"`
	JobSectors     string        `json:"sectors"`
	JobWage        int           `json:"wage"`
	Vacancy        int           `json:"vacancy"`
	JobDate        string        `json:"date"`
	WorkerName     string        `json:"name"`
	ContactNumber  string        `json:"contact_number"`
	WorkerEmail    string        `json:"email"`
	WorkerGender   string        `json:"gender"`
}
