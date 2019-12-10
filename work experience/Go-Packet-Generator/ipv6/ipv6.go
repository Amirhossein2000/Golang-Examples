package IPv6

import (
	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
)

var options gopacket.Serializeoptions

func send_udp(data []byte,
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

func send_tcp(data []byte,
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
