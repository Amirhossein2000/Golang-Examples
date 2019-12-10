package pktgen

import (
	"fmt"
	"log"
	"net"
	"strings"
	"time"

	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
)

const (
	udp = iota
	tcp = iota
)

type PacketGenerator struct {
	srcIPList []string
	dstIPList []string

	srcPortList [2]uint
	dstPortList [2]uint

	protocol      string
	srcMacAddress []string
	dstMacAddress []string

	rate int
}

//ProcessScenario start generatting the nping commands of scenario.
func (packetGenerator *PacketGenerator) ProcessScenario() {
	pktCount := len(packetGenerator.srcIPList) *
		len(packetGenerator.dstIPList) *
		len(packetGenerator.srcPortList) *
		len(packetGenerator.dstPortList)

	fmt.Printf("start sending %v packets\n", pktCount)

	generatePackets(packetGenerator)
}

func generatePackets(packetGenerator *PacketGenerator) {
	// rate for sleep
	sleepTimeForRate := 1000000000 / packetGenerator.rate

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
	switch strings.ToUpper(packetGenerator.protocol) {
	case "TCP":
		protocol = tcp
	case "UDP":
		protocol = udp
	}

	for _, srcIP := range packetGenerator.srcIPList {
		for _, dstIP := range packetGenerator.dstIPList {
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

		udpLayer.SetNetworkLayerForChecksum(IPv4Layer)
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
