+++
title = "Funktionsdeklarationen"
weight = 1
+++

# Funktionsdeklarationen

Jede Funktion muss deklariert werden bevor man sie benutzen kann. In DDP gibt es nur globale Funktionen, also solche, die sich nicht innerhalb eines Anweisungsblocks befinden.
Jede Funktion braucht einen einzigartigen **Namen**, eine Liste von **Parametern** und deren **Typen**, einen **Rückgabetyp** und einen oder mehrere **Aliase**.

Hier ist die allgemeine Form einer Funktionsdeklaration (optionale Bestandteile sind in Klammern gefasst):
```ddp
Die Funktion <Name> (mit den Parametern x, y und z vom Typ T1, T2 und T3,) gibt <Rückgabetyp> zurück, macht:
	<Funktionskörper>
Und kann so benutzt werden:
	"Alias mit Parameter <x> <y> und <z>" (,
	"<z> Noch ein Alias mit allen Parametern <y> <x>" oder
	...)
```

## Der Name

einer Funktion muss einmalig sein und sollte die Funktion gut beschreiben, ist aber ansonsten nicht weiter wichtig und wird im Code nicht direkt verwendet. Wichtiger wird er bei Externen Funktionen, die später beschrieben werden.

## Die Parameter

einer Funktion sind optional. Es können zwischen 0 und unendlich viele vorhanden sein. Jeder Parameter benötigt einen Typ.
Beim Schreiben der Parameter sind der Numerus (Singular, Plural) und die Aufzählungsregeln zu beachten.

Bei einem einzigen Parameter schreibt man 
```ddp
... mit dem Parameter x vom Typ T, ...
```
Bei zweien schreibt man 
```ddp
... mit den Parametern x und y vom Typ T1 und T2, ...
```
Und bei N Parameter 
```ddp
... mit den Parametern a, b, ... und z vom Typ T1, T2, ... und T26, ...
```

## Der Rückgabetyp

einer Funktion muss vorhanden sein, darf aber `nichts` sein.
`nichts` ist kein Typ, sondern ein Platzhalter für Funktionen, die keinen Wert erzeugen sondern nur Nebeneffekte haben.
Das Ergebnis einer Funktion die `nichts` zurückgibt kann nicht verwendet werden, weil, nun ja, es nicht vorhanden ist.

Sollte eine Funktion einen Wert zurückgeben, muss eine Rückgabe Anweisung am Ende des Funktionskörpers stehen.
Eine Rückgabe Anweisung sieht so aus:
```ddp
Gib <Wert vom Rückgabetyp> zurück.
```
Eine Rückgabe Anweisung kann überall im Funktionskörper auftauchen, und beendet die Funktion mit dem Ergebnis des zurückgegebenen Wertes.

Um Funktionen, die `nichts` zurückgeben frühzeitig zu verlassen, kann man `Verlasse die Funktion` verwenden.

## Aliase

sind die Art wie Funktionen aufgerufen werden.<br>
Da es das Ziel von DDP ist, so gutes Deutsch wie möglich zu sein ist es schwer Funktionen in jedem grammatischen Kontext korrekt aufzurufen. Mit Aliasen haben wir einen Weg gefunden das möglich zu machen.

Ein Alias ist ein Wort, Satzbaustein, Satz o.ä. in Form eines Text Literals, der alle Parameter seiner Funktion enthält.
Aliase werden vom Lexer wie gewöhnlicher DDP Code interpretiert, jedes DDP Schlüsselwort und jeder valide DDP-Name kann also darin verwendet werden.<br>

Jeder Alias muss einzigartig sein.
Einzigartigkeit ist in diesem Fall allerdings nicht so eng zu sehen, denn die Parameter in Aliasen sind Typsensitiv.
Das heißt, zwei Funktionen, die beide einen Parameter nehmen können dieselben Aliase haben, wenn die Parameter verschiedenen Typs sind.<br>
Das beste Beispiel ist der Duden. Er definiert 5 Schreibe_X Funktionen, eine für jeden Basistyp.
Die Aliase sind alle in der Form `"Schreibe <p1>"`.
Das ist möglich, da p1 immer einen anderen Typ besitzt, der beim Funktionsaufruf erkannt wird.

Jede Funktion braucht mindestens einen Alias, es wird aber empfohlen mehrere zu erstellen um grammatischen Kontexten gerecht zu werden.<br>
Die allgemeine Form der Alias Deklaration am Ende einer Funktionsdeklaration sieht so aus:
```ddp
...
Und kann so benutzt werden:
	"Alias1",
	"Alias2" oder
	"Alias3"
```
Das Komma und 'oder' können darin synonym verwendet werden.
Wenn eine Funktion Parameter besitzt, müssen diese alle im Alias vorhanden sein in der Form `<parameter-name>`.

Aliase können auch außerhalb der Funktionsdeklaration definiert werden, mit der Form:
```ddp
Der Alias "<Alias>" steht für die Funktion <Funktionsname>.
```
Das ist nützlich um Funktionen aus Bibliotheken (z.B. dem Duden) Aliase für grammatische Kontexte hinzuzufügen, an die der Ersteller der Funktion nicht denken konnte.<br>

**Aber Achtung!!**<br>
In solchen späten Aliasdeklarationen, müssen eventuelle Parameter denselben Namen wie in der Funktion besitzen.<br>
Möchte man also z.B. einer Funktion aus dem Duden einen Alias hinzufügen, muss man nachschauen wie die Parameter der Funktion heißen.


# Funktionsaufrufe

Funktionen werden ausschließlich durch ihre Aliase aufgerufen.<br>
Ein Funktionsaufruf ist technisch gesehen ein Ausdruck, und kann daher überall wo ein Ausdruck stehen kann verwendet werden.

Beim Übergeben von Parametern gibt es einige simple Regeln zu beachten:
- Parameter, die nur 1 Token lang sind (z.B. Literale wie "hi", 22 oder wahr) können einfach an der Stelle des Parameters stehen
- Parameter, die mehr als 1 Token lang sind (z.B. ein anderer Funktionsaufruf oder ein Listen Literal) müssen in Klammern gefasst werden
- Eine Ausnahme bildet der `-` Operator, da es schöner aussieht `-22` nicht in Klammern fassen zu müssen.

Bei Funktionen, die den gleichen Alias mit Parametern von unterschiedlichem Typ besitzen, werden die Typen der übergebenen Argumente erkannt, und die richtige Funktion ausgewählt.

## Beispiele

```ddp
Schreibe "Hi" auf eine Zeile.
Schreibe -22 auf eine Zeile.
Schreibe (einen Funktionsaufruf) auf eine Zeile.
```