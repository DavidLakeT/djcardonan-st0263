package main

import (
	"github.com/tkanos/gonfig"
)

type Configuration struct {
	IP   string
	PORT string
	DIR  string
}

func GetConfig() Configuration {
	configuration := Configuration{}
	fileName := "config.json"
	gonfig.GetConf(fileName, &configuration)
	return configuration
}
