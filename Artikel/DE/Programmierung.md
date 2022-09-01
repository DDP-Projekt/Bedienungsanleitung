# Programmierung

Die nachfolgenden Artikel beschreiben die einzelnen Features und Eigenschaften der Deutschen Programmiersprache und wie man sie benutzt.
In diesen Artikeln wird eine Fachbezogene Terminologie benutzt, die hier erklärt werden soll.

## Ausdrücke

Ein Ausdruck (Englisch *expression*) ist ein Quellcode Abschnitt der einen Wert produziert.
Der mathematische Ausdruck `2 * pi * radius` produziert den Umfang eines Kreises.
In DDP sähe dieser Ausdruck so aus: `2 mal pi mal radius`.
Ausdrücke können mehrere Argumente von verschiedenen Typen annehmen und produzieren draus einen Wert
vom gleichen oder anderen Typ.
Die verschiedenen Datentypen sind im Artikel [Datentypen](?p=Datentypen) beschrieben.

## Anweisungen

Eine Anweisung (Englisch *statement*) ist ein Quellcode Abschnitt der etwas tut.
Anweisungen werden immer mit einem Punkt (.) beendet.
Die Anweisung `Schreibe den Text "Hallo".` gibt zum Beispiel "Hallo" in die Konsole aus.
Ausdrücke mit einem nachfolgenden Punkt sind auch Anweisungen. Ihr Ergebnis wird verworfen.
`1 plus 1.` beispielsweise wird ausgewertet, aber das Ergebnis wird nie benutzt.

Ein DDP Programm besteht aus beliebig vielen Anweisungen, die nacheinander (im Quelltext von Links nach Rechts 
und von Oben nach Unten) ausgeführt werden.

## Anweisungsblock

Ein Anweisungsblock ist eine Folge von Answeisungen, die nacheinander ausgeführt werden.
Er beginnt mit einem Doppelpunkt (:) und einer neuen Zeile, und alle zugehörigen Anweisungen
müssen mindestens eine Einrückung (ein Tab-Zeichen oder 4 Leerzeichen) tiefer eingerückt sein als der Doppelpunkt.
```ddp
:
	Schreibe den Text "Ich bin in einem Anweisungsblock!".
	Schreibe den Text "Ich auch!".
Schreibe den Text "Ich bin nicht mehr in einem Anweisungsblock".
```
Der Anweisungsblock endet wenn das nächste Zeichen genauso oder weniger tief wie der Doppelpunkt eingerückt ist.

Anweisungsblöcke werden häufig in [Verzweigungen](?p=Verzweigungen%20&%20Schleifen#verzweigungen), [Schleifen](?p=Verzweigungen%20&%20Schleifen#schleifen) oder [Funktionen](?p=Funktionen) benutzt.
Mehr dazu in den nachfolgenden Artikeln.

## Kommentare

Um den Quellcode eines DDP Programms zu erklären und verständlicher zu machen kann man Kommentare benutzen.
Kommentare werden zwischen eckigen Klammern geschrieben ([]) und sind nur zum erklären des Quellcodes da, 
sie verändern nicht was er tut.

```ddp
Schreibe den Text "test". [schreibt den Text "test" in die Konsole]
```

Kommentare können auch mitten in Anweisungen und Ausdrücken stehen.
```ddp
Wenn x [wahr] ist, Schreibe den Text "x ist wahr".
```

Außerdem können Kommentare über mehrere Zeilen gehen.
```ddp
[
	Aufgabe: Diese Funktion schreibt "Hallo Welt!" in die Konsole.
	Parameter: keine Parameter.
	Rückgabe: keine Rückgabe.
]
Die Funktion hi gibt nichts zurück, macht:
	Schreibe den Text "Hallo Welt!" auf eine Zeile.
und kann so benutzt werden:
	"Schreibe Hallo Welt"
```

## Literale

Ein Literal ist eine Schreibweise im Quellcode die einen (in DDP konstanten) Wert repräsentiert.
Eine Liste von Literalen für die ensprechenden Datentypen kann im Artikel [Datentypen](?p=Datentypen) gefunden werden.


## Bemerkung

Viele, wenn nicht gar alle, der Features die in den nachfolgenden Artikeln beschrieben werden sollten jedem der schon einmal programmiert hat bereits bekannt sein.
Dennoch lohnt es sich, wenn auch nur für die Syntax, alles genaustens zu lesen, da DDP oftmals einige Details abändert um der Deutschen Sprache gerecht zu werden.