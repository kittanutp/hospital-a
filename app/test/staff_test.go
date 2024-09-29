package test

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/kittanutp/hospital-app/database"
	"github.com/kittanutp/hospital-app/service"
	"github.com/stretchr/testify/assert"
)

const (
	testStaffUsername = "TestStaff"
	testStaffPassword = "Password1#"
	testHospitalName  = "TestHospital"
)

func basicAuth(username, password string) string {
	auth := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(auth))
}

func TestCreateStaff(t *testing.T) {
	cfg, db := setUp(t)

	app := newTestServer(cfg, db)

	// Create staff payload
	payload := `{
		"username": "testAddStaff",
		"password": "testpassword",
		"hospital_name": "testAddStaff"
	}`

	// Test POST /staff/create
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/staff/create", strings.NewReader(payload))
	req.Header.Set("Authorization", "Basic "+basicAuth(cfg.Server.ServiceUsername, cfg.Server.ServicePassword)) // Assuming Basic Auth
	app.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	var responseBody map[string]interface{}
	err := json.NewDecoder(w.Body).Decode(&responseBody)
	assert.NoError(t, err)
	assert.Contains(t, responseBody, "id")
	assert.Equal(t, "testAddStaff", responseBody["username"])

	// Check if the staff was actually created in the database
	var createdStaff database.Staff
	db.GetSession().Where("username = ?", "testAddStaff").First(&createdStaff)
	assert.Equal(t, "testAddStaff", createdStaff.Username)
	resetDatabase(db)
}

func TestStaffLogin(t *testing.T) {
	cfg, db := setUp(t)

	app := newTestServer(cfg, db)

	addTestStaff(t, db)

	loginRequest := map[string]string{
		"username": testStaffUsername,
		"password": testStaffPassword,
	}
	body, _ := json.Marshal(loginRequest)

	// Test POST /staff/login
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/staff/login", bytes.NewBuffer(body))
	req.Header.Set("Authorization", "Basic "+basicAuth(cfg.Server.ServiceUsername, cfg.Server.ServicePassword))
	app.ServeHTTP(w, req)

	// Check response
	assert.Equal(t, 200, w.Code)
	var responseBody map[string]interface{}
	err := json.NewDecoder(w.Body).Decode(&responseBody)
	assert.NoError(t, err)
	assert.Contains(t, responseBody, "token")
	assert.Contains(t, responseBody, "token_type")
	resetDatabase(db)
}

func addTestStaff(t *testing.T, db database.Database) database.Staff {
	pwd, err := service.EncryptPassword(testStaffPassword, "TestSalt")
	assert.NoError(t, err)

	staff := database.Staff{
		Username: testStaffUsername,
		Password: pwd,
		Salt:     "TestSalt",
	}
	db.GetSession().Create(&staff)
	return staff
}
