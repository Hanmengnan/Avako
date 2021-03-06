package loadBalancer

import (
	"errors"
	"log"
	"strconv"
)

type ByRequestBalancer struct {
	Servers []*Server
	Index   int64
	Weight  int64
}

func (balancer *ByRequestBalancer) NewBalancer(s []*Server) {
	balancer.Servers = s
}
func (balancer *ByRequestBalancer) DoBalance(key ...string) (*Server, error) {
	serverNum := len(balancer.Servers)
	log.Println("keys:" + key[0] + key[1])
	id, err := strconv.Atoi(key[1])
	log.Println("choose sever", id)
	if err != nil {
		log.Println(err)
	}
	if serverNum == 0 {
		return nil, errors.New("no instance found")
	}
	s := balancer.Servers[id%serverNum]
	return s, nil
}
