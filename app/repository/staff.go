package repository

import "github.com/kittanutp/hospital-app/database"

type StaffRepository interface {
	GetStaffByUsername(username string) StaffResponse
	CreateStaff(staff *database.Staff) StaffResponse
}

type StaffAuthRepository interface {
	GetStaffById(id uint) StaffResponse
}

type StaffResponse struct {
	Staff *database.Staff
	Err   error
}
