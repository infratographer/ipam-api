package validator

import (
	"errors"
	"fmt"
	"net"
	"net/netip"
)

// IPAddr returns error if IP address is NOT valid
func IPAddr(ip string) error {
	if net.ParseIP(ip) != nil {
		return nil
	}

	return InvalidIPAddrError(ip)
}

// ErrInvalidIPAddr is an error raised when provided IP Address is invalid
var ErrInvalidIPAddr = errors.New("provided IP Address is invalid")

// InvalidIPAddrError returns Error Invalid IP Address
func InvalidIPAddrError(ip string) error {
	return fmt.Errorf("error %w: %s", ErrInvalidIPAddr, ip)
}

// IPBlockPref returns error if IP Block Prefix is NOT valid
func IPBlockPref(prefix string) error {
	_, err := netip.ParsePrefix(prefix)
	if err != nil {
		return InvalidIPPrefError(prefix)
	}

	return nil
}

// ErrInvalidIPPref is an error raised when provided IP Block Prefix is invalid
var ErrInvalidIPPref = errors.New("provided IP Block Prefix is invalid")

// InvalidIPPrefError returns Error Invalid IP Block Prefix
func InvalidIPPrefError(prefix string) error {
	return fmt.Errorf("error %w: %s", ErrInvalidIPPref, prefix)
}
