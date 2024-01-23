package app

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/rakesh-gupta29/json-countries/cmd/routers"
	"github.com/rakesh-gupta29/json-countries/pkg/config"
)

func LoadAndRun() error {
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime) // migrate to the logger package

	app_config, err := config.LoadConfig()
	if err != nil {
		logger.Fatal("Error loading config file")
	}

	server := &http.Server{
		Addr:         app_config.Address,
		Handler:      routers.MountAllRoutes(),
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
		IdleTimeout:  time.Minute,
	}
	log.Printf("Starting the server on %s", app_config.Address)
	server.ListenAndServe()
	return nil
}
