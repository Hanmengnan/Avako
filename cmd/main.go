package main

import (
	config "Avako/pkg/config"
	proxyserver "Avako/pkg/proxyServer"

	"log"
)

func main() {
	filepath := ""
	cfg, err := config.LoadConfig(filepath)
	if err != nil {
		log.Fatal("config exist error!")
	}
	server := proxyserver.NewProxyServer(cfg)
	server.StartServer()
}
