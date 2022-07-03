package ws

import (
	"markettracker.com/pkg/command"
)

type MessageAdapter interface {
	Adapt(buf []byte) (command.Command, error)
}
