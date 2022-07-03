package factory

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"markettracker.com/pkg/command"
	"markettracker.com/tracker/internal/replicate"
)

type TiingoAdapter struct{}

func NewTiingoAdapter() *TiingoAdapter {
	return &TiingoAdapter{}
}

// TiingoMsg interface of the tiingo api in the websocket
type TiingoMsg struct {
	MsgType string `json:"messageType"`
	Service string `json:"service"`
	// is an array with [Msg Type, Ticker, Date, Exchange, LastSize, LastPrice]
	Data [6]interface{} `json:"data"`
}

// Must implements IMsgAdapter
// TODO: The validations here are part of the domain. Refactor to domain
func (TiingoAdapter) Adapt(buf []byte) (command.Command, error) {
	msg := &TiingoMsg{}
	if err := json.Unmarshal(buf, msg); err != nil {
		return replicate.ReplicateCommand{}, err
	}
	// [Msg Type, Ticker, Date, Exchange, LastSize, LastPrice]
	var values [6]interface{}
	for idx, el := range msg.Data {
		values[idx] = el
	}
	date, ok := values[2].(string)
	if !ok {
		return replicate.ReplicateCommand{}, fmt.Errorf("date is nil")
	}
	dateTime, err := time.Parse(time.RFC3339, date)
	if err != nil {
		log.Println(err)
		return replicate.ReplicateCommand{}, fmt.Errorf("data doesnot have right format")
	}
	exchange, ok := values[3].(string)
	if !ok {
		return replicate.ReplicateCommand{}, fmt.Errorf("exchange is nil")
	}
	price, ok := values[5].(float64)
	if !ok {
		return replicate.ReplicateCommand{}, fmt.Errorf("lastprice is nil")
	}
	marketData := replicate.NewReplicateCommand(
		dateTime,
		exchange,
		float32(price),
	)
	return marketData, nil
}
