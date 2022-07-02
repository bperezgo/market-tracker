package tiingo

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/google/uuid"
	domain "markettracker.com/tracker/internal"
	"markettracker.com/tracker/internal/platform/ws"
	"markettracker.com/tracker/internal/replicate"
	"nhooyr.io/websocket"
)

type WsTiingo struct {
	conn       *websocket.Conn
	wsWrapper  *ws.WsWrapper
	opts       TiingoOptions
	replicator *replicate.Replicator
	adapter    ws.MessageAdapter
}

// Constructor of WsTiingo
func New(ctx context.Context, replicator *replicate.Replicator, opts TiingoOptions) (*WsTiingo, error) {
	dialOps := &websocket.DialOptions{}
	c, _, err := websocket.Dial(ctx, opts.Url, dialOps)
	if err != nil {
		return nil, fmt.Errorf("failed connecting to websocket; %s", err.Error())
	}
	wsWrapper := ws.New(16)
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
			tiingoMsg := &TiingoMsg{}
			marketMsg, err := w.createAssetDTO(message, tiingoMsg)
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

func (w *WsTiingo) createAssetDTO(message []byte, v *TiingoMsg) (domain.AssetDTO, error) {
	if err := json.Unmarshal(message, v); err != nil {
		return domain.AssetDTO{}, err
	}
	marketMsg, err := TiingoAdapter(v)
	if err != nil {
		return domain.AssetDTO{}, err
	}
	aggregateId := uuid.New().String()
	marketMsg.ID = aggregateId
	return marketMsg, nil
}
