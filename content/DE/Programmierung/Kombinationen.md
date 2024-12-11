+++
title = "Kombinationen"
weight = 7
+++

# Kombinationen

Verbundtypen, auch Strukturen oder Records genannt, sind einer der wichtigsten Datentypen der meisten Programmiersprachen.
Sie ermöglichen es zusammengesetzte Datentypen zu definieren und so auch komplexere Daten modellieren zu können.
In DDP nennen wir sie Kombinationen.

Ähnlich wie bei Funktionen steht bei Kombinationen in DDP auch die Lesbarkeit im Vordergrund, allerdings nicht ganz so extrem, da sie auch Nutzbar sein sollen.

## Deklaration

Eine Kombinationsdeklaration sieht im Allgemeinen so aus:

```ddp
Wir nennen die (öffentliche) Kombination aus
    (der/dem (öffentlichen) <Typ-Name> <Feld-Name> mit Standardwert <Ausdruck>),
    ...
ein/eine/einen <Kombinations-Name>, und erstellen sie so:
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
Die Kombination selbst und das Feld y sind öffentlich, können also auch in anderen Modulen benutzt werden (mehr dazu im Artikel [Module](/de/Programmierung/Module)).

Diese Schreibweise kling zwar sehr mathematisch, aber ermöglicht etwas sehr wichtiges.
Durch den unbestimmten Artikel (einer, eine oder ein) wird das grammatikalische Geschlecht (maskulin, feminin oder neutrum) des Typ-Namens erkennbar.
Die oben definierte Vektor2 Kombination, muss man also als grammatikalisch maskulin behandeln:

```ddp
Der Vektor2 x ist der Nullvektor2. [korrekt]
Die Vektor2 y ist der Nullvektor2. [Fehler: falscher Artikel]

Die Vektor2 Liste vektoren ist eine leere Vektor2 Liste.
Für jeden [nicht jede oder jedes!!] Vektor2 vek in vektoren, mache:
    ...
```

### Aliase

Kombinationen werden genau wie Funktionen über Aliase erstellt.
Anders als bei Funktionen muss ein Kombinationssalias jedoch nicht alle Felder enthalten.
Die Felder, die bei einem Kombinationsalias fehlen werden einfach auf den Standardwert gesetzt.

## Zugriff auf Felder

Nehmen wir an wir haben die obige Vektor2 Kombination deklariert und einen Vektor2 erstellt:

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

## Kombinationslisten

Natürlich kann man auch Listen von Kombinationen haben.
Angenommen, man hat eine Kombination `Vektor` definiert, dann sieht eine Vektoren Liste so aus:

```ddp
Die Vektor Liste vektoren ist eine leere Vektoren Liste.

Der Vektor vek1 ist vektoren an der Stelle 1.
Die Zahl x ist x von (vektoren an der Stelle 1).
```

Bei Benutzerdefinierten Kombinationen ist es leider (noch) nicht möglich den Typnamen entsprechend zu deklinieren, also ist es anders als bei eingebauten Typen (Zahl -> Zahl*en* Liste vs. Vektor -> Vektor Liste).

Wie man sieht hat der `von` Operator auch Vorrang vor dem `an der Stelle` Operator, so wie es in der [Priorisierung von Operatoren](/de/Programmierung/Operatoren/#operator-priorisierung) festgelegt ist.