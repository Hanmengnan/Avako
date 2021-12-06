package proxyserver

import (
	"Avako/pkg/config"
	"Avako/pkg/loadbalancer"
	"log"

	"net/http"
	"net/http/httputil"
	"net/url"
)

type ProxyServer struct {
	Host     string
	Port     string
	Balancer loadbalancer.Balancer
}

func NewProxyServer(*config.Config) *ProxyServer {
	return nil
}

func (s *ProxyServer) StartServer() {
	server := http.Server{
		Addr:    s.Host + ":" + s.Port,
		Handler: s,
	}
	server.ListenAndServe()
}

func (s *ProxyServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	remoteServer, err := s.Balancer.DoBalance()
	if err != nil {
		log.Println("")
	}
	remoteURL, err := url.Parse(s.Host + ":" + s.Port)
	if err != nil {
		log.Println("")
	}
	proxy := httputil.NewSingleHostReverseProxy(remoteURL)
	r.Host = remoteServer.Host
	proxy.ServeHTTP(w, r)
}
