package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"

	"ads-tracker/pkg/db"
	"ads-tracker/pkg/kafka"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	_ = godotenv.Load()

	db.InitPostgres()

	go func() {
		http.Handle("/metrics", promhttp.Handler())
		log.Println("Metrics at :2112/metrics")
		log.Fatal(http.ListenAndServe(":2112", nil))
	}()

	kafkaBroker := os.Getenv("KAFKA_BROKER")
	kafkaTopic := os.Getenv("KAFKA_TOPIC")
	kafkaGroup := os.Getenv("KAFKA_GROUP_ID")

	kafka.StartClickConsumer(kafkaBroker, kafkaTopic, kafkaGroup)
}
