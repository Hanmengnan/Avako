package config

import (
	"Avako/pkg/loadBalancer"
	"Avako/pkg/proxyServer"
	"encoding/json"
	"log"
	"os"
)

type Config struct {
	ProxyServers []proxyserver.ProxyServer
	Servers      []loadBalancer.Server
}

func LoadConfig(configFile string) (*Config, error) {
	jsonFile, err := os.ReadFile(configFile)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}
	config := new(Config)
	err = json.Unmarshal(jsonFile, config)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}
	return config, nil
}
