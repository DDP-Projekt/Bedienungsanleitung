+++
title = "Operatoren" 
weight = 3
+++

# Mathematische Operatoren

Zur Vereinfachung werden in diesem Artikel die Typen *Zahl*, *Byte* und *Kommazahl* als *Numerisch* zusammengefasst.

Um das Finden eines bestimmten Operators für Leser, die bereits eine Programmiersprache beherrschen einfacher zu machen ist
in diesen und späteren Tabellen jeweils der C-Operator dabei.

## Unäre Operatoren

| Funktion                              | Verwendung                                   | C Equivalent | Typ vom Operanden | Rückgabetyp                 | Beispiel                              | Ergebnis |
| ------------------------------------- | -------------------------------------------- | ------------ | ----------------- | --------------------------- | ------------------------------------- | -------- |
| Logische NICHT verknüpfung            | `logisch nicht a`                            | `~a`         | Zahl, Byte        | Zahl, Byte                  | `logisch nicht 1`                     | -2       |
| Listen/Text Element/Buchstaben Anzahl | `die Länge von a`                            | -            | Liste, Text       | Zahl                        | `Die Länge von "Hallo"`               | 5        |
| Byte Größe                            | `die Größe von einem/einer <Typname>`        | `sizeof(a)`  | Typename          | Zahl                        | `Die Größe von einer Zahl`            | 8        |
| Standardwert                          | `der Standardwert von einem/einer <Typname>` | -            | Typename          | passend zum Typnamen        | `Der Standardwert von einer Zahl`     | 0        |
| Betrag                                | `der Betrag von a`                           | `abs(a)`     | numerisch         | numerisch                   | `Der Betrag von -5`                   | 5        |

## Binäre Operatoren

| Funktion                     | Verwendung                              | C Equivalent          | Typ vom 1. Operand | Typ vom 2. Operand | Rückgabetyp | Beispiel                               | Ergebnis |
| ---------------------------- | --------------------------------------- | --------------------- | ------------------ | ------------------ | ----------- | -------------------------------------- | -------- |
| Addition                     | `a plus b`                              | `a + b`               | numerisch          | numerisch          | numerisch   | `1 plus 1`                             | 2        |
| Subtraktion                  | `a minus b`                             | `a - b`               | numerisch          | numerisch          | numerisch   | `1 minus 2`                            | -1       |
| Multiplikation               | `a mal b`                               | `a * b`               | numerisch          | numerisch          | numerisch   | `5 mal 3`                              | 15       |
| Division                     | `a durch b`                             | `a / b`               | numerisch          | numerisch          | Kommazahl   | `6 durch 2`                            | 3,0      |
| Rest                         | `a modulo b`                            | `a % b`               | Zahl, Byte         | Zahl, Byte         | Zahl, Byte  | `16 modulo 12`                         | 4        |
| Potenzieren                  | `a hoch b`                              | `pow(a, b)`           | numerisch          | numerisch          | Kommazahl   | `2 hoch 8`                             | 256,0    |
| Wurzelziehen                 | `die a. Wurzel von b`                   | `pow(a, 1/b)`         | numerisch          | numerisch          | Kommazahl   | `die 2. Wurzel von 9`                  | 3,0      |
| Logarithmus                  | `der Logarithmus von b zur Basis a`     | `log10(b) / log10(a)` | numerisch          | numerisch          | Kommazahl   | `der Logarithmus von 100 zur Basis 10` | 2,0      |
| Bit-Verschiebung nach links  | `a um b Bit nach links verschoben`      | `a << b`              | Zahl, Byte         | Zahl, Byte         | Zahl, Byte  | `7 um 3 Bit nach links verschoben`     | 56       |
| Bit-Verschiebung nach rechts | `a um b Bit nach rechts verschoben`     | `a >> b`              | Zahl, Byte         | Zahl, Byte         | Zahl, Byte  | `70 um 2 Bit nach rechts verschoben`   | 17       |
| Logische UND verknüpfung     | `a logisch und b`                       | `a&b`                 | Zahl, Byte         | Zahl, Byte         | Zahl, Byte  | `5 logisch und 2`                      | 0        |
| Logische ODER verknüpfung    | `a logisch oder b`                      | `a\| b`               | Zahl, Byte         | Zahl, Byte         | Zahl, Byte  | `5 logisch oder 2`                     | 7        |
| Logische XOR verknüpfung     | `a logisch kontra b`                    | `a^b`                 | Zahl, Byte         | Zahl, Byte         | Zahl, Byte  | `8 logisch kontra 5`                   | 13       |

# Bool'sche Operatoren

Mithilfe von Bool'schen Operatoren können komplexe Bedingungen ausgedrückt und zusammengefasst werden und z.B. für Verzweigungen oder Schleifen verwendet werden.

| Operator            | Beschreibung                                          | C Equivalent      | Beispiel                                                                                                                  | Ergebnis                                   |
| ------------------- | ----------------------------------------------------- | ----------------- | ------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------ |
| und                 | Wahr, wenn beide Argumente wahr sind.                 | `true && false`   | `wahr und wahr`<br>`wahr und falsch`<br>`falsch und wahr`<br>`falsch und falsch`                                          | `wahr`<br>`falsch`<br>`falsch`<br>`falsch` |
| oder                | Wahr, wenn eines der beiden Argumente wahr ist.       | `true \|\| false` | `wahr oder wahr`<br>`wahr oder falsch`<br>`falsch oder wahr`<br>`falsch oder falsch`                                      | `wahr`<br>`wahr`<br>`wahr`<br>`falsch`     |
| `entweder ... oder` | Wahr, wenn *nur* eines der beiden Argumente wahr ist. | `true != false`   | `entweder wahr oder wahr`<br>`entweder wahr oder falsch`<br>`entweder falsch oder wahr`<br>` entweder falsch oder falsch` | `falsch`<br>`wahr`<br>`wahr`<br>`falsch`   |
| nicht               | Der Wert des Arguments wird umgekehrt.                | `!true`           | `nicht wahr` <br>`nicht falsch`                                                                                           | `falsch`<br>`wahr`                         |

# Vergleichsoperatoren

| Operator          | Beschreibung                                                                          | C Equivalent           | Beispiel                                                                                     | Ergebnis                     |
| ----------------- | ------------------------------------------------------------------------------------- | ---------------------- | -------------------------------------------------------------------------------------------- | ---------------------------- |
| gleich            | Wahr, wenn beide Argumente den gleichen Wert haben.                                   | `1 == 1`               | `1 gleich 1 ist`<br>`1 gleich 2 ist`                                                         | `wahr`<br>`falsch`           |
| ungleich          | Wahr, wenn die beiden Argumente verschiedene Werte haben.                             | `1 != 1`               | `1 ungleich 1 ist`<br>`1 ungleich 2 ist`                                                     | `wahr`<br>`falsch`           |
| kleiner als       | Wahr, wenn das Linke Argument einen kleineren Wert als das Rechte hat.                | `5 < 10`               | `5 kleiner als 10 ist`<br>`30 kleiner als 15 ist`                                            | `wahr`<br>`falsch`           |
| größer als        | Wahr, wenn das Linke Argument einen größeren Wert als das Rechte hat.                 | `7 > 3`                | `7 größer als 3 ist`<br>`5 größer als 8 ist`                                                 | `wahr`<br>`falsch`           |
| kleiner als, oder | Wahr, wenn das Linke Argument einen kleineren oder denselben Wert wie das Rechte hat. | `5 <= 10`              | `5 kleiner als, oder 10 ist`<br>`30 kleiner als, oder 15 ist`<br>`5 kleiner als, oder 5 ist` | `wahr`<br>`falsch`<br>`wahr` |
| größer als, oder  | Wahr, wenn das Linke Argument einen größeren oder denselben Wert wie das Rechte hat.  | `7 >= 3`               | `7 größer als, oder 3 ist`<br>`5 größer als, oder 8 ist`<br>`5 größer als, oder 5 ist`       | `wahr`<br>`falsch`<br>`wahr` |
| zwischen          | Wahr, wenn das Linke Argument zwischen den beiden Rechten liegt.                      | `(4 < 5 && 4 > 3)`     | `4 zwischen 3 und 5 ist`                                                                     | `wahr`<br>`falsch`           |

Vergleichsoperatoren haben alle ein "ist" am Ende um der Grammatik in jedem Kontext gerecht zu werden.
Falls das zu mehreren "ist"s hintereinander führen sollte reicht ein einziges aus:
```ddp
2 größer als 2 ist gleich 4 größer als 4 ist ist. [hässlich]
2 größer als 2 ist gleich 4 größer als 4 ist. [schön]
```

# Listen und Text Operatoren

| Funktion              | Verwendung                 | C Equivalent | Typ vom 1. Operanden | Typ vom 2. Operanden | Typ vom 3. Operanden | Rückgabetyp           | Beispiel                              | Ergebnis       |
| --------------------- | -------------------------- | ------------ | -------------------- | -------------------- | -------------------- | --------------------- | ------------------------------------- | -------------- |
| Verkettung            | `a verkettet mit b`        | -            | Text/Liste/Buchstabe | Text/Liste/Buchstabe | -                    | Text/Liste            | `"Hallo" verkettet mit " Welt"`       | `"Hallo Welt"` |
| Indizierung           | `a an der Stelle b`        | `a[b]`       | Text/Liste           | Zahl, Byte           | -                    | Buchstabe/Element Typ | `"Hallo" an der Stelle 1`             | 'H'            |
| Bereich               | `a im Bereich von b bis c` | -            | Text/Liste           | Zahl, Byte           | Zahl, Byte           | Text/Liste            | `"Hallo Welt" im Bereich von 1 bis 5` | "Hallo"        |
| `... ab dem ...`      | `a ab dem b. Element`      | -            | Text/Liste           | Zahl, Byte           | -                    | Text/Liste            | `"Hallo Welt" ab dem 7. Element`      | "Welt"         |
| `... bis zum ...`     | `a bis zum b. Element`     | -            | Text/Liste           | Zahl, Byte           | -                    | Text/Liste            | `"Hallo Welt" bis zum 5. Element`     | "Hallo"        |

## Bemerkungen

Die Operanden Typen einer Verkettung können nicht beliebig kombiniert werden.

- Verkettet man 2 beliebige Werte desselben Typs entsteht eine Liste.
- Verkettet man 2 Texte entsteht ein neuer Text.
- Verkettet man einen Text mit einem Buchstaben oder umgekehrt entsteht ein Text.
- Verkettet man eine Liste und einen beliebigen Wert vom Element Typ der Liste oder umgekehrt entsteht eine Liste.

Bei Indizierungen und `im Bereich von ... bis` sind die Indizes immer inklusiv und fangen bei 1 an.
Eine Liste von 1 bis 5 sind also das erste, das fünfte und alle Element dazwischen.
Eine Liste an der Stelle 0 wäre ein Laufzeitfehler, genauso wie ein Index, der die Länge der Liste überschreitet.

Bei `im Bereich von ... bis` gibt es keine Laufzeitfehler bei zu kleinen oder zu großen Indizes.
Die Indizes werden automatisch in den Bereich [1, Länge der Liste/Text] gebracht.
Sollte danach der 2. Index kleiner als der 1. sein, so gibt es einen Laufzeitfehler.
Wenn also beide Indizes größer als die Länge der Liste sind, so ist das Ergebnis also das Letzte Element der Liste.

`... ab dem ...` und `... bis zum ...` sind nur Kurzschreibweisen von `im Bereich von ... bis` mit dem zweiten bzw. ersten Operanden
auf die Länge der Liste bzw. 1 gesetzt ist.

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

# Falls

Der `Falls` Operator entspricht dem Ternary Operator (`a ? b : c`) in anderen Sprachen.
Seine Parameter sind ein Wahrheitswert und 2 Ausdrücke beliebigen (aber des gleich) Typs:
```ddp
Der Text t ist "2 > 3", falls 2 größer als 3 ist, ansonsten "2" verkettet mit " < " verkettet mit "3".
```

# Operator Priorisierung

Verschiedene Operatoren werden verschieden Priorisiert.
Das bedeutet, dass manche Operatoren bei der Auswertung Vorrang vor anderen haben.
In der Mathematik haben z.B. Multiplikation und Division Vorrang vor Addition und Subtraktion, und
Potenzen haben wiederrum Vorrang vor Multiplikationen und Divisionen.
Bei den mathemathischen Operatoren hält DDP sich weitestgehend an die mathemathische Reihenfolge.
Hier ist eine Tabelle mit allen Operatoren und ihrer Priorisierung (hoch priorisierte Operatoren oben):

| Rang | Operator                                                        |
| ---- | --------------------------------------------------------------- |
| 1    | Funktionsaufruf                                                 |
| 2    | Literale, Konstante                                             |
| 3    | Klammern                                                        |
| 4    | Typkonvertierungen                                              |
| 5    | Feld Zugriff von Kombinationen                                  |
| 6    | Indizierung                                                     |
| 7    | `von ... bis`, `ab ... dem`, `bis ... zum`                      |
| 8    | Potenzieren, Wurzelziehen, Logarithmus                          |
| 9    | Negation,                                                       |
| 10   | Betrag, Größe, Länge, Standardwert, Logisches-/Bool'sches nicht |
| 11   | Multiplikation, Division, Rest,                                 |
| 12   | Addition, Subtraktion, Verkettung,                               |
| 13   | Bit-Verschiebung                                                |
| 14   | Größen Vergleich (kleiner, größer, ect.)                        |
| 15   | Gleichheit (gleich und ungleich)                                |
| 16   | Logische UND Verknüpfung                                        |
| 17   | Logische XOR Verknüpfung                                        |
| 18   | Logische ODER Verknüpfung                                       |
| 19   | Bool'sches UND                                                  |
| 20   | Bool'sches ODER                                                 |
| 21   | `entweder ... oder`                                             |
| 22   | Falls                                                           |

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

# Operator Überladung

Operatoren können auch Überladen werden um Funktionen anstelle von Operatoren auszuführen und trotzdem
von Operator Priorisierung zu Profitieren.
Mehr dazu im Artikel [Operatoren Überladung](/Bedienungsanleitung/de/Programmierung/Funktionen/Operatoren-Ueberladung).