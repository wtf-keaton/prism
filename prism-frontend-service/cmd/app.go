package main

import (
	"log"
	"prism-frontend-service/internal/router"
)

func main() {
	if err := router.Serve().Listen(":3000"); err != nil {
		log.Fatalf("Failed to launch frontend-service cause: %s\n", err.Error())
	}
}
