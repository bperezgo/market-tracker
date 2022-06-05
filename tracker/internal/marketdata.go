package domain

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
