package test

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_compress_simple_string_without_duplicate(t *testing.T) {
	compressData := Lzss.Compress(D_simple_string_without_duplicate)

	// original length is 11 + 2 for flag (11 flag is storage on 2 bytes)
	assert.Equal(t, len(C_simple_string_without_duplicate), len(compressData))
	assert.Equal(t, C_simple_string_without_duplicate, compressData)
}

func Test_compress_simple_string_with_duplicate(t *testing.T) {
	compressData := Lzss.Compress(D_simple_string_with_duplicate)

	// original length is 23 but a duplicate of 11 so 12 + 2 for flag + 2 for encoded pattern so 16
	assert.Equal(t, len(C_simple_string_with_duplicate), len(compressData))
	assert.Equal(t, C_simple_string_with_duplicate, compressData)
}
