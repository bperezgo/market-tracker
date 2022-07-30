package replicate

import (
	"context"
	"log"
	"time"

	"markettracker.com/pkg/event"
	"markettracker.com/tracker/internal/domain"
)

// This service take the data from the real data source, and is returned a desired structure
// to be sent to the replicators
type Replicator struct {
	assetRepository domain.AssetRepository
	eventBus        event.Bus
}

func New(assetRepository domain.AssetRepository, eventBus event.Bus) *Replicator {
	return &Replicator{
		assetRepository: assetRepository,
		eventBus:        eventBus,
	}
}

func (a *Replicator) Replicate(ctx context.Context, id string, date time.Time, exchange string, price float32) error {
	log.Printf("[INFO] replicating the message with exchange '%s'", exchange)
	asset, err := domain.NewAsset(id, date, exchange, price)
	if err != nil {
		return err
	}
	if err := a.assetRepository.Save(ctx, asset); err != nil {
		return err
	}
	return a.eventBus.Publish(ctx, asset.PullEvents())
}
