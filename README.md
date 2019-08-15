# Simple-LZSS

Simple LZSS implementation in Go

## Usage
```go
import "github.com/fbonhomm/Simple-LZSS/source"

var Lzss = source.LZSS{}

dataCompress := Lzss.Compress([]byte("..."))

rawData := Lzss.Decompress(dataCompress)
```

## LZSS Explications

 - [Explications en francais](FR-EXPLICATION.md)
