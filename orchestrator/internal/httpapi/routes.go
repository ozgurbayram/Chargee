package httpapi

import "github.com/gin-gonic/gin"

func (s *Server) registerRoutes(r *gin.Engine) {
	v1 := r.Group("/v1")
	{
		v1.GET("/health", s.handler.Health)
		v1.POST("/charge/start", s.handler.StartCharge)
		v1.POST("/charge/stop", s.handler.StopCharge)
	}
}
