package bootstrap

import (
	"context"

	"markettracker.com/pkg/command"
	"markettracker.com/tracker/configs"
	"markettracker.com/tracker/internal/platform/bus/kafka"
	"markettracker.com/tracker/internal/platform/factory"
	"markettracker.com/tracker/internal/replicate"
)

func EstablishRealTimeConnections(ctx context.Context, commandBus command.Bus) error {
	c, err := configs.GetConfiguration()
	if err != nil {
		return err
	}
	// TODO: define strategy of initialization of different kafka channels
	eventBusConfig := kafka.EventBusConfig{
		Brokers:  c.Events[0].Brokers,
		Topic:    c.Events[0].Topic,
		ClientID: "clientID",
	}
	eventBus, err := kafka.NewEventBus(eventBusConfig)
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
