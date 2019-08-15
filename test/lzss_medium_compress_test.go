package test

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_compress_mode0_medium_string(t *testing.T) {
	compressData := LzssM0.Compress(DecMediumString)

	assert.Equal(t, len(ComMediumStringMode0), len(compressData))
	assert.Equal(t, ComMediumStringMode0, compressData)
}

func Test_compress_mode1_medium_string(t *testing.T) {
	compressData := LzssM1.Compress(DecMediumString)

	assert.Equal(t, len(ComMediumStringMode1), len(compressData))
	assert.Equal(t, ComMediumStringMode1, compressData)
}
