package bootstrap

import (
	"context"

	"markettracker.com/tracker/configs"
	"markettracker.com/tracker/internal/platform/server"
)

func Run() error {
	// Containers of global dependencies
	ctx := context.Background()
	c, err := configs.GetConfiguration()
	if err != nil {
		return err
	}
	EstablishRealTimeConnections(ctx)
	s := server.New(c.Host, c.Port)
	return s.Start(ctx)
}
