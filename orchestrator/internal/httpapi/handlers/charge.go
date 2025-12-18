package handlers

import (
	"net/http"
	"orchestrator/internal/httpapi/response"
	"orchestrator/internal/usecase/charge"

	"github.com/gin-gonic/gin"
)

func (h *Handler) StartCharge(c *gin.Context) {
	_, err := h.startCharge.Execute(c.Request.Context(), charge.StartChargeRequest{})
	if err != nil {
		response.Fail(c, http.StatusNotImplemented, response.APIError{
			Code:    "NOT_IMPLEMENTED",
			Message: err.Error(),
		})
		return
	}

	response.NoContent(c)
}
