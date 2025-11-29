package main

import (
	"fmt"
	"log"
	"net/http"

	"api/src/config"
	"api/src/router"
)

func main() {
	config.LoadEnv()

	router := router.NewRouter()

	fmt.Printf("Server running at %s\n", config.Port)

	log.Fatal(http.ListenAndServe(config.Port, router))
}
