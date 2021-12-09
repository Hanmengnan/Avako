package config

import (
	"encoding/json"
	"log"
	"os"
)

type Proxyserver struct {
	Ip        string
	Port      string
	Algorithm string
}
type Server struct {
	Ip   string
	Port string
}

type Config struct {
	Nginx   []Proxyserver
	Servers []Server
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
