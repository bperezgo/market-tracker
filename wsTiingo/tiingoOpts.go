package wsTiingo

import (
	"markettracker.com/pkg/wsWrapper"
	"markettracker.com/replicators"
	"nhooyr.io/websocket"
)

// TODO: dependency injection strategy
type WsTiingo struct {
	conn      *websocket.Conn
	wsWrapper *wsWrapper.WsWrapper
	opts      *TiingoOptions
}

type EventDataTiingo struct {
	ThresholdLevel int `json:"thresholdLevel"`
}

type SubTiingoOpts struct {
	EventName     string           `json:"eventName"`
	Authorization string           `json:"authorization"`
	EventData     *EventDataTiingo `json:"eventData"`
}

type TiingoOptions struct {
	Url       string
	SubEvent  *SubTiingoOpts
	Consumers []replicators.Replicator
}
