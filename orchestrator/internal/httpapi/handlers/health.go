package handlers

import (
	"net/http"
	"orchestrator/internal/httpapi/response"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Health(c *gin.Context) {
	response.OK(c, http.StatusOK, "ok", gin.H{"status": "up"})
}
