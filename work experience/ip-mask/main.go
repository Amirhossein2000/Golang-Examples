package main

import (
	"fmt"
	"net"
)

func main() {
	fmt.Println(hostsCount("192.167.0.1/28"))
	fmt.Println((hosts("192.167.0.1/28")))
}

func inc(ip net.IP) {
	for j := len(ip) - 1; j >= 0; j-- {
		ip[j]++
		if ip[j] > 0 {
			break
		}
	}
}

func hostsCount(cidr string) (int, error) {
	ip, ipnet, err := net.ParseCIDR(cidr)
	if err != nil {
		return -1, err
	}

	var ips int

	for ip := ip.Mask(ipnet.Mask); ipnet.Contains(ip); inc(ip) {
		ips++
	}
	// remove network address and broadcast address
	return ips, nil
}

//hosts func will take a ip and range then will return all of the submasks.
func hosts(cidr string) []string {
	ip, ipnet, _ := net.ParseCIDR(cidr)
	// if err != nil {
	// 	return nil, err
	// }

	var ips []string

	for ip := ip.Mask(ipnet.Mask); ipnet.Contains(ip); inc(ip) {
		ips = append(ips, ip.String())
	}
	// remove network address and broadcast address
	return ips
}
