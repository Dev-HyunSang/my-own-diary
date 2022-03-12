package config

import (
	"log"
	"os"

	"github.com/fatih/color"
	"github.com/joho/godotenv"
)

func GetEnv(env string) string {
	err := godotenv.Load()
	if err != nil {
		log.Println(color.RedString("ERROR"), "Failed to Load .env File")
		log.Fatalln(err)
	}

	getEnv := os.Getenv(env)
	return getEnv
}
