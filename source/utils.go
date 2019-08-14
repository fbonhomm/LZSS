/**
 * Created by fbonhomm
 * Email: flo-github@outlook.fr
 * Licence: MIT
 */

package source

import "encoding/binary"

// PutUint32ToBuf convert uint32 to []byte
func (c *LZSS) PutUint32ToBuf(x uint32) ([]byte, int) {
	i := 0
	buf := make([]byte, binary.MaxVarintLen32)

	for x > 0xFF {
		buf[i] = byte(x)
		x >>= 8
		i++
	}
	buf[i] = byte(x)

	return buf, i + 1
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
