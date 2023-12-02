+++
title = "Type conversion"
weight = 2
type = "article"
+++

# Type conversion
A type conversion is the conversion from one type to another. In DDP, a type conversion looks like this:

```ddp
<expression of one type> als <other type>.
```

For example: `x als Text.`

## Conversion table
Only certain types can be converted to others.

| input type  | output type                                     | Notes                                                         |
|-------------|-------------------------------------------------|---------------------------------------------------------------|
| Zahl        | Kommazahl <br> Text <br> Boolean <br> Buchstabe | -<br>-<br> 0 => falsch; nicht 0 => wahr <br> uses ASCII value |
| Kommazahl   | Zahl <br> Text                                  | truncated <br> -                                              |
| Boolean     | Zahl                                            | falsch => 0; wahr => 1                                        |
| Text        | Zahl <br> Kommazahl <br> Buchstaben             | -<br>-<br>-<br>                                               |
| Buchstabe   | Zahl <br> Text                                  | utf8 bytes in decimal <br> -                                  |
| Buchstaben  | Text                                            | -                                                             |
