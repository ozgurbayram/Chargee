package main

import (
	"fmt"
	"log"
	"net/http"
	"ocpp/internal/config"
	"ocpp/internal/ocpp"
)

func main() {

	cfg := config.NewConfig()

	ocppService := ocpp.NewOcppService()

	http.HandleFunc("/ocpp/", ocppService.WsHandler)

	fmt.Println("Webserver started at", cfg.Port)

	if err := http.ListenAndServe(cfg.Port, nil); err != nil {
		log.Fatal(err)
	}
}
