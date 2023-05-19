# Module

Größere und komplexere Programme bestehen oft aus mehreren Quelldateien, auf die der Programmcode verteilt wird.

DDP stellt hierfür Module zur Verfügung.

## Prinzip

Jede DDP-Datei stellt ein DDP-Modul dar.
Jedes DDP-Module kann von einem anderen Eingebunden werden.
Dafür muss es aber ein nach außen sichtbares (öffentliches) Interface preisgeben.
Dies ist mit dem Schlüsselwort "öffentliche" möglich.
Das öffentliche Interface eines DDP-Modules ist die Menge aller als "öffentliche" deklarierten Variablen und Funktionen.

Wenn ein DDP-Modul ein anderes einbindet erhält es Zugriff auf dessen öffentliche Funktionen und Variablen und kann diese
selber benutzen.

In anderen Sprachen ist dieses Feature oft mit dem "import" Schlüsselwort (oder "#include" in C) umgesetzt, und die Beschränkung der Sichtbarkeit von Namen wird als Kapselung (engl. "Encapsulation") bezeichnet.

## Syntax

```ddp
Binde "<relativer Pfad>" ein.
Binde "Duden/<Datei aus der Standardbibliothek>" ein.
Binde a aus "<relativer Pfad>" ein.
Binde a, b und c aus "<relativer Pfad>" ein.
```

### Konkretes Beispiel

A.ddp:
```ddp
Binde "Duden/Ausgabe" ein.

Die öffentliche Zahl z ist 1.

Die öffentliche Funktion foo gibt nichts zurück, macht:
	Schreibe "foo" auf eine Zeile.
Und kann so benutzt werden:
	"foo"

Die Zahl zz ist 1.

Die Funktion bar gibt nichts zurück, macht:
	Schreibe "bar" auf eine Zeile.
Und kann so benutzt werden:
	"bar"
```

B.ddp:
```ddp
Binde "A" ein.

foo. [ruft foo aus A.ddp auf]
Speichere 2 in z. [setzt z aus A.ddp auf 2]

[
	bar und zz werden nicht als Funktion bzw. Variable erkannt,
	da sie in A.ddp nicht als öffentlich deklariert sind.
]
bar. 
Speichere 2 in zz. [setzt z aus A.ddp auf 2]

[
	Fehler:
	"Duden/Ausgabe" wurde zwar in A.ddp eingebunden,
	aber nicht in B.ddp, somit steht die Schreibe Funktion
	nicht zur Verfügung.
]
Schreibe "Hallo Welt" auf eine Zeile.
```

C.ddp:
```ddp
Binde foo aus "A" ein.

foo. [ruft foo aus A.ddp auf]
[
	Fehler:
	Es wurde nur foo eingebunden, aber nicht z.
]
Speichere 2 in z.

[
	Fehler:
	bar und zz werden nicht als Funktion bzw. Variable erkannt,
	da sie in A.ddp nicht als öffentlich deklariert sind.
]
bar. 
Speichere 2 in zz. [setzt z aus A.ddp auf 2]
```

## Erläuterung

Wie man im Beispiel sieht gibt man bei Einbindungen entweder einen relativen Pfad zu einer anderen .DDP-Datei an (ohne die .ddp Endung),
oder einen Pfad zu einer Datei aus der Standardbibliothek.

Es können entweder alle öffentlichen Namen, oder bloß einige spezielle eingebunden werden.
Das ist nützlich um nicht zu viele unnötige Funktions-Aliase einzubinden, die Eventuell Ärger bereiten.

Einbindungen werden zwischen Dateien auch nicht "vererbt".
Wenn B.ddp also A.ddp einbindet, so übernimmt B.ddp nicht A.ddps Einbindungen (wie im Beispiel "Duden/Ausgabe").

Um durch Verzeichnisse zu navigieren kann man unix-Dateipfade verwenden.

Bei dieser Datei-Struktur:
- root
	- Ordner1
		- A.ddp
	- Ordner2
		- B.ddp

müsste in B.ddp also `Binde "../Ordner1/A" ein.` stehen, um A.ddp einzubinden.