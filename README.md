# ASCII-ART-FS

Ascii-art is a program which consists in receiving a string as an argument and outputting the string in a graphic representation using ASCII.

**Graphic representation**

- standard
- shadow
- thinkertoy

**Instructions**

>The usage must respect this format:

```
go run . [STRING] [BANNER]
```

>Any other formats will return the following usage message:

```
Usage: go run . [STRING] [BANNER]

EX: go run . something standard
```

**Usage**

```
$ go run . "hello" standard | cat -e
 _              _   _          $
| |            | | | |         $
| |__     ___  | | | |   ___   $
|  _ \   / _ \ | | | |  / _ \  $
| | | | |  __/ | | | | | (_) | $
|_| |_|  \___| |_| |_|  \___/  $
                               $
                               $
$ go run . "Hello There!" shadow | cat -e
                                                                                      $
_|    _|          _| _|                _|_|_|_|_| _|                                  $
_|    _|   _|_|   _| _|   _|_|             _|     _|_|_|     _|_|   _|  _|_|   _|_|   $
_|_|_|_| _|_|_|_| _| _| _|    _|           _|     _|    _| _|_|_|_| _|_|     _|_|_|_| $
_|    _| _|       _| _| _|    _|           _|     _|    _| _|       _|       _|       $
_|    _|   _|_|_| _| _|   _|_|             _|     _|    _|   _|_|_| _|         _|_|_| $
                                                                                      $
                                                                                      $
$ go run . "Hello There!" thinkertoy | cat -e
                                              $
o  o     o o           o-O-o o                $
|  |     | |             |   |                $
O--O o-o | | o-o         |   O--o o-o o-o o-o $
|  | |-' | | | |         |   |  | |-' |   |-' $
o  o o-o o o o-o         o   o  o o-o o   o-o $
                                              $
                                              $
```
