package validator

import (
	"errors"
	"fmt"
	"net"

	"github.com/3th1nk/cidr"
)

// IPAddr returns error if IP address is NOT valid
func IPAddr(ip string) error {
	if net.ParseIP(ip) != nil {
		return nil
	}

	return InvalidIPAddrError(ip)
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

// ErrInvalidIPAddr is an error raised when provided IP Address is invalid
var ErrInvalidIPAddr = errors.New("provided IP Address is invalid")

// InvalidIPAddrError returns Error Invalid IP Address
func InvalidIPAddrError(ip string) error {
	return fmt.Errorf("error %w: %s", ErrInvalidIPAddr, ip)
}

// ErrIPAddrOutsideBlock is an error raised when provided IP Address is not part of the IP Block
var ErrIPAddrOutsideBlock = errors.New("provided IP Address is not part of the IP Block - Prefix")

// IPAddrOutsideBlockError returns Error IP Address doesn't belong to the IP Block
func IPAddrOutsideBlockError(block string, ip string) error {
	return fmt.Errorf("error %w: %s; IP Address: %s", ErrIPAddrOutsideBlock, block, ip)
}
