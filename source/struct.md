
```go
type LZSS struct {
	Mode          int // Define le mode utiliser (0 par defaut)
	DictSize      int // La taille du dictionnaire (4096 par default)
	NumByteEncode int // Le nombre d'octets utilisee (2 par default) 
	MaxMatch      int // La taille maximal du pattern a encoder (18 par default)
	MinMatch      int // La taille minimal du pattern a encoder (3 par default)
	Length        int // Le nombres de bits a utilisee pour le stokage de la longueur du pattern (4 par default) 
	Position      int // Le nombres de bits a utilisee pour le stockage de la position du pattern (12 par default)
	MaskLength    uint32 // Mask utitlisee pour trouver la position de la longueur du pattern dans uint32 (0xF000 soit 61440  par default)
	MaskPosition  uint32 // Mask utitlisee pour trouver la position de la longueur du pattern dans uint32 (0x0FFF soit 4095  par default)
}
```

##Init

Cette fonction initialise la struct `LZSS`.

```go
c.DictSize = int(math.Pow(2, float64(c.Position)))
```
2^12 - 1 = 4096, 12 bits
```go
c.MaxMatch = int(math.Pow(2, float64(c.Length))-1) + c.MinMatch
```
2^4 - 1 = 15, 4 bits et on ajoute 3 pour faire 18
```go 
c.NumByteEncode = int(math.Ceil(float64(c.Position+c.Length) / 8))
```
(12 + 4) / 8 = 2, le `math.Ceil` arrondie au-dessus
```go
totalBits := int(math.Pow(2, float64(8*c.NumByteEncode)))

c.MaskPosition = uint32(math.Pow(2, float64(c.Position))) - 1
c.MaskLength = uint32(totalBits) - c.MaskPosition - 1
```
Creation des Mask dinamiquement.

# Compress

Selection le bon algorithme de compression en fonction du mode.

# Decompress

Selection le bon algorithme de decompression en fonction du mode.
