package loadBalancer

import (
	"Avako/pkg/config"
	"log"
	"testing"
)

func TestByRequestBalancer_NewBalancer(t *testing.T) {
	cfg, err := config.LoadConfig("../../config/config.json")
	if err != nil {
		return
	}
	balancer := new(ByRequestBalancer)
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

func TestByRequestBalancer_DoBalance(t *testing.T) {
	cfg, err := config.LoadConfig("../../config/config.json")
	if err != nil {
		return
	}
	balancer := new(ByRequestBalancer)
	servers := make([]*Server, 0)
	for _, item := range cfg.Servers {
		servers = append(servers, &Server{
			Host:   item.Host,
			Port:   item.Port,
			Weight: item.Weight,
		})
	}
	balancer.NewBalancer(servers)
	log.Println(balancer.DoBalance("0", "1"))
}
