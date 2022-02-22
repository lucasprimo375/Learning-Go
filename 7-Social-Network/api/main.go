package main

import (
	"api/src/config"
	"api/src/router"
	"fmt"
	"log"
	"net/http"
)

/*func init() {
	key := make([]byte, 64)

	_, error := rand.Read(key)
	if error != nil {
		log.Fatal(error)
	}

	base64String := base64.StdEncoding.EncodeToString(key)
}*/

func main() {
	config.Load()

	router := router.Generate()

	fmt.Println("Running API on port", config.Port)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Port), router))
}
