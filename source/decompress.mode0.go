/**
 * Created by fbonhomm.
 * Email: flo-github@outlook.fr
 * License: MIT
 */

package source

// DecompressMode0 decompress in lzss with mode 0
func (c *LZSS) DecompressMode0(compressData []byte) []byte {
	var compressDataSize = len(compressData)
	var rawData []byte
	var spaceTaken, chunk, flag uint32
	var sizeRaw int

	for i := 0; i < compressDataSize; i++ {
		chunk = c.PutByteToUint32(c.GetChunkByte(compressData, i, i+4))
		flag = (chunk >> spaceTaken) & 1
		chunk >>= spaceTaken + 1

		if flag == 1 {
			rawData = append(rawData, byte(chunk))
			sizeRaw++
		} else {
			if i+1 >= compressDataSize {
				break
			}
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
				if sizeRaw > c.DictSize {
					position += sizeRaw - c.DictSize
				}
				for j := 0; j < length; j++ {
					rawData = append(rawData, rawData[position+j])
				}
				sizeRaw += length
			}
			i++
		}

		spaceTaken++
		if spaceTaken == 8 {
			spaceTaken = 0
			i++
		}
	}

	return rawData
}
