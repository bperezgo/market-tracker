package tiingo

import (
	"fmt"
	"log"
	"time"

	domain "markettracker.com/tracker/internal"
)

// TiingoMsg interface of the tiingo api in the websocket
type TiingoMsg struct {
	MsgType string `json:"messageType"`
	Service string `json:"service"`
	// is an array with [Msg Type, Ticker, Date, Exchange, LastSize, LastPrice]
	Data [6]interface{} `json:"data"`
}

// Must implements IMsgAdapter
// TODO: The validations here are part of the domain. Refactor to domain
func TiingoAdapter(msg *TiingoMsg) (domain.AssetDTO, error) {
	// [Msg Type, Ticker, Date, Exchange, LastSize, LastPrice]
	var values [6]interface{}
	for idx, el := range msg.Data {
		values[idx] = el
	}
	date, ok := values[2].(string)
	if !ok {
		return domain.AssetDTO{}, fmt.Errorf("date is nil")
	}
	dateTime, err := time.Parse(time.RFC3339, date)
	if err != nil {
		log.Println(err)
		return domain.AssetDTO{}, fmt.Errorf("data doesnot have right format")
	}
	exchange, ok := values[3].(string)
	if !ok {
		return domain.AssetDTO{}, fmt.Errorf("exchange is nil")
	}
	price, ok := values[5].(float64)
	if !ok {
		return domain.AssetDTO{}, fmt.Errorf("lastprice is nil")
	}
	marketData := domain.AssetDTO{
		Date:     dateTime,
		Exchange: exchange,
		Price:    float32(price),
	}
	return marketData, nil
}
