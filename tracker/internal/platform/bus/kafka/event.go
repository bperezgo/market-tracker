package kafka

import (
	"context"
	"encoding/json"
	"time"

	"github.com/google/uuid"
	"github.com/segmentio/kafka-go"
	"github.com/segmentio/kafka-go/snappy"
	"markettracker.com/tracker/pkg/event"
)

// TODO: Change the strategy to use the events array with many queues of kafka
type EventBus struct {
	writer   *kafka.Writer
	topic    string
	brokers  []string
	clientID string
}

func NewEventBus(bootstrapBrokerAddr string, brokers []string, topic string) *EventBus {
	clientID := uuid.New().String()
	localAddr := kafka.TCP(bootstrapBrokerAddr)
	dialer := &kafka.Dialer{
		Timeout:   10 * time.Second,
		ClientID:  clientID,
		LocalAddr: localAddr,
	}
	c := kafka.WriterConfig{
		Brokers:          brokers,
		Topic:            topic,
		Balancer:         &kafka.LeastBytes{},
		Dialer:           dialer,
		WriteTimeout:     10 * time.Second,
		ReadTimeout:      10 * time.Second,
		CompressionCodec: snappy.NewCompressionCodec(),
	}

	return &EventBus{
		writer:  kafka.NewWriter(c),
		topic:   topic,
		brokers: brokers,
	}
}

func (eb *EventBus) Publish(ctx context.Context, events []event.Event) error {
	for _, event := range events {
		message, err := eb.encondeMessage(event)
		if err != nil {
			return err
		}
		return eb.writer.WriteMessages(ctx, message)
	}
	return nil
}

func (eb *EventBus) encondeMessage(event event.Event) (kafka.Message, error) {
	m, err := json.Marshal(event)
	if err != nil {
		return kafka.Message{}, err
	}
	return kafka.Message{
		Value: m,
	}, nil
}
