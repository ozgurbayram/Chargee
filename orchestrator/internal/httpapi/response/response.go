package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const TraceIDKey = "trace_id"

type APIError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type Base[T any] struct {
	Success bool      `json:"success"`
	Message string    `json:"message,omitempty"`
	Data    *T        `json:"data,omitempty"`
	Error   *APIError `json:"error,omitempty"`
	TraceID string    `json:"traceId,omitempty"`
}

func OK[T any](c *gin.Context, status int, message string, data T) {
	c.JSON(status, Base[T]{
		Success: true,
		Message: message,
		Data:    &data,
		TraceID: c.GetString(TraceIDKey),
	})
}

func NoContent(c *gin.Context) {
	c.Status(http.StatusNoContent)
}

func Fail(c *gin.Context, status int, err APIError) {
	c.JSON(status, Base[any]{
		Success: false,
		Error:   &err,
		TraceID: c.GetString(TraceIDKey),
	})
}
