package server

import (
	"fmt"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/kittanutp/hospital-app/app/config"
	"github.com/kittanutp/hospital-app/app/database"
	"github.com/kittanutp/hospital-app/app/handler"
	"github.com/kittanutp/hospital-app/app/repository"
	"github.com/kittanutp/hospital-app/app/service"
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
	g.initializePatientHttpHandler()
	g.initializeStaffHttpHandler()
	serverUrl := fmt.Sprintf(":%d", g.config.Server.Port)
	g.app.Run(serverUrl)
}

func (g *ginServer) initializePatientHttpHandler() {
	repo := repository.NewPatientPostgresRepository(g.db)
	service := service.NewPatientService(repo)
	handler := handler.NewPatientHTTPHandler(service)

	routes := g.app.Group("patient")
	{
		routes.GET("search/:id", handler.GetPatient)
		routes.POST("search", handler.GetPatients)
	}
}

func (g *ginServer) initializeStaffHttpHandler() {
	repo := repository.NewStaffPostgresRepository(g.db)
	service := service.NewStaffService(repo, *g.config.Server)
	handler := handler.NewStaffHTTPHandler(service)
	routes := g.app.Group("staff")
	{
		routes.POST("login", handler.LogIn)
		routes.POST("create", handler.CreateStaff)
	}

}
