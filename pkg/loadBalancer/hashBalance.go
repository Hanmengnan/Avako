package loadBalancer

import (
	"errors"
	"fmt"
	"hash/crc32"
	"math/rand"
)

type HashBalance struct {
	Servers []*Server
	Index   *int64
	Weight  *int64
}

func (balancer HashBalance) DoBalance(key ...string) (*Server, error) {
	serverNum := len(balancer.Servers)
	var defKey string = fmt.Sprintf("%d", rand.Int())
	if len(key) > 0 {
		defKey = key[0]
	}
	if serverNum == 0 {
		return nil, errors.New("no instance found")
	}

	crcTable := crc32.MakeTable(crc32.IEEE)
	hashVal := crc32.Checksum([]byte(defKey), crcTable)
	index := int(hashVal) % serverNum
	s := balancer.Servers[index]

	return s, nil
}
