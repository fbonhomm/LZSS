/**
 * Created by fbonhomm
 * Email: flo-github@outlook.fr
 * Licence: MIT
 */

package source

import (
	"encoding/binary"
	"strings"
)

// SearchPattern search the maximal substring length
func (c *LZSS) SearchPattern(data []byte, pattern []byte) (int, int) {
	window := string(data)
	max := len(pattern)

	for ; max >= c.MinMatch; max-- {
		idx := strings.Index(window, string(pattern[:max]))

		if idx != -1 {
			return idx, max
		}
	}
	return 0, 0
}

// AddToCompressData add encoded data in final byte array
func (c *LZSS) AddToCompressData(compressData []byte, tmpI uint32, spaceTaken uint32, encoded bool) []byte {
	length := len(compressData)

	if spaceTaken != 0 {
		lastElement := compressData[length-1]
		compressData = compressData[:length-1]

		tmpI <<= spaceTaken
		tmpI += uint32(lastElement)
	}

	tmpB, _ := c.PutUint32ToBuf(tmpI)

	if encoded == true {
		tmpB = tmpB[:3]
	} else {
		tmpB = tmpB[:2]
	}
	compressData = append(compressData, tmpB...)

	return compressData
}

// CompressMode0 compress in lzss with mode 0
func (c *LZSS) CompressMode0(rawData []byte) []byte {
	rawDataSize := binary.Size(rawData)
	var compressData []byte
	var position, length int
	var tmpI, spaceTaken uint32

	c.Init()

	for i := 0; i < rawDataSize; i++ {
		tmpI = 0
		length = 0

		if spaceTaken == 8 {
			spaceTaken = 0
		}

		if i >= c.MinMatch {
			position, length = c.SearchPattern(
				c.GetChunkByte(rawData, i - c.DictSize, i),
				c.GetChunkByte(rawData, i, i + c.MaxMatch))
		}

		if length == 0 {
			tmpI = uint32(rawData[i])
			tmpI <<= 1
			tmpI++
			compressData = c.AddToCompressData(compressData, tmpI, spaceTaken, false)
		} else {
			//fmt.Println("Index: ", i, ", Position: ", position, ", Length: ", length, ", SpaceTaken: ", spaceTaken)
			tmpI = uint32(length - c.MinMatch) << uint32(c.Position)
			tmpI += uint32(position)
			tmpI <<= 1
			compressData = c.AddToCompressData(compressData, tmpI, spaceTaken, true)
			i += length - 1
		}

		spaceTaken++
	}

	return compressData
}
