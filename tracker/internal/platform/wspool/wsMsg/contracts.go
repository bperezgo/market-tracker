package wsMsg

import "time"

// MarketTrackerMsg struct is the representation of the output data.
// It will saved in the database with this structure
// i.e. all the implementation of the websocket must result in this struct
//
// Exchange field will be used to find a table where is neede to save the data
// The other fields, will be used to analyze the behavior of the market
type MarketTrackerMsg struct {
	Ticker    string
	Date      time.Time
	Exchange  string
	LastSize  float32
	LastPrice float32
}

// TiingoMsg interface of the tiingo api in the websocket
type TiingoMsg struct {
	MsgType string `json:"messageType"`
	Service string `json:"service"`
	// is an array with [Msg Type, Ticker, Date, Exchange, LastSize, LastPrice]
	Data [6]interface{} `json:"data"`
}
