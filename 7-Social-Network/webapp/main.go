package main

import (
	"fmt"
	"log"
	"net/http"

	"webapp/src/config"
	"webapp/src/cookies"
	"webapp/src/router"
	"webapp/src/utils"
)

/*func init() {
	hashKey := hex.EncodeToString(securecookie.GenerateRandomKey(16))
	fmt.Println(hashKey)

	blockKey := hex.EncodeToString(securecookie.GenerateRandomKey(16))
	fmt.Println(blockKey)
}*/

func main() {
	config.Load()

	cookies.Configure()

	utils.LoadTemplates()

	r := router.Generate()

	fmt.Printf("Running WebApp at port %d\n", config.Port)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Port), r))
}
