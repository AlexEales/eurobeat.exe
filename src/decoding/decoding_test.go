package decoding

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIsLapDataPacket(t *testing.T) {
	t.Run("packet bytes size incorrect returns false", func(t *testing.T) {
		bytes := make([]byte, 16)
		rand.Read(bytes)

		isLapData, err := IsLapDataPacket(bytes)
		require.False(t, isLapData)
		require.NoError(t, err)
	})

	t.Run("incorrect packet type returns false", func(t *testing.T) {
		bytes := make([]byte, lapDataPacketByteCount)
		rand.Read(bytes)
		bytes[packetIdentifierByteOffset] = byte(sessionPacketIdentifier)

		isLapData, err := IsLapDataPacket(bytes)
		require.False(t, isLapData)
		require.NoError(t, err)
	})

	t.Run("lap data packet returns true", func(t *testing.T) {
		bytes := make([]byte, lapDataPacketByteCount)
		rand.Read(bytes)
		bytes[packetIdentifierByteOffset] = byte(lapDataPacketidentifier)

		isLapData, err := IsLapDataPacket(bytes)
		require.True(t, isLapData)
		require.NoError(t, err)
	})
}

func TestIsSessionPacket(t *testing.T) {
	t.Run("packet bytes size incorrect returns false", func(t *testing.T) {
		bytes := make([]byte, 16)
		rand.Read(bytes)

		isLapData, err := IsSessionPacket(bytes)
		require.False(t, isLapData)
		require.NoError(t, err)
	})

	t.Run("incorrect packet type returns false", func(t *testing.T) {
		bytes := make([]byte, sessionPacketByteCount)
		rand.Read(bytes)
		bytes[packetIdentifierByteOffset] = byte(lapDataPacketidentifier)

		isLapData, err := IsSessionPacket(bytes)
		require.False(t, isLapData)
		require.NoError(t, err)
	})

	t.Run("lap data packet returns true", func(t *testing.T) {
		bytes := make([]byte, sessionPacketByteCount)
		rand.Read(bytes)
		bytes[packetIdentifierByteOffset] = byte(sessionPacketIdentifier)

		isLapData, err := IsSessionPacket(bytes)
		require.True(t, isLapData)
		require.NoError(t, err)
	})
}

func TestGetTotalLaps(t *testing.T) {
	t.Run("error returned when incorrect number of bytes provided", func(t *testing.T) {
		bytes := make([]byte, 16)
		rand.Read(bytes)

		totalLaps, err := GetTotalLaps(bytes)
		require.Equal(t, totalLaps, uint8(0))
		require.Error(t, err)
		require.EqualError(t, err, "Not enough bytes to be a session packet to determine the total lap count")
	})

	t.Run("correct total number of laps returned when packet bytes are valid", func(t *testing.T) {
		bytes := make([]byte, sessionPacketByteCount)
		rand.Read(bytes)
		bytes[totalNumberOfLapsByteOffset] = 42

		totalLaps, err := GetTotalLaps(bytes)
		require.Equal(t, totalLaps, uint8(42))
		require.NoError(t, err)
	})
}

func TestGetPlayerCurrentLap(t *testing.T) {
	t.Run("error returned when incorrect number of bytes provided", func(t *testing.T) {
		bytes := make([]byte, 16)
		rand.Read(bytes)

		currentLap, err := GetPlayerCurrentLap(bytes)
		require.Equal(t, currentLap, uint8(0))
		require.Error(t, err)
		require.EqualError(t, err, "Not enough bytes to be a lap data packet to determine the players current lap")
	})

	t.Run("correct current lap number returned when packet bytes are valid", func(t *testing.T) {
		bytes := make([]byte, lapDataPacketByteCount)
		rand.Read(bytes)
		bytes[playerCarIndexByteOffset] = byte(0)
		bytes[packetHeaderByteOffset+playerCurrentLapByteOffset] = byte(21)

		currentLap, err := GetPlayerCurrentLap(bytes)
		require.Equal(t, currentLap, uint8(21))
		require.NoError(t, err)
	})
}
