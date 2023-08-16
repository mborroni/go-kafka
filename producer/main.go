package main

import (
	"context"
	"github.com/segmentio/kafka-go"
	"log"
	"time"
)

func main() {
	w := &kafka.Writer{
		Addr:         kafka.TCP("localhost:9092"),
		Topic:        "my-topic",
		RequiredAcks: kafka.RequireAll,
		MaxAttempts:  5,
		WriteTimeout: 5 * time.Second,
		Logger:       log.Default(),
		ErrorLogger:  log.Default(),
	}
	err := w.WriteMessages(context.Background(),
		kafka.Message{Value: []byte("Hello Kafka!")},
	)
	if err != nil {
		log.Fatal("Failed to send message", err)
	}

	w.Close()
}
