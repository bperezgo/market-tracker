package replicate

import (
	"context"
	"fmt"
	"time"
)

// Used to manage multiple replicators depending on the exchange
// Each replicator has a different event bus with different topic
type ReplicatorStrategy struct {
	replicators map[string]*Replicator
}

func NewReplicatorStrategy() *ReplicatorStrategy {
	replicators := make(map[string]*Replicator)
	return &ReplicatorStrategy{
		replicators: replicators,
	}
}

func (s *ReplicatorStrategy) AppendReplicator(exchange string, replicator *Replicator) {
	s.replicators[exchange] = replicator
}

func (s *ReplicatorStrategy) Replicate(ctx context.Context, id string, date time.Time, exchange string, price float32) error {
	replicator, ok := s.replicators[exchange]
	if !ok {
		return fmt.Errorf("replicator with exchange '%s' is not defined", exchange)
	}
	return replicator.Replicate(ctx, id, date, exchange, price)
}
