package pktgen

import (
	"fmt"
	"log"
	"net"
	"rebin-test-bed/macutils"
	"strings"
	"time"

	"rebin-test-bed/ipmaputils"
	"rebin-test-bed/iputils"
	"rebin-test-bed/scenario"

	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
)

const (
	udp = iota
	tcp = iota
)

//ProcessScenario start generatting the nping commands of scenario.
func (packetGenerator *PacketGenerator) ProcessScenario(scenarioNum int) int {
	dstIPCount, _ := iputils.SubnetCount(scenario.Scenarios.Scenario[scenarioNum].Params.DestIP)
	srcIPCount, _ := iputils.SubnetCount(scenario.Scenarios.Scenario[scenarioNum].Params.SrcIP)

	allPkt := dstIPCount * srcIPCount *
		(scenario.Scenarios.Scenario[scenarioNum].Params.DestPortEnd -
			scenario.Scenarios.Scenario[scenarioNum].Params.DestPortStart + 1) *
		(scenario.Scenarios.Scenario[scenarioNum].Params.SrcPortEnd -
			scenario.Scenarios.Scenario[scenarioNum].Params.SrcPortStart + 1)

	fmt.Printf("start sending %v packets\n", allPkt)

	generatePackets(scenarioNum)

	return allPkt
}

func generatePackets(scenarioNum int) {
	srcIps, _ := iputils.Hosts(scenario.Scenarios.Scenario[scenarioNum].Params.SrcIP)
	dstIPs, _ := iputils.Hosts(scenario.Scenarios.Scenario[scenarioNum].Params.DestIP)

	// increment mac Address for all of the src ips
	macList := macutils.GenerateMacList(len(srcIps))

	// generate IPMapCSV.
	ipmaputils.WriteIPMapToRedis(srcIps, macList)

	// rate for sleep
	sleepTimeForRate := 1000000000 / scenario.Scenarios.Scenario[scenarioNum].Params.Rate

	// Set other options (false or true)
	options.FixLengths = true
	options.ComputeChecksums = true

	// Open device
	handle, err = pcap.OpenLive(device, snapshotLen, promiscuous, timeout)
	if err != nil {
		log.Fatal(err)
	}

	defer handle.Close()

	var protocol uint8
	switch strings.ToUpper(scenario.Scenarios.Scenario[scenarioNum].Params.Protocol) {
	case "TCP":
		protocol = tcp
	case "UDP":
		protocol = udp
	}

	for userID, srcIP := range srcIps {
		for _, dstIP := range dstIPs {
			for srcPort := scenario.Scenarios.Scenario[scenarioNum].Params.SrcPortStart; srcPort <= scenario.Scenarios.Scenario[scenarioNum].Params.SrcPortEnd; srcPort++ {
				for dstPort := scenario.Scenarios.Scenario[scenarioNum].Params.DestPortStart; dstPort <= scenario.Scenarios.Scenario[scenarioNum].Params.DestPortEnd; dstPort++ {
					sendPackets(srcIP, dstIP, macList[userID], srcPort, dstPort, protocol)
					time.Sleep(time.Nanosecond * time.Duration(sleepTimeForRate))
				}
			}
		}
	}
}

func sendPackets(srcIP, dstIP, mac string, srcPort, dstPort int, protocol uint8) {
	rawBytes := []byte("JustForTest")
	srcMac, err := net.ParseMAC("02:00:00:00:00:01")
	dstMac, _ := net.ParseMAC("ee:00:00:00:00:ee")
	if err != nil {
		panic(err)
	}

	IPv4Layer := &layers.IPv4{
		Version:    4,   //uint8
		IHL:        5,   //uint8
		TOS:        0,   //uint8
		Id:         0,   //uint16
		Flags:      0,   //IPv4Flag
		FragOffset: 0,   //uint16
		TTL:        255, //uint8
		SrcIP:      net.ParseIP(srcIP),
		DstIP:      net.ParseIP(dstIP),
	}

	ethernetLayer := &layers.Ethernet{
		SrcMAC:       srcMac,
		DstMAC:       dstMac,
		EthernetType: layers.EthernetTypeIPv4,
	}

	switch protocol {
	case udp:
		udpLayer := &layers.UDP{
			SrcPort: layers.UDPPort(srcPort),
			DstPort: layers.UDPPort(dstPort),
		}
		IPv4Layer.Protocol = layers.IPProtocolUDP
		sendUDP(rawBytes, udpLayer, IPv4Layer, ethernetLayer)
	case tcp:
		tcpLayer := &layers.TCP{
			SrcPort: layers.TCPPort(srcPort),
			DstPort: layers.TCPPort(dstPort),
		}
		IPv4Layer.Protocol = layers.IPProtocolTCP
		sendTCP(rawBytes, tcpLayer, IPv4Layer, ethernetLayer)
	}
}
