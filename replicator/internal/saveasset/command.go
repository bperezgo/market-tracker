package saveasset

import (
	"context"

	"markettracker.com/pkg/command"
	domain "markettracker.com/replicator/internal"
)

const SaveAssetCommandType = command.Type("save.asset.command")

type SaveAssetCommand struct {
	ID       string
	Date     string
	Exchange string
	Price    float32
}

// TODO: Command incomplete
func NewSaveAssetCommand(evt domain.AssetRecordedEventDTO) SaveAssetCommand {
	return SaveAssetCommand{
		ID:       evt.Data.AggregateId,
		Date:     evt.Data.Date,
		Exchange: evt.Data.Exchange,
		Price:    evt.Data.Price,
	}
}

func (SaveAssetCommand) Type() command.Type {
	return SaveAssetCommandType
}

type SaveAssetCommandHandler struct {
	service AssetService
}

func NewSaveAssetCommandHandler(service AssetService) SaveAssetCommandHandler {
	return SaveAssetCommandHandler{
		service: service,
	}
}

func (h SaveAssetCommandHandler) Handle(ctx context.Context, cmd command.Command) error {
	saveCmd, ok := cmd.(SaveAssetCommand)
	if !ok {
		return nil
	}
	return h.service.Save(
		ctx,
		saveCmd.ID,
		saveCmd.Date,
		saveCmd.Exchange,
		saveCmd.Price,
	)
}
