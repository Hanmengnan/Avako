package proxyserver

import (
	"Avako/pkg/config"
	"Avako/pkg/loadBalancer"
	"log"
	"sync"

	"net/http"
	"net/http/httputil"
	"net/url"
)

type ProxyServer struct {
	Host     string //监听地址和端口
	Port     string
	Balancer loadBalancer.Balancer //interface
}

func NewProxyServer(cfg *config.Config, indx int) *ProxyServer {

	// declare array of pointer about interface that are exposed to users.
	servers := make([]*loadBalancer.Server, 0)
	for _, item := range cfg.Servers {
		servers = append(servers, &loadBalancer.Server{
			Host:   item.Host,
			Port:   item.Port,
			Weight: item.Weight,
		})
	}
	//var serverAr = []*loadBalancer.Server{&s1, &s2}

	// 如何给接口赋值？
	var balancer loadBalancer.Balancer
	switch cfg.Nginx[indx].Algorithm {
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

	balancer.NewBalancer(servers, 0, 0)

	// construct a ProxyServer

	ser := ProxyServer{
		Host: cfg.Nginx[indx].Host,
		Port: cfg.Nginx[indx].Port,
		// 为什么给接口赋值特定的类
		Balancer: balancer,
	}
	return &ser
}

// method of ProxyServer
func (s *ProxyServer) StartServer(wg *sync.WaitGroup) {
	server := http.Server{
		Addr:    s.Host + ":" + s.Port,
		Handler: s,
	}
	log.Printf("Listening on %s", server.Addr)
	server.ListenAndServe()
	wg.Done()
}

func (s *ProxyServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Variable w means response from server.

	log.Println(r.RemoteAddr + " " + r.Method + " " + r.URL.String() + " " + r.Proto + " " + r.UserAgent())
	// Get server that truly provides service by loadBalancer
	remoteServer, err := s.Balancer.DoBalance()
	if err != nil {
		log.Println("")
	}
	remoteURL, err := url.Parse("http://" + remoteServer.Host + ":" + remoteServer.Port)
	if err != nil {
		log.Fatalln(err)
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

	//var wg sync.waitgroup
	//wg.add(1)
	//go func() {
	//	s.StartServer()
	//	wg.Done()
	//}()
	//wg.wait()
}
