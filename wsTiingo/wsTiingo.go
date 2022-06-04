package wsTiingo

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"

	"markettracker.com/pkg/errorHandler"
	"markettracker.com/pkg/wsWrapper"
	"markettracker.com/wsMsg"
	"nhooyr.io/websocket"
)

// Constructor of WsTiingo
func NewWsTiingo(ctx context.Context, opts *TiingoOptions) *WsTiingo {
	if opts == nil {
		err := errors.New("ERROR: opts is needed")
		errorHandler.PanicError(err)
	}
	dialOps := &websocket.DialOptions{}
	c, _, err := websocket.Dial(ctx, opts.Url, dialOps)
	errorHandler.PanicError(err)
	wsWrapper := wsWrapper.NewWsWrapper(16)
	return &WsTiingo{
		conn:      c,
		wsWrapper: wsWrapper,
		opts:      opts,
	}
}

func (w *WsTiingo) Close() error {
	w.wsWrapper.Close(w.conn)
	return nil
}

// Subscribe methos will connect with the respective ws api
func (w *WsTiingo) Subscribe(ctx context.Context) {
	// Subscription to the api
	msg, err := json.Marshal(w.opts.SubEvent)
	errorHandler.LogError(err)
	if err = w.conn.Write(ctx, websocket.MessageText, msg); err != nil {
		// TODO: Is it necesary to panic in this part if some websocket failed?
		errorHandler.PanicError(err)
	}
}

func (w *WsTiingo) Listen(ctx context.Context) {
	interrupt := make(chan os.Signal, 1)
	done := make(chan struct{})

	go func() {
		defer close(done)
		for {
			_, message, err := w.conn.Read(ctx)
			errorHandler.LogError(err)
			tiingoMsg := &wsMsg.TiingoMsg{}
			if err := json.Unmarshal(message, tiingoMsg); err != nil {
				errorHandler.LogError(err)
				continue
			}
			// TODO: Handle the error with more logic if failed
			marketMsg := wsMsg.TiingoAdapter(tiingoMsg)
			// Publish to all consumers, that was set in the setup
			w.publish(marketMsg)
		}
	}()

	select {
	case <-interrupt:
		log.Println("interrupt")
	case <-done:
		return
	}
	fmt.Println("message2")
	err := w.conn.Close(websocket.StatusNormalClosure, "")
	if err != nil {
		log.Println("write close:", err)
		return
	}
	<-done
}

func (w *WsTiingo) publish(marketMsg *wsMsg.MarketTrackerMsg) {
	for _, c := range w.opts.Consumers {
		c.Publish(marketMsg)
	}
}
