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
| Zahl                   | Byte <br> Kommazahl <br> Text <br> Wahrheitswert <br> Buchstabe | Only the lowest 8 bits are kept (numbers outside of 0 to 255 overflow)<br>-<br>-<br> 0 => falsch; nicht 0 => wahr <br> The number is interpreted as a Unicode code point |
| Byte                   | Zahl <br> Kommazahl <br> Text <br> Wahrheitswert <br> Buchstabe | -<br>-<br>-<br> 0 => falsch; nicht 0 => wahr <br> The byte is interpreted as a Unicode code point |
| Kommazahl              | Zahl, Byte <br> Text                       | truncated <br> -                                              |
| Wahrheitswert          | Zahl <br> Text                             | falsch => 0; wahr => 1 <br> -                                 |
| Text                   | Zahl <br> Kommazahl                        | The leading digits of the text are parsed, if the text is not a number the result is 0<br>The text must be a decimal number with a comma as decimal separator (e.g. `"3,14"`), otherwise the result is 0 |
| Buchstabe              | Zahl <br> Text                             | The Unicode code point of the character <br> -                |

Additionally, every value can be converted into a list of its type that contains only that value (e.g. `5 als Zahlen Liste`),
and every value can be converted into a `Variable` and a `Variable` back into its actual type.
Type definitions can be converted into their base type and back (see [Type aliases and type definitions](/en/Programmierung/Typ-Aliase-und-Typ-Definitionen)).
