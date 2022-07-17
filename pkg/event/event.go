package event

import (
	"time"

	"github.com/google/uuid"
)

// Type represents a domain event type.
// Expected struct is <application>.<bounded-context>.<version>.<type-of-message>.<aggregate>.<what-happened>
type Type string

type EventDataDTO struct {
	EventId    string                 `json:"eventId"`
	OccurredOn string                 `json:"occurredOn"`
	Type       string                 `json:"type"`
	Attributes interface{}            `json:"attributes"`
	Meta       map[string]interface{} `json:"meta"`
}

type EventDTO struct {
	Data EventDataDTO `json:"data"`
}

type Event interface {
	Id() string
	AggregateId() string
	OccurredOn() time.Time
	Type() Type
	// Method to get the event with the struct necessary to send to he other services
	DTO() interface{}
	Data() interface{}
	Meta() map[string]interface{}
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
