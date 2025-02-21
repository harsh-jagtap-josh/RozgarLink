package repo

import (
	"time"
)

type Gender string

const (
	Male    Gender = "male"
	Female  Gender = "female"
	Unknown Gender = "unknown"
)

type Worker struct {
	ID              int       `db:"id"`
	Name            string    `db:"name"`
	ContactNumber   string    `db:"contact_number"`
	Email           string    `db:"email"`
	Gender          Gender    `db:"gender"`
	Password        string    `db:"password"`
	Sectors         string    `db:"sectors"`
	Skills          string    `db:"skills"`
	Location        int       `db:"location"`
	IsAvailable     bool      `db:"is_available"`
	Rating          float64   `db:"rating"`
	TotalJobsWorked int       `db:"total_jobs_worked"`
	CreatedAt       time.Time `db:"created_at"`
	UpdatedAt       time.Time `db:"updated_at"`
	Language        string    `db:"language"`
	Details         string    `db:"details"`
	Street          string    `db:"street"`
	City            string    `db:"city"`
	State           string    `db:"state"`
	Pincode         int       `db:"pincode"`
}

type EmployerType string

const (
	Organization EmployerType = "organization"
	EmployerP    EmployerType = "employer"
)

type Employer struct {
	ID           int          `db:"id"`
	Name         string       `db:"name"`
	ContactNo    string       `db:"contact_number"`
	Email        string       `db:"email"`
	Type         EmployerType `db:"type"`
	Password     string       `db:"password"`
	Sectors      string       `db:"sectors"`
	Location     int          `db:"location"`
	IsVerified   bool         `db:"is_verified"`
	Rating       float64      `db:"rating"`
	WorkersHired int          `db:"workers_hired"`
	CreatedAt    time.Time    `db:"created_at"`
	UpdatedAt    time.Time    `db:"updated_at"`
	Language     string       `db:"language"`
	Details      string       `db:"details"`
	Street       string       `db:"street"`
	City         string       `db:"city"`
	State        string       `db:"state"`
	Pincode      int          `db:"pincode"`
}

type Address struct {
	ID      int    `db:"id"`
	Details string `db:"details"`
	Street  string `db:"street"`
	City    string `db:"city"`
	State   string `db:"state"`
	Pincode int    `db:"pincode"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginUserData struct {
	ID       int    `db:"id"`
	Name     string `db:"name"`
	Email    string `db:"email"`
	Password string `db:"password"`
	Role     string `db:"role"`
}

type LoginResponse struct {
	Token string
	User  LoginUserData
}

// Job Structs

type Job struct {
	ID              int       `db:"id"`
	EmployerID      int       `db:"employer_id"`
	Title           string    `db:"title" `
	RequiredGender  string    `db:"required_gender"`
	Description     string    `db:"description"`
	DurationInHours int       `db:"duration_in_hours"`
	SkillsRequired  string    `db:"skills_required"`
	Sectors         string    `db:"sectors"`
	Wage            int       `db:"wage"`
	Vacancy         int       `db:"vacancy"`
	Location        int       `db:"location"`
	Date            string    `db:"date"`
	StartHour       string    `db:"start_hour"`
	EndHour         string    `db:"end_hour"`
	CreatedAt       time.Time `db:"created_at"`
	UpdatedAt       time.Time `db:"updated_at"`
	Details         string    `db:"details"`
	Street          string    `db:"street"`
	City            string    `db:"city"`
	State           string    `db:"state"`
	Pincode         int       `db:"pincode"`
}

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
	ID             int           `db:"id"`
	JobID          int           `db:"job_id"`
	WorkerID       int           `db:"worker_id"`
	Status         Status        `db:"status"`
	ExpectedWage   int           `db:"expected_wage"`
	ModeOfArrival  ModeOfArrival `db:"mode_of_arrival"`
	PickUpLocation int           `db:"pick_up_location"`
	WorkerComment  string        `db:"worker_comments"`
	AppliedAt      time.Time     `db:"applied_at"`
	UpdatedAt      time.Time     `db:"updated_at"`
	Details        string        `db:"details"`
	Street         string        `db:"street"`
	City           string        `db:"city"`
	State          string        `db:"state"`
	Pincode        int           `db:"pincode"`
}

type Sector struct {
	ID          int    `db:"id"`
	Name        string `db:"name"`
	Description string `db:"description"`
}
