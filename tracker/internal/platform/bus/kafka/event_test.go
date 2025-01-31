package kafka

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"markettracker.com/pkg/event"
)

const DummyType = event.Type("events.dummy.tested")

type DummyEvent struct {
	event.BaseEvent
}

func NewDummyEvent() DummyEvent {
	aggregateId := uuid.NewString()
	return DummyEvent{
		BaseEvent: event.NewBaseEvent(aggregateId),
	}
}

func (DummyEvent) Type() event.Type {
	return DummyType
}

func (DummyEvent) DTO() interface{} {
	return "{}"
}

func (DummyEvent) Data() interface{} {
	return ""
}

func (DummyEvent) Meta() map[string]interface{} {
	return nil
}

func Test_Connection_With_Kafka(t *testing.T) {
	err := conn("localhost:9093", "events.dummy.type")
	require.NoError(t, err)
}

func Test_Should_Publish_Event(t *testing.T) {
	config := EventBusConfig{
		Brokers:  []string{"localhost:9093"},
		Topic:    "events.dummy.type",
		ClientID: "someId",
	}
	kafkaPublisher, err := NewEventBus(config)
	require.NoError(t, err, "no connected")
	ctx := context.Background()
	dummyEvent := NewDummyEvent()
	err = kafkaPublisher.Publish(ctx, []event.Event{dummyEvent})
	assert.NoError(t, err, "error was not expected")
}
