package main

import (
	"github.com/kittanutp/hospital-app/app/config"
	"github.com/kittanutp/hospital-app/app/database"
	"github.com/kittanutp/hospital-app/app/server"
)

func main() {
	config := config.GetConfig()
	db := database.NewPostgresDatabase(config)
	server.NewGinServer(config, db).Start()
}
