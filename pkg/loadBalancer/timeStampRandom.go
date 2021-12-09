package loadBalancer

import (
	"errors"
	"time"
)

type TimeStampRandomBalancer struct {
	Servers []*Server
	Index   *int64
	Weight  *int64
}

func NewTimeStampRandomBalancer(s []*Server, i *int64, w *int64) *TimeStampRandomBalancer {
	return &TimeStampRandomBalancer{
		Servers: s,
		Index:   i,
		Weight:  w,
	}
}
func (balancer TimeStampRandomBalancer) DoBalance(key ...string) (*Server, error) {
	serverNum := len(balancer.Servers)
	if serverNum == 0 {
		return nil, errors.New("no instance found")
	}
	timeStamp := time.Now().Unix()
	s := balancer.Servers[timeStamp%int64(serverNum)]

	return s, nil
}
