package kafka

import (
	"context"
	"encoding/json"
	"time"

	"ads-tracker/pkg/models"

	"github.com/segmentio/kafka-go"
)

type ClickProducer struct {
	writer *kafka.Writer
	topic  string
}

func NewClickProducer(broker, topic string) *ClickProducer {
	writer := &kafka.Writer{
		Addr:         kafka.TCP(broker),
		Topic:        topic,
		Balancer:     &kafka.LeastBytes{},
		RequiredAcks: kafka.RequireAll,
	}
	return &ClickProducer{writer: writer, topic: topic}
}

func (p *ClickProducer) SendClickEvent(event models.ClickEvent) error {
	msgBytes, err := json.Marshal(event)
	if err != nil {
		return err
	}

	msg := kafka.Message{
		Key:   []byte(event.AdID),
		Value: msgBytes,
		Time:  time.Now(),
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	return p.writer.WriteMessages(ctx, msg)
}
