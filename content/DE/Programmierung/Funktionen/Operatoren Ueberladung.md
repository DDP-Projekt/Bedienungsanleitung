+++
title = "Operatoren Überladung" 
weight = 4
type = "article"
+++

# Operatoren Überladung

Die eingebauten Operatoren wie `plus` oder `als` funktionieren nur mit bestimmten Typen.
Mit Operatoren Überladung kann man sie auch für andere Typen benutzen, zum Beispiel für eigene [Kombinationen](/Programmierung/Kombinationen).
Dafür schreibt man eine Funktion, die aufgerufen wird, wenn der Operator mit den passenden Typen benutzt wird.

## Syntax

Eine Funktion, die einen Operator überlädt, hat keine Aliase.
Stattdessen steht am Ende `Und überlädt den "<Operator>" Operator.`:

```ddp
Wir nennen die Kombination aus
	der Zahl x mit Standardwert 0,
	der Zahl y mit Standardwert 0,
einen Vektor, und erstellen sie so:
	"ein Vektor aus <x> und <y>"

Die Funktion Vektor_Plus mit den Parametern a und b vom Typ Vektor und Vektor, gibt einen Vektor zurück, macht:
	Gib ein Vektor aus (x von a plus x von b) und (y von a plus y von b) zurück.
Und überlädt den "plus" Operator.

Der Vektor a ist ein Vektor aus 1 und 2.
Der Vektor b ist ein Vektor aus 2 und 2.
Der Vektor c ist a plus b. [x = 3, y = 4]
```

Wenn jetzt zwei Vektoren mit `plus` addiert werden, wird `Vektor_Plus` aufgerufen.
Die [Operator Priorisierung](/Programmierung/Operatoren#operator-priorisierung) bleibt dabei gleich.

Man kann Operatoren auch für eingebaute Typen überladen.
Eine Überladung wird dabei immer vor dem eingebauten Operator benutzt.

```ddp
Die Funktion Text_Plus_Zahl mit den Parametern t und z vom Typ Text und Zahl, gibt einen Text zurück, macht:
	Gib t verkettet mit z als Text zurück.
Und überlädt den "plus" Operator.

Schreibe ("Nummer " plus 5) auf eine Zeile. [Nummer 5]
```

## Regeln

- Die Anzahl der Parameter muss zum Operator passen. Ein unärer Operator wie `Betrag` braucht einen Parameter, ein binärer Operator wie `plus` braucht zwei und ein ternärer Operator wie `zwischen` braucht drei.
- Die Funktion muss einen Wert zurückgeben. Der Rückgabetyp darf also nicht `nichts` sein.
- Ein Operator kann für dieselben Parametertypen nur einmal überladen werden.
- Der `als` Operator ist eine Ausnahme. Er kann für denselben Parametertyp mehrmals überladen werden, solange der Rückgabetyp jedes Mal anders ist.
- Parameter dürfen auch [Referenzen](/Programmierung/Funktionen/Referenzen) sein.
- Innerhalb der Funktion selbst wird beim Benutzen des Operators nicht die Überladung aufgerufen, sondern der normale Operator.
- Öffentliche Überladungen gelten auch in anderen Modulen, die das Modul einbinden.

## Der als Operator

Mit einer Überladung des `als` Operators kann man eigene Typumwandlungen schreiben.
Der Rückgabetyp der Funktion ist dabei der Typ, in den umgewandelt wird:

```ddp
Die Funktion Vektor_Als_Text mit dem Parameter v vom Typ Vektor, gibt einen Text zurück, macht:
	Gib "(" verkettet mit (x von v) als Text verkettet mit ", " verkettet mit (y von v) als Text verkettet mit ")" zurück.
Und überlädt den "als" Operator.

Schreibe (ein Vektor aus 3 und 4 als Text) auf eine Zeile. [(3, 4)]
```

## Generische Überladungen

Auch [generische Funktionen](/Programmierung/Funktionen/Generische-Funktionen) können Operatoren überladen.
Die Überladung wird dann für alle Typen benutzt, für die der Funktionskörper funktioniert.

## Liste der Operatoren

Diese Namen können zwischen den Anführungszeichen stehen:

| Art    | Operatoren |
| ------ | ---------- |
| unär   | `"Betrag"`, `"Länge"`, `"unäres minus"`, `"nicht"`, `"logisch nicht"` |
| binär  | `"plus"`, `"minus"`, `"mal"`, `"durch"`, `"modulo"`, `"hoch"`, `"logarithmus"`, `"verkettet mit"`, `"an der Stelle"`, `"ab dem"`, `"bis zum"`, `"und"`, `"oder"`, `"entweder ... oder"`, `"logisch und"`, `"logisch oder"`, `"logisch kontra"`, `"links verschiebung"`, `"rechts verschiebung"`, `"gleich"`, `"ungleich"`, `"kleiner als"`, `"größer als"`, `"kleiner als, oder"`, `"größer als, oder"` |
| ternär | `"von bis"` (für `im Bereich von ... bis`), `"zwischen"` |
| Typumwandlung | `"als"` |
