package loadBalancer

type Server struct {
	Host   string
	Port   string
	Weight int64
}

type Balancer interface {
	NewBalancer(s []*Server)
	DoBalance(key ...string) (*Server, error)
}
