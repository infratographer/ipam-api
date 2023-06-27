package ipamclient

import (
	"errors"
)

var (
	// ErrIPAMHTTPUnauthorized is returned when the request is not authorized
	ErrIPAMHTTPUnauthorized = errors.New("ipam api received unauthorized request")

	// ErrIPAMHTTPNotfound is returned when the load balancer ID not found
	ErrIPAMHTTPNotfound = errors.New("ipam ID not found")

	// ErrIPAMHTTPError is returned when the http response is an error
	ErrIPAMHTTPError = errors.New("ipam api http error")

	// ErrNoAvailableIPs is returned when there are no available IPs within the specified block
	ErrNoAvailableIPs = errors.New("no available IPS within specified block")
)
