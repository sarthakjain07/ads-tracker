package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"

	"ads-tracker/pkg/db"
	"ads-tracker/pkg/handlers"
	"ads-tracker/pkg/kafka"
	"ads-tracker/pkg/routes"
)

func main() {
	_ = godotenv.Load()

	port := os.Getenv("PORT")
	kafkaBroker := os.Getenv("KAFKA_BROKER")
	kafkaTopic := os.Getenv("KAFKA_TOPIC")

	if port == "" {
		port = "8080"
	}
	if kafkaBroker == "" {
		kafkaBroker = "localhost:9092"
	}
	if kafkaTopic == "" {
		kafkaTopic = "ad_clicks"
	}

	producer := kafka.NewClickProducer(kafkaBroker, kafkaTopic)
	handlers.InitClickProducer(producer)

	// init db
	db.InitPostgres()

	// setting upo routers
	r := routes.NewRouter()

	log.Printf("Listening on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
