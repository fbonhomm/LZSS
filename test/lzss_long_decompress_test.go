package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_decompress_mode0_long_string(t *testing.T) {
	rawData, err := LzssM0.Decompress(ComLongStringMode0)

	assert.Nil(t, err)
	assert.Equal(t, len(DecLongString), len(rawData))
	assert.Equal(t, DecLongString, rawData)
}

func Test_decompress_mode1_long_string(t *testing.T) {
	rawData, err := LzssM1.Decompress(ComLongStringMode1)

	assert.Nil(t, err)
	assert.Equal(t, len(DecLongString), len(rawData))
	assert.Equal(t, DecLongString, rawData)
}
