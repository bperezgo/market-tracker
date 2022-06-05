package bootstrap

import (
	"context"
	"log"

	"markettracker.com/config"
	"markettracker.com/internal/platform/server"
	"markettracker.com/internal/platform/wspool/wsTiingo"
)

func Run() error {
	// Containers of global dependencies
	log.SetFlags(0)
	ctx := context.Background()
	c := config.GetConfiguration()
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
	ws := wsTiingo.New(ctx, tiingoOpts)
	// run in a go rutine because in the subscription, the subscriber is waiting
	// for msgs
	ws.Subscribe(ctx)
	ws.Listen(ctx)
	s := server.New(c.Host, c.Port)
	s.Start(ctx)
	return nil
}
