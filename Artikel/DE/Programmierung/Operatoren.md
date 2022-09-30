# Mathematische Operatoren
Zur Vereinfachung werden in diesem Artikel die Typen *Zahl* und *Kommazahl* als *Numerisch* zusammengefasst.

Um das Finden eines bestimmten Operators für Leser, die bereits eine Programmiersprache beherschen einfacher zu machen ist
in diesen und späteren Tabellen jeweils der C-Operator dabei.

## Unäre Operatoren
| Funktion                              | Verwendung        | C Equivalent | Typ vom Operanden | Rückgabetyp | Beispiel                | Ergebnis |
|---------------------------------------|-------------------|--------------|-------------------|-------------|-------------------------|----------|
| Logische NICHT verknüpfung            | `logisch nicht a` | `~a`         | Zahl              | Zahl        | `logisch nicht 1`       | -2       |
| Listen/Text Element/Buchstaben Anzahl | `die Länge von a` | -            | Liste, Text       | Zahl        | `Die Länge von "Hallo"` | 5        |
| Byte Größe                            | `die Größe von a` | `sizeof(a)`  | alles             | Zahl        | `Die Größe von 1`       | 8        |

## Binäre Operatoren
| Funktion                     | Verwendung                              | C Equivalent          | Typ vom 1. Operand | Typ vom 2. Operand | Rückgabetyp    | Beispiel                              | Ergebnis |
|------------------------------|-----------------------------------------|-----------------------|--------------------|--------------------|----------------|---------------------------------------|----------|
| Addition                     | `a plus b`                              | `a + b`               | numerisch          | numerisch          | numerisch      | `1 plus 1`                            | 2    |
| Subtraktion                  | `a minus b`                             | `a - b`               | numerisch          | numerisch          | numerisch      | `1 minus 2`                           | -1   |
| Multiplikation               | `a mal b`                               | `a * b`               | numerisch          | numerisch          | numerisch      | `5 mal 3`                             | 15   |
| Division                     | `a durch b`                             | `a / b`               | numerisch          | numerisch          | Kommazahl      | `6 durch 2`                           | 3    |
| Rest                         | `a modulo b`                            | `a % b`               | Zahl               | Zahl               | Zahl           | `16 modulo 12`                        | 4    |
| Potenzieren                  | `a hoch b`                              | `pow(a, b)`           | numerisch          | numerisch          | Kommazahl      | `2 hoch 8`                            | 256  |
| Wurzelziehen                 | `die a. Wurzel von b`                   | `pow(a, 1/b)`         | numerisch          | numerisch          | Kommazahl      | `die 2. Wurzel von 9`                 | 3    |
| Logarithmus                  | `der Logarithmus mit der Basis a von b` | `log10(b) / log10(a)` | numerisch          | numerisch          | Kommazahl      | `der Logarithmus von 100 zur Basis 10` | 2    |
| Bit-Verschiebung nach links  | `a um b Bit nach links verschoben`      | `a << b`              | Zahl               | Zahl               | Zahl           | `7 um 3 Bit nach links verschoben`    | 56   |
| Bit-Verschiebung nach rechts | `a um b Bit nach rechts verschoben`     | `a >> b`              | Zahl               | Zahl               | Zahl           | `70 um 2 Bit nach rechts verschoben`  | 17   |
| Logische UND verknüpfung     | `a logisch und b`                       | `a&b`                 | Zahl               | Zahl               | Zahl           | `5 logisch und 2`                     | 0  |
| Logische ODER verknüpfung    | `a logisch oder b`                      | `a\| b`                 | Zahl               | Zahl           | Zahl             | `5 logisch oder 2`                  |7   |
| Logische XOR verknüpfung     | `a logisch kontra b`                    | `a^b`                 | Zahl               | Zahl               | Zahl           | `8 logisch kontra 5`                  | 13   |

## Zuweisungsoperatoren

Mathematische Zuweisungsoperatoren sind Operatoren, die das Ergebnis einer Mathematischen operation direkt einer Variable zuweist.

### Addition
- `Addiere <a> mit <b> und speichere das Ergebnis in <variable>`
- `Erhöhe <variable> um <a>`

### Subktraktion
- `Subtrahiere <a> von <b> und speichere das Ergebnis in <variable>`
- `Verringere <variable> um <a>`

### Multiplikation

- `Multipliziere <a> mit <b> und speichere das Ergebnis in <variable>.`
- `Vervielfache <variable> um <a>.`

### Division
- `Dividiere <a> durch <b> und speichere das Ergebnis in <variable>.`
- `Teile <variable> durch <a>.`

# Bool'sche Operatoren

Mithilfe von Bool'schen Operatoren können komplexe Bedingungen ausgedrückt und zusammengefasst werden und z.B. für Verzweigungen oder Schleifen verwendet werden.

| Operator | Beschreibung | C Equivalent | Beispiel | Ergebnis |
| -------- | ------------ | ------------ | -------- | -------- |
| und | Wahr, wenn beide Argumente wahr sind. | `true && false` | `wahr und wahr`<br>`wahr und falsch`<br>`falsch und wahr`<br>`falsch und falsch` | `wahr`<br>`falsch`<br>`falsch`<br>`falsch` |
| oder | Wahr, wenn eines der Beiden Argumente wahr ist. | `true \|\| false` | `wahr oder wahr`<br>`wahr oder falsch`<br>`falsch oder wahr`<br>`falsch oder falsch` | `wahr`<br>`wahr`<br>`wahr`<br>`falsch` |
| nicht | Der Wert des Arguments wird umgekehrt. | `!true` | `nicht wahr` <br>`nicht falsch` | `falsch`<br>`wahr` |

# Vergleichsoperatoren

| Operator          | Beschreibung                                                                          | C Equivalent | Beispiel                                                                         | Ergebnis                     |
|-------------------|---------------------------------------------------------------------------------------|--------------|----------------------------------------------------------------------------------|------------------------------|
| gleich            | Wahr, wenn beide Argumente den gleichen Wert haben.                                   | `1 == 1`     | `1 gleich 1`<br>`1 gleich 2`                                                     | `wahr`<br>`falsch`           |
| ungleich          | Wahr, wenn die Beiden Argumente verschiedene Werte haben.                             | `1 != 1`     | `1 ungleich 1`<br>`1 ungleich 2`                                                 | `wahr`<br>`falsch`           |
| kleiner als       | Wahr, wenn das Linke Argument einen kleineren Wert als das Rechte hat.                | `5 < 10`     | `5 kleiner als 10`<br>`30 kleiner als 15`                                        | `wahr`<br>`falsch`           |
| größer als        | Wahr, wenn das Linke Argument einen größeren Wert als das Rechte hat.                 | `7 > 3`      | `7 größer als 3`<br>`5 größer als 8`                                             | `wahr`<br>`falsch`           |
| kleiner als, oder | Wahr, wenn das Linke Argument einen kleineren oder denselben Wert wie das Rechte hat. | `5 <= 10`    | `5 kleiner als, oder 10`<br>`30 kleiner als, oder 15`<br>`5 kleiner als, oder 5` | `wahr`<br>`falsch`<br>`wahr` |
| größer als, oder  | Wahr, wenn das Linke Argument einen größeren oder denselben Wert wie das Rechte hat.  | `7 >= 3`     | `7 größer als, oder 3`<br>`5 größer als, oder 8`<br>`5 größer als, oder 5`       | `wahr`<br>`falsch`<br>`wahr` |

# Listen und Text Operatoren
<to-do></to-do>

# Operator Reihenfolge

| Rang | Operator                                          |
|------|---------------------------------------------------|
| 1    | Funktionsaufruf                                   |
| 2    | Klammern                                          |
| 3    | Typkonvertierungen                                |
| 4    | Indizierung, `von ... bis`                        |
| 5    | Literale, Konstante                               |
| 6    | Potenzieren, Wurzelziehen, Logarithmus            |
| 7    | Negation,                                         |
| 8    | Betrag, Größe, Länge, Logisches-/Bool'sches nicht |
| 9    | Multiplikation, Division, Rest,                   |
| 10   | Addition, Subtrakion, Verkettung,                 |
| 11   | Bit-Verschiebung                                  |
| 12   | Größen Vergleich (kleiner, größer, ect.)          |
| 13   | Gleichheit (gleich und ungleich)                  |
| 14   | Logische UND Verknüpfung                          |
| 15   | Logische XOR Verknüpfung                          |
| 16   | Logische ODER Verknüpfung                         |
| 17   | Bool'sches UND                                    |
| 18   | Bool'sches ODER                                   |