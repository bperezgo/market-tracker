package wsMsg

import (
	"log"
	"time"

	domain "markettracker.com/tracker/internal"
)

// Must implements IMsgAdapter
func TiingoAdapter(msg *TiingoMsg) domain.MarketTrackerDTO {
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
		return domain.MarketTrackerDTO{}
	}
	dateTime, err := time.Parse(time.RFC3339, date)
	if err != nil {
		log.Println(err)
		return domain.MarketTrackerDTO{}
	}
	exchange, ok := values[3].(string)
	if !ok {
		return domain.MarketTrackerDTO{}
	}
	lastSize, ok := values[4].(float64)
	if !ok {
		return domain.MarketTrackerDTO{}
	}
	lastPrice, ok := values[5].(float64)
	if !ok {
		return domain.MarketTrackerDTO{}
	}
	marketData := domain.MarketTrackerDTO{
		Ticker:    ticker,
		Date:      dateTime,
		Exchange:  exchange,
		LastSize:  float32(lastSize),
		LastPrice: float32(lastPrice),
	}
	return marketData
}
