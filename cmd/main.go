package main

import (
	config "Avako/pkg/config"
	proxyserver "Avako/pkg/proxyServer"
	"sync"

	"log"
)

func main() {
	filepath := "./config/config.json"
	cfg, err := config.LoadConfig(filepath)
	if err != nil {
		log.Fatal("Config file don't exist!")
	}
	var wg sync.WaitGroup
	wg.Add(len(cfg.ProxyServers))
	for i := 0; i < len(cfg.ProxyServers); i++ {
		go func(index int) {
			server := proxyserver.NewProxyServer(cfg, index)
			server.StartServer(&wg)
		}(i)
	}
	wg.Wait()
}
