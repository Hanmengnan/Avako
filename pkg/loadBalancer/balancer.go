package loadBalancer

type Server struct {
	Host   string
	Port   string
	Weight int64
}

type Balancer interface {
	DoBalance(key ...string) (*Server, error)
}

/* type NewBalancer interface {
	NewBalance(Servers []*Server, Index int64, Weight int64) (*Balancer, error)
} */
