package wsWrapper

import (
	"context"

	"nhooyr.io/websocket"
)

type IWsWrapper interface {
	Subscribe(ctx context.Context, conn *websocket.Conn) error
	Unsubscribe(conn *websocket.Conn, event string) error
	Publish(message interface{}) error
	Close() error
}

type Subscriber struct {
	msgs      chan []byte
	closeSlow func()
}

type SubscribeOptions struct {
	Event string
}
