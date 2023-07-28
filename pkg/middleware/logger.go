package middleware

import (
	"github.com/gin-gonic/gin"
)

type loggerResponseWriter struct {
	gin.ResponseWriter
	statusCode int
}

func (w *loggerResponseWriter) WriteHeader(code int) {
	w.statusCode = code
	w.ResponseWriter.WriteHeader(code)
}

// Logger is a middleware that logs the request as it comes in and the response as it goes out.
func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		w := &loggerResponseWriter{c.Writer, 200}
		c.Writer = w

		// TODO: Finish implementing this middleware

		c.Next()
	}
}
