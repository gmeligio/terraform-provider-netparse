package netparse

import (
	"fmt"
	"net/netip"
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

	prefix, err := netip.ParsePrefix(cidr)
	if err != nil {
		return nil, err
	}
	network := prefix.Masked()
	ip := prefix.Addr()

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
	prefix, err := netip.ParsePrefix(network)
	if err != nil {
		return false, fmt.Errorf("failed to parse IP network: %s", network)
	}
	
	addr, err := netip.ParseAddr(ip)
	if err != nil {
		return false, fmt.Errorf("failed to parse IP address: %s", ip)
	}

	return prefix.Contains(addr), nil
}
