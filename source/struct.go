/**
 * Created by fbonhomm.
 * Email: flo-github@outlook.fr
 * License: MIT
 */

package source

import (
	"fmt"
	"log"
	"math"
)

// LZSS config struct for lzss
type LZSS struct {
	Mode          int
	DictSize      int
	PositionMode  string
	NumByteEncode int
	NumByteFlags  int
	NumByteRaw    int
	MaxMatch      int
	MinMatch      int
	Length        int
	Position      int
	MaskLength    uint32
	MaskPosition  uint32
}

// Init initialize the config
func (c *LZSS) Init() {
	if c.PositionMode != "relative" {
		c.PositionMode = "absolute"
	}
	if c.Position == 0 {
		c.Position = 12
	}
	if c.Length == 0 {
		c.Length = 4
	}
	if c.MinMatch == 0 {
		c.MinMatch = 3
	}
	if c.NumByteFlags == 0 {
		c.NumByteFlags = 1
	}
	if c.NumByteRaw == 0 {
		c.NumByteRaw = 1
	}

	c.DictSize = int(math.Pow(2, float64(c.Position)))
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
	fmt.Println("Compress...")

	c.Init()
	if c.Mode == 1 {
		return c.CompressMode1(rawData)
	}

	return c.CompressMode0(rawData)
}

// Decompress choose good mode for decompression
func (c *LZSS) Decompress(compressData []byte) []byte {
	fmt.Println("Decompress...")

	c.Init()
	if c.Mode == 1 {
		return c.DecompressMode1(compressData)
	}
	if c.Mode == 2 {
		return c.DecompressMode2(compressData)
	}

	return c.DecompressMode0(compressData)
}
