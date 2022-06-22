package kafka

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/segmentio/kafka-go"
	domain "markettracker.com/replicator"
)

type consumer struct {
	reader *kafka.Reader
}

func NewConsumer(brokers []string, topic string) consumer {
	c := kafka.ReaderConfig{
		Brokers:         brokers,
		Topic:           topic,
		MinBytes:        10e3,            // 10KB
		MaxBytes:        10e6,            // 10MB
		MaxWait:         1 * time.Second, // Maximum amount of time to wait for new data to come when fetching batches of messages from kafka.
		ReadLagInterval: -1,
		GroupID:         "uuid",
		StartOffset:     kafka.LastOffset,
	}
	return consumer{
		reader: kafka.NewReader(c),
	}
}

func (c *consumer) Read(ctx context.Context, chMsg chan domain.Message, chErr chan error) {
	defer c.reader.Close()
	for {

		m, err := c.reader.ReadMessage(ctx)
		if err != nil {
			chErr <- errors.New(fmt.Sprintf("error while reading a message: %v", err))
			continue
		}

		var message domain.Message
		err = json.Unmarshal(m.Value, &message)
		if err != nil {
			chErr <- err
		}

		chMsg <- message
	}
}
