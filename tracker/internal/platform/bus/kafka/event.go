package kafka

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/segmentio/kafka-go"
	"github.com/segmentio/kafka-go/snappy"
	"markettracker.com/tracker/pkg/event"
)

// TODO: Change the strategy to use the events array with many queues of kafka
type EventBus struct {
	writer   *kafka.Writer
	clientID string
}

func NewEventBus(brokers []string, topic string) *EventBus {
	clientID := uuid.New().String()
	dialer := &kafka.Dialer{
		Timeout:  10 * time.Second,
		ClientID: clientID,
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
		writer: kafka.NewWriter(c),
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

func (EventBus) encondeMessage(event event.Event) (kafka.Message, error) {
	return kafka.Message{}, nil
}
