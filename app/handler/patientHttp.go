package handler

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kittanutp/hospital-app/app/database"
	"github.com/kittanutp/hospital-app/app/schema"
	"github.com/kittanutp/hospital-app/app/service"
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
		c.AbortWithStatusJSON(401, gin.H{
			"error": err,
		})
		return
	}

	res := h.patientService.ProcessGetPatient(id, *staff)
	if res.Err != nil {
		c.AbortWithStatusJSON(400, res.Err)
		return
	}
	c.SecureJSON(http.StatusOK, res.Patient)
}

func (h *patientHTTPHandler) GetPatients(c *gin.Context) {
	var filter schema.PatientFilterSchema
	if err := c.ShouldBindJSON(&filter); err != nil {
		if c.Request.Body != http.NoBody {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error": fmt.Sprintf("invalid json request: %v", err),
			})
			return
		}
	}

	staff, err := processStaffFromCtx(c)

	if err != nil {
		c.AbortWithStatusJSON(401, gin.H{
			"error": err,
		})
		return
	}

	res := h.patientService.ProcessGetPatients(filter, *staff)
	if res.Err != nil {
		c.AbortWithStatusJSON(400, res.Err)
		return
	}
	c.SecureJSON(http.StatusOK, res.Patients)
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
