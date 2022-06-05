package replicate

import (
	"time"

	"markettracker.com/tracker/pkg/event"
)

type Exchange string

const AssetRecordedEventType event.Type = "events.asset.recorded"

type AssetRecordedEvent struct {
	event.BaseEvent
	Data Data
}
type Data struct {
	Date     time.Time
	Exchange Exchange
	Price    float32
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
