package cmdbus

import (
	"context"

	"markettracker.com/pkg/command"
)

type CommandBus struct {
	handlers map[command.Type]command.Handler
}

func NewCommandBus() *CommandBus {
	handlers := make(map[command.Type]command.Handler)
	return &CommandBus{
		handlers: handlers,
	}
}

func (c *CommandBus) Register(cmdType command.Type, handler command.Handler) {
	c.handlers[cmdType] = handler
}

func (c *CommandBus) Dispatch(ctx context.Context, cmd command.Command) error {
	cmdType := cmd.Type()
	handler, ok := c.handlers[cmdType]
	if !ok {
		return nil
	}
	return handler.Handle(ctx, cmd)
}
