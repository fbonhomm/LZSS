package test

import (
	"github.com/fbonhomm/Simple-LZSS/source"
)

var D_simple_string_without_duplicate = []byte("simple test")                                                       // length 11
var C_simple_string_without_duplicate = []byte{231, 166, 109, 11, 151, 109, 89, 144, 116, 203, 206, 165, 3}         // length 13
var D_simple_string_with_duplicate = []byte("simple test simple test")                                              // length 23
var C_simple_string_with_duplicate = []byte{231, 166, 109, 11, 151, 109, 89, 144, 116, 203, 206, 165, 11, 2, 0, 16} // length 16

var Lzss = source.LZSS{}
