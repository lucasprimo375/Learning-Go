package main

import (
	"command-line-app/app"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("Starting")

	application := app.Generate()
	error := application.Run(os.Args)
	if error != nil {
		log.Fatal(error)
	}
}
