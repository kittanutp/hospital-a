package server

import (
	"fmt"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/kittanutp/hospital-app/app/config"
	"github.com/kittanutp/hospital-app/app/database"
)

type ginServer struct {
	app    *gin.Engine
	db     database.Database
	config *config.Config
}

func NewGinServer(config *config.Config, db database.Database) Server {
	app := gin.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins: config.Server.CORS,
		AllowMethods: []string{"*"},
		AllowHeaders: []string{"Content-Type,access-control-allow-origin, access-control-allow-headers"},
	}))
	app.SetTrustedProxies(config.Server.CORS)

	app.GET("/healthcheck", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "OK",
		})
	})

	return &ginServer{
		app:    app,
		db:     db,
		config: config,
	}
}

func (g *ginServer) Start() {
	g.app.Use(gin.Recovery())
	g.app.Use(gin.Logger())
	serverUrl := fmt.Sprintf(":%d", g.config.Server.Port)
	g.app.Run(serverUrl)
}
