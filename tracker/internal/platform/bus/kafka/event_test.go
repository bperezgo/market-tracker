package kafka

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"markettracker.com/tracker/pkg/event"
)

const DummyType = event.Type("events.dummy.tested")

type DummyEvent struct {
	event.BaseEvent
}

func NewDummyEvent() DummyEvent {
	return DummyEvent{
		BaseEvent: event.NewBaseEvent(""),
	}
}

func (DummyEvent) Type() event.Type {
	return DummyType
}

func Test_Ok_Publish_A_Market_Asset_Message_To_Kafka_Broker(t *testing.T) {
	kafkaPublisher := NewEventBus("localhost:19092", []string{"1", "2", "3"}, "events.dummy.type")
	ctx := context.Background()
	dummyEvent := NewDummyEvent()
	err := kafkaPublisher.Publish(ctx, []event.Event{dummyEvent})
	assert.NoError(t, err, "error was not expected")
}
