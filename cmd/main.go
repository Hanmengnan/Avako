package main

import (
	config "Avako/pkg/config"
	proxyserver "Avako/pkg/proxyServer"

	"log"
)

func main() {
	filepath := ""
	config, err := config.LoadConfig(filepath)
	if err != nil {
		log.Fatal("config exist error!")
	}
	proxyserver.NewProxyServer(config)
}
