Le [LZSS](https://fr.wikipedia.org/wiki/LZSS) (Lempel-Ziv-Storer-Szymanski) est un algorithme de [compression sans perte](https://fr.wikipedia.org/wiki/Algorithme_de_compression_sans_perte).

L'algorithme est une amélioration du LZ77.

Il utilise la méthode de fenêtre glissante.

Il y a plusieurs modes, le mode 0 explique le fonctionnement du LZSS par défaut.

Les autres modes sont des variantes.

## Table of Contents
1. [Mode0](#mode0)
2. [Mode1](#mode1)
3. [PositionMode](#positionmode)

## Mode0

#### Example

`je suis ici et je suis là-bas => je suis ici et je suis la-bas`

Exemple avec la phrase : `je suis ici et je suis la-bas` de 29.

La phrase encodée : `je suis ici et (0,8)la-bas` fait 23. (+3 octets de flags donc 26)

La recherche de schéma peut chevaucher la partie tampon de lecture.


![alt tag](assets/image.gif)


L'encodage se fait la plupart de temps sur 2 Octets :

    - 12 bits pour la position
    - 4 bits pour la longueur

Un bit de flag est positionnée devant chaque caractère ou encodage pour le reconnaitre pendant la décompression.

Un caractère fait 8 bits mais avec le bit de flag, il en fait 9 bits.

Donc le non-encodage est plus volumineux que le fichier original, il faut donc se rattraper sur l'encodage.

L'encodage fait 16 bits(12 bits de position et 4 bits de longueur) + 1 bit de flag soit 17 bits.


![alt tag](assets/image1.png)


Les bits de position définie la taille du dictionnaire de recherche.
```
12 bits = 4095
```

Les bits de longueur définie la taille du tampon de lecture.
```
4 bits = 15
```
Mais le LZSS a un minimum de match qui est 3 par défaut, alors le tampon de lecture passe de [0;15] a [3;18].

La taille du tampon passe a 18, 3 sera enlevée de la compression puis 3 ajouter lors de la décompression.

Compression:
```
18 = 01001 - 3 = 1111 = 4 bits 
```

Decompression:
```
15 = 1111 + 3 = 01001 = 18 
```


![alt tag](assets/image2.png) 


## Mode1

Voici une variante qui demande moins de ressources, donc plus rapide a compressée et décompressée.

Pas besoin de garder de nombre de bits pris par l'ancien octet.

Le principe est le même, mais au lieu de mettre un flag a chaque octet ou encodage, on les regroupe.

```
11111111 si cette octet est present au debut du fichier, alors les 8 prochains octets sont non-encoded
``` 

![alt tag](assets/image3.png)


## PositionMode

Il y a principalement 2 formes de comptage de la position.

Ce que j'appelle `absolute` compte à partir du début du dictionnaire comme l'exemple plus haut.

```
Exemple avec la phrase : `je suis ici et je suis la-bas` de 29.

                     | - 8 - |
La phrase encodée : `je suis ici et (0,8)la-bas` fait 23. (+3 octets de flags donc 26)
                     |
                     0
```

Et le format `relative` compte à partir de l'index courant.

```
Exemple avec la phrase : `je suis ici et je suis la-bas` de 29.

                     | - 8 - |
La phrase encodée : `je suis ici et (15,8)la-bas` fait 23. (+3 octets de flags donc 26)
                     | ---- 15 ---- |
```
