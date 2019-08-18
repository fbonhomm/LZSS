package test

import (
	"github.com/fbonhomm/Simple-LZSS/source"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"log"
	"testing"
)

func Test_binary_cat_mode0(t *testing.T) {
	var Lzss = source.LZSS{}

	file, err := ioutil.ReadFile("./fixture/cat")

	if err != nil {
		log.Fatal(err)
	}

	compressData := Lzss.Compress(file)
	rawData := Lzss.Decompress(compressData)

	assert.Equal(t, file, rawData)
}

func Test_binary_cat_mode0_and_relative(t *testing.T) {
	var Lzss = source.LZSS{PositionMode: 1}

	file, err := ioutil.ReadFile("./fixture/cat")

	if err != nil {
		log.Fatal(err)
	}

	compressData := Lzss.Compress(file)
	rawData := Lzss.Decompress(compressData)

	assert.Equal(t, file, rawData)
}

func Test_binary_cat_mode1(t *testing.T) {
	var Lzss = source.LZSS{Mode: 1}

	file, err := ioutil.ReadFile("./fixture/cat")

	if err != nil {
		log.Fatal(err)
	}

	compressData := Lzss.Compress(file)
	rawData := Lzss.Decompress(compressData)

	assert.Equal(t, file, rawData)
}

func Test_binary_cat_mode1_and_relative(t *testing.T) {
	var Lzss = source.LZSS{Mode: 1, PositionMode: 1}

	file, err := ioutil.ReadFile("./fixture/cat")

	if err != nil {
		log.Fatal(err)
	}

	compressData := Lzss.Compress(file)
	rawData := Lzss.Decompress(compressData)

	assert.Equal(t, file, rawData)
}
