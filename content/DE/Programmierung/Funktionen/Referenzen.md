+++
title = "Referenzen"
weight = 2
+++

# Referenzen

Normalerweise werden Argumente von Funktionen als Werte übergeben.
Das heißt, wenn man z.B. `Schreibe den Text t.` aufruft, wird eine Kopie der Variable t erstellt und der Funktion übergeben. Die Funktion kann `t` also nicht verändern.

In DDP können Funktionen nun aber auch sog. *Referenzen* als Parameter-Typen angeben.
Das sieht so aus:

```ddp
Die Funktion foo mit dem Parameter param vom Typ Text Referenz, gibt nichts zurück, macht:
    Speichere "neu!" in param.
Und kann so benutzt werden:
    "Verändere <t>"
```

Und man kann die Funktion nun so aufrufen:

```ddp
Der Text t ist "alt".
Schreibe t. [Zeigt "alt" in der Konsole]
Verändere t.
Schreibe t. [Zeigt "neu!" in der Konsole]
```

Wie man sieht, nehmen Referenz-Parameter Variablen als Argumente.
Wenn man ihnen einen Wert, z.B. ein Literal, übergeben würde, gäbe es einen Fehler:

```ddp
Verändere "Fehler". [Kompilierfehler: foo hat eine Variable erwartet]
```

Man kann aus jedem Typnamen einen Referenz-Typ machen, indem man "Referenz" an den Typnamen anhängt und ihn entsprechend dekliniert.
Aus `Zahlen Liste` wird also `Zahlen Listen Referenz`, aus `Wahrheitswert` wird `Wahrheitswert Referenz`, aus `Zahl` wird `Zahlen Referenz`, etc..

Normale Variablen können keine Referenz-Typen sein, nur Funktions-Parameter.

## Typumwandlungen als Referenzen

Bei [Typ-Aliasen und Typ-Definitionen](/Programmierung/Typ-Aliase-und-Typ-Definitionen) kann man eine Variable mit `als` umwandeln und trotzdem als Referenz übergeben.
Das funktioniert nur, wenn der Typ dabei eigentlich gleich bleibt, also zum Beispiel bei einer Typ-Definition und ihrem Basistyp:

```ddp
Wir definieren eine Hausnummer als eine Zahl.

Die Funktion Erhöhe_Zahl mit dem Parameter z vom Typ Zahlen Referenz, gibt nichts zurück, macht:
	Erhöhe z um 1.
Und kann so benutzt werden:
	"Zähle <z> hoch"

Die Hausnummer h ist 41 als Hausnummer.
Zähle (h als Zahl) hoch.
Schreibe (h als Zahl) auf eine Zeile. [42]
```

Andere Umwandlungen, wie zum Beispiel von einer Zahl in eine Kommazahl, können nicht als Referenz übergeben werden.