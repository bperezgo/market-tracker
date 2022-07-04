package storage

import (
	"context"
	"log"

	domain "markettracker.com/replicator/internal"
)

type InMemory struct{}

func NewInMemory() *InMemory {
	return &InMemory{}
}

func (s *InMemory) Save(ctx context.Context, asset domain.Asset) error {
	log.Printf("Asset saved id: '%s'; date: '%+s'", asset.ID(), asset.Date())
	return nil
}
