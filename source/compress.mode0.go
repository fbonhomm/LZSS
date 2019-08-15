/**
 * Created by fbonhomm
 * Email: flo-github@outlook.fr
 * License: MIT
 */

package source

import (
	"encoding/binary"
	"math"
)

// AddToCompressData add encoded data in final byte array
func (c *LZSS) AddToCompressData(compressData []byte, tmpU, spaceTaken uint32, encoded bool) []byte {
	length := len(compressData)

	if spaceTaken != 0 {
		lastElement := compressData[length-1]
		compressData = compressData[:length-1]

		tmpU <<= spaceTaken
		tmpU += uint32(lastElement)
	}

	tmpB, _ := c.PutUint32ToBuf(tmpU)

	if encoded {
		tmpB = tmpB[:3]
	} else {
		tmpB = tmpB[:2]
	}

	return append(compressData, tmpB...)
}

// CompressMode0 compress in lzss with mode 0
func (c *LZSS) CompressMode0(rawData []byte) []byte {
	rawDataSize := binary.Size(rawData)
	var compressData []byte
	var spaceTaken uint32

	for i := 0; i < rawDataSize; i++ {
		var position, length int
		var tmpU uint32

		if spaceTaken == uint32(8 * c.NumByteFlags) {
			spaceTaken = 0
		}

		if i >= c.MinMatch {
			position, length = c.SearchBytePattern(
				c.GetChunkByte(rawData, i-c.DictSize, i+c.MaxMatch),
				c.GetChunkByte(rawData, i, i+c.MaxMatch),
				i - int(math.Max(float64(i-c.DictSize), 0)))
		}

		if length == 0 {
			tmpU = uint32(rawData[i])
			tmpU <<= 1
			tmpU++
			compressData = c.AddToCompressData(compressData, tmpU, spaceTaken, false)
		} else {
			if c.PositionMode == "relative" {
				if i > c.DictSize {
					position = int(math.Abs(float64(c.DictSize - position)))
				} else {
					position = int(math.Abs(float64(i - position)))
				}
			}
			tmpU = uint32(length-c.MinMatch) << uint32(c.Position)
			tmpU += uint32(position)
			tmpU <<= 1
			compressData = c.AddToCompressData(compressData, tmpU, spaceTaken, true)
			i += length - 1
		}

		spaceTaken++
	}

	return compressData
}
