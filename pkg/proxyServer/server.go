package proxyserver

import (
	"Avako/pkg/config"
	"Avako/pkg/loadBalancer"
	"log"

	"net/http"
	"net/http/httputil"
	"net/url"
)

var Num int64 = -1   //记录调取负载均衡的次数
var Weight int64 = 0 //记录调取负载均衡的次数
type ProxyServer struct {
	Host     string //监听地址和端口
	Port     string
	Balancer loadBalancer.Balancer
}

func NewProxyServer(*config.Config) *ProxyServer {
	s1 := loadBalancer.Server{
		Host:   "127.0.0.1",
		Port:   "8001",
		Weight: 1,
	}
	s2 := loadBalancer.Server{
		Host:   "127.0.0.1",
		Port:   "8002",
		Weight: 2,
	}
	var serverAr = []*loadBalancer.Server{&s1, &s2}
	ser := ProxyServer{
		Host: "0.0.0.0",
		Port: "8888",
		//在此更改策略，支持TimeStampRandomBalancer.RandomBalance.HashBalance
		Balancer: loadBalancer.ByRequestBalancer{
			Servers: serverAr,
			Index:   &Num,
			Weight:  &Weight,
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
	//提取url参数serverid，http://127.0.0.1:8888/?serverid=1
	serverid := ""
	r.ParseForm()
	if len(r.Form["serverid"]) > 0 {
		serverid = r.Form["serverid"][0]
	}
	//获取负载均衡地址
	remoteServer, err := s.Balancer.DoBalance(r.RemoteAddr, serverid) //r.RemoteAddr作为标识，相同的r.RemoteAddr会代理到同一个IP
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
