package main

import (
	config "Avako/pkg/config"
	proxyserver "Avako/pkg/proxyServer"
	"sync"

	"log"
)

func main() {
	filepath := ""
	cfg, err := config.LoadConfig(filepath)
	if err != nil {
		log.Fatal("config exist error!")
	}

	wg := sync.WaitGroup{}
	wg.Add()
	for {
		go func() {
			server := proxyserver.NewProxyServer(cfg)
			server.StartServer(&wg)
		}()
	}
	wg.Wait()

	//server := proxyserver.NewProxyServer(cfg)
	//server.StartServer()
}
