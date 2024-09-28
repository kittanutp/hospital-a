package service

import (
	"github.com/kittanutp/hospital-app/app/database"
	"github.com/kittanutp/hospital-app/app/repository"
	"github.com/kittanutp/hospital-app/app/schema"
)

type PatientServiceInterface interface {
	ProcessGetPatients(filter schema.PatientFilterSchema, staff database.Staff) repository.PatientsResponse
	ProcessGetPatient(id string, staff database.Staff) repository.PatientResponse
}

type StaffServiceInterface interface {
	ProcessNewStaff(data schema.CreateStaffSchema) repository.StaffResponse
	ProcessLogIn(data schema.LogInSchema) (schema.TokenResponseSchema, error)
}
