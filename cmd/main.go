package main

import (
	config "Avako/pkg/config"
	proxyserver "Avako/pkg/proxyServer"
	"log"
	"sync"
)

func main() {
	filepath := "./config/config.json"
	cfg, err := config.LoadConfig(filepath)
	if err != nil {
		log.Fatal("Config file don't exist!")
	}
	start(cfg)
}

func start(cfg *config.Config) {
	var wg sync.WaitGroup
	wg.Add(len(cfg.ProxyServers))
	for i := 0; i < len(cfg.ProxyServers); i++ {
		proxyServerCfg := cfg.ProxyServers[i]
		servers := cfg.Servers
		go func(index int) {
			server := proxyserver.NewProxyServer(&proxyServerCfg, servers)
			server.StartServer(&wg)
		}(i)
	}
	wg.Wait()
}
