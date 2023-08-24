# Data types

Since the DDP is statically typed, every expression (e.g. mathematical expressions), every variable and every function has a fixed data type.

The type of variables, functions and expressions cannot change at runtime, it is determined at compile time.

## Simple data types

| type name  | Description                         | range of values                                                                    | literal                                                           | Example                                                       |
|------------|-------------------------------------|------------------------------------------------------------------------------------|-------------------------------------------------------------------|---------------------------------------------------------------|
| Zahl       | A 64-bit integer                    | *-2,147,483,648* to *2,147,483,647*                                                | A sequence of digits, e.g. 42                                     | `Die Zahl x ist 69.`, <br>`1 plus -7`                         |
| Kommazahl  | A 64-bit floating point number      | *-1.79769313486232x10^308* to <br>*1.79769313486232x10^308* with 16 decimal places | A number literal with decimal places, e.g. 3.1415                 | `Die Kommazahl x ist 6,5.`, <br>`2 durch 0,5`                 |
| Boolean    | A logical value (8 bits in size)    | *true* or *false*                                                                  | *true* or *false*                                                 | `Der Boolean x ist wahr.`, <br>`1 plus 1 gleich 2`            |
| Buchstabe  | A 4-byte UTF-8 encoded character    | *0* - *65535*                                                                      | A utf8 character between single quotes, e.g. 'a' or '\n'          | `Der Buchstabe x ist 'd'.`                                    |
| Text       | A utf-8 encoded sequence of letters | *any size*                                                                         | Any number of letters between <br>quotation marks, e.g. "Hello\n" | `Der Text x ist "abc".`, <br>`"Hallo" verkettet mit " du da"` |

### Unlike in other programming languages:

* **The decimal separator is not a period, but a comma!**
* **There is no null/nil type**

### Escape sequences
Escape sequences are character combinations that stand for characters
which one cannot normally write in text or letter literals.

| name | escape sequence | ASCII code |
|-----------------|------| ---|
| newline         | `\n` | 10 |
| carriage return | `\r` | 13 |
| regression      | `\b` |  8 |
| Tab             | `\t` |  9 |
| Acoustic signal | `\a` |  7 |
| backslash       | `\\` | 92 |
| Single quote*   | `\'` | 39 |
| double quote**  | `\"` | 34 |

*: Only within a letter literal.\
**: Only within a text literal.

***

## Lists

Lists are collections of values of any size.
Since DDP is statically typed, a list can only contain values of the same data type.
The type name of a list is generally the element type name declined with *list* appended (number -> number list, text -> text list).
A list can grow and shrink at runtime.
How to work with lists is described in the article Operators under [List Operators](?p=Programmierung/Operatoren#listen-und-text-operatoren).

### List literals

A list literal looks similar to an ordinary enumeration, but wrapped in a small phrase to avoid programming ambiguity.
The general form is: `eine[r] Liste, die aus x[, y, z] besteht`.
`eine` or `ein` can be used depending on the grammatical context and after the `aus` one or more must be used
Expressions of the same type follow.

Empty lists can be created like this: `eine[r] leere[n] (Typname) Liste`.

Lists with a single element can also be created simply by using `(Wert) als (Typname)`.

### Examples
```ddp
Die Zahlen Liste z ist eine Liste, die aus 1, 2, 3 besteht.
Die Text Liste t ist eine Liste, die aus "Hallo", "Welt" besteht.

Die Zahlen Liste z2 ist eine leere Zahlen Liste.
Die Text Liste t2 ist "Hallo" als Text Liste.
```

| type name         | Description          | literal                           | Example                                                                     |
|-------------------|----------------------|-----------------------------------|-----------------------------------------------------------------------------|
| Zahlen Liste      | A list of numbers    | A list literal as described above | `Die Zahlen Liste z ist eine Liste, die aus 1, 2, 3 besteht.`               |
| Kommazahlen Liste | A list of floats     | A list literal as described above | `Die Kommazahlen Liste z ist eine Liste, die aus 1,2, 3,2, 3,1415 besteht.` |
| Boolean Liste     | A list of booleans   | A list literal as described above | `Die Boolean Liste b ist eine Liste, die aus wahr, falsch, wahr besteht.`   |
| Buchstaben Liste  | A list of characters | A list literal as described above | `Die Buchstaben Liste b ist eine Liste, die aus 'b', 'h', 'z' besteht.`     |
| Text Liste        | A list of strings	   | A list literal as described above | `Die Text Liste t ist eine Liste, die aus "Hallo", "du", "da" besteht.`     |

***

## Remark

Actually, one would expect that an 'and' would have to occur in the enumeration of a list literal (`eine Liste, die aus 1, 2 und 3 besteht`). However, this would lead to ambiguity in Boolean expressions (`eine Liste, die aus wahr und falsch besteht`), and since the enumeration grammatically does not need an 'and', literals are omitted in lists.