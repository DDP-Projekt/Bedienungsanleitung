+++
title = "Type conversion"
weight = 2
+++

# Type conversion
A type conversion is the conversion from one type to another. In DDP, a type conversion looks like this:

```ddp
<expression of one type> als <other type>.
```

For example: `x als Text.`

## Conversion table
Only certain types can be converted to others.

| input type             | output type                                | Notes                                                         |
|------------------------|--------------------------------------------|---------------------------------------------------------------|
| Zahl                   | Byte <br> Kommazahl <br> Text <br> Wahrheitswert <br> Buchstabe | Numbers outside the range get mapped accordingly<br>-<br>-<br> 0 => falsch; nicht 0 => wahr <br> uses ASCII value |
| Kommazahl              | Zahl, Byte <br> Text                       | truncated <br> -                                              |
| Wahrheitswert          | Zahl, Byte                                 | falsch => 0; wahr => 1                                        |
| Text                   | Zahl, Byte <br> Kommazahl <br> Buchstaben  | Numerical value gets parsed<br>Text has to be in `\d+(,\d+)?` format<br>-<br>                                               |
| Buchstabe              | Zahl, Byte <br> Text                       | utf8 bytes in decimal <br> -                                  |
| Buchstaben Liste       | Text                                       | -                                                             |
