+++
title = "Operators"
weight = 3
+++

# Mathematical operators

For the sake of simplicity, the *Nummer* and *Kommazahl* types are grouped together as *Numeric* in this article.

To make finding a specific operator easier for readers who already know a programming language the C operator is included in this and later tables.

## Unary operators

| Function                  | Usage              | C equivalent | Type of operand   | return type | Example                 | Result   |
|---------------------------|--------------------|--------------|-------------------|-------------|-------------------------|----------|
| NOT gate                  | `logisch nicht a`  | `~a`         | Zahl              | Zahl        | `logisch nicht 1`       | -2       |
| list/Text/Character count | `die Länge von a`  | -            | Liste, Text       | Zahl        | `Die Länge von "Hallo"` | 5        |
| Byte Size                 | `die Größe von a`  | `sizeof(a)`  | all               | Zahl        | `Die Größe von 1`       | 8        |
| Absolute value            | `der Betrag von a` | `abs(a)`     | numeric           | numeric     | `der Betrag von -5`     | 5        |

## Binary operators

| Function        | Usage                                   | C equivalent          | Type of 1st operand | Type of 2nd operand | return type | Example                               | Result |
|-----------------|-----------------------------------------|-----------------------|---------------------|---------------------|-------------|---------------------------------------|--------|
| Addition        | `a plus b`                              | `a + b`               | numeric             | numeric             | numeric     | `1 plus 1`                            | 2      |
| Subtraction     | `a minus b`                             | `a - b`               | numeric             | numeric             | numeric     | `1 minus 2`                           | -1     |
| Multiplication  | `a mal b`                               | `a * b`               | numeric             | numeric             | numeric     | `5 mal 3`                             | 15     |
| Division        | `a durch b`                             | `a / b`               | numeric             | numeric             | Kommazahl   | `6 durch 2`                           | 3      |
| Remainder       | `a modulo b`                            | `a % b`               | Zahl                | Zahl                | Zahl        | `16 modulo 12`                        | 4      |
| Exponentiation  | `a hoch b`                              | `pow(a, b)`           | numeric             | numeric             | Kommazahl   | `2 hoch 8`                            | 256    |
| Root            | `die a. Wurzel von b`                   | `pow(a, 1/b)`         | numeric             | numeric             | Kommazahl   | `die 2. Wurzel von 9`                 | 3      |
| Logarithm       | `der Logarithmus von b zur Basis a` | `log10(b) / log10(a)` | numeric             | numeric             | Kommazahl   | `der Logarithmus von 100 zur Basis 10`| 2      |
| Left Bit-Shift  | `a um b Bit nach links verschoben`      | `a << b`              | Zahl                | Zahl                | Zahl        | `7 um 3 Bit nach links verschoben`    | 56     |
| Right Bit-Shift | `a um b Bit nach rechts verschoben`     | `a >> b`              | Zahl                | Zahl                | Zahl        | `70 um 2 Bit nach rechts verschoben`  | 17     |
| AND gate        | `a logisch und b`                       | `a&b`                 | Zahl                | Zahl                | Zahl        | `5 logisch und 2`                     | 0      |
| OR gate         | `a logisch oder b`                      | `a\| b`               | Zahl                | Zahl                | Zahl        | `5 logisch oder 2`                    | 7      |
| XOR gate        | `a logisch kontra b`                    | `a^b`                 | Zahl                | Zahl                | Zahl        | `8 logisch kontra 5`                  | 13     |

# Boolean operators

With the help of Boolean operators, complex conditions can be expressed and summarized and for example used for branches or loops.

| Operator | Description | C equivalent | Example | Result |
| -------- | ----------- | ------------ | ------- | ------ |
| und      | True if both arguments are true. | `true && false` | `wahr und wahr`<br>`wahr und falsch`<br>`falsch und wahr`<br>`falsch und falsch` | `wahr`<br>`falsch`<br>`falsch`<br>`falsch` |
| oder     | True if either argument is true. | `true \|\| false` | `wahr oder wahr`<br>`wahr oder falsch`<br>`falsch oder wahr`<br>`falsch oder falsch` | `wahr`<br>`wahr`<br>`wahr`<br>`falsch` |
| nicht    | The value of the argument is reversed. | `!true` | `nicht wahr` <br>`nicht falsch` | `falsch`<br>`wahr` |

# Comparison operators

| Operator          | Description                                                               | C equivalent | Example                                          | Result |
|-------------------|---------------------------------------------------------------------------|-----------|----------------------------------------------------------------------------------------------|------------------------------|
| gleich            | True if both arguments have the same value.                               | `1 == 1`  | `1 gleich 1 ist`<br>`1 gleich 2 ist`                                                         | `wahr`<br>`falsch`           |
| ungleich          | True if the two arguments have different values.                          | `1 != 1`  | `1 ungleich 1 ist`<br>`1 ungleich 2 ist`                                                     | `wahr`<br>`falsch`           |
| kleiner als       | True if the left argument has a smaller value than the right.             | `5 < 10`  | `5 kleiner als 10 ist`<br>`30 kleiner als 15 ist`                                            | `wahr`<br>`falsch`           |
| größer als        | True if the left argument has a greater value than the right.             | `7 > 3`   | `7 größer als 3 ist`<br>`5 größer als 8 ist`                                                 | `wahr`<br>`falsch`           |
| kleiner als, oder | True if the left argument is less than or the same value as the right.    | `5 <= 10` | `5 kleiner als, oder 10 ist`<br>`30 kleiner als, oder 15 ist`<br>`5 kleiner als, oder 5 ist` | `wahr`<br>`falsch`<br>`wahr` |
| größer als, oder  | True if the left argument is greater than or the same value as the right. | `7 >= 3`  | `7 größer als, oder 3 ist`<br>`5 größer als, oder 8 ist`<br>`5 größer als, oder 5 ist`       | `wahr`<br>`falsch`<br>`wahr` |

Relational operators all have an "ist" at the end to conform to the grammar in any context.

# List and text operators

| Function          | Usage                      | C equivalent | Type of 1st operand  | Type of 2nd operand  | Type of 3rd operand | return type           | Example                               | Result         |
|-------------------|----------------------------|--------------|----------------------|----------------------|---------------------|-----------------------|---------------------------------------|----------------|
| Concationation    | `a verkettet mit b`        | -            | Text/Liste/Buchstabe | Text/Liste/Buchstabe | -                   | Text/Liste            | `"Hallo" verkettet mit " Welt"`       | `"Hallo Welt"` |
| Indexing          | `a an der Stelle b`        | `a[b]`       | Text/Liste           | Zahl                 | -                   | Buchstabe/Element Typ | `"Hallo" an der Stelle 1`             | 'H'            |
| Range             | `a im Bereich von b bis c` | -            | Text/Liste           | Zahl                 | Zahl                | Text/Liste            | `"Hallo Welt" im Bereich von 1 bis 5` | "Hallo"        |
| `... ab dem ...`  | `a ab dem b. Element`      | -            | Text/Liste           | Zahl                 | -                   | Text/Liste            | `"Hallo Welt" ab dem 7. Element`      | "Welt"         |
| `... bis zum ...` | `a bis zum b. Element`     | -            | Text/Liste           | Zahl                 | -                   | Text/Liste            | `"Hallo Welt" bis zum 5. Element`     | "Hallo"        |

## Remarks

The operand types of a chain cannot be combined in any way.

- If you concatenate any 2 values of the same type, a list is created.
- If you concatenate 2 texts, a new text is created.
- If you concatenate a text with a letter or vice versa, a text is created.
- Concatenating a list and any value of the element type of the list or vice versa creates a list.

With indices and `von ... bis` the indices are always inclusive and start with 1.
So a list from 1 to 5 is the first, the fifth and everything in between.
A list at position 0 would be a run-time error, as would an index that exceeds the length of the list.

With `von ... bis` there are no runtime errors if indices are too small or too large.
The indices are automatically brought into the range [1, length of the list/text].
If the 2nd index is then smaller than the 1st, there is a runtime error.
So if both indices are greater than the length of the list, the result is the last element of the list.

## Examples

```ddp
Die Zahlen Liste z ist eine Liste, die aus 1, 2, 3, 4, 5, 6, 7, 8, 9, 10 besteht.
Schreibe (z an der Stelle 5). [Zeigt 5 in der Konsole]
Schreibe (z von 3 bis 7). [Zeigt "3, 4, 5, 6, 7" in der Konsole]
Schreibe (z von 3 bis 7 verkettet mit z von 1 bis 4). [Zeigt "3, 4, 5, 6, 7, 1, 2, 3, 4" in der Konsole]

Schreibe ("Hallo" verkettet mit 'Ü'). [Zeigt "HalloÜ" in der Konsole]
Schreibe ("Hallo" von 1 bis 3 verkettet mit 'ö'). [Zeigt "Halö" in der Konsole]
Schreibe ('b' verkettet mit 'a'). [Zeigt "b, a" in der Konsole]
```

# Operator prioritization

Different operators are prioritized differently.
This means that some operators take precedence over others when evaluated.
For example, in mathematics, multiplication and division take precedence over addition and subtraction, and
Powers take precedence over multiplication and division.
With the mathematical operators, DDP largely adheres to the mathematical order.
Here is a table with all operators and their prioritization (high priority operators above):

| Rank | Operator                                   |
| ---- | ------------------------------------------ |
| 1    | Function call                              |
| 2    | Literals, Constants                        |
| 3    | Grouping                                   |
| 4    | Type conversions                           |
| 5    | Field access on Combinations               |
| 6    | Indexing                                   |
| 7    | `von ... bis`, `ab ... dem`, `bis ... zum` |
| 8    | Exponentiation, Root, Logarithm            |
| 9    | Negation,                                  |
| 10   | Absolute Value, Size, Length, NOT Gate     |
| 11   | Multiplication, Division, Modulo           |
| 12   | Addition, Subtraction, Concatination       |
| 13   | Bit-Shifting                               |
| 14   | Comparison (`kleiner`, `größer`, ect.)     |
| 15   | Equivalence (`gleich` und `ungleich`)      |
| 16   | Logical AND Gate                           |
| 17   | Logical XOR Gate                           |
| 18   | Logical OR Gate                            |
| 19   | Boolean AND                                |
| 20   | Boolean OR                                 |

## Operator prioritization example

```ddp
Die Zahlen Liste z ist eine Liste, die aus 1, 2, 3 besteht.
Schreibe (z an der Stelle 2 hoch 3). [writes 8 to the Konsole]
```

Here you can see the prioritization of the `an der Stelle` operator over the `hoch` operator.
First `z an der Stelle 2` and then the result is raised to the power of 3.
If the `hoch` operator took precedence over `an der Stelle`, `2 hoch 3` would be calculated first
and then uses the element at position 8 of z, which would result in a runtime error because z
has only 3 elements.