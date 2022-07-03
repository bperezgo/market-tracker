package bootstrap

import (
	"context"
	"log"

	"markettracker.com/tracker/config"
	"markettracker.com/tracker/internal/platform/adapter"
	"markettracker.com/tracker/internal/platform/bus/inmemory"
	"markettracker.com/tracker/internal/platform/bus/kafka"
	"markettracker.com/tracker/internal/platform/server"
	"markettracker.com/tracker/internal/platform/ws"
	"markettracker.com/tracker/internal/platform/ws/tiingo"
	"markettracker.com/tracker/internal/replicate"
)

func Run() error {
	// Containers of global dependencies
	log.SetFlags(0)
	ctx := context.Background()
	c := config.GetConfiguration()
	commandBus := inmemory.NewCommandBus()
	// TODO: define strategy of initialization of different kafka channels
	eventBus, err := kafka.NewEventBus(c.Events[0].BootstrapBrokerAddr, c.Events[0].Topic)
	if err != nil {
		return err
	}
	replicator := replicate.New(eventBus)
	tiingoOpts := tiingo.Options{
		Url: c.TiingoApiUrl,
		SubEvent: &tiingo.SubOpts{
			EventName:     "subscribe",
			Authorization: c.TiingoApiToken,
			EventData: &tiingo.EventData{
				ThresholdLevel: 5,
			},
		},
	}
	replicateCmdHandler := replicate.NewReplicateCommandHandler(replicator)
	tiingoAdapter := adapter.NewTiingo()
	commandBus.Register(replicate.ReplicateCommandType, replicateCmdHandler)
	wsWrapper, err := ws.New(ctx, tiingoAdapter, commandBus, ws.Opts{
		Url: tiingoOpts.Url,
	})
	if err != nil {
		return err
	}
	err = wsWrapper.Subscribe(ctx, tiingoOpts.SubEvent)
	if err != nil {
		return err
	}
	s := server.New(c.Host, c.Port)
	return s.Start(ctx)
}
