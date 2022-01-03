package middleware

import (
	"github.com/gin-gonic/gin"
	"medicineApp/pkg/logger"
	"medicineApp/pkg/trace"
)

// TraceMiddleware 追踪 ID
func TraceMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		traceID := c.GetHeader("X-Request-Id")
		if traceID == "" {
			traceID = trace.NewTraceID()
		}

		ctx := logger.NewTraceIDContext(c.Request.Context(), traceID)
		c.Request = c.Request.WithContext(ctx)
		c.Writer.Header().Set("X-Trace-Id", traceID)

		c.Next()
	}
}
