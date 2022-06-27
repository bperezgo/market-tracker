package replicate

import (
	"context"
	"log"

	"github.com/google/uuid"
	"markettracker.com/pkg/event"
	domain "markettracker.com/tracker/internal"
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

func (a *Replicator) Replicate(ctx context.Context, marketMsg domain.AssetDTO) error {
	log.Printf("[INFO] replicating the message with exchanger '%s'", marketMsg.Exchange)
	aggregateId := uuid.New().String()
	assetRecorded, err := NewAssetRecordedEvent(aggregateId, marketMsg.Date, marketMsg.Exchange, marketMsg.Price)
	if err != nil {
		return err
	}
	return a.eventBus.Publish(ctx, []event.Event{assetRecorded})
}
