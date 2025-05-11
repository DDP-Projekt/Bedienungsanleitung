+++
title = "Typkonvertierung"
weight = 2
+++

# Typkonvertierung
Eine Typkonvertierung ist die Umwandlung von einem Typ zu einem anderen. In DDP sieht eine Typkonvertierung wie folgt aus:

```ddp
<Ausdruck von einem Typ> als <anderer typ>.
```

Zum Beispiel: `x als Text.`

## Konvertierungstabelle
Es können nur bestimmte Typen in andere Umgewandelt werden.

| Eingangstyp        | Ausgangstyp                                           | Besonderheiten                                                  |
|--------------------|-------------------------------------------------------|-----------------------------------------------------------------|
| Zahl               | Byte <br> Kommazahl <br> Text <br> Wahrheitswert <br> Buchstabe | Zahlen außerhalb des Wertebereichs werden angepasst<br>-<br>-<br> 0 => falsch; nicht 0 => wahr <br> benutzt ASCII wert |
| Kommazahl          | Zahl, Byte <br> Text                                  | Nachkommastellen werden abgeschnitten <br> -                    |
| Wahrheitswert      | Zahl, Byte                                            | falsch => 0; wahr => 1                                          |
| Text               | Zahl, Byte <br> Kommazahl <br> Buchstaben             | Numerischer Text Wert wird als Ziffer interpretiert<br>Text muss im Format `\d+(,\d+)?` sein<br>-<br>                                                 |
| Buchstabe          | Zahl, Byte <br> Text                                  | utf8 bytes in Dezimal <br> -                                    |
| Buchstaben Liste   | Text                                                  | -                                                               |
