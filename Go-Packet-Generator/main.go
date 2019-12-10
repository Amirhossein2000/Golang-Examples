package main

import (
	"flag"
	"log"
	"net"
	"time"

	"Go-Packet-Generator/iputils"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
)

var (
	device       string = "lo"
	snapshot_len int32  = 1024
	promiscuous  bool   = false
	err          error
	timeout      time.Duration = 30 * time.Second
	handle       *pcap.Handle
	buffer       gopacket.SerializeBuffer
	options      gopacket.SerializeOptions
	// layer options
	srcMac, dstMac net.HardwareAddr
	//srcIp, dstIp    net.IP
	srcPort, dstPort int
	count            int
)

func main() {
	// command-line flags
	_device := flag.String("device", "lo", "device name")
	_count := flag.Int("count", 1, "repeat count")
	_srcMac := flag.String("smac", "02:00:00:00:00:01", "source MAC")
	_dstMac := flag.String("dmac", "06:00:00:00:00:01", "destination MAC")
	// _srcIp := flag.String("sip", "127.0.0.1", "source IPv4 address range")
	// _dstIp := flag.String("dip", "10.0.3-1.11", "destination IPv4 address range")
	// _srcPort := flag.String("sport", "9000", "source udp port range")
	// _dstPort := flag.String("dport", "41-43", "destination udp port range")

	srcIP := "170.0.0.1/28"
	dstIP := "168.0.0.1/28"
	// parse and set command-line options
	flag.Parse()

	device = *_device
	count = *_count
	srcMac, _ = net.ParseMAC(*_srcMac)
	dstMac, _ = net.ParseMAC(*_dstMac)

	// Set other options (false or true)
	options.FixLengths = true
	options.ComputeChecksums = true

	// Open device
	handle, err = pcap.OpenLive(device, snapshot_len, promiscuous, timeout)
	if err != nil {
		log.Fatal(err)
	}
	defer handle.Close()

	rawBytes := make([]byte, 200)

	IPv4Layer := &layers.IPv4{
		Version:    4,   //uint8
		IHL:        5,   //uint8
		TOS:        0,   //uint8
		Id:         0,   //uint16
		Flags:      0,   //IPv4Flag
		FragOffset: 0,   //uint16
		TTL:        255, //uint8
		Protocol:   17,  //IPProtocol UDP(17)
	}

	//IPv6Layer := &layers.IPv6{
	//	Version:    6,
	//	NextHeader: 41,
	//	SrcIP:      net.ParseIP("170.0.0.1"),
	//	DstIP:      net.ParseIP("168.0.0.1"),
	//}

	udpLayer := &layers.UDP{
		SrcPort: 9000,
		DstPort: 8000,
	}

	tcpLayer := &layers.TCP{
		SrcPort: 9000,
		DstPort: 8000,
	}

	ethernetLayer := &layers.Ethernet{
		SrcMAC:       srcMac,
		DstMAC:       dstMac,
		EthernetType: 0x800,
	}

	srcIPList, _ := iputils.Hosts(srcIP)
	dstIPList, _ := iputils.Hosts(dstIP)
	udpLayer.SetNetworkLayerForChecksum(IPv4Layer)

	// log.Println(srcIPList)
	// log.Println(dstIPList)

	//send_tcp_v6(rawBytes, tcpLayer, IPv6Layer, ethernetLayer)
	//send_udp_v6(rawBytes, udpLayer, IPv6Layer, ethernetLayer)

	for _, sIP := range srcIPList {
		for _, dIP := range dstIPList {
			IPv4Layer.SrcIP = net.ParseIP(sIP)
			IPv4Layer.DstIP = net.ParseIP(dIP)
			send_udp(rawBytes, udpLayer, IPv4Layer, ethernetLayer)
			time.Sleep(time.Nanosecond * 3000)
		}
	}

	IPv4Layer.Protocol = 6

	for _, sIP := range srcIPList {
		for _, dIP := range dstIPList {
			IPv4Layer.SrcIP = net.ParseIP(sIP)
			IPv4Layer.DstIP = net.ParseIP(dIP)
			send_tcp(rawBytes, tcpLayer, IPv4Layer, ethernetLayer)
			time.Sleep(time.Nanosecond * 3000)
		}
	}
}

func send_udp(data []byte,
	udpLayer *layers.UDP,
	IPv4Layer *layers.IPv4,
	ethernetLayer *layers.Ethernet) (err error) {

	buffer := gopacket.NewSerializeBuffer()
	gopacket.SerializeLayers(buffer, options,
		udpLayer,
		gopacket.Payload(data),
	)
	return send_IPv4(buffer.Bytes(), IPv4Layer, ethernetLayer)
}

func send_tcp(data []byte,
	tcpLayer *layers.TCP,
	IPv4Layer *layers.IPv4,
	ethernetLayer *layers.Ethernet) (err error) {

	buffer := gopacket.NewSerializeBuffer()
	gopacket.SerializeLayers(buffer, options,
		tcpLayer,
		gopacket.Payload(data),
	)
	return send_IPv4(buffer.Bytes(), IPv4Layer, ethernetLayer)
}

func send_IPv4(data []byte,
	IPv4Layer *layers.IPv4,
	ethernetLayer *layers.Ethernet) (err error) {

	buffer_IPv4 := gopacket.NewSerializeBuffer()
	gopacket.SerializeLayers(buffer_IPv4, options,
		IPv4Layer,
		gopacket.Payload(data),
	)
	return send_ethernet(buffer_IPv4.Bytes(), ethernetLayer)
}

func send_ethernet(data []byte,
	ethernetLayer *layers.Ethernet) (err error) {

	buffer_ethernet := gopacket.NewSerializeBuffer()
	gopacket.SerializeLayers(buffer_ethernet, options,
		ethernetLayer,
		gopacket.Payload(data),
	)
	err = handle.WritePacketData(buffer_ethernet.Bytes())
	if err != nil {
		log.Fatal(err)
	}
	return err
}

func send_udp_v6(data []byte,
	udpLayer *layers.UDP,
	IPv6Layer *layers.IPv6,
	ethernetLayer *layers.Ethernet) (err error) {

	buffer := gopacket.NewSerializeBuffer()
	gopacket.SerializeLayers(buffer, options,
		udpLayer,
		gopacket.Payload(data),
	)
	return send_IPv6(buffer.Bytes(), IPv6Layer, ethernetLayer)
}

func send_tcp_v6(data []byte,
	tcpLayer *layers.TCP,
	IPv6Layer *layers.IPv6,
	ethernetLayer *layers.Ethernet) (err error) {

	buffer := gopacket.NewSerializeBuffer()
	gopacket.SerializeLayers(buffer, options,
		tcpLayer,
		gopacket.Payload(data),
	)
	return send_IPv6(buffer.Bytes(), IPv6Layer, ethernetLayer)
}

func send_IPv6(data []byte,
	IPv6Layer *layers.IPv6,
	ethernetLayer *layers.Ethernet) (err error) {

	buffer_IPv6 := gopacket.NewSerializeBuffer()
	gopacket.SerializeLayers(buffer_IPv6, options,
		IPv6Layer,
		gopacket.Payload(data),
	)
	return send_ethernet(buffer_IPv6.Bytes(), ethernetLayer)
}

func send_ethernet_v6(data []byte,
	ethernetLayer *layers.Ethernet) (err error) {

	buffer_ethernet := gopacket.NewSerializeBuffer()
	gopacket.SerializeLayers(buffer_ethernet, options,
		ethernetLayer,
		gopacket.Payload(data),
	)
	err = handle.WritePacketData(buffer_ethernet.Bytes())
	if err != nil {
		log.Fatal(err)
	}
	return err
}
