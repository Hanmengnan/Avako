package loadBalancer

import (
	"errors"
	"math/rand"
)

type RandomBalance struct {
	Servers []*Server
}

func (balancer RandomBalance) DoBalance(key ...string) (*Server, error) {
	serverNum := len(balancer.Servers)
	if serverNum == 0 {
		return nil, errors.New("no instance found")
	}
	index := rand.Intn(serverNum)
	s := balancer.Servers[index]

	return s, nil
}
