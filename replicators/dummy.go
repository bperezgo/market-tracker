package replicators

import (
	"log"

	"markettracker.com/wsMsg"
)

// Dummy follows the Replicator interface of
type Dummy struct{}

func (d *Dummy) Publish(msg *wsMsg.MarketTrackerMsg) {
	log.Printf("Market: %v", msg)
}
