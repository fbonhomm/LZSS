package source

import (
	"encoding/binary"
	"log"
	"math"
)

// LZSS config struct for lzss
type LZSS struct {
	DictSize int
	MinMatch int
	MaxMatch int
	NumBytes int
	Position int // set in bits, default: 12, decimal 2048
	Length   int // set in bits, default: 4, decimal 15
}

// Init initialize the config
func (c *LZSS) Init() {
	if c.Position == 0 {
		c.Position = 12
	}
	if c.Length == 0 {
		c.Length = 4
	}
	if c.MinMatch == 0 {
		c.MinMatch = 3
	}

	c.DictSize = int(math.Pow(2, float64(c.Position)))
	c.MaxMatch = int(math.Pow(2, float64(c.Length)) - 1) + c.MinMatch
	c.NumBytes = int(math.Ceil(float64(c.Position+c.Length) / 8))

	if c.MinMatch > c.MaxMatch {
		log.Fatal("The minimal value cannot be above the maximal value.")
	}
	if math.Mod(float64(c.Position+c.Length), 8) != 0 {
		log.Panic("Your Position and Length choose is not optimal.")
	}
	if c.NumBytes > 4 {
		log.Fatal("This program manage that 4 byte encoding.")
	}
}

// PutUint32ToBuf convert uint32 to []byte
func (c *LZSS) PutUint32ToBuf(x uint32) ([]byte, int) {
	i := 0
	buf := make([]byte, binary.MaxVarintLen32)

	for x > 0xFF {
		buf[i] = byte(x)
		x >>= 8
		i++
	}
	buf[i] = byte(x)

	return buf, i + 1
}

// PutByteToUint32 convert []byte to uint32
func (c *LZSS) PutByteToUint32(data []byte) uint32 {
	var tmp uint32

	if len(data) > 3 {
		tmp = uint32(data[3]) << 24
	}
	if len(data) > 2 {
		tmp |= uint32(data[2]) << 16
	}
	if len(data) > 1 {
		tmp |= uint32(data[1]) << 8
	}
	if len(data) > 0 {
		tmp |= uint32(data[0])
	}

	return tmp
}

// GetChunkByte slice byte array without error bound range
func (c *LZSS) GetChunkByte(data []byte, min, max int) []byte {
	size := len(data)

	if min < 0 {
		min = 0
	}

	if max >= size {
		return data[min:size]
	}
	return data[min:max]
}
