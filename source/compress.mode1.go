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

// CompressMode1 compress in lzss with mode 1
func (c *LZSS) CompressMode1(rawData []byte) []byte {
	rawDataSize := binary.Size(rawData)
	var compressData, buffer []byte
	var flags, spaceTaken uint32
	var begin int

	for i := 0; i < rawDataSize; i++ {
		var position, length int

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
			if c.PositionMode == "relative" {
				position = (i-begin) - position
			}
		// 	if i < 200 {
		// 		fmt.Println(position, length)
		// 	}
			tmpU := uint32(length-c.MinMatch) << uint32(c.Position)
			tmpU += uint32(position)
			tmpB, _ := c.PutUint32ToBuf(tmpU)
			buffer = append(buffer, tmpB[:2]...)
			i += length - 1
		}

		spaceTaken++
		if spaceTaken == 8 * uint32(c.NumByteFlags) || i + 1 >= rawDataSize {
			tmpB, _ := c.PutUint32ToBuf(flags)
			buffer = append(tmpB[:c.NumByteFlags], buffer...)
			compressData = append(compressData, buffer...)
			buffer = buffer[:0]
			flags = 0
			spaceTaken = 0
		}
	}

	return compressData
}
