// Package iputils implements some IPRange functions.
package iputils

import (
	"math"
	"net"
	"strconv"
	"strings"
)

//Hosts func takes a ip and range then returns all of the subnet as a slice of strings.
func Hosts(cidr string) ([]string, error) {
	ip, ipnet, err := net.ParseCIDR(cidr)
	if err != nil {
		return nil, err
	}

	var ips []string

	for ip := ip.Mask(ipnet.Mask); ipnet.Contains(ip); inc(ip) {
		ips = append(ips, ip.String())
	}
	return ips, nil
}

//inc fun is used for hosts func.
func inc(ip net.IP) {
	for j := len(ip) - 1; j >= 0; j-- {
		ip[j]++
		if ip[j] > 0 {
			break
		}
	}
}

//SubnetCount takes a range ip and returns the subnet count.
func SubnetCount(ipStrRange string) (int, error) {
	rangeNumber, _ := strconv.Atoi(strings.Split(ipStrRange, "/")[1])
	return int(math.Pow(2, float64(32-rangeNumber))), nil
}
