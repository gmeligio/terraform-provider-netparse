package netparse

import (
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

func ContainsIp(network string, ip string) (*DomainModel, error) {
	return nil, nil
}