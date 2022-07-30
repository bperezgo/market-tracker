package replicate

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
	"markettracker.com/pkg/command"
)

const ReplicateCommandType command.Type = "command.replicate"

type ReplicateCommand struct {
	ID       string    `json:"id"`
	Date     time.Time `json:"date"`
	Exchange string    `json:"exchange"`
	Price    float32   `json:"price"`
}

func NewReplicateCommand(date time.Time, exchange string, price float32) ReplicateCommand {
	id := uuid.New().String()
	return ReplicateCommand{
		ID:       id,
		Date:     date,
		Exchange: exchange,
		Price:    price,
	}
}

func (ReplicateCommand) Type() command.Type {
	return ReplicateCommandType
}

type ReplicateCommandHandler struct {
	service *ReplicatorStrategy
}

func NewReplicateCommandHandler(service *ReplicatorStrategy) ReplicateCommandHandler {
	return ReplicateCommandHandler{
		service: service,
	}
}

// Handle manage the particular command of the handler
func (ch ReplicateCommandHandler) Handle(ctx context.Context, cmd command.Command) error {
	replicateCmd, ok := cmd.(ReplicateCommand)
	if !ok {
		return errors.New("unexpected command")
	}
	return ch.service.Replicate(
		ctx,
		replicateCmd.ID,
		replicateCmd.Date,
		replicateCmd.Exchange,
		replicateCmd.Price,
	)
}
