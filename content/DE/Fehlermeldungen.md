+++
title = "Fehlermeldungen"
weight = 5
+++

# Fehlermeldungen

Wenn der Kompilierer einen Fehler im Quellcode findet, gibt er eine Fehlermeldung aus.
Eine Fehlermeldung sieht zum Beispiel so aus:

```terminal
Typ Fehler (3001) in test.ddp (Z: 1, S: 16)

1 |  Die Zahl z ist "hallo".
  |                 ^^^^^^^

Ein Wert vom Typ Text kann keiner Variable vom Typ Zahl zugewiesen werden.
```

Die erste Zeile enthält:
- die Art des Fehlers (hier `Typ Fehler`),
- den Fehlercode in Klammern (hier `3001`),
- die Datei, in der der Fehler gefunden wurde,
- die Zeile (`Z`) und Spalte (`S`) des Fehlers.

Darunter wird die fehlerhafte Zeile angezeigt. Die `^` Zeichen markieren die genaue Stelle.
Am Ende steht eine Beschreibung des Fehlers.

Manche Meldungen sind nur Warnungen. Bei einer Warnung wird das Programm trotzdem kompiliert.

Die erste Ziffer des Fehlercodes verrät die Art des Fehlers:

| Fehlercodes | Art                     | Bedeutung                                                         |
| ----------- | ----------------------- | ----------------------------------------------------------------- |
| 0000        | Sonstiger Fehler        | Fehler, die in keine andere Kategorie passen                      |
| 1000 - 1999 | Syntax Fehler           | Der Quellcode ist nicht richtig aufgebaut                         |
| 2000 - 2999 | Semantischer Fehler     | Der Quellcode ist richtig aufgebaut, ergibt aber keinen Sinn       |
| 3000 - 3999 | Typ Fehler              | Ein Wert hat den falschen Typ                                     |

## Liste der Fehlercodes

### Sonstige Fehler

| Code | Beschreibung |
| ---- | ------------ |
| 0000 | Ein Modul konnte nicht eingebunden werden, zum Beispiel weil die Datei nicht existiert oder zwei Module sich gegenseitig einbinden |

### Syntax Fehler

| Code | Beschreibung |
| ---- | ------------ |
| 1000 | Ein unerwartetes Wort oder Zeichen wurde gefunden. Oft fehlt ein Punkt oder ein Komma |
| 1001 | Es wurde ein Literal erwartet, aber ein Ausdruck gefunden |
| 1002 | Es wurde ein Name erwartet, zum Beispiel ein Parametername |
| 1003 | Es wurde ein Typname erwartet |
| 1004 | Es wurde ein Großbuchstabe erwartet |
| 1005 | Ein Literal ist fehlerhaft, zum Beispiel wegen einer ungültigen Escape-Sequenz |
| 1006 | Der Pfad einer Einbindung ist fehlerhaft |
| 1007 | Ein Alias ist fehlerhaft aufgebaut |
| 1008 | Der Text ist kein gültiges UTF-8 |
| 1009 | Der Artikel oder das Pronomen passt nicht zum grammatikalischen Geschlecht des Typs, zum Beispiel `Der Zahl z ist 1.` |
| 1010 | Der Name eines überladenen Operators ist ungültig |

### Semantische Fehler

| Code | Beschreibung |
| ---- | ------------ |
| 2000 | Der Name wurde bereits benutzt |
| 2001 | Der Name wurde noch nicht deklariert |
| 2002 | Die Anzahl der Parameternamen und Parametertypen einer Funktion ist unterschiedlich |
| 2003 | Bei einer externen Funktion wurde kein Pfad zu einer .c, .lib, .a oder .o Datei angegeben |
| 2004 | Eine Funktion wurde nicht im globalen Bereich deklariert |
| 2005 | Einer Funktion fehlt die Rückgabe Anweisung |
| 2006 | Ein Alias ist fehlerhaft, zum Beispiel weil Parameter fehlen |
| 2007 | Der Alias wird bereits von einer anderen Funktion oder Kombination benutzt |
| 2008 | Ein Alias aus einem eingebundenen Modul existiert bereits |
| 2009 | Ein nachträglicher Alias wurde nicht im globalen Bereich deklariert |
| 2010 | Die Argumente eines nachträglichen Alias passen nicht zur Funktion |
| 2011 | Eine Rückgabe Anweisung steht außerhalb einer Funktion |
| 2012 | Ein Funktionsname wurde statt eines Variablennamens benutzt oder umgekehrt |
| 2013 | Eine nicht globale Deklaration wurde als öffentlich markiert |
| 2014 | Ein Typ wurde nicht im globalen Bereich deklariert |
| 2015 | Ein Feld existiert in dieser Kombination nicht |
| 2016 | Das Wort öffentlich fehlt oder steht an einer falschen Stelle |
| 2017 | `Verlasse die Schleife` oder `Fahre mit der Schleife fort` steht außerhalb einer Schleife |
| 2018 | Ein Typ ist unbekannt oder wurde nicht eingebunden |
| 2019 | Eine externe Funktion wurde unnötigerweise auch als extern sichtbar markiert |
| 2020 | Die Anzahl der Parameter passt nicht zum überladenen Operator |
| 2021 | Der Operator ist für diese Parametertypen bereits überladen |
| 2022 | (Warnung) Ein `...` Platzhalter wurde gefunden |
| 2023 | Ein Typ kann nicht als `Variable` definiert werden |
| 2024 | Eine Funktion wurde mit `wird später definiert` deklariert, aber nie definiert |
| 2025 | Eine Funktion aus einem anderen Modul wurde definiert |
| 2026 | Eine Funktion wurde bereits definiert |
| 2027 | Eine generische Funktion hat keinen Typparameter |
| 2028 | Eine generische Kombination hat keinen Typparameter |
| 2029 | Eine generische Funktion wurde als extern sichtbar markiert |
| 2030 | Eine generische Funktion wurde ohne Definition deklariert |
| 2031 | Beim Erstellen einer generischen Funktion für bestimmte Typen gab es Fehler |
| 2032 | Ein nicht generischer Typ wurde mit Typparametern benutzt |
| 2033 | Bei einem Alias einer generischen Kombination konnten nicht alle Typparameter herausgefunden werden |
| 2034 | Einer Konstanten kann kein Wert zugewiesen werden |

### Typ Fehler

| Code | Beschreibung |
| ---- | ------------ |
| 3000 | Ein Operator oder eine Funktion hat einen Wert vom falschen Typ bekommen |
| 3001 | Ein Wert vom falschen Typ wurde einer Variable zugewiesen |
| 3002 | Eine Indizierung ist fehlerhaft, zum Beispiel weil der Index keine Zahl ist |
| 3003 | Die Werte in einem Listen Literal haben nicht alle denselben Typ |
| 3004 | Eine Typumwandlung mit `als` ist nicht möglich |
| 3005 | Ein Referenz-Parameter hat keine Variable bekommen |
| 3006 | Ein Wert kann nicht als Referenz übergeben werden, zum Beispiel ein Buchstabe in einem Text |
| 3007 | Die Bedingung ist kein Wahrheitswert |
| 3008 | Ein Wert in einer zählenden Schleife hat den falschen Typ |
| 3009 | Der zurückgegebene Wert passt nicht zum Rückgabetyp der Funktion |
| 3010 | Der `von` Operator wurde mit einem Wert benutzt, der keine Kombination ist |
| 3011 | Ein privates Feld wurde aus einem anderen Modul benutzt |
| 3012 | Ein überladener Operator gibt `nichts` zurück |
| 3013 | Ein Typparameter wurde als Referenz angegeben |
| 3014 | Der Typ eines generischen Feldes konnte nicht herausgefunden werden |
| 3015 | Ein generischer Typ konnte mit den angegebenen Typen nicht erstellt werden |
| 3016 | Ein generischer Parameter einer externen Funktion ist keine Liste oder Referenz |
