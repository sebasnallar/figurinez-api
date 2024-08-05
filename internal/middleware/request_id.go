package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func RequestID() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestID := uuid.New().String()
		c.Writer.Header().Set("X-Request-ID", requestID)
		c.Set("X-Request-ID", requestID)
		c.Next()
	}
}