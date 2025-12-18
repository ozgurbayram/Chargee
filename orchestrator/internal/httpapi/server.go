package httpapi

import (
	"net/http"
	"orchestrator/internal/config"
	"orchestrator/internal/httpapi/handlers"
	"orchestrator/internal/usecase/charge"

	"github.com/gin-gonic/gin"
)

type Server struct {
	httpServer *http.Server
	handler    *handlers.Handler
}

func New(cfg *config.Config) *Server {
	engine := gin.New()
	engine.Use(gin.Recovery())
	engine.Use(requestID())

	s := &Server{
		httpServer: &http.Server{
			Addr:    cfg.HttpAddress + cfg.HttpPort,
			Handler: engine,
		},
		handler: handlers.New(handlers.Deps{
			StartCharge: charge.NewStartChargeService(),
			StopCharge:  charge.NewStopChargeService(),
		}),
	}

	s.registerRoutes(engine)
	return s
}

func (s *Server) Run() error {
	return s.httpServer.ListenAndServe()
}
