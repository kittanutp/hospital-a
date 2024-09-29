package service

import (
	"errors"

	"github.com/kittanutp/hospital-app/database"
	"github.com/kittanutp/hospital-app/repository"
	"github.com/kittanutp/hospital-app/schema"
)

type PatientService struct {
	patientRepository repository.PatientRepository
}

func NewPatientService(repo repository.PatientRepository) PatientServiceInterface {
	return &PatientService{
		patientRepository: repo,
	}
}

func (s *PatientService) ProcessGetPatients(filter schema.PatientFilterSchema, staff database.Staff) repository.PatientsResponse {
	return s.patientRepository.GetPatients(filter, staff.HospitalName)
}

func (s *PatientService) ProcessGetPatient(id string, staff database.Staff) repository.PatientResponse {
	res := s.patientRepository.GetPatient(id)
	if res.Patient.PatientHN != staff.HospitalName {
		res.Err = errors.New("invalid data")
	}
	return res
}
