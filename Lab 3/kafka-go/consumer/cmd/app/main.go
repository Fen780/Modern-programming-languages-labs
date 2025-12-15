package main

import (
	"consumer/internal/database"
	"consumer/internal/entities"
	"consumer/internal/handlers"
	"consumer/internal/kafka"
	"encoding/json"
	"log"
	"net/http"
)

func main() {
	db, err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	if err := kafka.InitConsumer([]string{"kafka:9092"}); err != nil {
		log.Fatal(err)
	}
	defer kafka.CloseConsumer()

	catHandler := &handlers.CatHandler{DB: db}

	go func() {
		kafka.StartConsumer("my_topic", func(msg []byte) error {
			var newCat entities.Cat
			if err := json.Unmarshal(msg, &newCat); err != nil {
				return err
			}

			if err := catHandler.AddCat(&newCat); err != nil {
				return err
			}

			log.Printf("Cat saved: %s", newCat.Name)
			return nil
		})
	}()

	mux := http.NewServeMux()

	mux.HandleFunc("GET /api/v1/cat/", catHandler.GetAllCats)

	log.Println("Server started at http://localhost:8081")
	http.ListenAndServe(":8081", mux)
}
