package database

import (
	"fmt"
	"sync"

	"github.com/kittanutp/hospital-app/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type postgresDatabase struct {
	Db *gorm.DB
}

var (
	once       sync.Once
	dbInstance *postgresDatabase
)

func NewPostgresDatabase(conf *config.Config) Database {
	once.Do(func() {

		psqlInfo := fmt.Sprintf(
			"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
			conf.Db.Host,
			conf.Db.Port,
			conf.Db.User,
			conf.Db.Password,
			conf.Db.DBName,
		)

		db, err := gorm.Open(postgres.Open(psqlInfo), &gorm.Config{})
		if err != nil {
			panic("failed to connect database")
		}

		// Migrate the schema
		extensionStm := `CREATE EXTENSION IF NOT EXISTS "uuid-ossp";`
		db.Statement.Exec(extensionStm)
		db.AutoMigrate(&Patient{}, &Staff{})

		dbInstance = &postgresDatabase{Db: db}
	})

	return dbInstance
}

func (p *postgresDatabase) GetSession() *gorm.DB {
	return dbInstance.Db
}
