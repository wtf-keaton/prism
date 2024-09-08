package main

import (
	"log"
	"prism-auth-service/internal/database"
)

func init() {
	err := database.Connect()

	if err != nil {
		log.Fatalf("Failed to connect to database cause: %s\n", err.Error())
	}
}

func main() {

}
