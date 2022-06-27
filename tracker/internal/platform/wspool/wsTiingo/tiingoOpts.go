package wsTiingo

type EventDataTiingo struct {
	ThresholdLevel int `json:"thresholdLevel"`
}

type SubTiingoOpts struct {
	EventName     string           `json:"eventName"`
	Authorization string           `json:"authorization"`
	EventData     *EventDataTiingo `json:"eventData"`
}

type TiingoOptions struct {
	Url      string
	SubEvent *SubTiingoOpts
}
