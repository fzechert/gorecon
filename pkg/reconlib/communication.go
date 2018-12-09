package reconlib

import "fmt"

func parsePacket(buffer []byte) (*string, error) {
	packetSize := buffer[0]
	controlData := buffer[1]
	data := buffer[2 : packetSize+2]

	debugMessage(fmt.Sprintf("received packet with %d bytes and control byte %#x, data: %#x",
		packetSize, controlData, data))

	return nil, nil
}
