package kafka

import (
	"context"
	"crypto/rand"
	"encoding/json"
	"log"
	"time"

	"github.com/oklog/ulid/v2"
	"github.com/segmentio/kafka-go"
	"github.com/segmentio/kafka-go/snappy"
	"markettracker.com/pkg/event"
)

// TODO: Change the strategy to use the events array with many queues of kafka
type EventBus struct {
	config EventBusConfig

	writer *kafka.Writer
}

type EventBusConfig struct {
	Brokers  []string
	Topic    string
	ClientID string
}

func NewEventBus(config EventBusConfig) (*EventBus, error) {
	dialer := &kafka.Dialer{
		Timeout:  10 * time.Second,
		ClientID: config.ClientID,
	}

	c := kafka.WriterConfig{
		Brokers:          config.Brokers,
		Topic:            config.Topic,
		Balancer:         &kafka.LeastBytes{},
		Dialer:           dialer,
		WriteTimeout:     10 * time.Second,
		ReadTimeout:      10 * time.Second,
		CompressionCodec: snappy.NewCompressionCodec(),
	}
	// TODO: Review this implementation of https://github.com/friendsofgo/kafka-example/blob/master/pkg/kafka/publisher.go
	// and use the defer conn.Close()

	return &EventBus{
		config: config,
		writer: kafka.NewWriter(c),
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
	msgs := []kafka.Message{}
	for _, evt := range events {
		log.Println("[INFO] publishing event with id ", evt.AggregateId())
		msg, err := eb.encondeMessage(evt)
		if err != nil {
			log.Println("[INFO] failed sending the event with id ", evt.AggregateId())
			continue
		}
		msgs = append(msgs, msg)
	}
	return eb.writer.WriteMessages(ctx, msgs...)
}

func (eb *EventBus) encondeMessage(event event.Event) (kafka.Message, error) {
	m, err := json.Marshal(event.DTO())
	if err != nil {
		return kafka.Message{}, err
	}
	// TODO: Define topic with the type of the event
	return kafka.Message{
		Key:   []byte(Ulid()),
		Value: m,
	}, nil
}

// Ulid encapsulate the way to generate ulids
func Ulid() string {
	t := time.Now().UTC()
	id := ulid.MustNew(ulid.Timestamp(t), rand.Reader)

	return id.String()
}
