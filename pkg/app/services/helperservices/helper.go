package helperservices

import (
	"fmt"
	"net"
	"net/http"
	"strings"
)

func GetIP(r *http.Request) (string, error) {

	myIP := r.Header.Get("X-REAL-IP")
	netIP := net.ParseIP(myIP)
	if netIP != nil {
		return myIP, nil
	}

	myIPs := r.Header.Get("X-FORWARDED-FOR")
	myIPList := strings.Split(myIPs, ",")
	for _, ip := range myIPList {
		netIP := net.ParseIP(ip)
		if netIP != nil {
			return ip, nil
		}
	}

	myIP, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		return "", err
	}
	netIP = net.ParseIP(myIP)
	if netIP != nil {
		return myIP, nil
	}
	return "", fmt.Errorf("invalid ip address")
}
