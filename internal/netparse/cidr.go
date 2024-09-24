package netparse

import (
	"fmt"
	"net"
)

// CidrModel describes the CIDR model.
// References used.
// https://pkg.go.dev/net#ParseCIDR
type CidrModel struct {
	CIDR    string
	IP      string
	Network string
}

func ParseCIDR(c string) (*CidrModel, error) {
	cidr := c

	ip, network, err := net.ParseCIDR(cidr)
	if err != nil {
		return nil, err
	}

	return &CidrModel{
		CIDR:    cidr,
		IP:      ip.String(),
		Network: network.String(),
	}, nil
}

func (u *CidrModel) Validate() error {
	return nil
}

func CidrValidate(u string) error {
	return nil
}

func ContainsIP(network string, ip string) (bool, error) {
	_, parsedNetwork, err := net.ParseCIDR(network)
	if err != nil {
		return false, err
	}
	
	parsedIP := net.ParseIP(ip)
	if parsedIP == nil {
		return false, fmt.Errorf("failed to parse IP address: %s", ip)
	}

	return parsedNetwork.Contains(parsedIP), nil
}
