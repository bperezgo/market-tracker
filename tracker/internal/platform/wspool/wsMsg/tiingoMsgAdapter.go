package wsMsg

import (
	"log"
	"time"
)

// Must implements IMsgAdapter
func TiingoAdapter(msg *TiingoMsg) *MarketTrackerMsg {
	// [Msg Type, Ticker, Date, Exchange, LastSize, LastPrice]
	var values [6]interface{}
	for idx, el := range msg.Data {
		values[idx] = el
	}
	ticker, ok := values[1].(string)
	if !ok {
		ticker = ""
	}
	date, ok := values[2].(string)
	if !ok {
		return &MarketTrackerMsg{}
	}
	dateTime, err := time.Parse(time.RFC3339, date)
	if err != nil {
		log.Println(err)
		return &MarketTrackerMsg{}
	}
	exchange, ok := values[3].(string)
	if !ok {
		return &MarketTrackerMsg{}
	}
	lastSize, ok := values[4].(float64)
	if !ok {
		return &MarketTrackerMsg{}
	}
	lastPrice, ok := values[5].(float64)
	if !ok {
		return &MarketTrackerMsg{}
	}
	marketData := &MarketTrackerMsg{
		Ticker:    ticker,
		Date:      dateTime,
		Exchange:  exchange,
		LastSize:  float32(lastSize),
		LastPrice: float32(lastPrice),
	}
	return marketData
}
