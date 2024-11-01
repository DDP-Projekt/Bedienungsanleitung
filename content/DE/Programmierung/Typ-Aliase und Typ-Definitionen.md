+++
title = "Typ-Aliase und Typ-Definitionen"
weight = 7
type = "article"
+++

# Typ-Aliase und Typ-Definitionen

Typ-Aliase und -Definitionen erlauben es neue Namen für bestehende Typen zu erstellen oder
neue Typen zu definieren.
Dieses Feature ermöglicht es mit möglichst niedrigem Aufwand DDP Programme deutlich lesbarer zu gestalten.

## Typ-Aliase

Ein Typ Alias ist einfach ein weiterer Name für einen bestehenden Typ.
Erstellt werden sie so:

```ddp
Wir nennen eine Zahl auch eine Hausnummer.
Wir nennen eine Zahl öffentlich auch eine Postleitzahl. 

Die Hausnummer h ist 22.
```

Wie man sieht können Typ-Aliase öffentlich oder privat sein, genauso wie andere Kombinationen.
Typ-Aliase erstellen einfach nur einen neuen Namen für denselben Typ.
Das heißt, alle Operationen bleiben erhalten und der Typ kann zu funktionen übergeben werden, die den Basistyp erwarten:

```
Binde "Duden/Ausgabe" ein.

Wir nennen eine Zahl auch eine Hausnummer.
Die Hausnummer h ist 22.

[ Funktioniert obwohl der Parameter von Schreibe_Zahl als Zahl deklariert wurde ]
Schreibe die Zahl h auf eine Zeile.

[ Funktioniert ohne 'als' Operator ]
Die Zahl z ist h.
```

## Typ-Definitionen

Typ-Definitionen sehen ähnlich aus wie Typ-Aliase, haben aber einen entscheidenden Unterschied: Sie erstellen *neue* Typen.
Das heißt Operatoren und Funktionen, die den Basistyp erwarten funktionieren nicht mit dem neuen Typ:

```ddp
Wir definieren einen Zeiger als eine Zahl.
Wir definieren einen String öffentlich als einen Text.

[ Fehler: Ein Wert vom Typ Zahl kann keiner Variable vom Typ Zeiger zugewiesen werden (3001) ]
Der Zeiger zeiger ist 2.
[ Fehler: Der plus Operator erwartet einen Ausdruck vom Typ 'Zahl', oder 'Kommazahl' aber hat 'Zeiger' bekommen (3000) ]
Der Zeiger zeiger2 ist (2 als Zeiger) plus (2 als Zeiger).

[ OK ]
Der Zeiger z ist 2 als Zeiger.
[ OK ]
Der Zeiger z2 ist (z als Zahl plus z als Zahl) als Zeiger.
```

Wie man sieht muss man Typ-Definitionen immer korrekt Umwandeln um damit arbeiten zu können.

Typdefinitionen können in ihren Basistyp und zurück umgewandelt werden. Dies Funktioniert aber nicht rekursiv:

```ddp
Wir definieren einen Zeiger als eine Zahl.
Wir definieren einen Index als einen Zeiger.

[ OK ]
Der Zeiger z ist 0 als Zeiger.
[ OK ]
Der Index i ist z als Index.
[ Fehler: Ein Ausdruck vom Typ Index kann nicht in den Typ Zahl umgewandelt werden (3004) ]
Die Zahl z2 ist 1 als Zeiger als Index als Zahl.
```
