package wsTiingo

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	domain "markettracker.com/tracker/internal"
	"markettracker.com/tracker/internal/platform/wspool/wsMsg"
	"markettracker.com/tracker/internal/replicate"
	"markettracker.com/tracker/pkg/wsWrapper"
	"nhooyr.io/websocket"
)

type WsTiingo struct {
	conn       *websocket.Conn
	wsWrapper  *wsWrapper.WsWrapper
	opts       TiingoOptions
	replicator *replicate.Replicator
}

// Constructor of WsTiingo
func New(ctx context.Context, replicator *replicate.Replicator, opts TiingoOptions) (*WsTiingo, error) {
	dialOps := &websocket.DialOptions{}
	c, _, err := websocket.Dial(ctx, opts.Url, dialOps)
	if err != nil {
		return nil, fmt.Errorf("failed connecting to kafka broker; %s", err.Error())
	}
	wsWrapper := wsWrapper.New(16)
	return &WsTiingo{
		conn:       c,
		wsWrapper:  wsWrapper,
		opts:       opts,
		replicator: replicator,
	}, nil
}

func (w *WsTiingo) Close() error {
	w.wsWrapper.Close(w.conn)
	return nil
}

// Subscribe methos will connect with the respective ws api
func (w *WsTiingo) Subscribe(ctx context.Context) error {
	// Subscription to the api
	msg, err := json.Marshal(w.opts.SubEvent)
	if err = w.conn.Write(ctx, websocket.MessageText, msg); err != nil {
		return err
	}
	return nil
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
			marketMsg, err := createAssetDTO(message)
			if err != nil {
				log.Println("[ERROR] failed adapting the message;", err)
				continue
			}
			w.replicator.Replicate(ctx, marketMsg)
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

func createAssetDTO(message []byte) (domain.AssetDTO, error) {
	tiingoMsg := &wsMsg.TiingoMsg{}
	if err := json.Unmarshal(message, tiingoMsg); err != nil {
		return domain.AssetDTO{}, err
	}
	marketMsg, err := wsMsg.TiingoAdapter(tiingoMsg)
	if err != nil {
		return domain.AssetDTO{}, err
	}
	return marketMsg, nil
}
