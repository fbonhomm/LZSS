package test

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_decompress_mode0_medium_string(t *testing.T) {
	rawData := LzssM0.Decompress(ComMediumStringMode0)

	assert.Equal(t, len(DecMediumString), len(rawData))
	assert.Equal(t, DecMediumString, rawData)
}

func Test_decompress_mode1_medium_string(t *testing.T) {
	rawData := LzssM1.Decompress(ComMediumStringMode1)

	assert.Equal(t, len(DecMediumString), len(rawData))
	assert.Equal(t, DecMediumString, rawData)
}
