package main

import (
	"log"
	"os"
	"os/signal"
	"prism-auth-service/internal/postgres"
	"syscall"
)

func init() {
	err := postgres.Connect()

	if err != nil {
		log.Fatalf("Failed to connect to database cause: %s\n", err.Error())
	}
}

func main() {

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)

	// Waiting for SIGINT
	<-stop

	postgres.Close()
}
