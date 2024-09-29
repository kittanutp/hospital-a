package handler

import (
	"errors"
	"fmt"
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

	res := h.patientService.ProcessGetPatient(id, *staff)
	if res.Err != nil {
		c.AbortWithStatusJSON(400, gin.H{"error": res.Err.Error()})
		return
	}
	c.SecureJSON(200, res.Patient)
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

	res := h.patientService.ProcessGetPatients(filter, *staff)
	if res.Err != nil {
		c.AbortWithStatusJSON(400, gin.H{"error": res.Err.Error()})
		return
	}
	c.SecureJSON(200, res.Patients)
}

func processStaffFromCtx(c *gin.Context) (*database.Staff, error) {
	data, exist := c.Get("staff")

	if !exist {
		return nil, errors.New("invalid staff")
	}
	staff, ok := data.(database.Staff)
	if !ok {
		return nil, errors.New("invalid staff format")
	}
	return &staff, nil
}
