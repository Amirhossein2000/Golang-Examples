package pktgen

//PacketGenerator has a channel used for generate nping commands,run nping commands and create goroutine
type PacketGenerator struct {
}

//New func returns a new PacketGenerator and set the packetChannel of the PacketGenerator.
func New(packetChannel chan string) *PacketGenerator {
	packetGenerator := new(PacketGenerator)
	return packetGenerator
}
