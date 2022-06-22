package domain

import "context"

type Message struct{}

type Consumer interface {
	Read(ctx context.Context, chMsg chan Message, chErr chan error)
}
