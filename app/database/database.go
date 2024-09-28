package database

import (
	_ "github.com/lib/pq"
	"gorm.io/gorm"
)

type Database interface {
	GetSession() *gorm.DB
}
