package producer

import (
	"log"

	"github.com/IBM/sarama"
)

var producer sarama.SyncProducer

func Connect() error {
	var err error

	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 5
	config.Producer.Return.Successes = true

	producer, err = sarama.NewSyncProducer([]string{"localhost:9092"}, config)
	if err != nil {
		log.Printf("Failed to connect to Kafka: %v", err)

		return err
	}

	log.Println("Kafka producer connected successfully")

	return nil
}

func Get() sarama.SyncProducer {
	if producer == nil {
		log.Fatalf("Failed to get producer cuase it nil")
	}

	return producer
}

func Close() {
	if producer != nil {
		if err := producer.Close(); err != nil {
			log.Printf("Failed to close producer: %v", err)
		} else {
			log.Println("Producer closed successfully")
		}
	}
}
