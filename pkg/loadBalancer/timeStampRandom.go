package loadBalancer

import (
	"time"
)

type TimeStampRandomBalancer struct {
	Servers []*Server
}

func (balancer TimeStampRandomBalancer) DoBalance(...string) (*Server, error) {
	serverNum := len(balancer.Servers)
	timeStamp := time.Now().Unix()
	s := balancer.Servers[timeStamp%int64(serverNum)]

	server := new(Server)
	server.Host = s.Host
	server.Port = s.Port
	server.Weight = s.Weight

	return server, nil
}
