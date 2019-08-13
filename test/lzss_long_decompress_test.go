package test

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_decompress_long_string(t *testing.T) {
	rawData := Lzss.Decompress(ComLongString)

	assert.Equal(t, len(DecLongString), len(rawData))
	assert.Equal(t, DecLongString, rawData)
}
