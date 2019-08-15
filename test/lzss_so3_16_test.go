package test

import (
	"github.com/fbonhomm/Simple-LZSS/source"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"log"
	"testing"
)

func Test_so3_16_mode1(t *testing.T) {
	var lzss = source.LZSS{
		Mode: 1,
		NumByteFlags: 2,
		PositionMode: "relative",
	}

	fileC, err := ioutil.ReadFile("./fixture/lzss16.so3.compress")
	fileD, err := ioutil.ReadFile("./fixture/lzss16.so3.decompress")

	if err != nil {
		log.Fatal(err)
	}

	rawData := lzss.Decompress(fileC)

	assert.Equal(t, fileD, rawData[:len(fileD)])
}
