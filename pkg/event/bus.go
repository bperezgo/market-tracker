package event

import (
	"context"
)

//go:generate mockgen -destination ./pkg/event/mock_bus.go -source ./pkg/event/bus.go -package event

type Bus interface {
	Publish(context.Context, []Event) error
}
