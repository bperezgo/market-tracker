package wsMsg

import domain "markettracker.com/tracker/internal"

type IMsgAdapter func(msg *interface{}) domain.MarketTrackerMsg
