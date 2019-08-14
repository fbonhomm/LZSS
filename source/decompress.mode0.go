/**
 * Created by fbonhomm.
 * Email: flo-github@outlook.fr
 * License: MIT
 */

package source

import (
	"encoding/binary"
	"fmt"
	"math"
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
func (c *LZSS) GetEncodedData(chunk uint32) (position, length int) {
	position = int(chunk & c.MaskPosition /* 0x0FFF */)
	length = int((chunk&c.MaskLength /*0xF000*/)>>uint32(c.Position) + uint32(c.MinMatch))

	return position, length
}

// DecompressMode0 decompress in lzss with mode 0
func (c *LZSS) DecompressMode0(compressData []byte) []byte {
	compressDataSize := binary.Size(compressData)
	var rawData []byte
	var spaceTaken uint32
	var sizeRaw int

	c.Init()

	fmt.Println("Decompress...")
	for i := 0; i < compressDataSize; i++ {
		if spaceTaken == 8 {
			spaceTaken = 0
			i++
		}

		chunk := c.PutByteToUint32(c.GetChunkByte(compressData, i, i+4))
		flag := GetFlag(chunk, spaceTaken)

		chunk >>= spaceTaken + 1

		if flag == 1 {
			rawData = append(rawData, byte(chunk))
			sizeRaw++
		} else {
			if i+1 >= compressDataSize {
				break
			}
			position, length := c.GetEncodedData(chunk)

			if sizeRaw > c.DictSize {
				position += sizeRaw - c.DictSize
			}
			rawData = append(rawData, c.GetChunkByte(rawData, position, position+length)...)
			sizeRaw += length
			i++
		}

		spaceTaken++
	}

	return rawData
}
