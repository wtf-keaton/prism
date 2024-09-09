package main

import (
	"log"
	"prism-api-gateway/internal/producer"
	"prism-api-gateway/internal/router"
)

func init() {
	err := producer.Connect()
	if err != nil {
		log.Fatalf("Failed to connect to producer cause: %s\n", err.Error())
	}
}

func main() {
	defer producer.Close()

	if err := router.Serve().Listen(":8080"); err != nil {
		log.Fatalf("Failed to start listen server with port: 8080")
	}
}
