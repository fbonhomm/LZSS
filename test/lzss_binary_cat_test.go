package test

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"log"
	"testing"
)

func Test_binary_cat_mode0(t *testing.T) {
	file, err := ioutil.ReadFile("./fixture/cat")

	if err != nil {
		log.Fatal(err)
	}

	compressData := LzssM0.Compress(file)
	rawData := LzssM0.Decompress(compressData)

	assert.Equal(t, file, rawData)
}

func Test_binary_cat_mode1(t *testing.T) {
	file, err := ioutil.ReadFile("./fixture/cat")

	if err != nil {
		log.Fatal(err)
	}

	compressData := LzssM1.Compress(file)
	rawData := LzssM1.Decompress(compressData)

	assert.Equal(t, file, rawData)
}
