package netparse

import (
	"net"
)

// cidrDataSourceModel describes the data source data model.
// References used.
// https://pkg.go.dev/net#ParseCIDR
type cidrModel struct {
	CIDR    string
	IP      string
	Network string
}

func ParseCIDR(c string) (*cidrModel, error) {
	cidr := c

	ip, network, err := net.ParseCIDR(cidr)
	if err != nil {
		return nil, err
	}

	return &cidrModel{
		CIDR:    cidr,
		IP:      ip.String(),
		Network: network.String(),
	}, nil
}

func (u *cidrModel) Validate() error {
	return nil
}
