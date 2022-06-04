package replicators

import "markettracker.com/wsMsg"

// Replciators must follow this interface to will be passed
// to the arguments in the setup of the instance of the
// market tracker websocket
//
// The main functionality is to pass the information of the websocket
// to services that will use the data for other purposes
type Replicator interface {
	Publish(msg *wsMsg.MarketTrackerMsg)
}
