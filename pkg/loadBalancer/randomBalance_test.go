package loadBalancer

import (
	"Avako/pkg/config"
	"log"
	"testing"
)

func TestRandomBalance_NewBalancer(t *testing.T) {
	cfg, err := config.LoadConfig("../../config/config.json")
	if err != nil {
		return
	}
	balancer := new(RandomBalance)
	servers := make([]*Server, 0)
	for _, item := range cfg.Servers {
		servers = append(servers, &Server{
			Host:   item.Host,
			Port:   item.Port,
			Weight: item.Weight,
		})
	}
	balancer.NewBalancer(servers)
}

func TestRandomBalance_DoBalance(t *testing.T) {
	cfg, err := config.LoadConfig("../../config/config.json")
	if err != nil {
		return
	}
	balancer := new(RandomBalance)
	servers := make([]*Server, 0)
	for _, item := range cfg.Servers {
		servers = append(servers, &Server{
			Host:   item.Host,
			Port:   item.Port,
			Weight: item.Weight,
		})
	}
	balancer.NewBalancer(servers)
	log.Println(balancer.DoBalance())
}
