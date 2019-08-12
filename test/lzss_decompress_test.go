package test

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_decompress_simple_string_without_duplicate(t *testing.T) {
	rawData := Lzss.Decompress(C_simple_string_without_duplicate)

	// original length is 11 + 2 for flag (11 flag is storage on 2 bytes)
	assert.Equal(t, len(D_simple_string_without_duplicate), len(rawData))
	assert.Equal(t, D_simple_string_without_duplicate, rawData)
}

func Test_decompress_simple_string_with_duplicate(t *testing.T) {
	rawData := Lzss.Decompress(C_simple_string_with_duplicate)

	// original length is 23 but a duplicate of 11 so 12 + 2 for flag + 2 for encoded pattern so 16
	assert.Equal(t, len(D_simple_string_with_duplicate), len(rawData))
	assert.Equal(t, D_simple_string_with_duplicate, rawData)
}
