package handler

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/kittanutp/hospital-app/schema"
	"github.com/kittanutp/hospital-app/service"
)

type StaffHTTPHandler struct {
	staffService service.StaffServiceInterface
}

func NewStaffHTTPHandler(service service.StaffServiceInterface) StaffHandler {
	return &StaffHTTPHandler{
		staffService: service,
	}
}

func (h *StaffHTTPHandler) LogIn(c *gin.Context) {
	var data schema.LogInSchema
	if err := c.BindJSON(&data); err != nil {
		c.AbortWithStatusJSON(400, fmt.Sprintf("invalid json request as %v", err.Error()))
		return
	}
	log.Printf("Login with user %v", data.Username)
	resp, err := h.staffService.ProcessLogIn(data)
	if err != nil {
		c.AbortWithStatusJSON(401, gin.H{"error": err.Error()})
		return
	}
	c.SecureJSON(200, resp)
}

func (h *StaffHTTPHandler) CreateStaff(c *gin.Context) {
	var data schema.CreateStaffSchema
	if err := c.BindJSON(&data); err != nil {
		c.AbortWithStatusJSON(400, fmt.Sprintf("invalid json request as %v", err.Error()))
		return
	}
	log.Printf("Creating staff with data: username=%v, hospitalName=%v", data.Username, data.HospitalName)
	resp := h.staffService.ProcessNewStaff(data)
	if resp.Err != nil {
		c.AbortWithStatusJSON(400, gin.H{"error": resp.Err.Error()})
		return
	}
	c.SecureJSON(200, gin.H{
		"id":            resp.Staff.ID,
		"username":      resp.Staff.Username,
		"hospital_name": resp.Staff.HospitalName,
	})
}
