package database

import (
	"time"

	"gorm.io/gorm"
)

type Patient struct {
	gorm.Model
	FirstNameTH  string    `gorm:"column:first_name_th;type:varchar(100);not null"`
	MiddleNameTH *string   `gorm:"column:middle_name_th;type:varchar(100)"`
	LastNameTH   string    `gorm:"column:last_name_th;type:varchar(100);not null"`
	FirstNameEN  string    `gorm:"column:first_name_en;type:varchar(100);not null"`
	MiddleNameEN *string   `gorm:"column:middle_name_en;type:varchar(100)"`
	LastNameEN   string    `gorm:"column:last_name_en;type:varchar(100);not null"`
	DateOfBirth  time.Time `gorm:"column:date_of_birth;type:date;not null"`
	PatientHN    string    `gorm:"column:patient_hospital_name;type:varchar(50);index;not null"`
	NationalID   *string   `gorm:"column:national_id;type:varchar(50);unique"`
	PassportID   *string   `gorm:"column:passport_id;type:varchar(50);unique"`
	PhoneNumber  *string   `gorm:"column:phone_number;type:varchar(20);index"`
	Email        *string   `gorm:"column:email;type:varchar(100);index"`
	Gender       string    `gorm:"column:gender;type:char(1);check:gender IN ('M', 'F');not null"`
}

type Staff struct {
	gorm.Model
	Username     string `gorm:"column:username;type:varchar(100);unique;not null"`
	Password     string `gorm:"column:password;type:varchar(255);not null"`
	Salt         string `gorm:"column:salt;type:varchar(255);not null"`
	HospitalName string `gorm:"column:hospital_name;type:varchar(255);not null"`
}
