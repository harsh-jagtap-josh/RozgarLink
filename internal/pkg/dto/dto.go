package dto

// data transfer object - dto

type Gender string

const (
	Male    Gender = "Male"
	Female  Gender = "Female"
	Unknown Gender = "Unknown"
)

type EmployerType string

const (
	Organization EmployerType = "Organization"
	EmployerP    EmployerType = "Employer"
)

type RequiredGender string

const (
	ReqMale   RequiredGender = "Male"
	ReqFemale RequiredGender = "Female"
	ReqAny    RequiredGender = "ANY"
)

type ApplicationStatus string

const (
	Pending     ApplicationStatus = "Pending"
	Shortlisted ApplicationStatus = "Shortlisted"
	Confirmed   ApplicationStatus = "Confirmed"
)

type ModeOfArrival string

const (
	Personal ModeOfArrival = "Personal"
	PickUp   ModeOfArrival = "Pick-Up"
)

type Address struct {
	ID      int    `json:"id"`
	Details string `json:"details"`
	Street  string `json:"street,omitempty"`
	City    string `json:"city"`
	State   string `json:"state"`
}
