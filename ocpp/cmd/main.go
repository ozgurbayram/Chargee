package main

import (
	"fmt"
	"log"
	"net/http"
	"ocpp/internal/config"
	"ocpp/internal/ws"
)

func main() {

	cfg := config.NewConfig()

	server := ws.NewServer()

	http.HandleFunc("/ocpp/", server.WsHandler)

	fmt.Println("Webserver started at", cfg.Port)

	if err := http.ListenAndServe(cfg.Port, nil); err != nil {
		log.Fatal(err)
	}
}
