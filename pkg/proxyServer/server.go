package proxyserver

import (
	"Avako/pkg/loadBalancer"
	"log"
	"sync"

	"net/http"
	"net/http/httputil"
	"net/url"
)

type ProxyServer struct {
	Host      string
	Port      string
	Algorithm string
	Balancer  loadBalancer.Balancer
}

func NewProxyServer(ps *ProxyServer, ss []loadBalancer.Server) *ProxyServer {
	// declare array of pointer about interface that are exposed to users.
	servers := make([]*loadBalancer.Server, 0)
	for _, item := range ss {
		servers = append(servers, &loadBalancer.Server{
			Host:   item.Host,
			Port:   item.Port,
			Weight: item.Weight,
		})
	}
	var balancer loadBalancer.Balancer
	switch ps.Algorithm {
	case "TimeStampRandomBalancer":
		balancer = new(loadBalancer.TimeStampRandomBalancer)
	case "HashBalance":
		balancer = new(loadBalancer.HashBalance)
	case "ByRequestBalancer":
		balancer = new(loadBalancer.ByRequestBalancer)
	case "RandomBalance":
		balancer = new(loadBalancer.RandomBalance)
	case "WeightRoundRobin":
		balancer = new(loadBalancer.WeightRoundRobin)
	}

	balancer.NewBalancer(servers)
	// construct a ProxyServer
	ser := ProxyServer{
		Host:     ps.Host,
		Port:     ps.Port,
		Balancer: balancer,
	}
	return &ser
}

func (s *ProxyServer) StartServer(wg *sync.WaitGroup) {
	server := http.Server{
		Addr:    s.Host + ":" + s.Port,
		Handler: s,
	}
	log.Printf("Listening on %s", server.Addr)
	err := server.ListenAndServe()
	if err != nil {
		log.Println(err)
		return
	}
	wg.Done()
}

func (s *ProxyServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Variable w means response from server.
	log.Println(r.RemoteAddr + " " + r.Method + " " + r.URL.String() + " " + r.Proto + " " + r.UserAgent())
	// Get server that truly provides service by loadBalancer
	var remoteServer *loadBalancer.Server
	var err error
	switch s.Balancer.(type) {
	case *loadBalancer.ByRequestBalancer:
		_ = r.ParseForm()
		serverId := ""
		if len(r.Form["serverId"]) > 0 {
			serverId = r.Form["serverId"][0]
		}
		remoteServer, err = s.Balancer.DoBalance(r.RemoteAddr, serverId)
	default:
		remoteServer, err = s.Balancer.DoBalance()
		if err != nil {
			log.Println(err)
			return
		}
	}

	remoteURL, err := url.Parse("http://" + remoteServer.Host + ":" + remoteServer.Port)
	if err != nil {
		log.Fatalln(err)
		return
	}
	log.Printf("forward to: " + remoteURL.Host)
	// Return a new ReverseProxy that routes
	// URLs to the scheme, host, and base path provided in target.
	// Truly ReverseProxy about remoteURL!
	proxy := httputil.NewSingleHostReverseProxy(remoteURL)
	// Modify host and port in the request from client.
	r.Host = remoteServer.Host + ":" + remoteServer.Port
	// Send request to the true server.
	proxy.ServeHTTP(w, r)
}
