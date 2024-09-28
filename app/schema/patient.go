package schema

import (
	"time"
)

type PatientFilterSchema struct {
	NationalID  *string    `json:"national_id,omitempty" binding:"omitempty,len=13"`
	PassportID  *string    `json:"passport_id,omitempty" binding:"omitempty"`
	FirstName   *string    `json:"first_name,omitempty" binding:"omitempty"`
	MiddleName  *string    `json:"middle_name,omitempty" binding:"omitempty"`
	LastName    *string    `json:"last_name,omitempty" binding:"omitempty"`
	DateOfBirth *time.Time `json:"date_of_birth,omitempty" binding:"omitempty"`
	PhoneNumber *string    `json:"phone_number,omitempty" binding:"omitempty,len=10"`
	Email       *string    `json:"email,omitempty" binding:"omitempty,email"`
}
