/**
 * Created by fbonhomm.
 * Email: flo-github@outlook.fr
 * Licence: MIT
 */

package source

import (
	"fmt"
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
	Mode	 int
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

// Compress choose good mode for compression
func (c *LZSS) Compress(rawData []byte) []byte {
	fmt.Println("Compress...")

	if c.Mode == 0 {
		return c.CompressMode0(rawData)
	}

	return c.CompressMode0(rawData)
}

// Decompress choose good mode for decompression
func (c *LZSS) Decompress(compressData []byte) []byte {
	fmt.Println("Decompress...")

	if c.Mode == 0 {
		return c.DecompressMode0(compressData)
	}

	return c.DecompressMode0(compressData)
}
