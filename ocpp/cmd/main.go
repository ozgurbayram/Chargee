package main

import (
	"fmt"
	"log"
	"net/http"
	"ocpp/internal/config"
	"ocpp/internal/infra"
	"ocpp/internal/ocpp"
)

func main() {
	cfg := config.NewConfig()

	cluster, bucket, err := infra.CouchbaseInitialization(cfg)
	if err != nil {
		log.Fatal(err)
	}

	defer cluster.Close(nil)

	repository := infra.NewCouchbaseChargePointRepository(bucket)
	ocppService := ocpp.NewOcppService(repository)

	http.HandleFunc("/ocpp/", ocppService.WsHandler)

	fmt.Println("Webserver started at", cfg.Port)

	if err := http.ListenAndServe(cfg.Port, nil); err != nil {
		log.Fatal(err)
	}
}
