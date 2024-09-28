package repository

import "github.com/kittanutp/hospital-app/app/database"

type StaffRepository interface {
	GetStaff(username string) StaffResponse
	CreateStaff(staff *database.Staff) StaffResponse
}

type StaffResponse struct {
	Staff *database.Staff
	Err   error
}
