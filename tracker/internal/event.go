package domain

import (
	"time"

	"markettracker.com/pkg/event"
)

type Exchange string

// TODO: define the different types allowed from configuration,
// TODO: Allow that the event define the type in runtime,
// because must exist event for exhange
const AssetRecordedEventType event.Type = "events.asset.recorded"

type AssetRecordedEventDTO struct {
	event.EventDTO
	Data Data                   `json:"data"`
	Meta map[string]interface{} `json:"meta"`
}

type AssetRecordedEvent struct {
	event.BaseEvent
	data Data
	meta map[string]interface{}
}

type Data struct {
	Date     time.Time `json:"date"`
	Exchange Exchange  `json:"exchange"`
	Price    float32   `json:"price"`
}

func NewAssetRecordedEvent(id string, date time.Time, exchange string, price float32) AssetRecordedEvent {
	meta := make(map[string]interface{})
	return AssetRecordedEvent{
		BaseEvent: event.NewBaseEvent(id),
		data: Data{
			Date:     date,
			Exchange: Exchange(exchange),
			Price:    price,
		},
		meta: meta,
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
			OccurredOn:  ar.OccurredOn().String(),
			Type:        string(ar.Type()),
		},
		Data: ar.data,
		Meta: ar.meta,
	}
}

func (ar AssetRecordedEvent) Meta() map[string]interface{} {
	return ar.meta
}

func (ar AssetRecordedEvent) Data() interface{} {
	return ar.data
}
