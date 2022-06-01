package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	APIURL   = ""
	Port     = 0
	HashKey  []byte // Used to authenticate cookie
	BlockKey []byte // Used to encrypt cokie data
)

func Load() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	Port, err = strconv.Atoi(os.Getenv("APP_PORT"))
	if err != nil {
		log.Fatal(err)
	}

	APIURL = os.Getenv("API_URL")

	HashKey = []byte(os.Getenv("HASH_KEY"))

	BlockKey = []byte(os.Getenv("BLOCK_KEY"))
}
