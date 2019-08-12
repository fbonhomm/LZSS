package source

import (
	"encoding/binary"
	"fmt"
	"strings"
)

// SearchPattern search the maximal substring length
func (c *LZSS) SearchPattern(data []byte, pattern []byte) (int, int) {
	window := string(data)
	max := len(pattern)

	for ; max >= c.MinMatch; max-- {
		idx := strings.Index(window, string(pattern[:max]))

		if idx != -1 {
			return idx, max
		}
	}
	return 0, 0
}

// AddToCompressData add encoded data in final byte array
func (c *LZSS) AddToCompressData(compressData []byte, tmpI uint32, spaceTaken uint32) []byte {
	tmpB := make([]byte, binary.MaxVarintLen32)
	length := len(compressData)

	if spaceTaken != 0 {
		lastElement := compressData[length-1]
		compressData = compressData[:length-1]

		tmpI <<= spaceTaken
		tmpI += uint32(lastElement)
	}

	n := c.PutUint32ToBuf(tmpB, tmpI)
	fmt.Println(tmpB, n, tmpB[:n])
	tmpB = tmpB[:n]
	compressData = append(compressData, tmpB...)

	return compressData
}

// RemoveLastElementUseless remove all final element that equal at 0
func RemoveLastElementUseless(compressData []byte) []byte {
	for len(compressData) > 0 {
		if compressData[len(compressData)-1] != 0 {
			break
		}

		compressData = compressData[:len(compressData)-1]
	}

	return compressData
}

// Compress compress in lzss
func (c *LZSS) Compress(rawData []byte) []byte {
	rawDataSize := binary.Size(rawData)
	var compressData []byte
	var position, length int
	var tmpI, spaceTaken uint32

	c.Init()

	fmt.Println(rawData)
	for i := 0; i < rawDataSize; i++ {
		tmpI = 0
		length = 0

		fmt.Printf("Index :%d, Character: %p\n", i, rawData[i])
		if spaceTaken == 8 {
			spaceTaken = 0
		}

		if i >= c.MinMatch {
			position, length = c.SearchPattern(
				c.GetChunkByte(rawData, i-c.DictSize, i),
				c.GetChunkByte(rawData, i, i+c.MaxMatch))
		}

		if length == 0 {
			tmpI = uint32(rawData[i])
			tmpI <<= 1
			tmpI++
			fmt.Printf("Character encoded: %p\n", tmpI)
			compressData = AddToCompressData(compressData, tmpI, spaceTaken)
			fmt.Println(compressData)
		} else {
			fmt.Println("FOUND PATTERN")
			tmpI = uint32(length-c.MinMatch) << uint32(c.Position)
			tmpI += uint32(position)
			tmpI <<= 1
			fmt.Printf("Character encoded: %p\n", tmpI)
			fmt.Printf("spaceTaken: %p\n", spaceTaken)
			compressData = AddToCompressData(compressData, tmpI, spaceTaken)
			fmt.Println(compressData)
			i += length - 1
		}

		spaceTaken++
		fmt.Printf("\n\n")
	}

	return RemoveLastElementUseless(compressData)
}
