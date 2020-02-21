package main

import (
	"fmt"
	"github.com/codekey/ipv4"
)

func main() {
	reqIP := []string{"192.168.1.2", "200.32.22.12", "192.32.22.11",
		"200.32.22.15", "12.23.22.4", "24.70.12.3", "192.68.3.90",
		"192.168.1.1", "62.3.44.2", "192.168.1.1", "24.70.12.2",
		"100.64.0.1", "100.121.0.3"}

	ipv4.Sort(reqIP)

	fmt.Printf("        IP        Pub        Priv       CGN\n")
	for _, ip := range reqIP {
		fmt.Printf("%17s %8v %8v %8v\n", 
		   ip,ipv4.IsPublic(ip), ipv4.IsPrivate(ip), ipv4.IsCGN(ip))
	}
}
