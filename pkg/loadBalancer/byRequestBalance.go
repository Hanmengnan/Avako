package loadBalancer

import (
	"errors"
	"log"
	"strconv"
)

type ByRequestBalancer struct {
	Servers []*Server
	Index   *int64
	Weight  *int64
}

func NewByRequestBalancer(s []*Server, i *int64, w *int64) *ByRequestBalancer {
	return &ByRequestBalancer{
		Servers: s,
		Index:   i,
		Weight:  w,
	}
}
func (balancer ByRequestBalancer) DoBalance(key ...string) (*Server, error) {
	serverNum := len(balancer.Servers)
	log.Println("keys:" + key[0] + key[1])
	id, err := strconv.Atoi(key[1])
	if err != nil {
		log.Println("wrong with extacting server id")
	}
	if serverNum == 0 {
		return nil, errors.New("no instance found")
	}
	s := balancer.Servers[id%int(serverNum)]

	return s, nil
}
