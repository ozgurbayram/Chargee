package httpapi

import (
	"crypto/rand"
	"encoding/hex"
	"orchestrator/internal/httpapi/response"

	"github.com/gin-gonic/gin"
)

func requestID() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.GetHeader("X-Request-Id")
		if id == "" {
			id = newTraceID()
		}

		c.Set(response.TraceIDKey, id)
		c.Header("X-Request-Id", id)
		c.Next()
	}
}

func newTraceID() string {
	var b [16]byte
	if _, err := rand.Read(b[:]); err != nil {
		return "trace_unavailable"
	}
	return hex.EncodeToString(b[:])
}
