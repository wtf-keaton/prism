package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"prism-auth-service/internal/consumerGroup"
	"prism-auth-service/internal/database"
	"sync"
	"syscall"

	"github.com/IBM/sarama"
)

func init() {
	err := database.Connect()

	if err != nil {
		log.Fatalf("Failed to connect to database cause: %s\n", err.Error())
	}
}

func main() {
	config := sarama.NewConfig()
	config.Consumer.Group.Rebalance.Strategy = sarama.NewBalanceStrategyRoundRobin()

	groupID := "auth-service-group"
	topics := []string{"register_request", "login_request", "validate_token", "refresh_token"}

	consumerGroups, err := sarama.NewConsumerGroup([]string{"localhost:9092"}, groupID, config)
	if err != nil {
		log.Fatalf("Failed to start consumer group: %v", err)
	}
	defer consumerGroups.Close()

	ctx, cancel := context.WithCancel(context.Background())

	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-sigchan
		log.Println("Received termination signal, shutting down...")
		cancel()
	}()

	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			if err := consumerGroups.Consume(ctx, topics, &consumerGroup.AuthServiceConsumer{}); err != nil {
				log.Printf("Error from consumer: %v", err)
			}

			if ctx.Err() != nil {
				return
			}
		}
	}()
	wg.Wait()

	log.Println("Service stopped.")
}
