package config

import (
	"log"

	"github.com/tkanos/gonfig"
)

type Configuration struct {
	Port int32
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
