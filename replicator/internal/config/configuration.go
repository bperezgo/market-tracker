package config

import (
	"log"

	"github.com/tkanos/gonfig"
)

type Configuration struct {
	Host   string
	Events []Event
}

type Event struct {
	Type                string
	BootstrapBrokerAddr string
	Brokers             []string
	Topic               string
	ConsumerGroup       string
}

var configuration *Configuration

func GetConfiguration() *Configuration {
	if configuration != nil {
		return configuration
	}
	config := Configuration{}
	err := gonfig.GetConf("configuration.json", &config)
	if err != nil {
		log.Panic(err)
	}
	return &config
}
