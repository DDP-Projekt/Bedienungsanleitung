+++
title = "Operatoren" 
weight = 3
+++

# Mathematische Operatoren

Zur Vereinfachung werden in diesem Artikel die Typen *Zahl* und *Kommazahl* als *Numerisch* zusammengefasst.

Um das Finden eines bestimmten Operators für Leser, die bereits eine Programmiersprache beherschen einfacher zu machen ist
in diesen und späteren Tabellen jeweils der C-Operator dabei.

## Unäre Operatoren

| Funktion                              | Verwendung        | C Equivalent | Typ vom Operanden | Rückgabetyp | Beispiel                | Ergebnis |
| ------------------------------------- | ----------------- | ------------ | ----------------- | ----------- | ----------------------- | -------- |
| Logische NICHT verknüpfung            | `logisch nicht a` | `~a`         | Zahl              | Zahl        | `logisch nicht 1`       | -2       |
| Listen/Text Element/Buchstaben Anzahl | `die Länge von a` | -            | Liste, Text       | Zahl        | `Die Länge von "Hallo"` | 5        |
| Byte Größe                            | `die Größe von a` | `sizeof(a)`  | alles             | Zahl        | `Die Größe von 1`       | 8        |

## Binäre Operatoren

| Funktion                     | Verwendung                              | C Equivalent          | Typ vom 1. Operand | Typ vom 2. Operand | Rückgabetyp | Beispiel                               | Ergebnis |
| ---------------------------- | --------------------------------------- | --------------------- | ------------------ | ------------------ | ----------- | -------------------------------------- | -------- |
| Addition                     | `a plus b`                              | `a + b`               | numerisch          | numerisch          | numerisch   | `1 plus 1`                             | 2        |
| Subtraktion                  | `a minus b`                             | `a - b`               | numerisch          | numerisch          | numerisch   | `1 minus 2`                            | -1       |
| Multiplikation               | `a mal b`                               | `a * b`               | numerisch          | numerisch          | numerisch   | `5 mal 3`                              | 15       |
| Division                     | `a durch b`                             | `a / b`               | numerisch          | numerisch          | Kommazahl   | `6 durch 2`                            | 3,0      |
| Rest                         | `a modulo b`                            | `a % b`               | Zahl               | Zahl               | Zahl        | `16 modulo 12`                         | 4        |
| Potenzieren                  | `a hoch b`                              | `pow(a, b)`           | numerisch          | numerisch          | Kommazahl   | `2 hoch 8`                             | 256,0    |
| Wurzelziehen                 | `die a. Wurzel von b`                   | `pow(a, 1/b)`         | numerisch          | numerisch          | Kommazahl   | `die 2. Wurzel von 9`                  | 3,0      |
| Logarithmus                  | `der Logarithmus mit der Basis a von b` | `log10(b) / log10(a)` | numerisch          | numerisch          | Kommazahl   | `der Logarithmus von 100 zur Basis 10` | 2,0      |
| Bit-Verschiebung nach links  | `a um b Bit nach links verschoben`      | `a << b`              | Zahl               | Zahl               | Zahl        | `7 um 3 Bit nach links verschoben`     | 56       |
| Bit-Verschiebung nach rechts | `a um b Bit nach rechts verschoben`     | `a >> b`              | Zahl               | Zahl               | Zahl        | `70 um 2 Bit nach rechts verschoben`   | 17       |
| Logische UND verknüpfung     | `a logisch und b`                       | `a&b`                 | Zahl               | Zahl               | Zahl        | `5 logisch und 2`                      | 0        |
| Logische ODER verknüpfung    | `a logisch oder b`                      | `a\| b`               | Zahl               | Zahl               | Zahl        | `5 logisch oder 2`                     | 7        |
| Logische XOR verknüpfung     | `a logisch kontra b`                    | `a^b`                 | Zahl               | Zahl               | Zahl        | `8 logisch kontra 5`                   | 13       |

# Bool'sche Operatoren

Mithilfe von Bool'schen Operatoren können komplexe Bedingungen ausgedrückt und zusammengefasst werden und z.B. für Verzweigungen oder Schleifen verwendet werden.

| Operator | Beschreibung                                    | C Equivalent      | Beispiel                                                                             | Ergebnis                                   |
| -------- | ----------------------------------------------- | ----------------- | ------------------------------------------------------------------------------------ | ------------------------------------------ |
| und      | Wahr, wenn beide Argumente wahr sind.           | `true && false`   | `wahr und wahr`<br>`wahr und falsch`<br>`falsch und wahr`<br>`falsch und falsch`     | `wahr`<br>`falsch`<br>`falsch`<br>`falsch` |
| oder     | Wahr, wenn eines der Beiden Argumente wahr ist. | `true \|\| false` | `wahr oder wahr`<br>`wahr oder falsch`<br>`falsch oder wahr`<br>`falsch oder falsch` | `wahr`<br>`wahr`<br>`wahr`<br>`falsch`     |
| nicht    | Der Wert des Arguments wird umgekehrt.          | `!true`           | `nicht wahr` <br>`nicht falsch`                                                      | `falsch`<br>`wahr`                         |

# Vergleichsoperatoren

| Operator          | Beschreibung                                                                          | C Equivalent | Beispiel                                                                                     | Ergebnis                     |
| ----------------- | ------------------------------------------------------------------------------------- | ------------ | -------------------------------------------------------------------------------------------- | ---------------------------- |
| gleich            | Wahr, wenn beide Argumente den gleichen Wert haben.                                   | `1 == 1`     | `1 gleich 1 ist`<br>`1 gleich 2 ist`                                                         | `wahr`<br>`falsch`           |
| ungleich          | Wahr, wenn die Beiden Argumente verschiedene Werte haben.                             | `1 != 1`     | `1 ungleich 1 ist`<br>`1 ungleich 2 ist`                                                     | `wahr`<br>`falsch`           |
| kleiner als       | Wahr, wenn das Linke Argument einen kleineren Wert als das Rechte hat.                | `5 < 10`     | `5 kleiner als 10 ist`<br>`30 kleiner als 15 ist`                                            | `wahr`<br>`falsch`           |
| größer als        | Wahr, wenn das Linke Argument einen größeren Wert als das Rechte hat.                 | `7 > 3`      | `7 größer als 3 ist`<br>`5 größer als 8 ist`                                                 | `wahr`<br>`falsch`           |
| kleiner als, oder | Wahr, wenn das Linke Argument einen kleineren oder denselben Wert wie das Rechte hat. | `5 <= 10`    | `5 kleiner als, oder 10 ist`<br>`30 kleiner als, oder 15 ist`<br>`5 kleiner als, oder 5 ist` | `wahr`<br>`falsch`<br>`wahr` |
| größer als, oder  | Wahr, wenn das Linke Argument einen größeren oder denselben Wert wie das Rechte hat.  | `7 >= 3`     | `7 größer als, oder 3 ist`<br>`5 größer als, oder 8 ist`<br>`5 größer als, oder 5 ist`       | `wahr`<br>`falsch`<br>`wahr` |

Vergleichsoperatoren haben alle ein "ist" am Ende um der Grammatik in jedem Kontext gerecht zu werden.

# Listen und Text Operatoren

| Funktion              | Verwendung          | C Equivalent | Typ vom 1. Operanden | Typ vom 2. Operanden | Typ vom 3. Operanden | Rückgabetyp           | Beispiel                        | Ergebnis       |
| --------------------- | ------------------- | ------------ | -------------------- | -------------------- | -------------------- | --------------------- | ------------------------------- | -------------- |
| Verkettung            | `a verkettet mit b` | -            | Text/Liste/Buchstabe | Text/Liste/Buchstabe | -                    | Text/Liste            | `"Hallo" verkettet mit " Welt"` | `"Hallo Welt"` |
| Indizierung           | `a an der Stelle b` | `a[b]`       | Text/Liste           | Zahl                 | -                    | Buchstabe/Element Typ | `"Hallo" an der Stelle 1`       | 'H'            |
| `... von ... bis ...` | `a von b bis c`     | -            | Text/Liste           | Zahl                 | Zahl                 | Text/Liste            | `"Hallo Welt" von 1 bis 5`      | "Hallo"        |

## Bemerkungen

Die Operanden Typen einer Verkettung können nicht beliebig kombiniert werden.

- Verkettet man 2 beliebige Werte desselben Typs entsteht eine Liste.
- Verkettet man 2 Texte entsteht ein neuer Text.
- Verkettet man einen Text mit einem Buchstaben oder umgekehrt entsteht ein Text.
- Verkettet man eine Liste und einen beliebigen Wert vom Element Typ der Liste oder umgekehrt entsteht eine Liste.

Bei Indizierungen und `von ... bis` sind die Indizes immer inklusiv und fangen bei 1 an.
Eine Liste von 1 bis 5 sind also das erste, das fünfte und alle Element dazwischen.
Eine Liste an der Stelle 0 wäre ein Laufzeitfehler, genauso wie ein Index, der die Länge der Liste überschreitet.

Bei `von ... bis` gibt es keine Laufzeitfehler bei zu kleinen oder zu großen Indizes.
Die Indizes werden automatisch in den Bereich [1, Länge der Liste/Text] gebracht.
Sollte danach der 2. Index kleiner als der 1. sein, so gibt es einen Laufzeitfehler.
Wenn also beide Indizes größer als die Länge der Liste sind, so ist das Ergebnis also das Letzte Element der Liste.

## Beispiele

```ddp
Die Zahlen Liste z ist eine Liste, die aus 1, 2, 3, 4, 5, 6, 7, 8, 9, 10 besteht.
Schreibe (z an der Stelle 5). [Zeigt 5 in der Konsole]
Schreibe (z von 3 bis 7). [Zeigt "3, 4, 5, 6, 7" in der Konsole]
Schreibe (z von 3 bis 7 verkettet mit z von 1 bis 4). [Zeigt "3, 4, 5, 6, 7, 1, 2, 3, 4" in der Konsole]

Schreibe ("Hallo" verkettet mit 'Ü'). [Zeigt "HalloÜ" in der Konsole]
Schreibe ("Hallo" von 1 bis 3 verkettet mit 'ö'). [Zeigt "Halö" in der Konsole]
Schreibe ('b' verkettet mit 'a'). [Zeigt "b, a" in der Konsole]
```

# Operator Priorisierung

Verschiedene Operatoren werden verschieden Priorisiert.
Das bedeutet, dass manche Operatoren bei der Auswertung Vorrang vor anderen haben.
In der Mathematik haben z.B. Multiplikation und Division Vorrang vor Addition und Subtraktion, und
Potenzen haben wiederrum Vorrang vor Multiplikationen und Divisionen.
Bei den mathemathischen Operatoren hält DDP sich weitestgehend an die mathemathische Reihenfolge.
Hier ist eine Tabelle mit allen Operatoren und ihrer Priorisierung (hoch priorisierte Operatoren oben):

| Rang | Operator                                          |
| ---- | ------------------------------------------------- |
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

## Operator Priorisierung Beispiel

```ddp
Die Zahlen Liste z ist eine Liste, die aus 1, 2, 3 besteht.
Schreibe (z an der Stelle 2 hoch 3). [Zeigt 8 in der Konsole]
```

Hier sieht man die Priorisierung des `an der Stelle` Operators über den `hoch` Operator.
Es wird erst `z an der Stelle 2` ausgewertet und dann das Ergebnis mit 3 potenziert.
Hätte der `hoch` Operator Vorrang vor `an der Stelle`, würde erst `2 hoch 3` ausgerechnet werden
und dann das Element an der Stelle 8 von z benutzt, was zu einem Laufzeitfehler führen würde, da z
nur 3 Elemente besitzt.