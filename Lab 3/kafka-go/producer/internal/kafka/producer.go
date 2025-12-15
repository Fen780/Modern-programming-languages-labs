package kafka

import (
	"encoding/json"
	"log"
	"producer-go/internal/entities"

	"github.com/IBM/sarama"
)

var producer sarama.SyncProducer

func InitProducer(brokers []string) error {
	var err error
	producer, err = sarama.NewSyncProducer(brokers, nil)
	if err != nil {
		log.Println(err)
		return err
	}
	log.Println("Producer initialized")
	return nil
}

func CloseProducer() {
	if producer != nil {
		if err := producer.Close(); err != nil {
			log.Println(err)
		}
	}
}

func SendMessage(cat entities.Cat) {

	catJSON, err := json.Marshal(cat)
	if err != nil {
		log.Println(err)
		return
	}

	if producer == nil {
		log.Fatal("Kafka producer is not initialized!")
		return
	}

	msg := &sarama.ProducerMessage{Topic: "my_topic", Value: sarama.ByteEncoder(catJSON)}
	partition, offset, err := producer.SendMessage(msg)
	if err != nil {
		log.Printf("FAILED to send message: %s\n", err)
	} else {
		log.Printf("> message sent to partition %d at offset %d\n", partition, offset)
	}
}
