package repository

import "github.com/kittanutp/hospital-app/app/database"

type staffAuthPostgresRepository struct {
	db database.Database
}

func NewStaffAuthPostgresRepository(db database.Database) StaffAuthRepository {
	return &staffAuthPostgresRepository{db: db}
}

func (r *staffAuthPostgresRepository) GetStaffById(id uint) StaffResponse {
	var staff database.Staff
	res := r.db.GetSession().First(&staff, id)
	var err error
	if res.Error != nil {
		err = res.Error
	}

	return StaffResponse{
		Staff: &staff,
		Err:   err,
	}
}
