package wsMsg

type IMsgAdapter func(msg *interface{}) MarketTrackerMsg
