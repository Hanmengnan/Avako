package loadBalancer

import (
	"errors"
	"math/rand"
)

type RandomBalance struct {
	Servers []*Server
	Index   *int64
	Weight  *int64
}

func NewRandomBalance(s []*Server, i *int64, w *int64) *RandomBalance {
	return &RandomBalance{
		Servers: s,
		Index:   i,
		Weight:  w,
	}
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
