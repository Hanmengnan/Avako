package loadBalancer

type Server struct {
	Host   string
	Port   string
	Weight string
}

type Balancer interface {
	DoBalance(key ...string) (*Server, error)
}
