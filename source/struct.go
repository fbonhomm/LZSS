package source

import (
	"log"
	"math"
)

type LZSS struct {
	DictSize	int
	MinMatch	int
	MaxMatch	int
	NumBytes	int
	Position	int // set in bits, default: 12, decimal 2048
	Length		int // set in bits, default: 4, decimal 15
}

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

	c.DictSize = 2^c.Position
	c.MaxMatch = 2^c.Length + c.MinMatch
	c.NumBytes = int(math.Ceil(float64(c.Position + c.Length) / 8))

	if c.MinMatch > c.MaxMatch {
		log.Fatal("The minimal value cannot be above the maximal value.")
	}
	if math.Mod(float64(c.Position + c.Length), 8) != 0 {
		log.Panic("Your Position and Length choose is not optimal.")
	}
	if c.NumBytes > 4 {
		log.Fatal("This program manage that 4 byte encoding.")
	}
}
