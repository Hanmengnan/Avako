package loadBalancer

import (
	"errors"
)

type WeightRoundRobin struct {
	Servers []*Server
	Index   int64
	Weight  int64
}

func (balancer *WeightRoundRobin) NewBalancer(s []*Server, i int64, w int64) {
	balancer.Servers = s
}
func (balancer *WeightRoundRobin) DoBalance(key ...string) (*Server, error) {
	serverNum := len(balancer.Servers)
	if serverNum == 0 {
		return nil, errors.New("no instance found")
	}
	s := balancer.GetInst()

	return s, nil
}
func (balancer *WeightRoundRobin) GetInst() *Server {
	gcd := getGCD(balancer.Servers)
	for {
		balancer.Index = (balancer.Index + 1) % int64(len(balancer.Servers))
		if balancer.Index == 0 {
			balancer.Weight = balancer.Weight - gcd
			if balancer.Weight <= 0 {
				balancer.Weight = getMaxWeight(balancer.Servers)
				if balancer.Weight == 0 {
					return &Server{}
				}
			}
		}
		if balancer.Servers[balancer.Index].Weight >= balancer.Weight {
			return balancer.Servers[balancer.Index]
		}
	}
}

//计算两个数的最大公约数
func gcd(a, b int64) int64 {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

//计算多个数的最大公约数
func getGCD(insts []*Server) int64 {
	var weights []int64

	for _, inst := range insts {
		weights = append(weights, inst.Weight)
	}

	g := weights[0]
	for i := 1; i < len(weights)-1; i++ {
		oldGcd := g
		g = gcd(oldGcd, weights[i])
	}
	return g
}

//获取最大权重
func getMaxWeight(insts []*Server) int64 {
	var max int64 = 0
	for _, inst := range insts {
		if inst.Weight >= int64(max) {
			max = inst.Weight
		}
	}
	return max
}
