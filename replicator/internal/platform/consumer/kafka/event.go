package kafka

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/segmentio/kafka-go"
	"markettracker.com/pkg/event"
)

type consumer struct {
	reader *kafka.Reader
}

func NewConsumer(bootstrapBrokerAddr string, topic string, consumerGroup string) (consumer, error) {
	// use this method to know if the broker exists, like a ping, but the connection wont be used
	conn, err := kafka.DialLeader(context.Background(), "tcp", bootstrapBrokerAddr, topic, 0)
	defer conn.Close()
	if err != nil {
		return consumer{}, err
	}
	c := kafka.ReaderConfig{
		Brokers:         []string{bootstrapBrokerAddr},
		Topic:           topic,
		MinBytes:        10e3,            // 10KB
		MaxBytes:        10e6,            // 10MB
		MaxWait:         1 * time.Second, // Maximum amount of time to wait for new data to come when fetching batches of messages from kafka.
		ReadLagInterval: -1,
		GroupID:         consumerGroup,
		StartOffset:     kafka.LastOffset,
	}
	return consumer{
		reader: kafka.NewReader(c),
	}, nil
}

func (c *consumer) Read(ctx context.Context, chMsg chan event.EventDTO, chErr chan error) {
	defer c.reader.Close()
	for {
		m, err := c.reader.ReadMessage(ctx)
		if err != nil {
			chErr <- errors.New(fmt.Sprintf("error while reading a message: %v", err))
			continue
		}

		var message event.EventDTO
		err = json.Unmarshal(m.Value, &message)
		if err != nil {
			chErr <- err
		}

		chMsg <- message
	}
}
