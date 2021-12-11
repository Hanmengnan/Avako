package loadBalancer

type Server struct {
	Host   string
	Port   string
	Weight int64
}

type Balancer interface {
	NewBalancer(s []*Server, i int64, w int64)
	DoBalance(key ...string) (*Server, error)
}
