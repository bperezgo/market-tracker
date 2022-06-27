package kafka

import (
	"context"
	"encoding/json"
	"log"

	"github.com/segmentio/kafka-go"
	"markettracker.com/pkg/event"
)

// TODO: Change the strategy to use the events array with many queues of kafka
type EventBus struct {
	client   *kafka.Conn
	topic    string
	clientID string
}

func NewEventBus(bootstrapBrokerAddr string, topic string) (*EventBus, error) {
	conn, err := kafka.DialLeader(context.Background(), "tcp", bootstrapBrokerAddr, topic, 0)
	if err != nil {
		return nil, err
	}

	return &EventBus{
		client: conn,
		topic:  topic,
	}, nil
}

func (eb *EventBus) Publish(ctx context.Context, events []event.Event) error {
	for _, event := range events {
		log.Println("[INFO] publishing event with id ", event.AggregateId())
		message, err := eb.encondeMessage(event)
		if err != nil {
			return err
		}
		_, err = eb.client.WriteMessages(message)
		return err
	}
	return nil
}

func (eb *EventBus) encondeMessage(event event.Event) (kafka.Message, error) {
	m, err := json.Marshal(event.DTO())
	if err != nil {
		return kafka.Message{}, err
	}
	return kafka.Message{
		Key:   []byte("string"),
		Value: m,
	}, nil
}
