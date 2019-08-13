package test

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_compress_medium_string(t *testing.T) {
	compressData := Lzss.Compress(DecMediumString)

	assert.Equal(t, len(ComMediumString), len(compressData))
	assert.Equal(t, ComMediumString, compressData)
}
