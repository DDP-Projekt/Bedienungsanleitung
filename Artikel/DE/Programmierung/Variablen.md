<to-do></to-do>

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
Der Boolean <Variablenname> ist <wahr oder falsch> wenn <Ausdruck> ist. 
```
Eine Liste von allen Datentypen findest du im Artikel [Datentypen](?p=Programmierung/Datentypen)

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

Um einer Variable das Ergebnis eines Ausdrucks zuzuweisen, muss `Speichere ... in` benutzen:
```ddp
Speichere pi durch 2 in b.
```

Es gibt noch besondere Zuweisungsoperatoren, die nur bei Variablen vom Typ Zahl oder Kommazahl benutzt werden können. Weiteres dazu erfährst du in [Mathematische Operatoren](?p=Programmierung/Mathematische%20Operatoren#zuweisungsoperatoren)