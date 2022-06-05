package domain

type AdapterAssistance interface {
	Adapt(inputData interface{}) (outputData MarketTrackerDTO)
}
