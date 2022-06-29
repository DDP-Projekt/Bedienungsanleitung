# Typkonvertierung
Eine Typkonvertierung ist die Umwandlung von einem Typ in einen anderen. In DDP sieht eine Typkonvertierung wie folgt aus:

```ddp
<Ausdruck von einem Typ> als <anderer typ>.
```

Zum Beispiel: `x als Text.`

## Konvertierungstabelle
Es können nur bestimmte Typen in andere Umgewandelt werden.

| Eingangstyp | Ausgangstyp | Besonderheiten |
|-------------|-------------|----------------|
| Zahl | Kommazahl <br> Text <br> Boolean <br> Buchstabe | auch implizit <br>-<br> 0 => falsch; nicht 0 => wahr <br> benutzt ASCII wert |
| Kommazahl | Zahl <br> Text | Kommazahl wird trunkiert <br> - |
| Boolean | Zahl | falsch => 0; wahr => 1 |
| Text | Zahl <br> Kommazahl <br> Buchstaben | -<br>-<br>-<br> |
| Buchstabe | Zahl <br> Text | benutzt ASCII wert <br> - |
| Buchstaben | Text | - |
| Ansammlung | Liste | auch implizit |
| Ansammlung | Ansammlung anderen Typs | gleiche Regeln wie oben |
| Liste | Ansammlung | - |
| Liste | Liste anderen Typs | gleiche Regeln wie oben |

## Implizite Typkonvertierung
Eine implizite Typkonvertierung ist die Umwandlung eines Datentyps ohne eine im Quelltext stehende Anweisung, die diese Umwandlung ausführt. Nur bestimmte Datentypen haben implizite Typkonvertierungen, da sie in manchen Fällen Fehlerquellen sein können.

### Beispiel:
```ddp
Die Zahl a ist 2.
Die Kommazahl b ist a. [a wird implizit zur Kommazahl umgewandelt]
``` 