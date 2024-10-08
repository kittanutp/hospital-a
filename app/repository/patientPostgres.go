package repository

import (
	"github.com/kittanutp/hospital-app/database"
	"github.com/kittanutp/hospital-app/schema"
)

type patientPostgresRepository struct {
	db database.Database
}

func NewPatientPostgresRepository(db database.Database) PatientRepository {
	return &patientPostgresRepository{db: db}
}

func (r *patientPostgresRepository) GetPatients(filter schema.PatientFilterSchema, hospitalName string) PatientsResponse {
	stm := r.db.GetSession().Where("patient_hospital_name = ?", hospitalName)

	filters := map[string]*string{
		"national_id":  filter.NationalID,
		"passport_id":  filter.PassportID,
		"phone_number": filter.PhoneNumber,
		"email":        filter.Email,
	}

	for column, value := range filters {
		if value != nil {
			stm = stm.Where(column+" = ?", *value)
		}
	}

	if filter.DateOfBirth != nil {
		stm = stm.Where("date_of_birth = ?", filter.DateOfBirth)
	}

	if filter.FirstName != nil {
		stm = stm.Where("(first_name_th = ? OR first_name_en = ?)", *filter.FirstName, *filter.FirstName)
	}

	if filter.MiddleName != nil {
		stm = stm.Where("(middle_name_th = ? OR middle_name_en = ?)", *filter.MiddleName, *filter.MiddleName)
	}

	if filter.LastName != nil {
		stm = stm.Where("(last_name_th = ? OR last_name_en = ?)", *filter.LastName, *filter.LastName)
	}

	var patients []database.Patient
	res := stm.Find(&patients)

	var err error
	if res.Error != nil {
		err = res.Error
	}

	return PatientsResponse{
		Patients: patients,
		Err:      err,
	}
}

func (r *patientPostgresRepository) GetPatient(id string) PatientResponse {
	var patient database.Patient
	res := r.db.GetSession().Where("national_id = ?", id).Or("passport_id = ?", id).First(&patient)
	var err error
	if res.Error != nil {
		err = res.Error
	}

	return PatientResponse{
		Patient: patient,
		Err:     err,
	}
}
