package loadBalancer

import (
	"errors"
	"math/rand"
)

type RandomBalance struct {
	Servers []*Server
	Index   int64
	Weight  int64
}

func (balancer *RandomBalance) NewBalancer(s []*Server) {
	balancer.Servers = s
}
func (balancer *RandomBalance) DoBalance(...string) (*Server, error) {
	serverNum := len(balancer.Servers)
	if serverNum == 0 {
		return nil, errors.New("no instance found")
	}
	index := rand.Intn(serverNum)
	s := balancer.Servers[index]

	return s, nil
}
