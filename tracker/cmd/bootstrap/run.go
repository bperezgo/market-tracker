package bootstrap

import (
	"context"

	"markettracker.com/tracker/configs"
	"markettracker.com/tracker/internal/platform/bus/inmemory"
	"markettracker.com/tracker/internal/platform/server"
)

func Run() error {
	// Containers of global dependencies
	ctx := context.Background()
	c, err := configs.GetConfiguration()
	if err != nil {
		return err
	}
	commandBus := inmemory.NewCommandBus()
	err = EstablishRealTimeConnections(ctx, commandBus)
	if err != nil {
		return err
	}
	s := server.New(c.Host, c.Port, commandBus)
	return s.Start(ctx)
}
