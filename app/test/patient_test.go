package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/kittanutp/hospital-app/database"
	"github.com/kittanutp/hospital-app/schema"
	"github.com/stretchr/testify/assert"
)

func TestSearchPatientByID(t *testing.T) {
	cfg, db := setUp(t)
	app := newTestServer(cfg, db)

	addTestStaff(t, db)
	token := loginTestStaff(app, t, cfg)

	// Test data setup
	nid := "123456789"
	phoneNumber := "0123456789"
	patient1 := database.Patient{
		NationalID:  &nid,
		FirstNameEN: "John",
		LastNameEN:  "Doe",
		DateOfBirth: time.Date(1997, 2, 24, 0, 0, 0, 0, time.Local),
		PhoneNumber: &phoneNumber,
		Gender:      "M",
		PatientHN:   testHospitalName,
	}
	db.GetSession().Create(&patient1)

	pid := "987654321"
	patient2 := database.Patient{
		PassportID:  &pid,
		FirstNameEN: "Jane",
		LastNameEN:  "Doe",
		DateOfBirth: time.Date(1997, 1, 25, 0, 0, 0, 0, time.Local),
		PhoneNumber: &phoneNumber,
		Gender:      "F",
		PatientHN:   testHospitalName,
	}
	db.GetSession().Create(&patient2)

	// Test GET /patient/search/:id for patient1 (NationalID)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/patient/search/"+nid, nil)
	req.Header.Set("Authorization", token)
	app.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)

	var responsePatient1 schema.PatientJsonResponse
	err := json.NewDecoder(w.Body).Decode(&responsePatient1)
	assert.NoError(t, err)
	assert.Equal(t, patient1.ID, responsePatient1.ID)

	// Test GET /patient/search/:id for patient2 (PassportID)
	req, _ = http.NewRequest("GET", "/patient/search/"+pid, nil)
	req.Header.Set("Authorization", token)
	fmt.Print("/patient/search/" + pid)
	app.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)

	var responsePatient2 schema.PatientJsonResponse
	err = json.NewDecoder(w.Body).Decode(&responsePatient2)
	assert.NoError(t, err)
	assert.Equal(t, patient2.ID, responsePatient2.ID)

	resetDatabase(db)
}

func TestSearchPatientByIDNotFound(t *testing.T) {
	cfg, db := setUp(t)
	app := newTestServer(cfg, db)

	addTestStaff(t, db)
	token := loginTestStaff(app, t, cfg)

	// Test data setup
	nid := "123456789"
	phoneNumber := "0123456789"
	patient := database.Patient{
		NationalID:  &nid,
		FirstNameEN: "John",
		LastNameEN:  "Doe",
		DateOfBirth: time.Date(1997, 2, 24, 0, 0, 0, 0, time.Local),
		PhoneNumber: &phoneNumber,
		Gender:      "M",
		PatientHN:   "IncorrectHospital",
	}
	db.GetSession().Create(&patient)

	// Test GET /patient/search/:id for patient1 (NationalID)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/patient/search/"+nid, nil)
	req.Header.Set("Authorization", token)
	app.ServeHTTP(w, req)
	assert.Equal(t, 400, w.Code)
	resetDatabase(db)
}

func addTestPatients(db database.Database) (database.Patient, database.Patient, database.Patient, database.Patient) {
	nid1 := "12345678901"
	patient1 := database.Patient{
		NationalID:  &nid1,
		FirstNameTH: "สมชาย",
		LastNameTH:  "คำดี",
		PatientHN:   testHospitalName,
		Gender:      "M",
	}
	db.GetSession().Create(&patient1)

	nid2 := "12345678902"
	patient2 := database.Patient{
		NationalID:  &nid2,
		FirstNameTH: "สมหญิง",
		LastNameTH:  "คำดี",
		PatientHN:   testHospitalName,
		Gender:      "F",
	}
	db.GetSession().Create(&patient2)

	pid3 := "98765432101"
	patient3 := database.Patient{
		PassportID:  &pid3,
		FirstNameEN: "Jane",
		LastNameEN:  "Doe",
		PatientHN:   testHospitalName,
		Gender:      "M",
	}
	db.GetSession().Create(&patient3)

	pid4 := "98765432102"
	patient4 := database.Patient{
		PassportID:  &pid4,
		FirstNameEN: "Jane",
		LastNameEN:  "Doe",
		PatientHN:   "IncorrectHospital",
		Gender:      "F",
	}
	db.GetSession().Create(&patient4)

	return patient1, patient2, patient3, patient4
}

type testJSONResponse struct {
	Data []schema.PatientJsonResponse `json:"data"`
}

func decodedJSONResponse(t *testing.T, w *httptest.ResponseRecorder) testJSONResponse {
	assert.Equal(t, 200, w.Code)
	var response testJSONResponse
	err := json.NewDecoder(w.Body).Decode(&response)
	assert.NoError(t, err)
	return response
}

func TestSearchPatientsWithNoFilter(t *testing.T) {
	cfg, db := setUp(t)
	app := newTestServer(cfg, db)

	addTestStaff(t, db)
	token := loginTestStaff(app, t, cfg)

	addTestPatients(db)
	var reqBody map[string]interface{}
	jsonBody, _ := json.Marshal(reqBody)
	req, _ := http.NewRequest("POST", "/patient/search", bytes.NewBuffer(jsonBody))
	req.Header.Set("Authorization", token)
	w := httptest.NewRecorder()
	app.ServeHTTP(w, req)
	response := decodedJSONResponse(t, w)
	assert.Len(t, response.Data, 3)
	resetDatabase(db)
}

func TestSearchPatientsWithSingleFilter(t *testing.T) {
	cfg, db := setUp(t)
	app := newTestServer(cfg, db)

	addTestStaff(t, db)
	token := loginTestStaff(app, t, cfg)

	patient1, _, _, _ := addTestPatients(db)
	reqBody := map[string]interface{}{
		"national_id": *patient1.NationalID,
	}
	jsonBody, _ := json.Marshal(reqBody)
	req, _ := http.NewRequest("POST", "/patient/search", bytes.NewBuffer(jsonBody))
	req.Header.Set("Authorization", token)
	w := httptest.NewRecorder()
	app.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)

	response := decodedJSONResponse(t, w)
	assert.NotEmpty(t, response.Data)
	assert.Len(t, response.Data, 1)
	assert.Equal(t, patient1.FirstNameTH, response.Data[0].FirstNameTH)
	resetDatabase(db)
}

func TestSearchPatientsWithORFilter(t *testing.T) {
	cfg, db := setUp(t)
	app := newTestServer(cfg, db)

	addTestStaff(t, db)
	token := loginTestStaff(app, t, cfg)

	patient1, _, _, _ := addTestPatients(db)

	pid5 := "98765432103"
	patient5 := database.Patient{
		PassportID:  &pid5,
		FirstNameEN: patient1.FirstNameTH,
		LastNameEN:  "Doe",
		PatientHN:   testHospitalName,
		Gender:      "M",
	}
	db.GetSession().Create(&patient5)

	reqBody := map[string]interface{}{
		"first_name": patient1.FirstNameTH,
	}
	jsonBody, _ := json.Marshal(reqBody)
	req, _ := http.NewRequest("POST", "/patient/search", bytes.NewBuffer(jsonBody))
	req.Header.Set("Authorization", token)
	w := httptest.NewRecorder()
	app.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)

	response := decodedJSONResponse(t, w)
	assert.NotEmpty(t, response.Data)
	assert.Len(t, response.Data, 2)
	assert.Equal(t, patient1.FirstNameTH, response.Data[0].FirstNameTH)
	assert.Equal(t, patient1.FirstNameTH, response.Data[1].FirstNameEN)
	resetDatabase(db)
}
