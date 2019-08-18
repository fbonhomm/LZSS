/**
 * Created by fbonhomm
 * Email: flo-github@outlook.fr
 * License: MIT
 */

package source

import (
	"math"
)

// CompressMode1 compress in lzss with mode 1
func (c *LZSS) CompressMode1(rawData []byte) []byte {
	var rawDataSize = len(rawData)
	var compressData, buffer []byte
	var flags, spaceTaken uint32
	var begin, position, length int

	for i := 0; i < rawDataSize; i++ {
		if i >= c.MinMatch {
			begin = int(math.Max(float64(i-(c.DictSize)), 0))

			position, length = c.SearchBytePattern(
				c.GetChunkByte(rawData, begin, i+c.MaxMatch),
				c.GetChunkByte(rawData, i, i+c.MaxMatch),
				i-begin)
		}

		if length == 0 {
			buffer = append(buffer, rawData[i])
			flags |= uint32(math.Pow(2, float64(spaceTaken)))
		} else {
			if c.PositionMode == RELATIVE {
				position = (i - begin) - position
			}
			tmpU := (uint32(length-c.MinMatch) << uint32(c.Position)) + uint32(position)
			tmpB, _ := c.PutUint32ToByte(tmpU)
			buffer = append(buffer, tmpB[:2]...)
			i += length - 1
		}

		spaceTaken++
		if spaceTaken == 8 || i+1 >= rawDataSize {
			tmpB, _ := c.PutUint32ToByte(flags)
			buffer = append(tmpB[:1], buffer...)
			compressData = append(compressData, buffer...)
			buffer = buffer[:0]
			flags, spaceTaken = 0, 0
		}
	}

	return compressData
}
