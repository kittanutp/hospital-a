package database

import (
	"time"

	"gorm.io/gorm"
)

type Patient struct {
	gorm.Model
	FirstNameTH  string    `gorm:"column:first_name_th;type:varchar(100);not null" json:"first_name_th"`
	MiddleNameTH *string   `gorm:"column:middle_name_th;type:varchar(100)" json:"middle_name_th,omitempty"`
	LastNameTH   string    `gorm:"column:last_name_th;type:varchar(100);not null" json:"last_name_th"`
	FirstNameEN  string    `gorm:"column:first_name_en;type:varchar(100);not null" json:"first_name_en"`
	MiddleNameEN *string   `gorm:"column:middle_name_en;type:varchar(100)" json:"middle_name_en,omitempty"`
	LastNameEN   string    `gorm:"column:last_name_en;type:varchar(100);not null" json:"last_name_en"`
	DateOfBirth  time.Time `gorm:"column:date_of_birth;type:date;not null" json:"date_of_birth"`
	PatientHN    string    `gorm:"column:patient_hospital_name;type:varchar(50);index;not null" json:"patient_hn"`
	NationalID   *string   `gorm:"column:national_id;type:varchar(50);unique" json:"national_id,omitempty"`
	PassportID   *string   `gorm:"column:passport_id;type:varchar(50);unique" json:"passport_id,omitempty"`
	PhoneNumber  *string   `gorm:"column:phone_number;type:varchar(20);index" json:"phone_number,omitempty"`
	Email        *string   `gorm:"column:email;type:varchar(100);index" json:"email,omitempty"`
	Gender       string    `gorm:"column:gender;type:char(1);check:gender IN ('M', 'F');not null" json:"gender"`
}

type Staff struct {
	gorm.Model
	Username     string `gorm:"column:username;type:varchar(100);unique;not null" json:"username"`
	Password     string `gorm:"column:password;type:varchar(255);not null" json:"password"`
	Salt         string `gorm:"column:salt;type:varchar(255);not null" json:"salt"`
	HospitalName string `gorm:"column:hospital_name;type:varchar(255);not null" json:"hospital_name"`
}
