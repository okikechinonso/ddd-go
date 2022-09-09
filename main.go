package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {

	app := fiber.New()
	port := os.Getenv("PORT")
	log.Println(float64(1) / 3)
	err := godotenv.Load()
	if err != nil {
		log.Println(err.Error())
	}

	app.Listen(":" + port)
}
