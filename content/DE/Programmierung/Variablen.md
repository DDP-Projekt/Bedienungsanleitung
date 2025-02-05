+++
title = "Variablen"
weight = 4
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

Diese Syntax funktioniert auch mit Rückgaben in [Funktionen](/Programmierung/Funktionen):
```ddp
Die öffentliche Funktion Ist_Leer_Text mit dem Parameter liste vom Typ Text Liste, gibt einen Wahrheitswert zurück, macht:
	Gib wahr, wenn die Länge von liste gleich 0 ist zurück.
Und kann so benutzt werden:
	"<liste> leer ist"
```

Eine Liste von allen Datentypen findest du im Artikel [Datentypen](/Programmierung/Datentypen)

## Beispiele:

```ddp
Die Zahl a ist 10.
Die Kommazahl b ist 4,32.
Der Text c ist "Hallo!".
Der Wahrheitswert d ist wahr.
Der Wahrheitswert e ist falsch wenn 1 gleich 1 ist. 
```

# Zuweisung

Die Zuweisung ist die Änderung des Wertes einer Variable. In DDP gibt es mehrere Wege Variablen zu ändern.

Mit dem Schlüsselwort `ist` kann man Variablen nur einem Literal zuweisen:
```ddp
a ist 30.
```

Um einer Variable das Ergebnis eines Ausdrucks zuzuweisen, muss `Speichere ... in` verwendet werden:
```ddp
Speichere pi durch 2 in b.
```

# Extern sichtbare Variablen

Der DDP Kompilierer benutzt eine Technik namens [name mangling](https://en.wikipedia.org/wiki/Name_mangling). Das heißt, die Namen von Funktionen und Variablen
sind im Quellcode nicht dieselben wie in der entstehenden Binärdatei.

Wenn man also eine Variable zwischen DDP und C Quellcode teilen will muss man das name mangling ausschalten, indem man die Variable als "extern sichtbar" markiert:
```ddp
[test.ddp]
Die extern sichtbare Zahl z ist 22.
```

```c
// test.c
#include "ddptypes.h"
#include <stdio.h>

extern ddpint z;

int main(void) {
    printf("%ld", z);
    return 0;
}
```

# Spezielle Zuweisungen

Es gibt noch einige Zuweisungsoperatoren mit deren Hilfe Variablen direkt verändert werden können,
ohne sie selbst in einem Ausdruck verwenden zu müssen.
In anderen Sprachen sind das sogenannte "compound assignements" also Operatoren wie `+=, -=, *=, etc.` .
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