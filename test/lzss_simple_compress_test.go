package test

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_compress_mode0_simple_string_without_duplicate(t *testing.T) {
	compressData := LzssM0.Compress(DecSimpleStringWithoutDuplicate)

	// original length is 11 + 2 for flag (11 flag is storage on 2 bytes)
	assert.Equal(t, len(ComSimpleStringWithoutDuplicateMode0), len(compressData))
	assert.Equal(t, ComSimpleStringWithoutDuplicateMode0, compressData)
}

func Test_compress_mode0_simple_string_with_duplicate(t *testing.T) {
	compressData := LzssM0.Compress(DecSimpleStringWithDuplicate)

	// original length is 23 but a duplicate of 11 so 12 + 2 for flag + 2 for encoded pattern so 16
	assert.Equal(t, len(ComSimpleStringWithDuplicateMode0), len(compressData))
	assert.Equal(t, ComSimpleStringWithDuplicateMode0, compressData)
}

func Test_compress_mode1_simple_string_without_duplicate(t *testing.T) {
	compressData := LzssM1.Compress(DecSimpleStringWithoutDuplicate)

	assert.Equal(t, len(ComSimpleStringWithoutDuplicateMode1), len(compressData))
	assert.Equal(t, ComSimpleStringWithoutDuplicateMode1, compressData)
}

func Test_compress_mode1_simple_string_with_duplicate(t *testing.T) {
	compressData := LzssM1.Compress(DecSimpleStringWithDuplicate)

	// // original length is 23 but a duplicate of 11 so 12 + 2 for flag + 2 for encoded pattern so 16
	assert.Equal(t, len(ComSimpleStringWithDuplicateMode1), len(compressData))
	assert.Equal(t, ComSimpleStringWithDuplicateMode1, compressData)
}
