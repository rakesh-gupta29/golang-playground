package main

import (
	"log"

	"github.com/rakesh-gupta29/json-countries/app"
)

func main () {
  err := app.LoadAndRun()
  if err != nil {
    log.Fatal("Error starting the server.")
  }

}

