package cmd

import (
	"log"
	"net/http"
	"orchestrator/internal/config"
)

func main() {
	cfg := config.NewConfig()

	if err := http.ListenAndServe(cfg.HttpPort, nil); err != nil {
		log.Fatal(err)
	}
}
