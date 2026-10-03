+++
title = "Generische Funktionen"
weight = 5
+++

# Generische Funktionen

Manchmal möchte man dieselbe Funktion für verschiedene Typen schreiben.
Eine Funktion, die das größere von zwei Werten zurückgibt, funktioniert zum Beispiel mit Zahlen genauso wie mit Kommazahlen.
Damit man die Funktion nicht für jeden Typ einzeln schreiben muss, gibt es generische Funktionen.

## Deklaration

Eine generische Funktion wird wie eine normale Funktion deklariert.
Es kommt nur das Wort `generische` dazu.
Als Typ eines Parameters kann man dann einen beliebigen neuen Namen benutzen, zum Beispiel `T`.
Dieser Name ist ein sogenannter Typparameter. Er steht für einen Typ, der erst beim Aufruf feststeht.

```ddp
Die generische Funktion Größtes mit den Parametern a und b vom Typ T und T, gibt ein T zurück, macht:
	Gib a, falls a größer als b ist, ansonsten b zurück.
Und kann so benutzt werden:
	"das Größere von <a> und <b>"

Schreibe (das Größere von 3 und 7) auf eine Zeile. [7]
Schreibe (das Größere von 2,5 und 1,5) auf eine Zeile. [2,5]
```

Beim ersten Aufruf ist `T` eine Zahl, beim zweiten Aufruf eine Kommazahl.
Der Kompilierer erkennt den Typ an den Argumenten und erstellt für jeden Typ eine eigene Version der Funktion.

Typparameter sind grammatikalisch sächlich. Man schreibt also `ein T` und `das T`.

Eine generische Funktion braucht mindestens einen Typparameter.
Es können aber auch mehrere sein, zum Beispiel `T` und `R`.

## Typparameter in anderen Typen

Typparameter können auch in Listen und Referenzen benutzt werden:

```ddp
Die generische Funktion Erstes mit dem Parameter liste vom Typ T Liste, gibt ein T zurück, macht:
	Gib liste an der Stelle 1 zurück.
Und kann so benutzt werden:
	"das erste Element von <liste>"

Schreibe (das erste Element von (eine Liste, die aus 4, 5, 6 besteht)) auf eine Zeile. [4]
Schreibe (das erste Element von (eine Liste, die aus "a", "b" besteht)) auf eine Zeile. [a]
```

```ddp
Die generische Funktion Tausche mit den Parametern a und b vom Typ T Referenz und T Referenz, gibt nichts zurück, macht:
	Das T temp ist a.
	Speichere b in a.
	Speichere temp in b.
Und kann so benutzt werden:
	"Tausche <a> und <b>"

Der Text x ist "links".
Der Text y ist "rechts".
Tausche x und y.
Schreibe x auf eine Zeile. [rechts]
```

Im Funktionskörper kann man `T` wie jeden anderen Typ benutzen, zum Beispiel um Variablen zu deklarieren.

## Fehler in generischen Funktionen

Der Funktionskörper wird erst überprüft, wenn die Funktion mit einem bestimmten Typ aufgerufen wird.
Ruft man `Größtes` zum Beispiel mit zwei Wahrheitswerten auf, gibt es einen Fehler, denn der `größer als` Operator funktioniert nicht mit Wahrheitswerten.
Der Fehler wird dann beim Aufruf gemeldet.

## Einschränkungen

- Eine generische Funktion kann nicht [extern sichtbar](/Programmierung/Funktionen/Externe-Funktionen#extern-sichtbare-funktionen) sein.
- Eine generische Funktion muss sofort definiert werden. Eine [Vorwärts-Deklaration](/Programmierung/Funktionen/Anwendung#vorwärts-deklarationen) ist nicht möglich.
- Generische Funktionen können auch [extern](/Programmierung/Funktionen/Externe-Funktionen#generische-externe-funktionen) definiert sein. Dann gibt es aber ein paar zusätzliche Regeln.

Generische Funktionen können auch Operatoren überladen. Mehr dazu im Artikel [Operatoren Überladung](/Programmierung/Funktionen/Operatoren-Ueberladung).
Wie man generische Kombinationen erstellt, steht im Artikel [Kombinationen](/Programmierung/Kombinationen#generische-kombinationen).
