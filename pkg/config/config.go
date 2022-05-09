package config

import (
	"fmt"

	"github.com/tkanos/gonfig"
)

type Configuration struct {
	DB_USERNAME string `json:"DB_USERNAME"`
	DB_PASSWORD string `json:"DB_PASSWORD"`
	DB_PORT     string `json:"DB_PORT"`
	DB_HOST     string `json:"DB_HOST"`
	DB_NAME     string `json:"DB_NAME"`
}

func GetConfig() Configuration {
	conf := Configuration{}
	err := gonfig.GetConf("pkg/config/config.json", &conf)
	fmt.Println(conf, err)
	return conf
}
