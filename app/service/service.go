package service

import (
	"github.com/kittanutp/hospital-app/database"
	"github.com/kittanutp/hospital-app/repository"
	"github.com/kittanutp/hospital-app/schema"
)

type PatientServiceInterface interface {
	ProcessGetPatients(filter schema.PatientFilterSchema, staff database.Staff) repository.PatientsResponse
	ProcessGetPatient(id string, staff database.Staff) repository.PatientResponse
}

type StaffServiceInterface interface {
	ProcessNewStaff(data schema.CreateStaffSchema) repository.StaffResponse
	ProcessLogIn(data schema.LogInSchema) (schema.TokenResponseSchema, error)
}

type StaffAuthServiceInterface interface {
	ProcessStaffToken(token string) repository.StaffResponse
}
