# Listen

Listen sind beliebig große Ansammlungen von Werten.
Da DDP statisch typisiert ist kann eine Liste nur Werte eines Datentyps enthalten.
Der Typname einer Liste ist im Allgemeinen der Element-Typname entsprechend dekliniert mit *Liste* angehängt (Zahl -> Zahlen Liste, Text -> Text Liste).
Eine Liste kann zur Laufzeit wachsen und schrumpfen.
Wie man mit Listen arbeitet, wird in dem Artikel [Listen Operatoren](?p=Listen%20Operatoren) beschrieben.

## Listen Literale

Ein Listen Literal sieht ähnlich wie eine gewöhnliche Aufzählung aus, allerdings in einen kleinen Satzbaustein verpackt um Mehrdeutigkeiten beim Programmieren zu verhindern.
Die allgemeine Form sieht so aus: `eine[r] Liste, die aus x[, y, z] besteht`.
Es kann `eine` oder `einer` je nach grammatischem Kontext benutzt werden und auf das `aus` müssen einer oder mehrere
Ausdrücke desselben Typs folgen.

Leere Listen können so erstellt werden: `eine[r] leere[n] (Typname) Liste`.

Listen mit einem einzigen Element können auch einfach durch `(Wert) als (Typname)` erstellt werden.

### Beispiele
```ddp
Die Zahlen Liste z ist eine Liste, die aus 1, 2, 3 besteht.
Die Text Liste t ist eine Liste, die aus "Hallo", "Welt" besteht.

Die Zahlen Liste z2 ist eine leere Zahlen Liste.
Die Text Liste t2 ist "Hallo" als Text Liste.
```

| Typname | Beschreibung | Literal | Beispiel |
| ------- | ------------ | ------- | -------- |
| Zahlen Liste | Eine Liste von Zahlen | Ein Listen Literal wie oben beschrieben | `Die Zahlen Liste z ist eine Liste, die aus 1, 2, 3 besteht.` |
| Kommazahlen Liste | Eine Liste von Kommazahlen | Ein Listen Literal wie oben beschrieben | `Die Kommazahlen Liste z ist eine Liste, die aus 1,2, 3,2, 3,1415 besteht.` | |
| Boolean Liste | Eine Liste von Booleans | Ein Listen Literal wie oben beschrieben | `Die Boolean Liste b ist eine Liste, die aus wahr, falsch, wahr besteht.` |
| Buchstaben Liste | Eine Liste von Buchstaben | Ein Listen Literal wie oben beschrieben | `Die Buchstaben Liste b ist eine Liste, die aus 'b', 'h', 'z' besteht.` |
| Text Liste | Eine Liste von Texten | Ein Listen Literal wie oben beschrieben | `Die Text Liste t ist eine Liste, die aus "Hallo", "du", "da" besteht.` |
***

## Bemerkung

Eigentlich würde man ja erwarten, dass in der Aufzählung eines Listen Literals noch ein 'und' vorkommen müsste (`eine Liste, die aus 1, 2 und 3 besteht`). Das würde aber zu Mehrdeutigkeiten in Boolschen Ausdrücken führen (`eine Liste, die aus wahr und falsch besteht`), und da die Aufzählung, grammatisch gesehen, kein 'und' braucht wird es in Listen Literalen weggelassen.
