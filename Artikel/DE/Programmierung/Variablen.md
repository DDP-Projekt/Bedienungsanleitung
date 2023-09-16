# Variablen

Variablen sind in DDP, wie in anderen Sprachen, "Container" mit Namen in denen Werte gespeichert werden.

# Deklaration

Variablen können folgendermaßen erstellt (oder deklariert) werden:

```ddp
<Datentyp mit Artikel> <Variablenname> ist <Ausdruck>.
```

Es gibt eine Besonderheit bei Variablen mit dem Datentyp "Boolean".\
Wenn man solche Variablen mit einem Ausdruck deklarieren will, sollte man stattdessen diese Syntax benutzen:
```ddp
Der Boolean <Variablenname> ist <wahr oder falsch>, wenn <Ausdruck>. 
```
Eine Liste von allen Datentypen findest du im Artikel [Datentypen](/Bedienungsanleitung/DE/Programmierung/Datentypen)

## Beispiele:

```dpp
Die Zahl a ist 10.
Die Kommazahl b ist 4,32.
Der Text c ist "Hallo!".
Der Boolean d ist wahr.
Der Boolean e ist falsch wenn 1 gleich 1 ist. 
```

# Zuweisung

Die Zuweisung ist die änderung des Wertes einer Variable. In DDP gibt es mehrere Wege Variablen zu ändern.

Mit dem Schlüsselwort `ist` kann man Variablen nur einem Literal zuweisen:
```ddp
a ist 30.
```

Um einer Variable das Ergebnis eines Ausdrucks zuzuweisen, muss `Speichere ... in` verwendet werden:
```ddp
Speichere pi durch 2 in b.
```

# Spezielle Zuweisungen

Es gibt noch eineige Zuweisungsoperatoren mit deren Hilfe Variablen direkt verändert werden können,
ohne sie selbst in einem Ausdruck verwenden zu müssen.
In anderen Sprachen sind das sogennatne "compound assignements" also Operatoren wie `+=, -=, *=, etc.` .
Diese Operatoren können verwendet werden um Code lesbarer zu gestalten.

## Addition

`Erhöhe <variable> um <a>`  
equivalent zu  
`Speichere <variable> plus <a> in <variable>`

## Subtraktion

`Verringere <variable> um <a>`  
equivalent zu  
`Speichere <variable> minus <a> in <variable>`

## Multiplikation

`Vervielfache <variable> um <a>.`  
equivalent zu  
`Speichere <variable> mal <a> in <variable>`

## Division

`Teile <variable> durch <a>.`  
equivalent zu  
`Speichere <variable> durch <a> in <variable>`

## Negation

`Negiere <variable>.`  
equivalent zu  
`Speichere -<variable> in <variable>` bzw.  
`Speichere nicht <variable> in <variable>`

## Bitshift

`Verschiebe <variable> um <a> Bit nach Links`  
`Verschiebe <variable> um <a> Bit nach Rechts`  
equivalent zu  
`Speichere <variable> um <a> Bit nach Links verschoben in <variable>`  
`Speichere <variable> um <a> Bit nach Rechts verschoben in <variable>`