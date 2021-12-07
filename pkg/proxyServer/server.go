package proxyserver

import (
	"Avako/pkg/config"
	"Avako/pkg/loadBalancer"
	"log"

	"net/http"
	"net/http/httputil"
	"net/url"
)

type ProxyServer struct {
	Host     string //监听地址和端口
	Port     string
	Balancer loadBalancer.Balancer
}

func NewProxyServer(*config.Config) *ProxyServer {
	s1 := loadBalancer.Server{
		Host: "127.0.0.1",
		Port: "8001",
	}
	s2 := loadBalancer.Server{
		Host: "127.0.0.1",
		Port: "8002",
	}
	var serverAr = []*loadBalancer.Server{&s1, &s2}
	ser := ProxyServer{
		Host: "0.0.0.0",
		Port: "8888",
		Balancer: loadBalancer.TimeStampRandomBalancer{
			Servers: serverAr,
		},
	}
	return &ser
}

func (s *ProxyServer) StartServer() {
	server := http.Server{
		Addr:    s.Host + ":" + s.Port,
		Handler: s,
	}
	log.Printf("Listening on %s", server.Addr)
	server.ListenAndServe()
}

func (s *ProxyServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println(r.RemoteAddr + " " + r.Method + " " + r.URL.String() + " " + r.Proto + " " + r.UserAgent())
	//获取负载均衡地址
	remoteServer, err := s.Balancer.DoBalance()
	if err != nil {
		log.Println("")
	}
	remoteURL, err := url.Parse("http://" + remoteServer.Host + ":" + remoteServer.Port)
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("forward to: " + remoteURL.Host)
	proxy := httputil.NewSingleHostReverseProxy(remoteURL)
	r.Host = remoteServer.Host + ":" + remoteServer.Port
	proxy.ServeHTTP(w, r)
}
