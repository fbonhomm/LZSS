/**
 * Created by fbonhomm
 * Email: flo-github@outlook.fr
 * License: MIT
 */

package source

import (
	"math"
)

// AddToCompressData add encoded data in final byte array
func (c *LZSS) AddToCompressData(compressData []byte, tmpU, flags uint32, encoded bool) []byte {
	var length = len(compressData)

	if flags != 0xFF {
		lastElement := compressData[length-1]
		compressData = compressData[:length-1]

		tmpU = (tmpU << c.CountBitToZero(byte(flags))) + uint32(lastElement)
	}

	tmpB, _ := c.PutUint32ToByte(tmpU)

	if encoded {
		tmpB = tmpB[:3]
	} else {
		tmpB = tmpB[:2]
	}

	return append(compressData, tmpB...)
}

// CompressMode0 compress in lzss with mode 0
func (c *LZSS) CompressMode0(rawData []byte) []byte {
	var rawDataSize = len(rawData)
	var compressData []byte
	var flags, tmpU uint32
	var position, length int

	for i := 0; i < rawDataSize; i++ {
		if flags == 0 {
			flags = 0xFF
		}

		if i >= c.MinMatch {
			begin := int(math.Max(float64(i-c.DictSize), 0))

			position, length = c.SearchBytePattern(
				c.GetChunkByte(rawData, begin, i+c.MaxMatch),
				c.GetChunkByte(rawData, i, i+c.MaxMatch),
				i-begin)
		}

		if length == 0 {
			tmpU = (uint32(rawData[i]) << 1) + 1
			compressData = c.AddToCompressData(compressData, tmpU, flags, false)
		} else {
			if c.PositionMode == RELATIVE {
				if i > c.DictSize {
					position = int(math.Abs(float64(c.DictSize - position)))
				} else {
					position = int(math.Abs(float64(i - position)))
				}
			}
			tmpU = uint32(length-c.MinMatch) << uint32(c.Position)
			tmpU = (tmpU + uint32(position)) << 1
			compressData = c.AddToCompressData(compressData, tmpU, flags, true)
			i += length - 1
		}
		flags >>= 1
	}

	return compressData
}
