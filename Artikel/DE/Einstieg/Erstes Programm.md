# Erstes Programm
Nun da wir die Programmiersprache erfolgreich installiert haben können wir anfangen ein paar Programme in ihr zu schreiben.

## Hallo Welt
Mit dem Befehl `Schreibe den Text` aus der Standard Bibliothek gefolgt von einem Text kann man diesen in der Konsole ausgeben.<br>
So sieht ein übliches "Hello World" Programm aus:
```ddp
Binde "Duden/Ausgabe" ein.

Schreibe den Text "Hallo Welt!".
```
<h6 style="margin-top: 0; margin-left: 10px">Datei: HalloWelt.ddp </h6>


Das muss nun als `.ddp` Datei gespeichert werden.

Um das Programm schließlich auszuführen, verwendet man diese Befehle:
```terminal
$ kddp kompiliere HalloWelt.ddp
$ ./HalloWelt
```

Alternativ kann man den Befehl `kddp starte` benutzen:
```terminal
$ kddp starte HalloWelt.ddp
```
welcher einfach die beiden obigen Befehle ausführt.