+++
title = "Datentypen"
weight = 1
type = "article"
+++

# Datentypen

Da die Deutsche Programmiersprache statisch typisiert ist, hat jeder Ausdruck (z.B. mathematische Ausdrücke), jede Variable und jede Funktion einen festen Datentyp.

Der Typ von Variablen, Funktionen und Ausdrücken kann sich nicht zur Laufzeit ändern, er wird zur Kompilierzeit festgelegt.

## Einfache Datentypen

| Typname   | Beschreibung                                              | Wertebereich                                                                       | Literal                                                                               | Beispiel                                                      |
| --------- | --------------------------------------------------------- | ---------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------- | ------------------------------------------------------------- |
| Zahl      | Eine 64 Bit große, ganze Zahl                             | *-2.147.483.648* bis *2.147.483.647*                                               | Eine Abfolge von Ziffern, z.B. 42                                                     | `Die Zahl x ist 69.`, <br>`1 plus -7`                         |
| Kommazahl | Eine 64 Bit große, gleitkomma Zahl                        | *-1,79769313486232x10^308* bis <br>*1,79769313486232x10^308* mit 16 Dezimalstellen | Ein Zahlenliteral mit Nachkommastellen, z.B. 3,1415                                   | `Die Kommazahl x ist 6,5.`, <br>`2 durch 0,5`                 |
| Wahrheitswert   | Ein Wahrheitswert (8 Bit groß)                            | *wahr* oder *falsch*                                                               | *wahr* oder *falsch*                                                                  | `Der Wahrheitswert x ist wahr.`, <br>`1 plus 1 gleich 2`            |
| Buchstabe | Ein 4 Byte großes, mit utf-8 kodiertes Zeichen            | *0* - *65535*                                                                      | Ein utf8 Zeichen zwischen einfachen Anführungszeichen, z.B. 'a' oder '\n'             | `Der Buchstabe x ist 'd'.`                                    |
| Text      | Eine utf-8 kodierte Aneinanderreihung mehrerer Buchstaben | *beliebig groß*                                                                    | Beliebig viele Buchstaben zwischen <br>(englischen) Anführungszeichen, z.B. "Hallo\n" | `Der Text x ist "abc".`, <br>`"Hallo" verkettet mit " du da"` |

### Anders als in anderen Programmiersprachen:

* **Das Dezimaltrennzeichen ist kein Punkt, sondern ein Komma!**
* **Es gibt keinen null/nil Typ**

### Escape-Sequenzen
Escape-Sequenzen sind Zeichenkombinationen, die für Zeichen stehen, 
welche man nicht normalerweise in Text- oder Buchstabenliterale schreiben kann.

| Name                          | Escape-Sequenz | ASCII Code |
| ----------------------------- | -------------- | ---------- |
| Zeilenvorschub                | `\n`           | 10         |
| Wagenrücklauf                 | `\r`           | 13         |
| Rückschritt                   | `\b`           | 8          |
| Tabulator                     | `\t`           | 9          |
| Akustisches Signal            | `\a`           | 7          |
| Rückstrich                    | `\\`           | 92         |
| Einfaches Anführungszeichen*  | `\'`           | 39         |
| Doppeltes Anführungszeichen** | `\"`           | 34         |

*: Nur innerhalb eines Buchstaben-Literals.\
**: Nur innerhalb eines Text-Literals.

***

## Listen

Listen sind beliebig große Ansammlungen von Werten.
Da DDP statisch typisiert ist kann eine Liste nur Werte desselben Datentyps enthalten.
Der Typname einer Liste ist im Allgemeinen der Element-Typname entsprechend dekliniert mit *Liste* angehängt (Zahl -> Zahlen Liste, Text -> Text Liste).
Bei Benutzerdefinierten Typen ([Kombinationen](/Bedienungsanleitung/de/Programmierung/Kombinationen)) kann die richtige Deklination (noch) nicht geparset werden, deshalb wird nicht dekliniert, sondern einfach nur *Liste* angehängt (siehe [Kombinationslisten](/Bedienungsanleitung/de/Programmierung/Kombinationen#kombinationslisten))
Eine Liste kann zur Laufzeit wachsen und schrumpfen.
Wie man mit Listen arbeitet, wird in dem Artikel Operatoren unter [Listen Operatoren](/Bedienungsanleitung/de/Programmierung/Operatoren#listen-und-text-operatoren) beschrieben.

### Listen Literale

Ein Listen Literal sieht ähnlich wie eine gewöhnliche Aufzählung aus, allerdings in einen kleinen Satzbaustein verpackt um Mehrdeutigkeiten beim Programmieren zu verhindern.
Die allgemeine Form sieht so aus: `eine[r] Liste, die aus x[, y, z] besteht`.
Es kann `eine` oder `einer` je nach grammatischem Kontext benutzt werden und auf das `aus` müssen einer oder mehrere
Ausdrücke desselben Typs folgen.

Leere Listen können so erstellt werden: `eine[r] leere[n] (Typname) Liste`.

Listen mit einem einzigen Element können auch einfach durch `(Wert) als (Typname)` erstellt werden.

Es gibt auch noch einen weiteren Sonderfall: Listen können bei der Variablen Deklaration auch mit der `Die (Typname) Liste (Variablenname) ist <Zahl> Mal <Wert>` Syntax erstellt werden.
Dabei wird die Liste mit dem Wert initialisiert, der so oft in der Liste vorkommt, wie die Zahl angibt.
Das ist besonders nützlich um relativ große Listen mit einem Standardwert zu erstellen, da bei dieser Syntax nur ein Mal der ganze Speicher allokiert wird, anstatt für jedes Element einzeln.

### Beispiele
```ddp
Die Zahlen Liste z ist eine Liste, die aus 1, 2, 3 besteht.
Die Text Liste t ist eine Liste, die aus "Hallo", "Welt" besteht.

Die Zahlen Liste z2 ist eine leere Zahlen Liste.
Die Text Liste t2 ist "Hallo" als Text Liste.
```

| Typname           | Beschreibung               | Literal                                 | Beispiel                                                                    |
| ----------------- | -------------------------- | --------------------------------------- | --------------------------------------------------------------------------- |
| Zahlen Liste      | Eine Liste von Zahlen      | Ein Listen Literal wie oben beschrieben | `Die Zahlen Liste z ist eine Liste, die aus 1, 2, 3 besteht.`               |
| Kommazahlen Liste | Eine Liste von Kommazahlen | Ein Listen Literal wie oben beschrieben | `Die Kommazahlen Liste z ist eine Liste, die aus 1,2, 3,2, 3,1415 besteht.` |
| Wahrheitswert Liste     | Eine Liste von Wahrheitswerten    | Ein Listen Literal wie oben beschrieben | `Die Wahrheitswert Liste b ist eine Liste, die aus wahr, falsch, wahr besteht.`   |
| Buchstaben Liste  | Eine Liste von Buchstaben  | Ein Listen Literal wie oben beschrieben | `Die Buchstaben Liste b ist eine Liste, die aus 'b', 'h', 'z' besteht.`     |
| Text Liste        | Eine Liste von Texten      | Ein Listen Literal wie oben beschrieben | `Die Text Liste t ist eine Liste, die aus "Hallo", "du", "da" besteht.`     |
***

### Bemerkung

Eigentlich würde man ja erwarten, dass in der Aufzählung eines Listen Literals noch ein 'und' vorkommen müsste (`eine Liste, die aus 1, 2 und 3 besteht`). Das würde aber zu Mehrdeutigkeiten in Boolschen Ausdrücken führen (`eine Liste, die aus wahr und falsch besteht`), und da die Aufzählung, grammatisch gesehen, kein 'und' braucht wird es in Listen Literalen weggelassen.

## Kombinationen

Kombinationen (structs in C) sind benutzerdefinierte zusammengesetzte Datentypen, die eine oder mehrere Variablen in einem Typ zusammenfassen.
Mehr zu Kombinationen ist im Artikel [Kombinationen](/Bedienungsanleitung/de/Programmierung/Kombinationen) zu finden.
