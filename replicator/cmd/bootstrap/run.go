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
	"markettracker.com/replicator/internal/platform/consumer/kafka"
)

func Run() error {
	c := config.GetConfiguration()
	chMsg := make(chan domain.Message)
	chErr := make(chan error)
	consumer := kafka.NewConsumer(c.Events[0].Brokers, c.Events[0].Topic)

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

	fmt.Println("\nyou have abandoned the room")
	return nil
}
