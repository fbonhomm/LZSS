package test

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_decompress_mode0_long_string(t *testing.T) {
	rawData := LzssM0.Decompress(ComLongStringMode0)

	assert.Equal(t, len(DecLongString), len(rawData))
	assert.Equal(t, DecLongString, rawData)
}

func Test_decompress_mode1_long_string(t *testing.T) {
	rawData := LzssM1.Decompress(ComLongStringMode1)

	assert.Equal(t, len(DecLongString), len(rawData))
	assert.Equal(t, DecLongString, rawData)
}
