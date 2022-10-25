package test

import (
	"io/ioutil"
	"log"
	"testing"

	"github.com/fbonhomm/LZSS/source"
	"github.com/stretchr/testify/assert"
)

func Test_so3_mode1(t *testing.T) {
	var lzss = source.LZSS{
		Mode:         1,
		PositionMode: 1,
	}

	fileC, err := ioutil.ReadFile("./fixture/lzss.so3.compress")
	fileD, err := ioutil.ReadFile("./fixture/lzss.so3.decompress")

	if err != nil {
		log.Fatal(err)
	}

	rawData, err := lzss.Decompress(fileC)

	assert.Nil(t, err)
	assert.Equal(t, fileD, rawData[:len(fileD)])
}

func Test_so3_cd_mode1(t *testing.T) {
	var lzss = source.LZSS{
		Mode:         1,
		PositionMode: 1,
	}

	fileD, err := ioutil.ReadFile("./fixture/lzss.so3.decompress")

	if err != nil {
		log.Fatal(err)
	}

	compressData := lzss.Compress(fileD)
	rawData, err := lzss.Decompress(compressData)

	assert.Nil(t, err)
	assert.Equal(t, fileD, rawData)
}
