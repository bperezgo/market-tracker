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
	// TODO: Review this implementation of https://github.com/friendsofgo/kafka-example/blob/master/pkg/kafka/publisher.go
	// and use the defer conn.Close()

	return &EventBus{
		client:   conn,
		topic:    topic,
		clientID: "clientID",
	}, nil
}

func conn(bootstrapBrokerAddr string, topic string) error {
	conn, err := kafka.DialLeader(context.Background(), "tcp", bootstrapBrokerAddr, topic, 0)
	if err != nil {
		return err
	}
	defer conn.Close()
	return nil
}

func (eb *EventBus) Publish(ctx context.Context, events []event.Event) error {
	log.Println("[INFO] Publishing events")
	for _, event := range events {
		log.Println("[INFO] publishing event with id ", event.AggregateId())
		message, err := eb.encondeMessage(event)
		if err != nil {
			return err
		}
		_, err = eb.client.WriteMessages(message)
		if err != nil {
			return err
		}
	}
	return nil
}

func (eb *EventBus) encondeMessage(event event.Event) (kafka.Message, error) {
	m, err := json.Marshal(event.DTO())
	if err != nil {
		return kafka.Message{}, err
	}
	// TODO: Define topic with the type of the event
	return kafka.Message{
		Key:   []byte("string"),
		Value: m,
	}, nil
}
