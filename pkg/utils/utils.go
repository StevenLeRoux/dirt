package utils

import (
	"fmt"
	"log"
	"net"
	"time"
)

func IsIPAddress(s string) {
	fmt.Println(s)
}

// Lookup outbound ip. Target is exepected to be in the form of "1.2.3.4:5678"
func LookupOutboundIP(target string) net.IP {
	conn, err := net.Dial("udp", target)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP
}

func Ok(d time.Duration, b bool) time.Duration {
	return d
}
