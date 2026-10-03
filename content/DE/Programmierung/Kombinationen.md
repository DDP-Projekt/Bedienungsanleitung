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
Die Kombination selbst und das Feld y sind öffentlich, können also auch in anderen Modulen benutzt werden (mehr dazu im Artikel [Module](/Programmierung/Module)).

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

Aus praktischen Gründen (vor allem in Hinsicht auf generische Typen) sind die Aliase und Standardwerte bei Kombinationen jeweils Optional.
Das heißt statt dem obigen Beispiel, könnte man einen Vektor auch so definieren:

```ddp
Wir nennen die öffentliche Kombination aus
	der Zahl x,
	der öffentlichen Zahl y,
einen Vektor2.

Der Vektor2 v ist der Standardwert von einem Vektor2. [ x = 0, y = 0 ]
```

Wie man sieht kann man eine Kombination in diesem Fall nur mithilfe des "Standardwert" Operators erstellen.
Felder, die keinen expliziten Standardwert haben, werden ebenfalls auf ihren Standardwert gesetzt.

### Aliase

Kombinationen werden genau wie Funktionen über Aliase erstellt.
Anders als bei Funktionen muss ein Kombinationssalias jedoch nicht alle Felder enthalten.
Die Felder, die bei einem Kombinationsalias fehlen werden einfach auf den Standardwert gesetzt.

## Zugriff auf Felder

Nehmen wir an wir haben die obige Vektor2 Kombination deklariert und einen Vektor2 erstellt:

```ddp
Der Vektor2 vek ist der Nullvektor2.
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
Angenommen, man hat eine Kombination `Vektor` definiert, dann sieht eine Vektor Liste so aus:

```ddp
Die Vektor Liste vektoren ist eine leere Vektor Liste.

Der Vektor vek1 ist vektoren an der Stelle 1.
Die Zahl x ist x von (vektoren an der Stelle 1).
```

Bei Benutzerdefinierten Kombinationen ist es leider (noch) nicht möglich den Typnamen entsprechend zu deklinieren, also ist es anders als bei eingebauten Typen (Zahl -> Zahl*en* Liste vs. Vektor -> Vektor Liste).

Wie man sieht hat der `von` Operator auch Vorrang vor dem `an der Stelle` Operator, so wie es in der [Priorisierung von Operatoren](/Programmierung/Operatoren/#operator-priorisierung) festgelegt ist.

## Generische Kombinationen

Genau wie es [generische Funktionen](/Programmierung/Funktionen/Generische-Funktionen) gibt, gibt es auch generische Kombinationen.
Dabei steht das Wort `generische` vor `Kombination`, und die Felder können Typparameter wie `T` als Typ haben:

```ddp
Wir nennen die generische Kombination aus
	dem T erstes,
	dem R zweites,
ein Paar, und erstellen sie so:
	"ein Paar aus <erstes> und <zweites>"
```

Typparameter sind grammatikalisch sächlich. Deshalb steht hier `dem T` und nicht `der T`.

Wenn man eine generische Kombination benutzt, schreibt man die Typen für die Typparameter mit Bindestrichen vor den Namen.
Die Reihenfolge ist dabei die Reihenfolge, in der die Typparameter in den Feldern zum ersten Mal vorkommen.
Bei `Paar` ist also zuerst `T` und dann `R` dran:

```ddp
Das Zahl-Text-Paar p ist ein Paar aus 1 und "eins".
Schreibe (erstes von p) auf eine Zeile. [1]
Schreibe (zweites von p) auf eine Zeile. [eins]

Die Zahl-Text-Paar Liste liste ist eine leere Zahl-Text-Paar Liste.
```

Ist einer der Typen selbst ein generischer Typ, setzt man ihn in Klammern, zum Beispiel `(Zahl-Paar)-Text-Paar`.

### Aliase von generischen Kombinationen

Bei einem Alias muss der Kompilierer alle Typparameter herausfinden können.
Das geht über die Argumente des Alias oder über die Standardwerte der Felder:

```ddp
Wir nennen die generische Kombination aus
	dem T x mit Standardwert 0,
	dem T y mit Standardwert 0,
einen Punkt, und erstellen sie so:
	"der Ursprung",
	"ein Punkt bei <x> und <y>"

Der Zahl-Punkt u ist der Ursprung. [T ist eine Zahl wegen der Standardwerte]
Der Kommazahl-Punkt k ist ein Punkt bei 1,5 und 2,5. [T ist eine Kommazahl wegen der Argumente]
```

Hätten die Felder keine Standardwerte, wäre der Alias `"der Ursprung"` nicht erlaubt, denn dann ist nicht klar, welcher Typ `T` ist.
Ohne Aliase kann man eine generische Kombination immer mit dem `Standardwert` Operator erstellen:

```ddp
Das Zahl-Text-Paar p ist der Standardwert von einem Zahl-Text-Paar.
```

### Generische Kombinationen in Funktionen

Generische Funktionen können generische Kombinationen als Parameter haben:

```ddp
Die generische Funktion Summe mit dem Parameter p vom Typ T-Punkt, gibt ein T zurück, macht:
	Gib x von p plus y von p zurück.
Und kann so benutzt werden:
	"die Summe von <p>"

Schreibe (die Summe von k) auf eine Zeile. [4]
```