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

func NewProxyServer(*config.Config) *ProxyServer {
	// Servers that truly provide service
	s1 := loadBalancer.Server{
		Host: "127.0.0.1",
		Port: "8001",
	}
	s2 := loadBalancer.Server{
		Host: "127.0.0.1",
		Port: "8002",
	}
	// declare array of pointer about interface that are exposed to users.
	var serverAr = []*loadBalancer.Server{&s1, &s2}
	// construct a ProxyServer
	ser := ProxyServer{
		Host: "0.0.0.0",
		Port: "8888",
		// 为什么给接口赋值特定的类
		Balancer: loadBalancer.TimeStampRandomBalancer{
			Servers: serverAr,
		},
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

	var wg sync.waitgroup
	wg.add(1)
	go func() {
		s.StartServer()
		wg.Done()
	}()
	wg.wait()
}
