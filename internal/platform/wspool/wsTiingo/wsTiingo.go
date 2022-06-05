package wsTiingo

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"markettracker.com/internal/platform/wspool/wsMsg"
	"markettracker.com/pkg/wsWrapper"
	"nhooyr.io/websocket"
)

// Constructor of WsTiingo
func New(ctx context.Context, opts TiingoOptions) *WsTiingo {
	dialOps := &websocket.DialOptions{}
	c, _, _ := websocket.Dial(ctx, opts.Url, dialOps)
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
	if err = w.conn.Write(ctx, websocket.MessageText, msg); err != nil {
		// TODO: Is it necesary to panic in this part if some websocket failed?
	}
}

func (w *WsTiingo) Listen(ctx context.Context) {
	interrupt := make(chan os.Signal, 1)
	done := make(chan struct{})

	go func() {
		defer close(done)
		for {
			_, message, err := w.conn.Read(ctx)
			if err != nil {
				continue
			}
			tiingoMsg := &wsMsg.TiingoMsg{}
			if err := json.Unmarshal(message, tiingoMsg); err != nil {
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
