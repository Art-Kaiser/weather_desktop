package configs

import (
	"github.com/joho/godotenv"
	"log"
)

func InitConfig() {
	if err := godotenv.Load("./configs/.env"); err != nil {
		log.Print("No .env file found")
	}
}
