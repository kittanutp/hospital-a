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

func TestSearchPatients(t *testing.T) {
	cfg, db := setUp(t)
	app := newTestServer(cfg, db)

	addTestStaff(t, db)
	token := loginTestStaff(app, t, cfg)

	// Setup patients for testing
	nid1 := "123456789"
	patient1 := database.Patient{
		NationalID:  &nid1,
		FirstNameEN: "John",
		LastNameEN:  "Doe",
		PatientHN:   testHospitalName,
	}
	db.GetSession().Create(&patient1)

	nid2 := "987654321"
	patient2 := database.Patient{
		NationalID:  &nid2,
		FirstNameEN: "Jane",
		LastNameEN:  "Doe",
		PatientHN:   testHospitalName,
	}
	db.GetSession().Create(&patient2)

	// Test POST /patient/search
	reqBody := map[string]interface{}{
		"national_id": nid1,
	}
	jsonBody, _ := json.Marshal(reqBody)
	req, _ := http.NewRequest("POST", "/patient/search", bytes.NewBuffer(jsonBody))
	req.Header.Set("Authorization", token)
	w := httptest.NewRecorder()
	app.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)

	var patients []schema.PatientJsonResponse
	err := json.NewDecoder(w.Body).Decode(&patients)
	assert.NoError(t, err)
	assert.Len(t, patients, 1) // Check that only one patient is returned
	assert.Equal(t, patient1.FirstNameEN, patients[0].FirstNameEN)

	resetDatabase(db)
}
