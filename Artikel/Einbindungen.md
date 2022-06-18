# Einbindungen

Größere und komplexere Programme bestehen oft aus mehreren Quelldateien, auf die der Programmcode verteilt wird.

DDP stellt hier für Einbindungen zur Verfügung.

## Syntax
  `Binde "<relativer Dateipfad>" ein.`

Wobei der angegebene Dateipfad relativ zur ausgeführten Datei sein muss.

An die Stelle der Einbindung wird nun einfach der Inhalt der angegebenen .ddp Datei gesetzt.

## Beispiel

#### Beispiel/test1.ddp
```ddp
Schreibe den Text "test1.ddp\n".
Die Zahl z ist 22.
```

#### Beispiel/test2.ddp
```ddp
Binde "test1" ein.

Schreibe die Zahl z.
schreibeZeile("\ntest2.ddp").
```

Wenn nun test2.ddp ausgeführt wird, ist die Ausgabe wie folgt.

#### Ausgabe
```
test1.ddp
22
test2.ddp
```

### Relative Pfade
Einbindungen erlauben relative Pfade.

Um eine Datei aus einem über- bzw. untergeordnetem Ordner einzubinden, benutzt man die folgende Syntax:

```ddp
[ Übergeordneter Ordner ]
Binde "../datei" ein.
[ Untergeordneter Ordner ]
Binde "SubOrdner/datei" ein.
```