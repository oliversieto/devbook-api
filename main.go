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

	fmt.Printf("Listening on port %s\n", config.ApiPort)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", config.ApiPort), router))
}
