package repository

import "github.com/kittanutp/hospital-app/app/database"

type staffPostgresRepository struct {
	db database.Database
}

func NewStaffPostgresRepository(db database.Database) StaffRepository {
	return &staffPostgresRepository{db: db}
}

func (r *staffPostgresRepository) GetStaff(username string) StaffResponse {
	var staff database.Staff
	res := r.db.GetSession().First(&staff, "username = ?", username)
	var err error
	if res.Error != nil {
		err = res.Error
	}

	return StaffResponse{
		Staff: &staff,
		Err:   err,
	}
}

func (r *staffPostgresRepository) CreateStaff(staff *database.Staff) StaffResponse {
	res := r.db.GetSession().Create(staff)
	var err error
	if res.Error != nil {
		err = res.Error
	}

	return StaffResponse{
		Staff: staff,
		Err:   err,
	}
}
