package bootstrap

import (
	"context"
	"log"

	"markettracker.com/tracker/config"
	"markettracker.com/tracker/internal/platform/bus/kafka"
	"markettracker.com/tracker/internal/platform/server"
	"markettracker.com/tracker/internal/platform/wspool/wsTiingo"
	"markettracker.com/tracker/internal/replicate"
)

func Run() error {
	// Containers of global dependencies
	log.SetFlags(0)
	ctx := context.Background()
	c := config.GetConfiguration()
	// TODO: define strategy of initialization of different kafka channels
	eventBus, err := kafka.NewEventBus(c.Events[0].BootstrapBrokerAddr, c.Events[0].Brokers, c.Events[0].Topic)
	if err != nil {
		return err
	}
	replicator := replicate.New(eventBus)
	tiingoOpts := wsTiingo.TiingoOptions{
		Url: c.TiingoApiUrl,
		SubEvent: &wsTiingo.SubTiingoOpts{
			EventName:     "subscribe",
			Authorization: c.TiingoApiToken,
			EventData: &wsTiingo.EventDataTiingo{
				ThresholdLevel: 5,
			},
		},
	}
	ws, err := wsTiingo.New(ctx, replicator, tiingoOpts)
	if err != nil {
		return err
	}
	// run in a go rutine because in the subscription, the subscriber is waiting
	// for msgs
	ws.Subscribe(ctx)
	ws.Listen(ctx)
	s := server.New(c.Host, c.Port)
	return s.Start(ctx)
}
