package schema

import (
	"fmt"
	"net"
)

// IsValidIPAddress returns error if IP address is NOT valid
func IsValidIPAddress(ip string) error {
	if net.ParseIP(ip) != nil {
		return nil
	}

	return fmt.Errorf("Provided IP Address is invalid: %s", ip)
}
