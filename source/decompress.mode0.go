/**
 * Created by fbonhomm.
 * Email: flo-github@outlook.fr
 * License: MIT
 */

package source

import (
	"encoding/binary"
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

// DecompressMode0 decompress in lzss with mode 0
func (c *LZSS) DecompressMode0(compressData []byte) []byte {
	compressDataSize := binary.Size(compressData)
	var rawData []byte
	var spaceTaken uint32
	var sizeRaw int

	for i := 0; i < compressDataSize; i++ {
		if spaceTaken == uint32(8 * c.NumByteFlags) {
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

			if c.PositionMode == "relative" {
				if position == 0 {
					position = 1
				}
				for j := 0; j < length ; j++ {
					rawData = append(rawData, rawData[(sizeRaw) - position])
					sizeRaw++
				}
			} else {
				if sizeRaw > c.DictSize {
					position += sizeRaw - c.DictSize
				}
				for j := 0; j < length ; j++ {
					rawData = append(rawData, rawData[position + j])
				}
				sizeRaw += length
			}
			i += c.NumByteEncode - 1
		}

		spaceTaken++
	}

	return rawData
}
