package bootstrap

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	domain "markettracker.com/replicator/internal"
	"markettracker.com/replicator/internal/config"
	"markettracker.com/replicator/internal/platform/cmdbus"
	"markettracker.com/replicator/internal/platform/consumer/kafka"
	"markettracker.com/replicator/internal/platform/storage"
	"markettracker.com/replicator/internal/saveasset"
)

func Run() error {
	log.Println("[INFO] Running the replicator instance")
	c := config.GetConfiguration()
	ctx := context.Background()
	cmdBus := cmdbus.NewCommandBus()

	inmemory := storage.NewInMemory()
	assetService := saveasset.NewAssetService(inmemory)
	assetServiceHandler := saveasset.NewSaveAssetCommandHandler(assetService)

	cmdBus.Register(saveasset.SaveAssetCommandType, assetServiceHandler)
	chMsg := make(chan domain.AssetRecordedEventDTO)
	chErr := make(chan error)
	consumer, err := kafka.NewConsumer(c.Events[0].BootstrapBrokerAddr, c.Events[0].Topic, c.Events[0].ConsumerGroup)
	if err != nil {
		return err
	}

	go func() {
		consumer.Read(context.Background(), chMsg, chErr)
	}()
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	for {
		select {
		case <-quit:
			goto end
		case m := <-chMsg:
			cmdBus.Dispatch(ctx, saveasset.NewSaveAssetCommand(m))
		case err := <-chErr:
			log.Println(err)
		}
	}
end:

	fmt.Println("\nconsumer is finished")
	return nil
}
