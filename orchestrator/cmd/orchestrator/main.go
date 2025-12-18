package main

import (
	"log"
	"orchestrator/internal/config"
	"orchestrator/internal/httpapi"

	"errors"
	"net/http"
)

func main() {
	cfg := config.NewConfig()

	srv := httpapi.New(cfg)
	if err := srv.Run(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Fatal(err)
	}
}
