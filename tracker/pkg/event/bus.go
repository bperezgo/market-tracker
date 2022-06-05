package event

import "context"

type Bus interface {
	Publish(context.Context, []Event) error
}

type Event interface{}
