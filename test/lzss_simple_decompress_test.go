package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_decompress_mode0_simple_string_without_duplicate(t *testing.T) {
	rawData, err := LzssM0.Decompress(ComSimpleStringWithoutDuplicateMode0)

	assert.Nil(t, err)
	assert.Equal(t, len(DecSimpleStringWithoutDuplicate), len(rawData))
	assert.Equal(t, DecSimpleStringWithoutDuplicate, rawData)
}

func Test_decompress_mode0_simple_string_with_duplicate(t *testing.T) {
	rawData, err := LzssM0.Decompress(ComSimpleStringWithDuplicateMode0)

	assert.Nil(t, err)
	assert.Equal(t, len(DecSimpleStringWithDuplicate), len(rawData))
	assert.Equal(t, DecSimpleStringWithDuplicate, rawData)
}

func Test_decompress_mode1_simple_string_without_duplicate(t *testing.T) {
	rawData, err := LzssM1.Decompress(ComSimpleStringWithoutDuplicateMode1)

	assert.Nil(t, err)
	assert.Equal(t, len(DecSimpleStringWithoutDuplicate), len(rawData))
	assert.Equal(t, DecSimpleStringWithoutDuplicate, rawData)
}

func Test_decompress_mode1_simple_string_with_duplicate(t *testing.T) {
	rawData, err := LzssM1.Decompress(ComSimpleStringWithDuplicateMode1)

	assert.Nil(t, err)
	assert.Equal(t, len(DecSimpleStringWithDuplicate), len(rawData))
	assert.Equal(t, DecSimpleStringWithDuplicate, rawData)
}

// in order to find crashes in Decompress()
func FuzzDecompressMode0(f *testing.F) {
	testcases := [][]byte{ComSimpleStringWithoutDuplicateMode0, ComSimpleStringWithDuplicateMode0}
	for _, tc := range testcases {
		f.Add(tc)
	}
	f.Fuzz(func(t *testing.T, orig []byte) {
		_, _ = LzssM0.Decompress(orig)
	})
}

// in order to find crashes in Decompress()
func FuzzDecompressMode1(f *testing.F) {
	testcases := [][]byte{ComSimpleStringWithoutDuplicateMode1, ComSimpleStringWithDuplicateMode1}
	for _, tc := range testcases {
		f.Add(tc)
	}
	f.Fuzz(func(t *testing.T, orig []byte) {
		_, _ = LzssM1.Decompress(orig)
	})
}
