// WsWrapper: wrapper of websocket
// The next code was based from the code
// https://github.com/nhooyr/websocket/blob/master/examples/chat/chat.go
package wsWrapper

import (
	"context"
	"fmt"
	"sync"
	"time"

	"golang.org/x/time/rate"

	"nhooyr.io/websocket"
)

// WsWrapper is agnostic for the connection of the websocket,
// for this reason in the methods, the connection is passed like parameter
type WsWrapper struct {
	subscriberMessageBuffer int
	publishLimiter          *rate.Limiter
	subscribersMu           sync.Mutex
	subscribers             map[*Subscriber]struct{}
	nSubscribers            int
}

func New(subscriberMessageBuffer int) *WsWrapper {
	if subscriberMessageBuffer < 0 {
		return nil
	}
	return &WsWrapper{
		subscriberMessageBuffer: subscriberMessageBuffer,
		publishLimiter:          rate.NewLimiter(rate.Every(time.Millisecond*100), 8),
		subscribers:             make(map[*Subscriber]struct{}),
		nSubscribers:            0,
	}
}

func (w *WsWrapper) Subscribe(ctx context.Context, conn *websocket.Conn, options *SubscribeOptions) error {
	ctx = conn.CloseRead(ctx)

	s := &Subscriber{
		msgs: make(chan []byte, w.subscriberMessageBuffer),
		closeSlow: func() {
			conn.Close(websocket.StatusPolicyViolation, "connection too slow to keep up with messages")
		},
	}
	w.addSubscriber(s, conn)
	defer w.deleteSubscriber(s)

	for {
		select {
		case msg := <-s.msgs:
			fmt.Println("msg")
			fmt.Println(msg)
			// if obtain a msg from the connection to the subscriber will be notified
			err := writeTimeout(ctx, time.Second*5, conn, msg)
			if err != nil {
				return err
			}
		case <-ctx.Done():
			return ctx.Err()
		}
	}
}

func (w *WsWrapper) Unsubscribe(event string) error {
	return nil
}

// TODO: This publish method only works for chat server. It is necessary to define it
// for send messages depending on the event of the index of the market
func (w *WsWrapper) Publish(msg []byte) {
	w.subscribersMu.Lock()
	defer w.subscribersMu.Unlock()

	w.publishLimiter.Wait(context.Background())

	for s := range w.subscribers {
		select {
		case s.msgs <- msg:
		default:
			go s.closeSlow()
		}
	}

}

func (w *WsWrapper) Close(conn *websocket.Conn) error {
	err := conn.Close(websocket.StatusNormalClosure, "")
	return err
}

func (w *WsWrapper) deleteSubscriber(s *Subscriber) {
	w.subscribersMu.Lock()
	delete(w.subscribers, s)
	w.nSubscribers -= 1
	w.subscribersMu.Unlock()
}

func writeTimeout(ctx context.Context, timeout time.Duration, c *websocket.Conn, msg []byte) error {
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	return c.Write(ctx, websocket.MessageText, msg)
}

// addSubscriber registers a subscriber.
func (w *WsWrapper) addSubscriber(s *Subscriber, conn *websocket.Conn) {
	w.subscribersMu.Lock()
	w.subscribers[s] = struct{}{}
	w.nSubscribers += 1
	w.subscribersMu.Unlock()
}
