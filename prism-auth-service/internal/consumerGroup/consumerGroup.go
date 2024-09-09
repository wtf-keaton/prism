package consumerGroup

import (
	"prism-auth-service/internal/handlers"

	"github.com/IBM/sarama"
)

type AuthServiceConsumer struct{}

func (c *AuthServiceConsumer) Setup(sarama.ConsumerGroupSession) error {
	return nil
}

func (c *AuthServiceConsumer) Cleanup(sarama.ConsumerGroupSession) error {
	return nil
}

func (c *AuthServiceConsumer) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for message := range claim.Messages() {
		switch message.Topic {
		case "register_request":
			go handlers.HandleRegister(message)
		}

		session.MarkMessage(message, "")
	}

	return nil
}
