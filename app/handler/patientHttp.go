package handler

import (
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kittanutp/hospital-app/database"
	"github.com/kittanutp/hospital-app/schema"
	"github.com/kittanutp/hospital-app/service"
)

type patientHTTPHandler struct {
	patientService service.PatientServiceInterface
}

func NewPatientHTTPHandler(service service.PatientServiceInterface) PatientHandler {
	return &patientHTTPHandler{
		patientService: service,
	}
}

func (h *patientHTTPHandler) GetPatient(c *gin.Context) {
	id := c.Param("id")
	staff, err := processStaffFromCtx(c)
	if err != nil {
		c.AbortWithStatusJSON(401, gin.H{"error": err.Error()})
		return
	}

	res := h.patientService.ProcessGetPatient(id, staff)
	if res.Err != nil {
		c.AbortWithStatusJSON(400, gin.H{"error": res.Err.Error()})
		return
	}

	c.SecureJSON(200, schema.ConvertJSONResponse(res.Patient))
}

func (h *patientHTTPHandler) GetPatients(c *gin.Context) {
	var filter schema.PatientFilterSchema
	if err := c.ShouldBindJSON(&filter); err != nil {
		if c.Request.Body != http.NoBody {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error": fmt.Sprintf("invalid json request: %v", err.Error()),
			})
			return
		}
	}

	staff, err := processStaffFromCtx(c)

	if err != nil {
		c.AbortWithStatusJSON(401, gin.H{"error": err.Error()})
		return
	}
	log.Printf("Get staff with data: username=%v, hospitalName=%v", staff.Username, staff.HospitalName)
	res := h.patientService.ProcessGetPatients(filter, staff)
	if res.Err != nil {
		c.AbortWithStatusJSON(400, gin.H{"error": res.Err.Error()})
		return
	}
	var patients []schema.PatientJsonResponse
	fmt.Println(len(res.Patients))
	for _, p := range res.Patients {
		patients = append(patients, schema.ConvertJSONResponse(p))
	}

	c.SecureJSON(200, gin.H{"data": patients})
}

func processStaffFromCtx(c *gin.Context) (database.Staff, error) {
	data, exist := c.Get("staff")

	if !exist {
		return database.Staff{}, errors.New("invalid staff")
	}
	staff, ok := data.(database.Staff)
	if !ok {
		return database.Staff{}, fmt.Errorf("invalid staff format, expected database.Staff, got=%T", data)
	}
	return staff, nil
}
