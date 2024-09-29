package middleware

import "github.com/gin-gonic/gin"

type StaffAuthMiddlewareInterface interface {
	AuthStaff() gin.HandlerFunc
}
