/**
 * Created by fbonhomm.
 * Email: flo-github@outlook.fr
 * License: MIT
 */

package source

import (
	"log"
	"math"
)

// ABSOLUTE define the started count to begin dictionary

// RELATIVE define the started count to current index
const RELATIVE = 1

// LZSS config struct for lzss
type LZSS struct {
	Mode          int
	DictSize      int
	PositionMode  int
	NumByteEncode int
	MaxMatch      int
	MinMatch      int
	Length        int
	Position      int
	MaskLength    uint32
	MaskPosition  uint32
}

// Init initialize the config
func (c *LZSS) Init() {
	c.Position = 12
	c.Length = 4
	c.MinMatch = 3
	c.DictSize = int(math.Pow(2, float64(c.Position))) - 1
	c.MaxMatch = int(math.Pow(2, float64(c.Length))-1) + c.MinMatch
	c.NumByteEncode = int(math.Ceil(float64(c.Position+c.Length) / 8))

	if c.MinMatch > c.MaxMatch {
		log.Fatal("The minimal value cannot be above the maximal value.")
	}
	if math.Mod(float64(c.Position+c.Length), 8) != 0 {
		log.Panic("Your Position and Length choose is not optimal.")
	}
	if c.NumByteEncode > 4 {
		log.Fatal("This program manage that 4 byte encoding.")
	}

	totalBits := int(math.Pow(2, float64(8*c.NumByteEncode)))

	c.MaskPosition = uint32(math.Pow(2, float64(c.Position))) - 1
	c.MaskLength = uint32(totalBits) - c.MaskPosition - 1
}

// Compress choose good mode for compression
func (c *LZSS) Compress(rawData []byte) []byte {
	c.Init()
	if c.Mode == 1 {
		return c.CompressMode1(rawData)
	}

	return c.CompressMode0(rawData)
}

// Decompress choose good mode for decompression
func (c *LZSS) Decompress(compressData []byte) ([]byte, error) {
	c.Init()
	if c.Mode == 1 {
		return c.DecompressMode1(compressData)
	}

	return c.DecompressMode0(compressData)
}
