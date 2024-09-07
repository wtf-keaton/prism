package main

import (
	"log"
	"prism-auth-service/internal/database"
	"prism-auth-service/internal/router"
)

func init() {
	err := database.Connect()

	if err != nil {
		log.Fatalf("Failed to connect to database cause: %s\n", err.Error())
	}
}

func main() {
	if err := router.Serve().Listen(":8080"); err != nil {
		log.Fatalf("Failed to start web listener cause: %s", err.Error())
	}
}
