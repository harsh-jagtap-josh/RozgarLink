package admin

import "time"

type Admin struct {
	Name      string    `json:"name"`
	ContactNo string    `json:"contact_no"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
