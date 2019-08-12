package source

import (
	"encoding/binary"
	"fmt"
	"math"
	//"unsafe"
)

// GetFlag return the encoded flag
func GetFlag(chunk, spaceTaken uint32) int {
	flag := chunk & uint32(math.Pow(2, float64(spaceTaken)))

	if flag > 0 {
		return 1
	}
	return 0
}

// GetEncodedData return feature encoded pattern
func (c *LZSS) GetEncodedData(chunk, spaceTaken uint32) (uint32, uint32) {
	chunk >>= spaceTaken + 1

	position := chunk & 0x0FFF
	length := (chunk&0xF000)>>uint32(c.Position) + uint32(c.MinMatch)

	return position, length
}

// Decompress decompress on using lzss
func (c *LZSS) Decompress(compressData []byte) []byte {
	compressDataSize := binary.Size(compressData)
	var rawData []byte
	var spaceTaken uint32

	c.Init()

	fmt.Println(compressData)
	for i := 0; i < compressDataSize; i++ {
		if spaceTaken == 8 {
			spaceTaken = 0
			i++
		}

		chunk := c.PutByteToUint32(c.GetChunkByte(compressData, i, i+4))
		flag := GetFlag(chunk, spaceTaken)

		if flag == 1 {
			rawData = append(rawData, byte(chunk>>(spaceTaken+1)))
		} else {
			if i+1 >= compressDataSize {
				break
			}
			position, length := c.GetEncodedData(chunk, spaceTaken)
			rawData = append(rawData, rawData[position:position+length]...)
		}
		fmt.Println(flag)

		if flag == 0 {
			i++
		}
		spaceTaken++
	}

	return rawData
}
