package wsMsg

// TiingoMsg interface of the tiingo api in the websocket
type TiingoMsg struct {
	MsgType string `json:"messageType"`
	Service string `json:"service"`
	// is an array with [Msg Type, Ticker, Date, Exchange, LastSize, LastPrice]
	Data [6]interface{} `json:"data"`
}
