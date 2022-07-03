package configs

import (
	"github.com/tkanos/gonfig"
)

type Configuration struct {
	Port                int32
	Host                string
	TiingoApiToken      string
	TiingoApiUrl        string
	Events              []Event
	RealTimeConnections map[string]map[string]interface{}
}

type Event struct {
	Type                string
	BootstrapBrokerAddr string
	Brokers             []string
	Topic               string
}

var configuration *Configuration

func GetConfiguration() (*Configuration, error) {
	if configuration != nil {
		return configuration, nil
	}
	config := Configuration{}
	err := gonfig.GetConf("configuration.json", &config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}
