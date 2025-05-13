package middleware

import (
	"net"
	"net/url"
	"strings"
)

/*
 *
 * Checks whether the URL isn't about local network IP
 *
 */
func IsUrlSafe(u *url.URL) bool {
	if u.Scheme != "https" {
		return false
	}

	var hostname string = u.Hostname()
	if hostname == "localhost" || strings.HasPrefix(hostname, "127.") || strings.HasPrefix(hostname, "::1") {
		return false
	}

	var networkIP = net.ParseIP(hostname)
	if networkIP != nil {
		var privateNetworkAddreses = [...]string{
			"10.0.0.0/8",     /* class A */
			"172.16.0.0/12",  /* class B */
			"192.168.0.0/16", /* class C */
			"127.0.0.0/8",    /* loopback */
			"::1/128",        /* IPv6 loopback */
			"fc00::/7",       /* IPv6 unique local */
			"fe80::/10",      /* IPv6 link-local */
		}

		for _, addr := range privateNetworkAddreses {
			_, blockedCIDR, _ := net.ParseCIDR(addr)

			if blockedCIDR.Contains(networkIP) {
				return false
			}
		}

		return true
	}

	return false
}
