package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"strings"
	"time"

	"github.com/go-routeros/routeros"
)

var (
	address    = flag.String("address", "192.168.0.1:8728", "Address")
	username   = flag.String("username", "admin", "Username")
	password   = flag.String("password", "admin", "Password")
	properties = flag.String("properties", "name,rx-byte,tx-byte,rx-packet,tx-packet", "Properties")
	interval   = flag.Duration("interval", 1*time.Second, "Interval")
)

func main() {
	flag.Parse()

	c, err := routeros.Dial(*address, *username, *password)

	if err != nil {
		log.Fatal(err)
	}

	for {
		reply, err := c.Run("/interface/print", "?disabled=false", "?running=true", "=.proplist="+*properties)
		if err != nil {
			log.Fatal(err)
		}

		for _, re := range reply.Re {
			for _, p := range strings.Split(*properties, ",") {
				fmt.Print(re.Map[p], "\t")
			}
			fmt.Print("\n")
		}
		fmt.Print("\n")

		time.Sleep(*interval)
	}
}

func Hosts(cidr string) ([]string, error) {
	ip, ipnet, err := net.ParseCIDR(cidr)
	if err != nil {
		return nil, err
	}

	var ips []string
	for ip := ip.Mask(ipnet.Mask); ipnet.Contains(ip); inc(ip) {
		ips = append(ips, ip.String())
	}
	// remove network address and broadcast address
	return ips[1 : len(ips)-1], nil
}

//  http://play.golang.org/p/m8TNTtygK0
func inc(ip net.IP) {
	for j := len(ip) - 1; j >= 0; j-- {
		ip[j]++
		if ip[j] > 0 {
			break
		}
	}
}
