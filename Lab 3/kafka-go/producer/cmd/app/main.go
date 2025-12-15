package main

import (
	"log"
	"net/http"
	"producer-go/internal/handlers"
	"producer-go/internal/kafka"
	"producer-go/internal/middleware"
	"time"
)

func main() {

	var err error
	for i := 0; i < 5; i++ {
		err = kafka.InitProducer([]string{"kafka:9092"})
		if err == nil {
			break
		}
		log.Printf("Failed to init Kafka (attempt %d/5): %v", i+1, err)
		time.Sleep(2 * time.Second)
	}

	if err != nil {
		log.Fatal("Failed to initialize Kafka producer after 5 attempts")
	}

	defer kafka.CloseProducer()

	cat := &handlers.CatHandler{}

	mux := http.NewServeMux()

	mux.Handle("POST /api/v1/cat/add",
		middleware.ValidateCat(http.HandlerFunc(cat.AddCat)),
	)

	log.Println("Server started at http://localhost:8080")
	http.ListenAndServe(":8080", mux)
}
