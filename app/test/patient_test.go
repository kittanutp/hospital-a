package test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/kittanutp/hospital-app/database"
	"github.com/stretchr/testify/assert"
)

func TestSearchPatientByID(t *testing.T) {
	cfg, db := setUp(t)
	app := newTestServer(cfg, db)

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
	}
	db.GetSession().Create(&patient2)

	// Test GET /patient/search/:id for patient1 (NationalID)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/patient/search/"+nid, nil)
	app.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)

	var responsePatient1 database.Patient
	err := json.NewDecoder(w.Body).Decode(&responsePatient1)
	assert.NoError(t, err)
	assert.Equal(t, patient1.ID, responsePatient1.ID)

	// Test GET /patient/search/:id for patient2 (PassportID)
	w = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "/patient/search/"+pid, nil)
	fmt.Print("/patient/search/" + pid)
	app.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)

	var responsePatient2 database.Patient
	err = json.NewDecoder(w.Body).Decode(&responsePatient2)
	assert.NoError(t, err)
	assert.Equal(t, patient2.ID, responsePatient2.ID)

	resetDatabase(db)
}

// func TestSearchPatients(t *testing.T) {
// 	cfg, db := setUp()
// 	app := newTestServer(cfg, db)
// 	defer db.Close()

// 	// Setup patients for testing
// 	patient1 := database.Patient{
// 		NationalID:  "123456789",
// 		PassportID:  "A1234567",
// 		FirstName:   "John",
// 		LastName:    "Doe",
// 		DateOfBirth: "1990-01-01",
// 		PhoneNumber: "0123456789",
// 		Email:       "john.doe@example.com",
// 	}
// 	patient2 := database.Patient{
// 		NationalID:  "987654321",
// 		PassportID:  "B7654321",
// 		FirstName:   "Jane",
// 		LastName:    "Doe",
// 		DateOfBirth: "1992-01-01",
// 		PhoneNumber: "0987654321",
// 		Email:       "jane.doe@example.com",
// 	}
// 	db.Create(&patient1)
// 	db.Create(&patient2)

// 	// Test POST /patient/search
// 	reqBody := map[string]interface{}{
// 		"national_id": patient1.NationalID,
// 	}
// 	jsonBody, _ := json.Marshal(reqBody)
// 	req, _ := http.NewRequest("POST", "/patient/search", bytes.NewBuffer(jsonBody))
// 	w := httptest.NewRecorder()
// 	app.ServeHTTP(w, req)

// 	assert.Equal(t, http.StatusOK, w.Code)

// 	var patients []database.Patient
// 	err := json.NewDecoder(w.Body).Decode(&patients)
// 	assert.NoError(t, err)
// 	assert.Len(t, patients, 1) // Check that only one patient is returned
// 	assert.Equal(t, patient1.FirstName, patients[0].FirstName)
// }
