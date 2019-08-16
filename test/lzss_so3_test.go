package test

import (
	"github.com/fbonhomm/Simple-LZSS/source"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"log"
	"testing"
)

func Test_so3_mode1(t *testing.T) {
	var lzss = source.LZSS{
		Mode: 1,
		PositionMode: "relative",
	}

	fileC, err := ioutil.ReadFile("./fixture/lzss.so3.compress")
	fileD, err := ioutil.ReadFile("./fixture/lzss.so3.decompress")

	if err != nil {
		log.Fatal(err)
	}

	rawData := lzss.Decompress(fileC)

	assert.Equal(t, fileD, rawData[:len(fileD)])
}

func Test_so3_cd_mode1(t *testing.T) {
	var lzss = source.LZSS{
		Mode: 1,
		PositionMode: "relative",
	}

	fileD, err := ioutil.ReadFile("./fixture/lzss.so3.decompress")

	if err != nil {
		log.Fatal(err)
	}

	compressData := lzss.Compress(fileD)
	rawData := lzss.Decompress(compressData)

	assert.Equal(t, fileD, rawData)
}
