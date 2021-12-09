package config

import (
	"log"
	"testing"
)

func TestLoadConfig(t *testing.T) {
	configFile := "../../config/config.json"
	config, err := LoadConfig(configFile)
	if err != nil {
		return
	}
	log.Println(config)
}
