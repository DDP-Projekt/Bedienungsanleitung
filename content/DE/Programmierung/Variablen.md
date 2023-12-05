+++
title = "Variablen"
weight = 4
type = "article"
+++

# Variablen

Variablen sind in DDP, wie in anderen Sprachen, "Container" mit Namen in denen Werte gespeichert werden.

# Deklaration

Variablen können folgendermaßen erstellt (oder deklariert) werden:

```ddp
<Datentyp mit Artikel> <Variablenname> ist <Ausdruck>.
```

Es gibt eine Besonderheit bei Variablen mit dem Datentyp "Wahrheitswert".\
Wenn man solche Variablen mit einem Ausdruck deklarieren will, sollte man stattdessen diese Syntax benutzen:
```ddp
Der Wahrheitswert <Variablenname> ist <wahr oder falsch>, wenn <Ausdruck>. 
```
Eine Liste von allen Datentypen findest du im Artikel [Datentypen](/Bedienungsanleitung/de/Programmierung/Datentypen)

## Beispiele:

```ddp
Die Zahl a ist 10.
Die Kommazahl b ist 4,32.
Der Text c ist "Hallo!".
Der Wahrheitswert d ist wahr.
Der Wahrheitswert e ist falsch wenn 1 gleich 1 ist. 
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

```ddp
Erhöhe <variable> um <a>
```  
equivalent zu  
```ddp
Speichere <variable> plus <a> in <variable>
```

## Subtraktion

```ddp
Verringere <variable> um <a>
```  
equivalent zu  
```ddp
Speichere <variable> minus <a> in <variable>
```

## Multiplikation

```ddp
Vervielfache <variable> um <a>.
```
equivalent zu  
```ddp
Speichere <variable> mal <a> in <variable>
```

## Division

```ddp
Teile <variable> durch <a>.
```
equivalent zu  
```ddp
Speichere <variable> durch <a> in <variable>
```

## Negation

```ddp
Negiere <variable>.
```
equivalent zu  
```ddp
Speichere -<variable> in <variable>
```
bzw.  
```ddp
Speichere nicht <variable> in <variable>
```

## Bitshift
```ddp
Verschiebe <variable> um <a> Bit nach Links
```
```ddp
Verschiebe <variable> um <a> Bit nach Rechts
```
equivalent zu  
```ddp
Speichere <variable> um <a> Bit nach Links verschoben in <variable>
```
```ddp
Speichere <variable> um <a> Bit nach Rechts verschoben in <variable>
```