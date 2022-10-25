/**
 * Created by fbonhomm.
 * Email: flo-github@outlook.fr
 * License: MIT
 */

package source

import (
	"math"
)

// DecompressMode1 decompress in lzss with mode 1
func (c *LZSS) DecompressMode1(compressData []byte) ([]byte, error) {
	var compressDataSize = len(compressData)
	var rawData []byte
	var flags uint32
	var sizeRaw int

	for i := 0; i < compressDataSize; i++ {
		flags >>= 1
		if flags == 0 || flags <= 0xFF {
			flags = 0xFF00 + uint32(compressData[i])
			i++
		}

		if (flags & 1) == 1 {
			rawData = append(rawData, compressData[i])
			sizeRaw++
		} else {
			if i+1 >= compressDataSize {
				break
			}
			chunk := c.PutByteToUint32(c.GetChunkByte(compressData, i, i+2))
			position, length := c.GetEncodedData(chunk)

			if c.PositionMode == RELATIVE {
				if position == 0 {
					position = 1
				}
				for j := 0; j < length; j++ {
					rawData = append(rawData, rawData[sizeRaw-position])
					sizeRaw++
				}
			} else {
				position += int(math.Max(float64(sizeRaw-c.DictSize), 0))
				for j := 0; j < length; j++ {
					rawData = append(rawData, rawData[position+j])
				}
				sizeRaw += length
			}
			i++
		}
	}

	return rawData, nil
}
