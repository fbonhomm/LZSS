[![CircleCI](https://circleci.com/gh/fbonhomm/LZSS/tree/master.svg?style=svg)](https://circleci.com/gh/fbonhomm/LZSS/tree/master)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](./LICENSE)

# LZSS

[EN] Implementation of the LZSS compression algorithm

[FR] Implémentation de l'algorithme de compression LZSS

## Usage
```go
import "github.com/fbonhomm/LZSS/source"

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

## Explanation

 - [English](documentation/explanation.en.md)
 - [Français](documentation/explanation.fr.md)

## Ressources
  - https://www.romhacking.net/utilities/819/
  - https://en.wikipedia.org/wiki/Lempel%E2%80%93Ziv%E2%80%93Storer%E2%80%93Szymanski
  - http://michael.dipperstein.com/lzss/
