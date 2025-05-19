package middleware

import (
	"fmt"
	"net"
	"net/url"
)

/*
 *
 * Checks whether the URL isn't about local network IP
 *
 */
func IsUrlSafe(u *url.URL) bool {
	if u.Scheme != "https" {
		fmt.Println("No http")
		return false
	}

	var hostname string = u.Hostname()
	if hostname == "localhost" {
		fmt.Println("hostname")
		return false
	}

	var privateNetworkAddreses = [...]string{
		"127.0.0.0/8",    /* loopback */
		"10.0.0.0/8",     /* class A */
		"172.16.0.0/12",  /* class B */
		"192.168.0.0/16", /* class C */
		"::1/128",        /* IPv6 loopback */
		"fc00::/7",       /* IPv6 unique local */
		"fe80::/10",      /* IPv6 link-local */
	}
	var isPrivate = func(ip net.IP) bool {
		for _, addr := range privateNetworkAddreses {
			_, block, _ := net.ParseCIDR(addr)
			if block.Contains(ip) {
				return true
			}
		}

		return false
	}
	if ip := net.ParseIP(hostname); ip != nil {
		if isPrivate(ip) {
			fmt.Println("Blocked : Private.Reserved IP")
			return false
		}

		return true
	}

	ips, err := net.LookupIP(hostname)
	if err != nil {
		fmt.Printf("Blocked : DNS resolution failed for host %s\n", hostname)
		return false
	}

	for _, ip := range ips {
		if isPrivate(ip) {
			fmt.Printf("Blocked : domain %s resolved to private IP %s\n", hostname, ip)
			return false
		}
	}

	return true
}
