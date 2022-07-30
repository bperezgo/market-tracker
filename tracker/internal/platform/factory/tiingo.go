package factory

import (
	"context"
	"fmt"

	"markettracker.com/pkg/command"
	"markettracker.com/tracker/configs"
	"markettracker.com/tracker/internal/platform/ws"
)

type TiingoEventData struct {
	ThresholdLevel int `json:"thresholdLevel"`
}

type TiingoSubOpts struct {
	EventName     string          `json:"eventName"`
	Authorization string          `json:"authorization"`
	EventData     TiingoEventData `json:"eventData"`
}

type TiingoOptions struct {
	Url string `json:"TiingoApiUrl"`
	// SubscriptionEvent TiingoSubOpts `json:"SubscriptionEvent"`
	SubscriptionEvent interface{} `json:"SubscriptionEvent"`
}

func NewTiingo(ctx context.Context, commandBus command.Bus, config configs.RealTimeConnection) error {
	url, ok := config.Data["TiingoApiUrl"].(string)
	if !ok {
		return fmt.Errorf("Tiingo Url is not defined")
	}
	subEvent, ok := config.Data["SubscriptionEvent"]
	if !ok {
		return fmt.Errorf("Tiingo SubscriptionEvent is not defined")
	}
	opts := TiingoOptions{
		Url:               url,
		SubscriptionEvent: subEvent,
	}
	tiingoAdapter := NewTiingoAdapter()
	wsWrapper, err := ws.New(ctx, tiingoAdapter, commandBus, ws.Opts{
		Url:               opts.Url,
		SubscriptionEvent: opts.SubscriptionEvent,
	})
	if err != nil {
		return err
	}
	return wsWrapper.Subscribe(ctx)
}
