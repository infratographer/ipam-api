package validator

import (
	"errors"
	"fmt"
)

// ErrInvalidIPAddr is an error raised when provided IP Address is invalid
var ErrInvalidIPAddr = errors.New("ip Address is invalid")

// InvalidIPAddrError returns Error Invalid IP Address
func InvalidIPAddrError(ip string) error {
	return fmt.Errorf("%w: %s", ErrInvalidIPAddr, ip)
}

// ErrInvalidIPPref is an error raised when provided IP Block Prefix is invalid
var ErrInvalidIPPref = errors.New("ip block prefix is invalid")

// InvalidIPPrefError returns Error Invalid IP Block Prefix
func InvalidIPPrefError(prefix string) error {
	return fmt.Errorf("%w: %s", ErrInvalidIPPref, prefix)
}

// ErrIPAddrOutsideBlock is an error raised when provided IP Address is not part of the IP Block
var ErrIPAddrOutsideBlock = errors.New("ip address is invalid for IP Block Prefix")

// IPAddrOutsideBlockError returns Error IP Address doesn't belong to the IP Block
func IPAddrOutsideBlockError(block string, ip string) error {
	return fmt.Errorf("%w: ip address: %s, ip block %s", ErrIPAddrOutsideBlock, ip, block)
}
