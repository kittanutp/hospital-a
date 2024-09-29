package schema

import (
	"time"

	"github.com/kittanutp/hospital-app/database"
)

type PatientFilterSchema struct {
	NationalID  *string    `json:"national_id,omitempty" binding:"omitempty"`
	PassportID  *string    `json:"passport_id,omitempty" binding:"omitempty"`
	FirstName   *string    `json:"first_name,omitempty" binding:"omitempty"`
	MiddleName  *string    `json:"middle_name,omitempty" binding:"omitempty"`
	LastName    *string    `json:"last_name,omitempty" binding:"omitempty"`
	DateOfBirth *time.Time `json:"date_of_birth,omitempty" binding:"omitempty"`
	PhoneNumber *string    `json:"phone_number,omitempty" binding:"omitempty,len=10"`
	Email       *string    `json:"email,omitempty" binding:"omitempty,email"`
}

type PatientJsonResponse struct {
	ID           uint      `json:"id"`
	FirstNameTH  string    `json:"first_name_th"`
	MiddleNameTH *string   `json:"middle_name_th"`
	LastNameTH   string    `json:"last_name_th"`
	FirstNameEN  string    `json:"first_name_en"`
	MiddleNameEN *string   `json:"middle_name_en"`
	LastNameEN   string    `json:"last_name_en"`
	DateOfBirth  time.Time `json:"date_of_birth"`
	PatientHN    string    `json:"patient_hn"`
	NationalID   *string   `json:"national_id"`
	PassportID   *string   `json:"passport_id"`
	PhoneNumber  *string   `json:"phone_number"`
	Email        *string   `json:"email"`
	Gender       string    `json:"gender"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

func ConvertJSONResponse(p database.Patient) PatientJsonResponse {
	return PatientJsonResponse{
		ID:           p.ID,
		FirstNameTH:  p.FirstNameTH,
		MiddleNameTH: p.MiddleNameTH,
		LastNameTH:   p.LastNameTH,
		FirstNameEN:  p.FirstNameEN,
		MiddleNameEN: p.MiddleNameEN,
		LastNameEN:   p.LastNameEN,
		DateOfBirth:  p.DateOfBirth,
		PatientHN:    p.PatientHN,
		NationalID:   p.NationalID,
		PassportID:   p.PassportID,
		PhoneNumber:  p.PhoneNumber,
		Email:        p.Email,
		Gender:       p.Gender,
	}
}
