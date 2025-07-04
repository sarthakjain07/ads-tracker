package kafka

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"ads-tracker/pkg/models"
	"ads-tracker/pkg/repository"

	"github.com/segmentio/kafka-go"
)

func StartClickConsumer(broker, topic, groupID string) {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:     []string{broker},
		Topic:       topic,
		GroupID:     groupID,
		StartOffset: kafka.LastOffset,
	})

	log.Printf("Kafka consumer started on topic %s", topic)

	for {
		msg, err := r.ReadMessage(context.Background())
		if err != nil {
			log.Printf("[ERROR] Kafka read failed: %v", err)
			continue
		}

		var click models.ClickEvent
		if err := json.Unmarshal(msg.Value, &click); err != nil {
			log.Printf("[ERROR] Invalid message: %v", err)
			continue
		}

		if err := repository.SaveClickEvent(click); err != nil {
			log.Printf("[RETRY] Failed to save click: %v. Retrying...", err)
			go retrySave(click)
		}
	}
}

func retrySave(click models.ClickEvent) {
	for i := 0; i < 5; i++ {
		time.Sleep(time.Duration(i*2) * time.Second)
		if err := repository.SaveClickEvent(click); err == nil {
			log.Println("[RETRY] Save succeeded")
			return
		}
		log.Printf("[RETRY] Attempt %d failed", i+1)
	}
	log.Println("[FAILURE] All retries failed for click:", click)
}
