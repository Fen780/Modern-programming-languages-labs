package kafka

import (
	"log"

	"github.com/IBM/sarama"
)

var consumer sarama.Consumer

func InitConsumer(brokers []string) error {
	var err error

	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true
	consumer, err = sarama.NewConsumer(brokers, config)
	if err != nil {
		return err
	}
	log.Println("Consumer initialized")
	return nil
}

func CloseConsumer() {
	if consumer != nil {
		if err := consumer.Close(); err != nil {
			log.Fatalln(err)
		}
	}
}

func StartConsumer(topic string, handleFunc func([]byte) error) {
	if consumer == nil {
		log.Println("Consumer not initialized")
		return
	}

	partitionConsumer, err := consumer.ConsumePartition(topic, 0, sarama.OffsetNewest)
	if err != nil {
		log.Printf("Failed to start partition consumer: %v", err)
		return
	}
	defer partitionConsumer.Close()

	log.Printf("Started consuming topic: %s", topic)

	for msg := range partitionConsumer.Messages() {
		log.Printf("Received message offset %d", msg.Offset)

		if err := handleFunc(msg.Value); err != nil {
			log.Printf("Error processing message: %v", err)
		}
	}
}
