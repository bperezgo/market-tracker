package kafka

import (
	"context"

	"markettracker.com/tracker/pkg/event"
)

type EventBus struct{}

func NewEventBus() *EventBus {
	return &EventBus{}
}

func (eb *EventBus) Publish(ctx context.Context, events []event.Event) error {
	return nil
}
