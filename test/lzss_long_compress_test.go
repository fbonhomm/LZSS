package test

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_compress_mode0_long_string(t *testing.T) {
	compressData := LzssM0.Compress(DecLongString)

	assert.Equal(t, len(ComLongStringMode0), len(compressData))
	assert.Equal(t, ComLongStringMode0, compressData)
}

func Test_compress_mode1_long_string(t *testing.T) {
	compressData := LzssM1.Compress(DecLongString)

	assert.Equal(t, len(ComLongStringMode1), len(compressData))
	assert.Equal(t, ComLongStringMode1, compressData)
}
