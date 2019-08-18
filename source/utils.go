/**
 * Created by fbonhomm
 * Email: flo-github@outlook.fr
 * License: MIT
 */

package source

import (
	"bytes"
)

// PutUint32ToByte convert uint32 to []byte
func (c *LZSS) PutUint32ToByte(x uint32) (buffer []byte, nbWrite int) {
	nbWrite = 0
	buffer = make([]byte, 4)

	for x > 0xFF {
		buffer[nbWrite] = byte(x)
		x >>= 8
		nbWrite++
	}
	buffer[nbWrite] = byte(x)

	return buffer, nbWrite + 1
}

// PutByteToUint32 convert []byte to uint32
func (c *LZSS) PutByteToUint32(data []byte) uint32 {
	var tmp uint32

	if len(data) > 3 {
		tmp = uint32(data[3]) << 24
	}
	if len(data) > 2 {
		tmp |= uint32(data[2]) << 16
	}
	if len(data) > 1 {
		tmp |= uint32(data[1]) << 8
	}
	if len(data) > 0 {
		tmp |= uint32(data[0])
	}

	return tmp
}

// CountBitToZero convert []byte to uint32
func (c *LZSS) CountBitToZero(n byte) uint32 {
	var count uint32

	for i := 0; i < 8; i++ {
		if (n & 1) == 0 {
			count++
		}
		n >>= 1
	}

	return count
}

// GetChunkByte slice byte array without error bound range
func (c *LZSS) GetChunkByte(data []byte, min, max int) []byte {
	size := len(data)

	if min < 0 {
		min = 0
	}

	if max >= size {
		return data[min:size]
	}
	return data[min:max]
}

// SearchBytePattern search the maximal substring length
func (c *LZSS) SearchBytePattern(data, pattern []byte, excludeIndex int) (position, length int) {
	for length = len(pattern); length >= c.MinMatch; length-- {
		position = bytes.Index(data, pattern[:length])

		if position < excludeIndex && position != -1 {
			return position, length
		}
	}
	return 0, 0
}

// GetEncodedData return feature encoded pattern
func (c *LZSS) GetEncodedData(chunk uint32) (position, length int) {
	position = int(chunk & c.MaskPosition /* 0x0FFF */)
	length = int((chunk&c.MaskLength /*0xF000*/)>>uint32(c.Position) + uint32(c.MinMatch))

	return position, length
}
