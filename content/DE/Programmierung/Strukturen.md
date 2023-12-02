+++
title = "Strukturen"
weight = 7
type = "article"
+++

# Strukturen

Strukturen sind ein wichtiges Feature der meisten Programmiersprachen.
Sie ermöglichen es zusammengesetzte Datentypen zu definieren und so auch komplexere Daten modellieren zu können.

Ähnlich wie bei Funktionen steht bei Strukturen in DDP auch die Lesbarkeit im Vordergrund, allerdings nicht ganz so extrem, da sie auch Nutzbar sein sollen

## Deklaration

Eine Strukturdeklaration sieht im Allgemeinen so aus:

```ddp
Wir nennen die (öffentliche) Kombination aus
    (der/dem (öffentlichen) <Typ-Name> <Feld-Name> mit Standardwert <Ausdruck>),
    ...
ein/eine/einen <Struktur-Name>, und erstellen sie so:
	"Alias mit Feld <x>" (,
	"Noch ein Alias mit den Feldern <x> und <y>" oder
	...)
```

Beispiel:

```ddp
Wir nennen die öffentliche Kombination aus
	der Zahl x mit Standardwert 0, [privates Feld]
	der öffentlichen Zahl y mit Standardwert 0, [öffentliches Feld]
einen Vektor2, und erstellen sie so:
	"Nullvektor2" oder
	"der Nullvektor2" oder [kein Parameter]
	"ein Vektor2 mit x gleich <x>" oder [1 Parameter]
	"ein Vektor2 mit x gleich <x> und y gleich <y>" [alle Parameter]
```

Hier wird ein zweidimensionaler Vektor, der aus zwei Zahlen (x und y) besteht definiert.
Die Struktur selbst und das Feld y sind öffentlich, können also auch in anderen Modulen benutzt werden (mehr dazu im Artikel [Module](Bedienungsanleitung/DE/Programmierung/Module)).

Diese Schreibweise kling zwar sehr mathematisch, aber ermöglicht etwas sehr wichtiges.
Durch den unbestimmten Artikel (einer, eine oder ein) wird das grammatikalische Geschlecht (maskulin, feminin oder neutrum) des Typ-Namens erkennbar.
Die oben definierte Vektor2 Struktur, muss man also als grammatikalisch maskulin behandeln:

```ddp
Der Vektor2 x ist der Nullvektor2. [korrekt]
Die Vektor2 y ist der Nullvektor2. [Fehler: falscher Artikel]

Die Vektor2 Liste vektoren ist eine leere Vektor2 Liste.
Für jeden [nicht jede oder jedes!!] Vektor2 vek in vektoren, mache:
    ...
```

### Aliase

Strukturen werden genau wie Funktionen über Aliase erstellt.
Anders als bei Funktionen muss ein Struktur-Alias jedoch nicht alle Felder enthalten.
Die Felder, die bei einem Struktur-Alias fehlen werden einfach auf den Standardwert gesetzt.

## Zugriff auf Felder

Nehmen wir an wir haben die obige Vektor2 Struktur deklariert und einen Vektor2 erstellt:

```ddp
Der Vektor2 vek ist der Nullvektor.
```

Auf die einzelnen Felder von vek wird mit dem `von` Operator (`.` in anderen Sprachen) zugegriffen:

```ddp
Schreibe (x von vek). [0]
Schreibe (y von vek). [0]

Speichere 2 in x von vek. [vek.x wird auf 2 gesetzt]

Schreibe (x von vek). [2]
Schreibe (y von vek). [0]
```

