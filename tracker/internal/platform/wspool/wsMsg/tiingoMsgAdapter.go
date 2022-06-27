package wsMsg

import (
	"fmt"
	"log"
	"time"

	domain "markettracker.com/tracker/internal"
)

// Must implements IMsgAdapter
// TODO: The validations here are part of the domain. Refactor to domain
func TiingoAdapter(msg *TiingoMsg) (domain.MarketTrackerDTO, error) {
	// [Msg Type, Ticker, Date, Exchange, LastSize, LastPrice]
	var values [6]interface{}
	for idx, el := range msg.Data {
		values[idx] = el
	}
	ticker, ok := values[1].(string)
	if !ok {
		return domain.MarketTrackerDTO{}, fmt.Errorf("ticker is nil")
	}
	date, ok := values[2].(string)
	if !ok {
		return domain.MarketTrackerDTO{}, fmt.Errorf("date is nil")
	}
	dateTime, err := time.Parse(time.RFC3339, date)
	if err != nil {
		log.Println(err)
		return domain.MarketTrackerDTO{}, fmt.Errorf("data doesnot have right format")
	}
	exchange, ok := values[3].(string)
	if !ok {
		return domain.MarketTrackerDTO{}, fmt.Errorf("exchange is nil")
	}
	lastSize, ok := values[4].(float64)
	if !ok {
		return domain.MarketTrackerDTO{}, fmt.Errorf("lastsize is nil")
	}
	lastPrice, ok := values[5].(float64)
	if !ok {
		return domain.MarketTrackerDTO{}, fmt.Errorf("lastprice is nil")
	}
	marketData := domain.MarketTrackerDTO{
		Ticker:    ticker,
		Date:      dateTime,
		Exchange:  exchange,
		LastSize:  float32(lastSize),
		LastPrice: float32(lastPrice),
	}
	return marketData, nil
}
