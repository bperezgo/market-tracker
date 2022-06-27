package replicate

import (
	"time"

	"markettracker.com/tracker/pkg/event"
)

type Exchange string

// TODO: review if it is better inject this event from the NewAssetRecordedEvent
const AssetRecordedEventType event.Type = "events.asset.recorded"

type AssetRecordedEvent struct {
	event.BaseEvent
	Data Data `json:"data"`
}
type Data struct {
	Date     time.Time `json:"date"`
	Exchange Exchange  `json:"exchange"`
	Price    float32   `json:"price"`
}

func NewAssetRecordedEvent(id string, date time.Time, exchange string, price float32) (AssetRecordedEvent, error) {
	return AssetRecordedEvent{
		BaseEvent: event.NewBaseEvent(id),
		Data: Data{
			Date:     date,
			Exchange: Exchange(exchange),
			Price:    price,
		},
	}, nil
}

func (AssetRecordedEvent) Type() event.Type {
	return AssetRecordedEventType
}
