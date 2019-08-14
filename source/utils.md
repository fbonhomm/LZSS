
## PutUint32ToBuf

Met est `uint32` dans un tableaux de `byte`

Un `uint32` fais 4 octets, donc le tableau de sortie fera 4 element

```go
for x > 0xFF
```
Je regarde si `uint32` passer en parametre est superieur a 255(0xFF).

```go
buffer[nbWrite] = byte(x)
x >>= 8
```

Je copie le 1er octet dans le buffer et je deplacer les bit de 8 sur la gauche(little endian).

En cessant ca, le 2eme octet deviens le 1er et il est pres a etre copier.   

## PutByteToUint32

Met est tableau `byte` dans un `uint32`

```go
if len(data) > 3
```
Si le tableau est superieur a 3, je prend le 4eme element.

```go
tmp = uint32(data[3]) << 24     (tmp = tmp | uint32(data[3]) << 24)
```
Je deplacer l'octet `data[3]` a la 4eme place du uint32

Si `data[3]` = 255
```
                                      debut (little endian)
                                    /
11111111 |---------- 24 ----------|
00000000 00000000 00000000 00000000
```
et ainsi de suite

```
11111111 11111111 |----- 16 ------|
00000000 00000000 00000000 00000000

11111111 11111111 11111111 |-  8 -|
00000000 00000000 00000000 00000000

11111111 11111111 11111111 11111111
00000000 00000000 00000000 00000000
```
le `|` me permet de ne pas perdre les anciennes infomations.

## GetChunkByte

Me renvoie la partie voulue sans l'erreur de depassement.
