package bootstrap

import (
	"context"

	"markettracker.com/pkg/command"
	"markettracker.com/tracker/configs"
	"markettracker.com/tracker/internal/domain"
	"markettracker.com/tracker/internal/platform/bus/kafka"
	"markettracker.com/tracker/internal/platform/factory"
	"markettracker.com/tracker/internal/platform/repository"
	"markettracker.com/tracker/internal/replicate"
)

func EstablishRealTimeConnections(ctx context.Context, commandBus command.Bus) error {
	c, err := configs.GetConfiguration()
	if err != nil {
		return err
	}
	allowAutoTopicCreation := true
	replicatorStrategy := replicate.NewReplicatorStrategy()
	for _, config := range c.Events {
		eventBusConfig := kafka.EventBusConfig{
			Brokers:                config.Brokers,
			Topic:                  config.Type,
			ClientID:               config.ClientID,
			AllowAutoTopicCreation: allowAutoTopicCreation,
		}
		eventBus, err := kafka.NewEventBus(eventBusConfig)
		if err != nil {
			return err
		}
		repoConfig := repository.PostgresqlConfig{
			Host:     config.Repository.Host,
			Port:     config.Repository.Port,
			User:     config.Repository.User,
			Password: config.Repository.Password,
			Dbname:   config.Repository.Dbname,
		}
		assetRepo, err := repository.NewPostgresql(config.Repository.Table, repoConfig)
		if err != nil {
			return err
		}
		replicator := replicate.New(assetRepo, eventBus)
		replicatorStrategy.AppendReplicator(config.Exchange, replicator)
	}

	replicateCmdHandler := replicate.NewReplicateCommandHandler(replicatorStrategy)
	commandBus.Register(replicate.ReplicateCommandType, replicateCmdHandler)

	for _, config := range c.RealTimeConnections {
		// TODO: Define strategy to create various factories invokations
		events := config.Events
		for _, evt := range events {
			eventBusConfig := kafka.EventBusConfig{
				Brokers:                evt.Brokers,
				Topic:                  evt.Type,
				ClientID:               evt.ClientID,
				AllowAutoTopicCreation: allowAutoTopicCreation,
			}
			eventBus, err := kafka.NewEventBus(eventBusConfig)
			if err != nil {
				return err
			}
			// TODO: Create a different instance of this repo
			repoMock := domain.AssetRepositoryMock{}
			replicator := replicate.New(&repoMock, eventBus)
			replicatorStrategy.AppendReplicator(evt.Exchange, replicator)
		}
		err := factory.NewTiingo(ctx, commandBus, config)
		if err != nil {
			return err
		}
	}
	return nil
}
