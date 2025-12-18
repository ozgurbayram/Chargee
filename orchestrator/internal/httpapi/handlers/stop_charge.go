package handlers

import (
	"net/http"
	"orchestrator/internal/httpapi/response"
	"orchestrator/internal/usecase/charge"

	"github.com/gin-gonic/gin"
)

func (h *Handler) StopCharge(c *gin.Context) {
	_, err := h.stopCharge.Execute(c.Request.Context(), charge.StopChargeRequest{})
	if err != nil {
		response.Fail(c, http.StatusNotImplemented, response.APIError{
			Code:    "NOT_IMPLEMENTED",
			Message: err.Error(),
		})
		return
	}

	response.NoContent(c)
}
