package internal

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func GetDotEnv(path string) string {
	err := godotenv.Load(path)
	if err != nil {
		log.Fatal(err)
	}

	return os.Getenv("APIKEY")

}