// ws: wrapper of websocket
// The next code was based from the code
// https://github.com/nhooyr/websocket/blob/master/examples/chat/chat.go
package ws

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"golang.org/x/time/rate"
	"markettracker.com/pkg/command"

	"nhooyr.io/websocket"
)

type Opts struct {
	Url               string
	SubscriptionEvent interface{}
}

// used to generalize the subscription to different websockets
type Ws struct {
	opts Opts
	// TODO: publisherChannel chan<- interface{}
	conn           *websocket.Conn
	publishLimiter *rate.Limiter
	messageAdapter MessageAdapter
	commandBus     command.Bus
}

func New(ctx context.Context, messageAdapter MessageAdapter, commandBus command.Bus, opts Opts) (*Ws, error) {
	dialOps := &websocket.DialOptions{}
	c, _, err := websocket.Dial(ctx, opts.Url, dialOps)
	if err != nil {
		return nil, fmt.Errorf("failed connecting to websocket; %s", err.Error())
	}
	if opts.SubscriptionEvent == nil {
		return nil, fmt.Errorf("Opts.SubscriptionEvent is nil")
	}
	return &Ws{
		messageAdapter: messageAdapter,
		commandBus:     commandBus,
		opts:           opts,
		conn:           c,
		publishLimiter: rate.NewLimiter(rate.Every(time.Millisecond*100), 8),
	}, nil
}

// Subscribe methos will connect with the respective ws api
//
// subEvent is the subscription event needed to connect to the websocket
func (w *Ws) Subscribe(ctx context.Context) error {
	msg, err := json.Marshal(w.opts.SubscriptionEvent)
	if err != nil {
		return err
	}
	if err = w.conn.Write(ctx, websocket.MessageText, msg); err != nil {
		return err
	}
	w.listen(ctx)
	return nil
}

func (w *Ws) listen(ctx context.Context) {
	interrupt := make(chan os.Signal, 1)
	done := make(chan struct{})

	go func() {
		defer close(done)
		for {
			_, message, err := w.conn.Read(ctx)
			if err != nil {
				continue
			}
			marketMsgCmd, err := w.messageAdapter.Adapt(message)
			if err != nil {
				log.Println("[ERROR] failed adapting the message;", err)
				continue
			}
			err = w.commandBus.Dispatch(ctx, marketMsgCmd)
			if err != nil {
				log.Println("[ERROR] failed dispatching the command;", err)
				continue
			}
		}
	}()

	select {
	case <-interrupt:
		log.Println("interrupt")
	case <-done:
		return
	}
	err := w.conn.Close(websocket.StatusNormalClosure, "")
	if err != nil {
		log.Println("write close:", err)
		return
	}
	<-done
}
