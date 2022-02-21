package main

import (
	"api/src/config"
	"api/src/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	config.Load()

	router := router.Generate()

	fmt.Println("Running API on port", config.Port)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Port), router))
}
