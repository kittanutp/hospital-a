package main

import (
	"github.com/kittanutp/hospital-app/config"
	"github.com/kittanutp/hospital-app/database"
	"github.com/kittanutp/hospital-app/server"
)

func main() {
	config := config.GetConfig()
	db := database.NewPostgresDatabase(config)
	server.NewGinServer(config, db).Start()
}
