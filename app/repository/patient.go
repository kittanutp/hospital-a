package repository

import (
	"github.com/kittanutp/hospital-app/database"
	"github.com/kittanutp/hospital-app/schema"
)

type PatientRepository interface {
	GetPatients(filter schema.PatientFilterSchema, hospitalName string) PatientsResponse
	GetPatient(id string) PatientResponse
}

type PatientsResponse struct {
	Patients []database.Patient
	Err      error
}

type PatientResponse struct {
	Patient database.Patient
	Err     error
}
