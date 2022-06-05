package adapting

import (
	"context"

	"markettracker.com/tracker/pkg/event"
)

// This service take the data from the real data source, and is returned a desired structure
// to be sent to the replicators
type Replicator struct {
	eventBus event.Bus
}

func New() *Replicator {
	return &Replicator{}
}

func (a *Replicator) Replicate(ctx context.Context) {
	a.eventBus.Publish(ctx, []event.Event{})
}
