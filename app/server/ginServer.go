package server

import (
	"fmt"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/kittanutp/hospital-app/config"
	"github.com/kittanutp/hospital-app/database"
	"github.com/kittanutp/hospital-app/handler"
	"github.com/kittanutp/hospital-app/middleware"
	"github.com/kittanutp/hospital-app/repository"
	"github.com/kittanutp/hospital-app/service"
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
		AllowHeaders: []string{
			"Content-Type",
			"access-control-allow-origin",
			"access-control-allow-headers",
			"X-Real-IP",
			"X-Forwarded-For",
			"X-Forwarded-Proto",
		},
		AllowCredentials: true,
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
	patientRepo := repository.NewPatientPostgresRepository(g.db)
	patientService := service.NewPatientService(patientRepo)
	handler := handler.NewPatientHTTPHandler(patientService)

	authRepo := repository.NewStaffAuthPostgresRepository(g.db)
	authService := service.NewStaffAuthService(authRepo, *g.config.Server)
	middleware := middleware.NewAuthMiddleware(authService)

	routes := g.app.Group("patient")
	routes.Use(middleware.AuthStaff())
	{
		routes.GET("search/:id", handler.GetPatient)
		routes.POST("search", handler.GetPatients)
	}
}

func (g *ginServer) initializeStaffHttpHandler() {
	staffRepo := repository.NewStaffPostgresRepository(g.db)
	service := service.NewStaffService(staffRepo, *g.config.Server)
	handler := handler.NewStaffHTTPHandler(service)
	routes := g.app.Group("staff")
	routes.Use(gin.BasicAuth(gin.Accounts{
		g.config.Server.ServiceUsername: g.config.Server.ServicePassword,
	}))
	{
		routes.POST("login", handler.LogIn)
		routes.POST("create", handler.CreateStaff)
	}

}
