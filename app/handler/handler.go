package handler

import "github.com/gin-gonic/gin"

type PatientHandler interface {
	GetPatient(c *gin.Context)
	GetPatients(c *gin.Context)
}

type StaffHandler interface {
	LogIn(c *gin.Context)
	CreateStaff(c *gin.Context)
}
