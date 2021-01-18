package decoding

import "errors"

const (
	lapDataPacketidentifier uint8 = 2
	sessionPacketIdentifier uint8 = 1

	lapDataPacketByteCount = 843
	sessionPacketByteCount = 149

	packetHeaderByteOffset     = 23
	packetIdentifierByteOffset = 5
	playerCarIndexByteOffset   = 22
	playerCurrentLapByteOffset = 32

	lapDataByteCount            = 41
	totalNumberOfLapsByteOffset = packetHeaderByteOffset + 3
)

// IsLapDataPacket takes in a slice of bytes representing a packet and returns if the
// packet is a lap data packet.
func IsLapDataPacket(bytes []byte) (bool, error) {
	if len(bytes) != lapDataPacketByteCount {
		return false, nil
	}

	packetType, err := getPacketType(bytes)
	if err != nil {
		return false, err
	}

	return packetType == lapDataPacketidentifier, nil
}

// IsSessionPacket takes in a slice of bytes representing a packet and returns if the
// packet is a session packet.
func IsSessionPacket(bytes []byte) (bool, error) {
	if len(bytes) != sessionPacketByteCount {
		return false, nil
	}

	packetType, err := getPacketType(bytes)
	if err != nil {
		return false, err
	}

	return packetType == sessionPacketIdentifier, nil
}

// GetTotalLaps takes a slice of bytes representing a session packet and returns the
// number of laps in the given session/race
func GetTotalLaps(bytes []byte) (uint8, error) {
	if len(bytes) != sessionPacketByteCount {
		return 0, errors.New("Not enough bytes to be a session packet to determine the total lap count")
	}

	return uint8(bytes[totalNumberOfLapsByteOffset]), nil
}

// GetPlayerCurrentLap takes a slice of bytes representing a lap data packet and uses the
// CalculatePlayerCurrentLapByteOffset and GetPlayerCarIndex to retrieve the current lap the player is on.
func GetPlayerCurrentLap(bytes []byte) (uint8, error) {
	if len(bytes) != lapDataPacketByteCount {
		return 0, errors.New("Not enough bytes to be a lap data packet to determine the players current lap")
	}

	playerCarIndex, err := getPlayerCarIndex(bytes)
	if err != nil {
		return 0, err
	}
	byteOffset := calculatePlayerCurrentLapByteOffset(playerCarIndex) + playerCurrentLapByteOffset

	return uint8(bytes[byteOffset]), nil
}

func getPacketType(bytes []byte) (uint8, error) {
	if len(bytes) <= packetIdentifierByteOffset {
		return 0, errors.New("Not enough bytes to determine packet type")
	}

	return uint8(bytes[packetIdentifierByteOffset]), nil
}

func getPlayerCarIndex(bytes []byte) (uint8, error) {
	if len(bytes) <= playerCarIndexByteOffset {
		return 0, errors.New("Not enough bytes to determine playercar index")
	}

	return uint8(bytes[playerCarIndexByteOffset]), nil
}

func calculatePlayerCurrentLapByteOffset(playerCarIndex uint8) int {
	return int(packetHeaderByteOffset + playerCarIndex*lapDataByteCount)
}
