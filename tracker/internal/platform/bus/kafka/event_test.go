package kafka

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"markettracker.com/tracker/pkg/event"
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

func Test_Ok_Publish_A_Market_Asset_Message_To_Kafka_Broker(t *testing.T) {
	kafkaPublisher, err := NewEventBus("localhost:9092", "events.dummy.type")
	require.NoError(t, err, "no connected")
	ctx := context.Background()
	dummyEvent := NewDummyEvent()
	err = kafkaPublisher.Publish(ctx, []event.Event{dummyEvent})
	assert.NoError(t, err, "error was not expected")
}
