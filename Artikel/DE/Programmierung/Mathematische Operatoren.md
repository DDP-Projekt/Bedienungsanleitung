# Mathematische Operatoren
Zur Vereinfachung werden in diesem Artikel die Typen *Zahl* und *Kommazahl* als *Numerisch* zusammengefasst.

Um das Finden eines bestimmten Operators für Leser, die bereits eine Programmiersprache beherschen einfacher zu machen ist
in diesen und späteren Tabellen jeweils der C-Operator dabei.

## Unäre Operatoren
| Funktion                              | Verwendung        | C Equivalent | Typ vom Operanden | Rückgabetyp | Beispiel                | Ergebnis |
|---------------------------------------|-------------------|--------------|-------------------|-------------|-------------------------|----------|
| Logische NICHT verknüpfung            | `logisch nicht a` | `~a`         | Zahl              | Zahl        | `logisch nicht 1        | -2       |
| Listen/Text Element/Buchstaben Anzahl | `die Länge von a` | -            | Liste, Text       | Zahl        | `Die Länge von "Hallo"` | 5        |
| Byte Größe                            | `die Größe von a` | `sizeof(a)`  | alles             | Zahl        | `Die Größe von 1`       | 8        |

## Binäre Operatoren
|Funktion|Verwendung|C Equivalent|Typ vom 1. Operand| Typ vom 2. Operand|Rückgabetyp|Beispiel|Ergebnis|
|--------|----------|------------|------------------|-------------------|-----------|--------|--|
|Addition|`a plus b`|`a + b`|numerisch|numerisch|numerisch|`1 plus 1`|2|
|Subtraktion|`a minus b`|`a - b`|numerisch|numerisch|numerisch|`1 minus 2`|-1|
|Multiplikation|`a mal b`|`a * b`|numerisch|numerisch|numerisch|`5 mal 3`|15|
|Division|`a durch b`|`a / b`|numerisch|numerisch|Kommazahl|`6 durch 2`|3|
|Rest|`a modulo b`|`a % b`|Zahl|Zahl|Zahl|`16 modulo 12`|4|
|Exponentiation|`a hoch b`|`pow(a, b)`|numerisch|numerisch|Kommazahl|`2 hoch 8`|256|
|Wurzelziehen|`die a. Wurzel von b`|`pow(a, 1/b)`|numerisch|numerisch|Kommazahl|`die 2. Wurzel von 9`|3|
|Logarithmus|`der Logarithmus mit der Basis a von b`|`log10(b) / log10(a)`|numerisch|numerisch|Kommazahl|`der Logarithmus von 100 zur Basis 10`|2|
|Bit verschiebung nach links|`a um b Bit nach links verschoben`|`a << b`|Zahl|Zahl|Zahl|`7 um 3 Bit nach links verschoben`|56|
|Bit verschiebung nach Rechts|`a um b Bit nach rechts verschoben`|`a >> b`|Zahl|Zahl|Zahl|`70 um 2 Bit nach rechts verschoben`|17|
|Logische UND verknüpfung|`a logisch und b`|`a&b`|Zahl|Zahl|Zahl|`5 logisch und 2`|0|
|Logische ODER verknüpfung|`a logisch oder b`|`a\|b`|Zahl|Zahl|Zahl|`5 logisch oder 2`|7|
|Logische XOR verknüpfung|`a logisch kontra b`|`a^b`|Zahl|Zahl|Zahl|`8 logisch kontra 5`|13|