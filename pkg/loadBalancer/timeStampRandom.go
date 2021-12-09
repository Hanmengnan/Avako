package loadBalancer

import (
	"errors"
	"time"
)

type TimeStampRandomBalancer struct {
	Servers []*Server
	Index   int64
	Weight  int64
}

func (balancer *TimeStampRandomBalancer) NewBalancer(s []*Server, i int64, w int64) {
	balancer.Servers = s
}
func (balancer *TimeStampRandomBalancer) DoBalance(key ...string) (*Server, error) {
	serverNum := len(balancer.Servers)
	if serverNum == 0 {
		return nil, errors.New("no instance found")
	}
	timeStamp := time.Now().Unix()
	s := balancer.Servers[timeStamp%int64(serverNum)]

	return s, nil
}
