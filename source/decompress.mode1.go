/**
 * Created by fbonhomm.
 * Email: flo-github@outlook.fr
 * License: MIT
 */

package source

import (
    "encoding/binary"
)

// DecompressMode1 decompress in lzss with mode 1
func (c *LZSS) DecompressMode1(compressData []byte) []byte {
	compressDataSize := binary.Size(compressData)
	var rawData []byte
	var flags, spaceTaken uint32
	var sizeRaw int

	for i := 0; i < compressDataSize; i++ {
		flags >>= 1
		if spaceTaken == 0 || spaceTaken == uint32(8 * c.NumByteFlags) {
			flags = c.PutByteToUint32(c.GetChunkByte(compressData, i, i+c.NumByteFlags))
			spaceTaken = 0
			i += c.NumByteFlags
		}

		if (flags & uint32(1)) == 1 {
			var j int

			for ; j < c.NumByteRaw; j++ {
				rawData = append(rawData, compressData[i+j])
			}
			i += j - 1
			sizeRaw += j
		} else {
			if i+1 >= compressDataSize {
				break
			}
			chunk := c.PutByteToUint32(compressData[i:i+c.NumByteEncode])
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
