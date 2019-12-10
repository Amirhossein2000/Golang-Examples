package main

import (
	"github.com/Amirhossein2000/Go-Packet-Generator/iputils"
	"log"
)

type config struct {
	srcIP *string
	dstIP *string

	srcPort *uint
	dstPort *uint

	srcMacAddress *string
	dstMacAddress *string

	protocol *string
	rate     *uint
}

func main() {
	//flagConf :=config{}
	//
	//flagConf.srcIP  = flag.String("src-ip", "", "Single-IP 192.180.10.2 or Range-IP 192.179.3.1/28")
	//flagConf.dstIP  = flag.String("dst-ip", "", "Single-IP 192.180.10.2 or Range-IP 192.179.3.1/28")
	//
	//flagConf.srcPort  = flag.Uint("src-port", 0, "Single-Port 9000 or Range-Port 9000-9010")
	//flagConf.dstPort  = flag.Uint("dst-port", 0, "Single-Port 9000 or Range-Port 9000-9010")

	//flagConf.srcMacAddress  = flag.String("src-mac-address", "random", "single 00:14:22:01:23:45 or random")
	//flagConf.dstMacAddress  = flag.String("dst-mac-address", "random", "single 00:14:22:01:23:45 or random")

	//
	//flagConf.protocol  = flag.String("protocol", "", "tcp or udp")

	//flag.Parse()
	//checkConfig(flagConf)

}

func checkConfig(flagconf config) *Gnerator {
	srcIPList, err := iputils.Hosts(*flagconf.srcIP)
	if err != nil {
		log.Fatal("Bad srcIP --> ", flagconf.srcIP)
	}

	generator := new(Gnerator)
	generator.srcIPList = srcIPList
	return generator
}
