# Datentypen

Da die Deutsche Programmiersprache statisch typisiert ist, hat jeder Ausdruck (z.B. mathematische Ausdrücke), jede Variable und jede Funktion einen festen Datentyp.

Der Typ von Variablen, Funktionen und Ausdrücken kann sich nicht zur Laufzeit ändern, er wird zur Kompilierzeit festgelegt.

## Einfache Datentypen

| Typname | Beschreibung | Wertebereich | Literal | Beispiel |
| ------- | ------------ | ------------ | ------- | -------- |
| Zahl | Eine 64 Bit große, ganze Zahl | *-2.147.483.648* bis *2.147.483.647* | Eine Abfolge von Ziffern, z.B. 42 | `Die Zahl x ist 69.`, <br>`1 plus -7` |
| Kommazahl | Eine 64 Bit große, reellen Zahl |*-1,79769313486232x10^308* bis <br>*1,79769313486232x10^308* mit 16 Dezimalstellen | Ein Zahlenliteral mit Nachkommastellen, z.B. 3,1415 | `Die Kommazahl x ist 6,5.`, <br>`2 durch 0,5` | |
| Boolean | Ein Wahrheitswert (8 Bit groß) | *wahr* oder *falsch* | *wahr* oder *falsch* | `Der Boolean x ist wahr.`, <br>`1 plus 1 gleich 2` |
| Buchstabe | Ein 1-4 Bit großes, mit utf-8 kodiertes Zeichen | *0* - *65535* | Ein utf8 Zeichen zwischen einfachen Anführungszeichen, z.B. 'a' oder '\n' | `Der Buchstabe x ist 'd'.` |
| Text | Eine Aneinanderreihung mehrerer Buchstaben | *beliebig groß* | Beliebig viele Buchstaben zwischen <br>(englischen) Anführungszeichen, z.B. "Hallo\n" | `Der Text x ist "abc".`, <br>`"Hallo" verkettet mit " du da"` |

## **Anders als in anderen Programmiersprachen:**

* **Das Dezimaltrennzeichen ist kein Punkt, sondern ein Komma!**
* **Es gibt keinen null/nil Typ**
***

## Listen

Listen sind beliebig große Ansammlungen von Werten.
Mehr zu Listen ist in dem Artikel [Listen](?p=Listen) zu finden.

## Strukturen

Strukturen sind zusammengesetzte Datentypen, das heißt sie bestehen aus mehreren Variablen, die zu einer Einheit zusammengefasst sind.
Wie man Strukturen benutzt wird in dem Artikel [Strukturen](?p=Strukturen) beschrieben.