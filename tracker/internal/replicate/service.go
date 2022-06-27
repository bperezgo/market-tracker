package replicate

import (
	"context"
	"log"

	"github.com/google/uuid"
	domain "markettracker.com/tracker/internal"
	"markettracker.com/tracker/pkg/event"
)

// This service take the data from the real data source, and is returned a desired structure
// to be sent to the replicators
type Replicator struct {
	eventBus event.Bus
}

func New(eventBus event.Bus) *Replicator {
	return &Replicator{
		eventBus: eventBus,
	}
}

func (a *Replicator) Replicate(ctx context.Context, marketMsg domain.MarketTrackerDTO) error {
	log.Printf("[INFO] replicating the message with ticker '%s'", marketMsg.Ticker)
	aggregateId := uuid.New().String()
	assetRecorded, err := NewAssetRecordedEvent(aggregateId, marketMsg.Date, marketMsg.Exchange, marketMsg.LastPrice)
	if err != nil {
		return err
	}
	return a.eventBus.Publish(ctx, []event.Event{assetRecorded})
}
