package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kittanutp/hospital-app/app/service"
)

type StaffAuthMiddleware struct {
	service service.StaffAuthServiceInterface
}

func NewAuthMiddleware(service service.StaffAuthServiceInterface) StaffAuthMiddlewareInterface {
	return &StaffAuthMiddleware{
		service: service,
	}
}

func (m *StaffAuthMiddleware) AuthStaff() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is required"})
			return
		}

		res := m.service.ProcessStaffToken(authHeader)
		if res.Err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": res.Err.Error()})
			return
		}

		c.Set("staff", res.Staff)
		c.Next()
	}
}
