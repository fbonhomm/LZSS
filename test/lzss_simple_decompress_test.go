package test

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_decompress_mode0_simple_string_without_duplicate(t *testing.T) {
	rawData := LzssM0.Decompress(ComSimpleStringWithoutDuplicateMode0)

	assert.Equal(t, len(DecSimpleStringWithoutDuplicate), len(rawData))
	assert.Equal(t, DecSimpleStringWithoutDuplicate, rawData)
}

func Test_decompress_mode0_simple_string_with_duplicate(t *testing.T) {
	rawData := LzssM0.Decompress(ComSimpleStringWithDuplicateMode0)

	assert.Equal(t, len(DecSimpleStringWithDuplicate), len(rawData))
	assert.Equal(t, DecSimpleStringWithDuplicate, rawData)
}

func Test_decompress_mode1_simple_string_without_duplicate(t *testing.T) {
	rawData := LzssM1.Decompress(ComSimpleStringWithoutDuplicateMode1)

	assert.Equal(t, len(DecSimpleStringWithoutDuplicate), len(rawData))
	assert.Equal(t, DecSimpleStringWithoutDuplicate, rawData)
}

func Test_decompress_mode1_simple_string_with_duplicate(t *testing.T) {
	rawData := LzssM1.Decompress(ComSimpleStringWithDuplicateMode1)

	assert.Equal(t, len(DecSimpleStringWithDuplicate), len(rawData))
	assert.Equal(t, DecSimpleStringWithDuplicate, rawData)
}
