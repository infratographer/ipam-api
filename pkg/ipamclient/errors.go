package ipamclient

import (
	"errors"
)

var (
	// ErrLBHTTPUnauthorized is returned when the request is not authorized
	ErrLBHTTPUnauthorized = errors.New("ipam api received unauthorized request")

	// ErrLBHTTPNotfound is returned when the load balancer ID not found
	ErrLBHTTPNotfound = errors.New("ipam ID not found")

	// ErrLBHTTPError is returned when the http response is an error
	ErrLBHTTPError = errors.New("ipam api http error")

	// ErrNoAvailableIPs is returned when there are no available IPs within the specified block
	ErrNoAvailableIPs = errors.New("no available IPS within specified block")
)
