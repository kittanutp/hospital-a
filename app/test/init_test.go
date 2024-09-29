package test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/kittanutp/hospital-app/config"
	"github.com/kittanutp/hospital-app/database"
	"github.com/kittanutp/hospital-app/handler"
	"github.com/kittanutp/hospital-app/middleware"
	"github.com/kittanutp/hospital-app/repository"
	"github.com/kittanutp/hospital-app/service"
	"github.com/stretchr/testify/assert"
)

func setUp(t *testing.T) (*config.Config, database.Database) {
	t.Setenv("DB_NAME", "hospital-test")
	t.Setenv("DB_PASSWORD", "testpassword")
	t.Setenv("DB_USER", "testuser")
	t.Setenv("DB_HOST", "localhost")
	t.Setenv("DB_PORT", "5433")
	t.Setenv("OAUTH_KEY", "test")
	t.Setenv("SERVICE_USER", "test")
	t.Setenv("SERVICE_PWD", "test")
	t.Setenv("CORS", "http://127.0.0.1:8910")
	t.Setenv("GIN_MODE", "test")

	cfg := config.GetConfig()
	db := database.NewPostgresDatabase(cfg)
	return cfg, db
}

func addPatientHandler(db database.Database) handler.PatientHandler {
	patientRepo := repository.NewPatientPostgresRepository(db)
	patientService := service.NewPatientService(patientRepo)
	handler := handler.NewPatientHTTPHandler(patientService)
	return handler
}

func getAuthMiddleware(cfg *config.Config, db database.Database) middleware.StaffAuthMiddlewareInterface {
	authRepo := repository.NewStaffAuthPostgresRepository(db)
	authService := service.NewStaffAuthService(authRepo, *cfg.Server)
	return middleware.NewAuthMiddleware(authService)
}

func addStaffHandler(cfg *config.Config, db database.Database) handler.StaffHandler {
	staffRepo := repository.NewStaffPostgresRepository(db)
	service := service.NewStaffService(staffRepo, *cfg.Server)
	return handler.NewStaffHTTPHandler(service)
}

func newTestServer(cfg *config.Config, db database.Database) *gin.Engine {
	app := gin.New()
	app.GET("/healthcheck", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "OK",
		})
	})

	patienthandler := addPatientHandler(db)
	middleware := getAuthMiddleware(cfg, db)
	patientRoutes := app.Group("patient")
	patientRoutes.Use(middleware.AuthStaff())
	{
		patientRoutes.GET("search/:id", patienthandler.GetPatient)
		patientRoutes.POST("search", patienthandler.GetPatients)
	}

	staffHandler := addStaffHandler(cfg, db)
	staffRoutes := app.Group("staff")
	staffRoutes.Use(gin.BasicAuth(gin.Accounts{
		cfg.Server.ServiceUsername: cfg.Server.ServicePassword,
	}))
	{
		staffRoutes.POST("login", staffHandler.LogIn)
		staffRoutes.POST("create", staffHandler.CreateStaff)
	}

	return app
}

func TestHealthCheck(t *testing.T) {
	cfg, db := setUp(t)
	app := newTestServer(cfg, db)

	resp := httptest.NewRequest(http.MethodGet, "/healthcheck", nil)
	w := httptest.NewRecorder()
	app.ServeHTTP(w, resp)
	assert.Equal(t, http.StatusOK, w.Code)
}

func resetDatabase(db database.Database) {
	db.GetSession().Exec("TRUNCATE TABLE patients RESTART IDENTITY CASCADE")
	db.GetSession().Exec("TRUNCATE TABLE staffs RESTART IDENTITY CASCADE")
}
