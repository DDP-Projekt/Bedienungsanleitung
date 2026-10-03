+++
title = "Data types"
weight = 1
+++

# Data types

Since DDP is statically typed, every expression (e.g. mathematical expressions), every variable and every function has a fixed data type.

The type of variables, functions and expressions cannot change at runtime, it is determined at compile time.

## Simple data types

| type name     | Description                         | range of values                                                      | literal                                                                  | Example                                                       |
| ------------- | ----------------------------------- | -------------------------------------------------------------------- | ------------------------------------------------------------------------ | ------------------------------------------------------------- |
| Zahl          | A 64-bit integer                    | *-2^63* to *2^63-1*                                                  | A sequence of digits, e.g. 42                                            | `Die Zahl x ist 75.`, <br>`1 plus -7`                         |
| Kommazahl     | A 64-bit floating point number      | approx. *-1,797x10^308* to <br>*1,797x10^308* with 16 decimal places | A number literal with decimal places, e.g. 3,1415                        | `Die Kommazahl x ist 6,5.`, <br>`2 durch 0,5`                 |
| Byte          | An 8-bit positive integer           | *0* to *255*                                                         | A sequence of digits, e.g. 16                                            | `Das Byte x ist 128.`, <br>`1 plus 5`                         |
| Wahrheitswert | A boolean value (8 bits in size)    | *wahr* or *falsch*                                                   | *wahr* or *falsch*                                                       | `Der Wahrheitswert x ist wahr.`, <br>`2 gleich 2`             |
| Buchstabe     | A 4-byte Unicode character          | *0* - *1114111*                                                      | A utf8 character between single quotes, e.g. 'a' or '\n'                 | `Der Buchstabe x ist 'd'.`                                    |
| Text          | A utf-8 encoded sequence of letters | *any size*                                                           | Any number of letters between (English) quotation marks, e.g. "Hello\n" | `Der Text x ist "abc".`, <br>`"Hello" verkettet mit " you there"` |

### Unlike in other programming languages:

* **The decimal separator is not a period, but a comma!**
* **There is no null/nil type**

### Escape sequences
Escape sequences are character combinations that stand for characters
which one cannot normally write in text or letter literals.

| name            | escape sequence | ASCII code |
| --------------- | --------------- | ---------- |
| newline         | `\n`            | 10         |
| carriage return | `\r`            | 13         |
| backspace       | `\b`            | 8          |
| Tab             | `\t`            | 9          |
| Acoustic signal | `\a`            | 7          |
| backslash       | `\`            | 92         |
| Single quote*   | `\'`            | 39         |
| double quote**  | `\"`            | 34         |

*: Only within a letter literal.\
**: Only within a text literal.

## Lists

Lists are collections of values of any size.
Since DDP is statically typed, a list can only contain values of the same data type.
The type name of a list is generally the element type name, declined accordingly, with *Liste* appended (Zahl -> Zahlen Liste, Text -> Text Liste).
For user-defined types ([combinations](/en/Programmierung/Kombinationen)) the correct declension can't be parsed (yet), so the name is not declined and *Liste* is simply appended (see [combination lists](/en/Programmierung/Kombinationen#combination-lists)).
A list can grow and shrink at runtime.
How to work with lists is described in the article Operators under [List Operators](/en/Programmierung/Operatoren#list-and-text-operators).

### List literals

A list literal looks similar to an ordinary enumeration, but wrapped in a small phrase to avoid ambiguities when programming.
The general form is: `eine[r] Liste, die aus x[, y, z] besteht`.
`eine` or `einer` can be used depending on the grammatical context, and the `aus` must be followed by one or more
expressions of the same type.

Empty lists can be created like this: `eine[r] leere[n] (type name) Liste`.

Lists with a single element can also be created simply by using `(value) als (type name)`.

There is one more special case: when declaring a variable, lists can also be created with the `Die (type name) Liste (variable name) ist <number> Mal <value>` syntax.
The list is then initialized with the value, which appears in the list as often as the number specifies.
This is especially useful to create relatively large lists with a default value, since with this syntax the whole memory is allocated only once, instead of once for every element.

### Examples
```ddp
Die Zahlen Liste z ist eine Liste, die aus 1, 2, 3 besteht.
Die Text Liste t ist eine Liste, die aus "Hello", "World" besteht.

Die Zahlen Liste z2 ist eine leere Zahlen Liste.
Die Text Liste t2 ist "Hello" als Text Liste.
```

| type name           | Example                                                                         |
| ------------------- | ------------------------------------------------------------------------------- |
| Zahlen Liste        | `Die Zahlen Liste z ist eine Liste, die aus 1, 2, 3 besteht.`                   |
| Kommazahlen Liste   | `Die Kommazahlen Liste z ist eine Liste, die aus 1,2, 3,2, 3,1415 besteht.`     |
| Byte Liste          | `Die Byte Liste by ist eine Liste, die aus 5, 12, 255 besteht.`                 |
| Wahrheitswert Liste | `Die Wahrheitswert Liste w ist eine Liste, die aus wahr, falsch, wahr besteht.` |
| Buchstaben Liste    | `Die Buchstaben Liste b ist eine Liste, die aus 'b', 'h', 'z' besteht.`         |
| Text Liste          | `Die Text Liste t ist eine Liste, die aus "Hello", "you", "there" besteht.`         |

### Remark

Actually, one would expect that an 'und' would have to occur in the enumeration of a list literal (`eine Liste, die aus 1, 2 und 3 besteht`). However, this would lead to ambiguities in boolean expressions (`eine Liste, die aus wahr und falsch besteht`), and since the enumeration grammatically does not need an 'und', it is omitted in list literals.

## Combinations

Combinations (structs in C) are user-defined composite data types that combine one or more variables into one type.
More about combinations can be found in the article [Combinations](/en/Programmierung/Kombinationen).
