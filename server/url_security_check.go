package yt_go_server

import (
	"net"
	"net/url"
	"strings"
)

/*
* I don't think this needs any kind of explanation
 */
func isUrlSafe(u *url.URL) bool {
	if u.Scheme != "http" && u.Scheme != "https" {
		return false
	}

	var hostname string = u.Hostname()
	if hostname == "localhost" || strings.HasPrefix(hostname, "127.") || strings.HasPrefix(hostname, "::1") {
		return false
	}

	var ip = net.ParseIP(hostname)
	if ip != nil {
		var privateNetworkAddreses []string = []string{"10.0.0.0.8", "172.16.0.0/12", "192.168.0.0/16", "127.0.0.0/8", "::1/128"}
		for _, addr := range privateNetworkAddreses {
			_, block, _ := net.ParseCIDR(addr)
			if block.Contains(ip) {
				return false
			}
		}
	}

	return true
}
