package schema

import (
	"errors"
	"fmt"
	"net"
)

// IsValidIPAddress returns error if IP address is NOT valid
func IsValidIPAddress(ip string) error {
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