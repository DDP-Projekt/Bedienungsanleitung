+++
title = "Operators"
weight = 3
+++

# Mathematical operators

For the sake of simplicity, the *Zahl*, *Byte* and *Kommazahl* types are grouped together as *numeric* in this article.

To make finding a specific operator easier for readers who already know a programming language, the C operator is included in this and later tables.

## Unary operators

| Function                | Usage                                        | C equivalent | Type of operand | return type            | Example                           | Result |
| ----------------------- | -------------------------------------------- | ------------ | --------------- | ---------------------- | --------------------------------- | ------ |
| NOT gate                | `logisch nicht a`                            | `~a`         | Zahl, Byte      | Zahl, Byte             | `logisch nicht 1`                 | -2     |
| List/Text element count | `die Länge von a`                            | -            | Liste, Text     | Zahl                   | `Die Länge von "Hello"`           | 5      |
| Byte size               | `die Größe von einem/einer <type name>`        | `sizeof(a)`  | type name       | Zahl                   | `Die Größe von einer Zahl`        | 8      |
| Default value           | `der Standardwert von einem/einer <type name>` | -            | type name       | matching the type name | `Der Standardwert von einer Zahl` | 0      |
| Negation                | `-a`                                         | `-a`         | numeric         | Zahl, Kommazahl        | `-(2 plus 3)`                     | -5     |
| Absolute value          | `der Betrag von a`                           | `abs(a)`     | numeric         | numeric                | `Der Betrag von -5`               | 5      |

## Binary operators

| Function        | Usage                               | C equivalent          | Type of 1st operand | Type of 2nd operand | return type | Example                                | Result |
| --------------- | ----------------------------------- | --------------------- | ------------------- | ------------------- | ----------- | -------------------------------------- | ------ |
| Addition        | `a plus b`                          | `a + b`               | numeric             | numeric             | numeric     | `1 plus 1`                             | 2      |
| Subtraction     | `a minus b`                         | `a - b`               | numeric             | numeric             | numeric     | `1 minus 2`                            | -1     |
| Multiplication  | `a mal b`                           | `a * b`               | numeric             | numeric             | numeric     | `5 mal 3`                              | 15     |
| Division        | `a durch b`                         | `a / b`               | numeric             | numeric             | Kommazahl   | `6 durch 2`                            | 3,0    |
| Remainder       | `a modulo b`                        | `a % b`               | Zahl, Byte          | Zahl, Byte          | Zahl, Byte  | `16 modulo 12`                         | 4      |
| Exponentiation  | `a hoch b`                          | `pow(a, b)`           | numeric             | numeric             | Kommazahl   | `2 hoch 8`                             | 256,0  |
| Root            | `die a. Wurzel von b`               | `pow(b, 1/a)`         | numeric             | numeric             | Kommazahl   | `die 2. Wurzel von 9`                  | 3,0    |
| Logarithm       | `der Logarithmus von b zur Basis a` | `log10(b) / log10(a)` | numeric             | numeric             | Kommazahl   | `der Logarithmus von 100 zur Basis 10` | 2,0    |
| Left bit shift  | `a um b Bit nach links verschoben`  | `a << b`              | Zahl, Byte          | Zahl, Byte          | Zahl, Byte  | `7 um 3 Bit nach links verschoben`     | 56     |
| Right bit shift | `a um b Bit nach rechts verschoben` | `a >> b`              | Zahl, Byte          | Zahl, Byte          | Zahl, Byte  | `70 um 2 Bit nach rechts verschoben`   | 17     |
| AND gate        | `a logisch und b`                   | `a&b`                 | Zahl, Byte          | Zahl, Byte          | Zahl, Byte  | `5 logisch und 2`                      | 0      |
| OR gate         | `a logisch oder b`                  | `a\| b`               | Zahl, Byte          | Zahl, Byte          | Zahl, Byte  | `5 logisch oder 2`                     | 7      |
| XOR gate        | `a logisch kontra b`                | `a^b`                 | Zahl, Byte          | Zahl, Byte          | Zahl, Byte  | `8 logisch kontra 5`                   | 13     |

# Boolean operators

With the help of boolean operators, complex conditions can be expressed and combined, and used for example in branches or loops.

| Operator            | Description                                  | C equivalent      | Example                                                                                                                  | Result                                     |
| ------------------- | -------------------------------------------- | ----------------- | ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------ |
| und                 | True if both arguments are true.             | `true && false`   | `wahr und wahr`<br>`wahr und falsch`<br>`falsch und wahr`<br>`falsch und falsch`                                         | `wahr`<br>`falsch`<br>`falsch`<br>`falsch` |
| oder                | True if either argument is true.             | `true \|\| false` | `wahr oder wahr`<br>`wahr oder falsch`<br>`falsch oder wahr`<br>`falsch oder falsch`                                     | `wahr`<br>`wahr`<br>`wahr`<br>`falsch`     |
| `entweder ..., oder` | True if *only* one of the arguments is true. | `true != false`   | `entweder wahr, oder wahr`<br>`entweder wahr, oder falsch`<br>`entweder falsch, oder wahr`<br>`entweder falsch, oder falsch` | `falsch`<br>`wahr`<br>`wahr`<br>`falsch`   |
| nicht               | The value of the argument is reversed.       | `!true`           | `nicht wahr` <br>`nicht falsch`                                                                                          | `falsch`<br>`wahr`                         |

# Comparison operators

| Operator          | Description                                                               | C equivalent       | Example                                                                                      | Result                       |
| ----------------- | ------------------------------------------------------------------------- | ------------------ | -------------------------------------------------------------------------------------------- | ---------------------------- |
| gleich            | True if both arguments have the same value.                               | `1 == 1`           | `1 gleich 1 ist`<br>`1 gleich 2 ist`                                                         | `wahr`<br>`falsch`           |
| ungleich          | True if the two arguments have different values.                          | `1 != 1`           | `1 ungleich 1 ist`<br>`1 ungleich 2 ist`                                                     | `falsch`<br>`wahr`           |
| kleiner als       | True if the left argument has a smaller value than the right.             | `5 < 10`           | `5 kleiner als 10 ist`<br>`30 kleiner als 15 ist`                                            | `wahr`<br>`falsch`           |
| größer als        | True if the left argument has a greater value than the right.             | `7 > 3`            | `7 größer als 3 ist`<br>`5 größer als 8 ist`                                                 | `wahr`<br>`falsch`           |
| kleiner als, oder | True if the left argument is less than or the same value as the right.    | `5 <= 10`          | `5 kleiner als, oder 10 ist`<br>`30 kleiner als, oder 15 ist`<br>`5 kleiner als, oder 5 ist` | `wahr`<br>`falsch`<br>`wahr` |
| größer als, oder  | True if the left argument is greater than or the same value as the right. | `7 >= 3`           | `7 größer als, oder 3 ist`<br>`5 größer als, oder 8 ist`<br>`5 größer als, oder 5 ist`       | `wahr`<br>`falsch`<br>`wahr` |
| zwischen          | True if the left argument lies between the two right ones.                | `(4 < 5 && 4 > 3)` | `4 zwischen 3 und 5 ist`<br>`6 zwischen 3 und 5 ist`                                         | `wahr`<br>`falsch`           |

Comparison operators all have an "ist" at the end to conform to the grammar in any context.
If that leads to several "ist"s in a row, a single one is enough:
```ddp
2 größer als 2 ist gleich 4 größer als 4 ist ist. [ugly]
2 größer als 2 ist gleich 4 größer als 4 ist. [nice]
```

# Type check

For values of type [Variable](/en/Programmierung/Datentypen#variable) you can check which type the stored value currently has.
To do this, you write `ein`, `eine`, `kein` or `keine` and the type name. The article must match the type.

| Operator           | Description                                    | Example                                                  | Result             |
| ------------------ | ---------------------------------------------- | -------------------------------------------------------- | ------------------ |
| `ein(e) ... ist`   | True if the value has this type.               | `(5 als Variable) eine Zahl ist`<br>`(5 als Variable) ein Text ist` | `wahr`<br>`falsch` |
| `kein(e) ... ist`  | True if the value does *not* have this type.   | `(5 als Variable) keine Zahl ist`<br>`(5 als Variable) kein Text ist` | `falsch`<br>`wahr` |

# List and text operators

| Function          | Usage                      | C equivalent | Type of 1st operand  | Type of 2nd operand  | Type of 3rd operand | return type            | Example                               | Result         |
| ----------------- | -------------------------- | ------------ | -------------------- | -------------------- | ------------------- | ---------------------- | ------------------------------------- | -------------- |
| Concatenation     | `a verkettet mit b`        | -            | Text/Liste/Buchstabe | Text/Liste/Buchstabe | -                   | Text/Liste             | `"Hello" verkettet mit " World"`       | `"Hello World"` |
| Indexing          | `a an der Stelle b`        | `a[b]`       | Text/Liste           | Zahl, Byte           | -                   | Buchstabe/element type | `"Hello" an der Stelle 1`             | 'H'            |
| Range             | `a im Bereich von b bis c` | -            | Text/Liste           | Zahl, Byte           | Zahl, Byte          | Text/Liste             | `"Hello World" im Bereich von 1 bis 5` | "Hello"        |
| `... ab dem ...`  | `a ab dem b. Element`      | -            | Text/Liste           | Zahl, Byte           | -                   | Text/Liste             | `"Hello World" ab dem 7. Element`      | "World"         |
| `... bis zum ...` | `a bis zum b. Element`     | -            | Text/Liste           | Zahl, Byte           | -                   | Text/Liste             | `"Hello World" bis zum 5. Element`     | "Hello"        |

## Remarks

The operand types of a concatenation cannot be combined arbitrarily.

- If you concatenate any 2 values of the same type, a list is created.
- If you concatenate 2 texts, a new text is created.
- If you concatenate a text with a letter or vice versa, a text is created.
- Concatenating a list and any value of the element type of the list or vice versa creates a list.

With indexing and `im Bereich von ... bis` the indices are always inclusive and start at 1.
So a list from 1 to 5 contains the first, the fifth and all elements in between.
A list at position 0 would be a runtime error, as would an index that exceeds the length of the list.

With `im Bereich von ... bis` there are no runtime errors if indices are too small or too large.
The indices are automatically brought into the range [1, length of the list/text].
If the 2nd index is then smaller than the 1st, there is a runtime error.
So if both indices are greater than the length of the list, the result is the last element of the list.

`... ab dem ...` and `... bis zum ...` are just short forms of `im Bereich von ... bis` with the second or first operand
set to the length of the list or 1 respectively.

## Examples

```ddp
Die Zahlen Liste z ist eine Liste, die aus 1, 2, 3, 4, 5, 6, 7, 8, 9, 10 besteht.
Schreibe (z an der Stelle 5). [writes 5 to the console]
Schreibe (z im Bereich von 3 bis 7). [writes "3, 4, 5, 6, 7" to the console]
Schreibe (z im Bereich von 3 bis 7 verkettet mit z im Bereich von 1 bis 4). [writes "3, 4, 5, 6, 7, 1, 2, 3, 4" to the console]

Schreibe ("Hello" verkettet mit 'Ü'). [writes "HelloÜ" to the console]
Schreibe ("Hello" im Bereich von 1 bis 3 verkettet mit 'ö'). [writes "Helö" to the console]
Schreibe ('b' verkettet mit 'a'). [writes "b, a" to the console]
```

# Falls

The `Falls` operator corresponds to the ternary operator (`a ? b : c`) in other languages.
Its parameters are a Wahrheitswert and 2 expressions of any (but the same) type:
```ddp
Der Text t ist "2 > 3", falls 2 größer als 3 ist, ansonsten "2" verkettet mit " < " verkettet mit "3".
```

# Operator prioritization

Different operators are prioritized differently.
This means that some operators take precedence over others when evaluated.
For example, in mathematics, multiplication and division take precedence over addition and subtraction, and
powers in turn take precedence over multiplication and division.
With the mathematical operators, DDP largely adheres to the mathematical order.
Here is a table with all operators and their prioritization (high priority operators at the top):

| Rank | Operator                                                              |
| ---- | --------------------------------------------------------------------- |
| 1    | Function call                                                         |
| 2    | Literals, constants                                                   |
| 3    | Parentheses                                                           |
| 4    | Type conversions                                                      |
| 5    | Field access on combinations                                          |
| 6    | Indexing                                                              |
| 7    | `im Bereich von ... bis`, `ab dem ... Element`, `bis zum ... Element` |
| 8    | Exponentiation, root, logarithm                                       |
| 9    | Negation                                                              |
| 10   | Absolute value, size, length, default value, logical/boolean NOT      |
| 11   | Multiplication, division, remainder                                   |
| 12   | Addition, subtraction, concatenation                                  |
| 13   | Bit shift                                                             |
| 14   | Comparison (`kleiner`, `größer`, etc.)                                |
| 15   | Equality (`gleich` and `ungleich`), type check (`ein ... ist`)        |
| 16   | Logical AND gate                                                      |
| 17   | Logical XOR gate                                                      |
| 18   | Logical OR gate                                                       |
| 19   | Boolean AND                                                           |
| 20   | Boolean OR                                                            |
| 21   | `entweder ..., oder`                                                  |
| 22   | Falls                                                                 |

## Operator prioritization example

```ddp
Die Zahlen Liste z ist eine Liste, die aus 1, 2, 3 besteht.
Schreibe (z an der Stelle 2 hoch 3). [writes 8 to the console]
```

Here you can see the prioritization of the `an der Stelle` operator over the `hoch` operator.
First `z an der Stelle 2` is evaluated and then the result is raised to the power of 3.
If the `hoch` operator took precedence over `an der Stelle`, `2 hoch 3` would be calculated first
and then the element at position 8 of z would be used, which would result in a runtime error, since z
only has 3 elements.

# Operator overloading

Operators can also be overloaded to execute functions instead of operators while still
profiting from operator prioritization.
More about this in the article [Operator overloading](/en/Programmierung/Funktionen/Operatoren-Ueberladung).
