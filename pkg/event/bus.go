package event

import (
	"context"
	"time"

	"github.com/google/uuid"
)

//go:generate mockgen -destination ./pkg/event/mock_bus.go -source ./pkg/event/bus.go -package event

type Bus interface {
	Publish(context.Context, []Event) error
}

// Type represents a domain event type.
type Type string

type EventDTO struct {
	EventId     string
	AggregateId string
	OccurredOn  time.Time
	Type        string
}

type Event interface {
	Id() string
	AggregateId() string
	OccurredOn() time.Time
	Type() Type
	// Method to get the event with the struct necessary to send to he other services
	DTO() interface{}
}

type BaseEvent struct {
	eventId     string
	aggregateId string
	occurredOn  time.Time
}

func NewBaseEvent(aggregateId string) BaseEvent {
	return BaseEvent{
		eventId:     uuid.New().String(),
		aggregateId: aggregateId,
		occurredOn:  time.Now(),
	}
}

func (b BaseEvent) Id() string {
	return b.eventId
}

func (b BaseEvent) OccurredOn() time.Time {
	return b.occurredOn
}

func (b BaseEvent) AggregateId() string {
	return b.aggregateId
}
