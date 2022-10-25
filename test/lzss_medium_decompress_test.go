package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_decompress_mode0_medium_string(t *testing.T) {
	rawData, err := LzssM0.Decompress(ComMediumStringMode0)

	assert.Nil(t, err)
	assert.Equal(t, len(DecMediumString), len(rawData))
	assert.Equal(t, DecMediumString, rawData)
}

func Test_decompress_mode1_medium_string(t *testing.T) {
	rawData, err := LzssM1.Decompress(ComMediumStringMode1)

	assert.Nil(t, err)
	assert.Equal(t, len(DecMediumString), len(rawData))
	assert.Equal(t, DecMediumString, rawData)
}
