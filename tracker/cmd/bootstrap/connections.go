package bootstrap

import (
	"context"

	"markettracker.com/tracker/configs"
	"markettracker.com/tracker/internal/platform/bus/inmemory"
	"markettracker.com/tracker/internal/platform/bus/kafka"
	"markettracker.com/tracker/internal/platform/factory"
	"markettracker.com/tracker/internal/replicate"
)

func EstablishRealTimeConnections(ctx context.Context) error {
	c, err := configs.GetConfiguration()
	if err != nil {
		return err
	}
	commandBus := inmemory.NewCommandBus()
	// TODO: define strategy of initialization of different kafka channels
	eventBus, err := kafka.NewEventBus(c.Events[0].BootstrapBrokerAddr, c.Events[0].Topic)
	if err != nil {
		return err
	}
	replicator := replicate.New(eventBus)
	replicateCmdHandler := replicate.NewReplicateCommandHandler(replicator)
	commandBus.Register(replicate.ReplicateCommandType, replicateCmdHandler)

	for _, config := range c.RealTimeConnections {
		// TODO: Define strategy to create various factories invokations
		return factory.NewTiingo(ctx, commandBus, config)
	}
	return nil
}
