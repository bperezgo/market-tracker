package replicate

import (
	"context"
	"log"

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

func (a *Replicator) Replicate(ctx context.Context, marketMsg domain.MarketTrackerDTO) {
	log.Printf("New Market Message")
	log.Printf("%+v", marketMsg)
	a.eventBus.Publish(ctx, []event.Event{})
}
