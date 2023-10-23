package validator

import (
	"net"

	"net/netip"

	"github.com/3th1nk/cidr"
)

// IPAddr returns error if IP address is NOT valid
func IPAddr(ip string) error {
	if net.ParseIP(ip) != nil {
		return nil
	}

	return InvalidIPAddrError(ip)
}

// IPBlockPref returns error if IP Block Prefix is NOT valid
func IPBlockPref(prefix string) error {
	_, err := netip.ParsePrefix(prefix)
	if err != nil {
		return InvalidIPPrefError(prefix)
	}

	return nil
}

// PartOfBlock returns error if IP address is NOT part of the block given block's prefix
func PartOfBlock(ipBlockPref string, ipAdrr string) error {
	c, _ := cidr.Parse(ipBlockPref)
	belongsToBlock := false

	c.Each(func(ip string) bool {
		if ip == ipAdrr {
			belongsToBlock = true
		}

		return true
	})

	if belongsToBlock {
		return nil
	}

	return IPAddrOutsideBlockError(ipBlockPref, ipAdrr)
}
