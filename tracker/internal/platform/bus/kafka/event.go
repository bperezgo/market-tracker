package kafka

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/segmentio/kafka-go"
	"github.com/segmentio/kafka-go/snappy"
	"markettracker.com/tracker/pkg/event"
)

// TODO: Change the strategy to use the events array with many queues of kafka
type EventBus struct {
	client   *kafka.Conn
	writer   *kafka.Writer
	topic    string
	brokers  []string
	clientID string
}

func NewEventBus(bootstrapBrokerAddr string, brokers []string, topic string) (*EventBus, error) {
	clientID := uuid.New().String()
	localAddr := kafka.TCP(bootstrapBrokerAddr)
	conn, err := kafka.DialLeader(context.Background(), "tcp", bootstrapBrokerAddr, topic, 0)
	if err != nil {
		return nil, err
	}
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
		client:  conn,
		writer:  kafka.NewWriter(c),
		topic:   topic,
		brokers: brokers,
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
	m, err := json.Marshal(event)
	if err != nil {
		return kafka.Message{}, err
	}
	return kafka.Message{
		Key:   []byte("string"),
		Value: m,
	}, nil
}
