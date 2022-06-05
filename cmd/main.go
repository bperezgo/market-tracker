package main

import (
	"context"
	"log"

	"markettracker.com/config"
	"markettracker.com/internal/platform/server"
	"markettracker.com/replicators"
	"markettracker.com/wsTiingo"
)

func init() {
	// Containers of global dependencies
	log.SetFlags(0)
	ctx := context.Background()
	c := config.GetConfiguration()
	dummyReplicator := &replicators.Dummy{}
	tiingoOpts := &wsTiingo.TiingoOptions{
		Url: c.TiingoApiUrl,
		SubEvent: &wsTiingo.SubTiingoOpts{
			EventName:     "subscribe",
			Authorization: c.TiingoApiToken,
			EventData: &wsTiingo.EventDataTiingo{
				ThresholdLevel: 5,
			},
		},
		Consumers: []replicators.Replicator{
			dummyReplicator,
		},
	}
	ws := wsTiingo.NewWsTiingo(ctx, tiingoOpts)
	// run in a go rutine because in the subscription, the subscriber is waiting
	// for msgs
	ws.Subscribe(ctx)
	ws.Listen(ctx)
}

func main() {
	c := config.GetConfiguration()
	s := server.InitServer(c.Port)
	s.Start()
}
