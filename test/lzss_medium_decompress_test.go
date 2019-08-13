package test

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_decompress_medium_string(t *testing.T) {
	rawData := Lzss.Decompress(ComMediumString)

	assert.Equal(t, len(DecMediumString), len(rawData))
	assert.Equal(t, DecMediumString, rawData)
}
