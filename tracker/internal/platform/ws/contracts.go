package ws

import (
	"context"

	"markettracker.com/pkg/command"
	"nhooyr.io/websocket"
)

type WsWrapper interface {
	Subscribe(ctx context.Context, conn *websocket.Conn) error
	Publish(message interface{}) error
}

type MessageAdapter interface {
	Adapt(buf []byte) (command.Command, error)
}
