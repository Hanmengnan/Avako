package loadBalancer

import (
	"errors"
)

type WeightRoundRobin struct {
	Servers []*Server
	Index   *int64
	Weight  *int64
}

func (balancer WeightRoundRobin) DoBalance(key ...string) (*Server, error) {
	serverNum := len(balancer.Servers)
	if serverNum == 0 {
		return nil, errors.New("no instance found")
	}
	s := balancer.GetInst()

	return s, nil
}
func (p *WeightRoundRobin) GetInst() *Server {
	gcd := getGCD(p.Servers)
	for {
		*p.Index = (*p.Index + 1) % int64(len(p.Servers))
		if *p.Index == 0 {
			*p.Weight = *p.Weight - gcd
			if *p.Weight <= 0 {
				*p.Weight = getMaxWeight(p.Servers)
				if *p.Weight == 0 {
					return &Server{}
				}
			}
		}
		if p.Servers[*p.Index].Weight >= *p.Weight {
			return p.Servers[*p.Index]
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
