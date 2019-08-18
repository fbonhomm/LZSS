[![CircleCI](https://circleci.com/gh/fbonhomm/Simple-LZSS/tree/master.svg?style=svg)](https://circleci.com/gh/fbonhomm/Simple-LZSS/tree/master)

# LZSS

LZSS implementation in Go

## Usage
```go
import "github.com/fbonhomm/Simple-LZSS/source"

var Lzss = source.LZSS{Mode: 1, PositionMode: 0}

dataCompress := Lzss.Compress([]byte("..."))

rawData := Lzss.Decompress(dataCompress)
```

## Configuration

`Mode`:
  - 0 is mode with one bit flag in front each byte or encoding (default) 
  - 1 is mode with one byte flag in front each 8 pattern of character or encoding 

`PositionMode`:
  - 0 set count position to started begin dictionary (default)
  - 1 set count position to started current index

## LZSS Explications

 - [Explications en francais](FR-EXPLICATION.md)
 - [English Explications](EN-EXPLICATION.md)

## Ressources
  - https://www.romhacking.net/utilities/819/
  - https://en.wikipedia.org/wiki/Lempel%E2%80%93Ziv%E2%80%93Storer%E2%80%93Szymanski
  - http://michael.dipperstein.com/lzss/
