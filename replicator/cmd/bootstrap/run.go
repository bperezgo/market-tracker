package bootstrap

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"markettracker.com/pkg/event"
	"markettracker.com/replicator/internal/config"
	"markettracker.com/replicator/internal/platform/consumer/kafka"
)

func Run() error {
	c := config.GetConfiguration()
	chMsg := make(chan event.EventDTO)
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
			log.Printf("message: %+v", m)
		case err := <-chErr:
			log.Println(err)
		}
	}
end:

	fmt.Println("\nconsumer is finished")
	return nil
}
