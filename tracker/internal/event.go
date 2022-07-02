package domain

import (
	"time"

	"markettracker.com/pkg/event"
)

type Exchange string

// TODO: review if it is better inject this event from the NewAssetRecordedEvent
const AssetRecordedEventType event.Type = "events.asset.recorded"

type AssetRecordedEventDTO struct {
	event.EventDTO
	Data Data `json:"data"`
}

type AssetRecordedEvent struct {
	event.BaseEvent
	Data Data
}
type Data struct {
	Date     time.Time `json:"date"`
	Exchange Exchange  `json:"exchange"`
	Price    float32   `json:"price"`
}

func NewAssetRecordedEvent(id string, date time.Time, exchange string, price float32) AssetRecordedEvent {
	return AssetRecordedEvent{
		BaseEvent: event.NewBaseEvent(id),
		Data: Data{
			Date:     date,
			Exchange: Exchange(exchange),
			Price:    price,
		},
	}
}

func (AssetRecordedEvent) Type() event.Type {
	return AssetRecordedEventType
}

func (ar AssetRecordedEvent) DTO() interface{} {
	return &AssetRecordedEventDTO{
		EventDTO: event.EventDTO{
			EventId:     ar.Id(),
			AggregateId: ar.AggregateId(),
			OccurredOn:  ar.OccurredOn(),
			Type:        string(ar.Type()),
		},
		Data: ar.Data,
	}
}
