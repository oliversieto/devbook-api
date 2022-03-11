package main

import (
	"devbook-api/src/config"
	"devbook-api/src/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	config.Load()
	router := router.Generate()

	fmt.Println("Listening on port 5000")

	log.Fatal(http.ListenAndServe(config.ApiPort, router))
}
