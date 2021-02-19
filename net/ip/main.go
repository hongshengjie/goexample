package main

import (
	"fmt"
	"net"
)

func main() {
	GetAvailableIP("10.8.0.1/24")
}
func GetAvailableIP(cidr string) {
	ip, net, _ := net.ParseCIDR(cidr)

	broadcastAddr := GetBroadcastIP(net).String()
	networkAddr := net.IP.String()
	fmt.Println(broadcastAddr, networkAddr)
	var i int
	for ip := ip.Mask(net.Mask); i < 255; inc(ip) {
		fmt.Println(ip.String())
		i++
	}

}

func inc(ip net.IP) {
	for j := len(ip) - 1; j >= 0; j-- {
		ip[j]++
		if ip[j] > 0 {
			break
		}
	}
}

// GetBroadcastIP func to get the broadcast ip address of a network
func GetBroadcastIP(n *net.IPNet) net.IP {
	var broadcast net.IP
	if len(n.IP) == 4 {
		broadcast = net.ParseIP("0.0.0.0").To4()
	} else {
		broadcast = net.ParseIP("::")
	}
	for i := 0; i < len(n.IP); i++ {
		broadcast[i] = n.IP[i] | ^n.Mask[i]
	}
	return broadcast
}
