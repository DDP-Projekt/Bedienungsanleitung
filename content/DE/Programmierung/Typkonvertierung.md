+++
title = "Typkonvertierung"
weight = 2
type = "article"
+++

# Typkonvertierung
Eine Typkonvertierung ist die Umwandlung von einem Typ zu einem anderen. In DDP sieht eine Typkonvertierung wie folgt aus:

```ddp
<Ausdruck von einem Typ> als <anderer typ>.
```

Zum Beispiel: `x als Text.`

## Konvertierungstabelle
Es können nur bestimmte Typen in andere Umgewandelt werden.

| Eingangstyp | Ausgangstyp                                     | Besonderheiten                                                  |
|-------------|-------------------------------------------------|-----------------------------------------------------------------|
| Zahl        | Kommazahl <br> Text <br> Boolean <br> Buchstabe | -<br>-<br> 0 => falsch; nicht 0 => wahr <br> benutzt ASCII wert |
| Kommazahl   | Zahl <br> Text                                  | Nachkommastellen werden abgeschnitten <br> -                    |
| Boolean     | Zahl                                            | falsch => 0; wahr => 1                                          |
| Text        | Zahl <br> Kommazahl <br> Buchstaben             | -<br>-<br>-<br>                                                 |
| Buchstabe   | Zahl <br> Text                                  | utf8 bytes in Dezimal <br> -                                    |
| Buchstaben  | Text                                            | -                                                               |
