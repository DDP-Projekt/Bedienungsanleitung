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
| Zahl               | Byte <br> Kommazahl <br> Text <br> Wahrheitswert <br> Buchstabe | Es werden nur die unteren 8 Bit behalten (Zahlen außerhalb von 0 bis 255 laufen über)<br>-<br>-<br> 0 => falsch; nicht 0 => wahr <br> Die Zahl wird als Unicode Codepoint interpretiert |
| Byte               | Zahl <br> Kommazahl <br> Text <br> Wahrheitswert <br> Buchstabe | -<br>-<br>-<br> 0 => falsch; nicht 0 => wahr <br> Der Byte wird als Unicode Codepoint interpretiert |
| Kommazahl          | Zahl, Byte <br> Text                                  | Nachkommastellen werden abgeschnitten <br> -                    |
| Wahrheitswert      | Zahl <br> Text                                        | falsch => 0; wahr => 1 <br> -                                   |
| Text               | Zahl <br> Kommazahl                                   | Die Ziffern am Anfang des Textes werden als Zahl gelesen, ist der Text keine Zahl ergibt sich 0<br>Der Text muss eine Kommazahl mit Komma als Dezimaltrennzeichen sein (z.B. `"3,14"`), ansonsten ergibt sich 0 |
| Buchstabe          | Zahl <br> Text                                        | Der Unicode Codepoint des Buchstabens <br> -                    |

Außerdem kann jeder Wert in eine Liste seines Typs umgewandelt werden, die nur diesen Wert enthält (z.B. `5 als Zahlen Liste`),
und jeder Wert kann in eine `Variable` und eine `Variable` zurück in ihren eigentlichen Typ umgewandelt werden.
Typ-Definitionen können in ihren Basistyp und zurück umgewandelt werden (siehe [Typ-Aliase und Typ-Definitionen](/Programmierung/Typ-Aliase-und-Typ-Definitionen)).
