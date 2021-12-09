package loadBalancer

import (
	"Avako/pkg/config"
	"log"
	"testing"
)

func TestHashBalance_NewBalancer(t *testing.T) {
	cfg, err := config.LoadConfig("../../config/config.json")
	if err != nil {
		return
	}
	balancer := new(HashBalance)
	servers := make([]*Server, 0)
	for _, item := range cfg.Servers {
		servers = append(servers, &Server{
			Host:   item.Host,
			Port:   item.Port,
			Weight: item.Weight,
		})
	}
	balancer.NewBalancer(servers, 0, 0)
}

func TestHashBalance_DoBalance(t *testing.T) {
	cfg, err := config.LoadConfig("../../config/config.json")
	if err != nil {
		return
	}
	balancer := new(HashBalance)
	servers := make([]*Server, 0)
	for _, item := range cfg.Servers {
		servers = append(servers, &Server{
			Host:   item.Host,
			Port:   item.Port,
			Weight: item.Weight,
		})
	}
	balancer.NewBalancer(servers, 0, 0)
	log.Println(balancer.DoBalance())
}
