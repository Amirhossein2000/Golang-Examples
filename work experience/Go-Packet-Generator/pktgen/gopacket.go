package pktgen

import (
	"log"
	"time"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
)

var (
	device                 = "lo"
	snapshotLen      int32 = 1024
	promiscuous            = false
	err              error
	timeout          = 30 * time.Second
	handle           *pcap.Handle
	options          gopacket.SerializeOptions
	srcPort, dstPort int
)

func sendUDP(data []byte,
	udpLayer *layers.UDP,
	IPv4Layer *layers.IPv4,
	ethernetLayer *layers.Ethernet) (err error) {

	buffer := gopacket.NewSerializeBuffer()
	gopacket.SerializeLayers(buffer, options,
		udpLayer,
		gopacket.Payload(data),
	)
	return sendIPv4(buffer.Bytes(), IPv4Layer, ethernetLayer)
}

func sendTCP(data []byte,
	tcpLayer *layers.TCP,
	IPv4Layer *layers.IPv4,
	ethernetLayer *layers.Ethernet) (err error) {

	buffer := gopacket.NewSerializeBuffer()
	gopacket.SerializeLayers(buffer, options,
		tcpLayer,
		gopacket.Payload(data),
	)
	return sendIPv4(buffer.Bytes(), IPv4Layer, ethernetLayer)
}

func sendIPv4(data []byte,
	IPv4Layer *layers.IPv4,
	ethernetLayer *layers.Ethernet) (err error) {

	bufferIPv4 := gopacket.NewSerializeBuffer()
	gopacket.SerializeLayers(bufferIPv4, options,
		IPv4Layer,
		gopacket.Payload(data),
	)
	return sendEthernet(bufferIPv4.Bytes(), ethernetLayer)
}

func sendEthernet(data []byte,
	ethernetLayer *layers.Ethernet) (err error) {

	bufferEthernet := gopacket.NewSerializeBuffer()
	gopacket.SerializeLayers(bufferEthernet, options,
		ethernetLayer,
		gopacket.Payload(data),
	)
	err = handle.WritePacketData(bufferEthernet.Bytes())
	if err != nil {
		log.Fatal(err)
	}
	return err
}
