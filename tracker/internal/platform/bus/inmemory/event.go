package inmemory

import (
	"context"
	"log"

	"markettracker.com/tracker/pkg/event"
)

type EventBus struct{}

func NewEventBus() *EventBus {
	return &EventBus{}
}

func (EventBus) Publish(ctx context.Context, events []event.Event) error {
	log.Println("Events")
	log.Println(events)
	return nil
}
