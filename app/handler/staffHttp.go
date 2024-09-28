package handler

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/kittanutp/hospital-app/app/schema"
	"github.com/kittanutp/hospital-app/app/service"
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
		c.AbortWithStatusJSON(400, fmt.Sprintf("invalid json request as %v", err))
		return
	}
	resp, err := h.staffService.ProcessLogIn(data)
	if err != nil {
		c.AbortWithStatus(401)
		return
	}
	c.SecureJSON(200, resp)
}

func (h *StaffHTTPHandler) CreateStaff(c *gin.Context) {
	var data schema.CreateStaffSchema
	if err := c.BindJSON(&data); err != nil {
		c.AbortWithStatusJSON(400, fmt.Sprintf("invalid json request as %v", err))
		return
	}
	resp := h.staffService.ProcessNewStaff(data)
	if resp.Err != nil {
		c.AbortWithStatusJSON(400, resp.Err)
		return
	}
	c.SecureJSON(200, gin.H{
		"id":       resp.Staff.ID,
		"username": resp.Staff.Username,
	})
}
