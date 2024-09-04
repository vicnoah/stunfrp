package client

import (
	"context"
	"net"

	"github.com/duke-git/lancet/v2/validator"
)

// hasDNS checks if the provided name is a valid domain name format.
func hasDNS(name string) bool {
	// Check if the name is a valid IP address
	if net.ParseIP(name) != nil {
		return false // If it's an IP address, it's not a domain
	}

	// Try to resolve the name to see if it's a valid domain
	isdns := validator.IsDns(name)
	return isdns
}

func lookupTXT(name string) (string, string) {
	addrs, err := net.DefaultResolver.LookupTXT(context.Background(), name)
	if err == nil {
		for _, addr := range addrs {
			addr_s, port_s, err := net.SplitHostPort(addr)
			if err != nil {
				continue
			}

			return addr_s, port_s
		}
	}

	return "127.0.0.1", "8080"
}
