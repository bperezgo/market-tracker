package configs

import (
	"github.com/tkanos/gonfig"
)

type Configuration struct {
	Port                int32
	Host                string
	Events              []Event
	RealTimeConnections []RealTimeConnection
}

type RealTimeConnection struct {
	Type   string
	Data   map[string]interface{}
	Events []Event
}

type Event struct {
	Type                string
	BootstrapBrokerAddr string
	Brokers             []string
	ClientID            string
	Exchange            string
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
