# Datentypen

Da Die Deutsche Programmiersprache statisch typisiert ist, hat jeder Ausdruck (z.B. mathematische Ausdrücke), jede Variable und jede Funktion einen festen Datentyp.

Der Typ von Variablen, Funktionen und Ausdrücken kann sich nicht zur Laufzeit ändern, er wird zur Kompilerzeit festgelegt.

## Einfache Datentypen

| Typname | Beschreibung | Wertebereich | Literal | Beispiel |
| ------- | ------------ | ------------ | ------- | -------- |
| Zahl | Eine 64 Bit große, ganze Zahl | *-2.147.483.648* bis *2.147.483.648* | Eine Abfolge von Ziffern, z.B. 42 | `Die Zahl x ist 69.`, <br>`1 plus -7` |
| Kommazahl | Eine 64 Bit große, rationale Zahl |*-1,79769313486232x10^308* bis <br>*1,79769313486232x10^308* mit 16 Dezimalstellen | Ein Zahlenliteral mit Nachkommastellen, z.B. 3,1415 | `Die Kommazahl x ist 6,5.`, <br>`2 durch 0,5` | |
| Boolean | Ein Wahrheitswert (8 Bit groß) | *wahr* oder *falsch* | *wahr* oder *falsch* | `Der Boolean x ist wahr.`, <br>`1 plus 1 gleich 2` |
| Buchstabe | Ein 1-4 Bit großes, mit utf-8 kodiertes Zeichen | *0* - *65535* | Ein utf8 Zeichen zwischen einfachen Anführungszeichen, z.B. 'a' oder '\n' | `Der Buchstabe x ist 'd'.` |
| Text | Eine Aneinanderreihung mehrerer Buchstaben | *beliebig groß* | Beliebig viele Buchstaben zwischen <br>(englischen) Anführungszeichen, z.B. "Hallo\n" | `Der Text x ist "abc".`, <br>`"Hallo" verkettet mit " du da"` |

## **Anders als in anderen Programmiersprachen:**

* **Das Dezimaltrennzeichen ist kein Punkt, sondern ein Komma!**
* **Es gibt keinen null/nil Typ**
***

## Listen

Listen sind beliebig große Ansammlungen von Werten.
Da DDP statisch typisiert ist kann eine Liste nur Werte eines Datentyps enthalten.
Der Typname einer Liste ist im Allgemeinen der Plural des Typs den sie enhält (z.B. Zahl -> Zahlen).
Eine Liste kann zur Laufzeit wachsen und schrumpfen.
Wie man mit Listen arbeitet wird in dem Artikel [Listen Operatoren](?p=Listen%20Operatoren) beschrieben.

| Typname | Beschreibung | Literal | Beispiel |
| ------- | ------------ | ------- | -------- |
| Zahlen | Eine Liste von Zahlen | Eine Aufzählung von Zahlen, z.B. 1, 2 und 3 | `Die Zahlen x sind 1, 2 und 3.` |
| Kommazahlen | Eine Liste von Kommazahlen | Eine Aufzählung von Kommazahlen, z.B. 1,2, 3,1415 und 42,42 | `Die Kommazahlen x sind 6,5 und 2,5.` | |
| Booleans | Eine Liste von Booleans | Eine Aufzählung von Booleans, z.B. wahr, falsch und wahr | `Die Booleans x sind wahr und falsch.` |
| Buchstaben | Eine Liste von Buchstaben | Eine Aufzählung von Buchstaben, z.B. '1' und 'h' | `Die Buchstaben x sind 'd', und '\n'.` |
| Texte | Eine Liste von Texten | Eine Aufzählung von Texten, z.B. "Hallo", "test" und "Tschüss" | `Die Texte t sind "hi" und "bye"` |
***

## Strukturen

Strukturen sind zusammengesetzte Datentypen, das heißt sie bestehen aus mehreren Variablen, die zu einer Einheit zusammengefasst sind.
Wie man Strukturen benutzt wird in dem Artikel [Strukturen](?p=Strukturen) beschrieben.