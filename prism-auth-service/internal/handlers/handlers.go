package handlers

import (
	"log"
	"strings"

	"github.com/IBM/sarama"
)

func HandleRegister(message *sarama.ConsumerMessage) error {
	log.Printf("Message received: %s", string(message.Value))

	credentials := strings.Split(string(message.Value), " ")
	if len(credentials) != 2 {
		log.Println("Invalid message format")

		return nil
	}
	return nil
}
