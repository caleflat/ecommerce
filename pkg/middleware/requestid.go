package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

const (
	RequestIdKey = "request_id"
)

// RequestID is a middleware that injects a request ID into the context of each request.
// If X-Request-ID header is already present in the request, it is copied into the context
// and response headers.
// Otherwise, a random UUID is generated.
func RequestID() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestId := c.Request.Header.Get(RequestIdKey)
		if requestId == "" {
			requestId = uuid.New().String()
		}
		c.Set(RequestIdKey, requestId)
		c.Writer.Header().Set(RequestIdKey, requestId)
		c.Next()
	}
}
