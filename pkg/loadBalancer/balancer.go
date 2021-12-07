package loadBalancer

type Server struct {
	Host   string
	Port   string
	Weight string
}

type Balancer interface {
	DoBalance(...string) (*Server, error)
}
