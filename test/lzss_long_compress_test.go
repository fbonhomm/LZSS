package test

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"fmt"
)

func Test_compress_long_string(t *testing.T) {
	compressData := Lzss.Compress(DecLongString)

	assert.Equal(t, len(ComLongString), len(compressData))
	assert.Equal(t, ComLongString, compressData)
}
