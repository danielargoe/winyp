package internal

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func GetDotEnv() string {
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatal(err)
	}

	return os.Getenv("APIKEY")

}