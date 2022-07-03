package tiingo

type EventData struct {
	ThresholdLevel int `json:"thresholdLevel"`
}

type SubOpts struct {
	EventName     string     `json:"eventName"`
	Authorization string     `json:"authorization"`
	EventData     *EventData `json:"eventData"`
}

type Options struct {
	Url      string
	SubEvent *SubOpts
}
