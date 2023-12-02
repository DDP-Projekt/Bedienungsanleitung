+++
title = "Updates"
weight = 4
type = "article"
+++

# Updates

Wenn du kddp installiert hast und eine neue Version herauskommt musst du es nicht komplett neu installieren.
Der `kddp update` Befehl übernimmt das für dich.

## Standardfall

Wenn du einfach nur `kddp update` ausführst wird kddp nachschauen ob eine neuere Version.
Sollte dies der fall sein wirst du gefragt ob du diese installieren möchtest.
Wenn du mit ja ('j') antwortest, wird das update beginnen.

kddp wird dann:
* die neuste Version für dein System herunterladen,
* die `kddp` executable updaten,
* die Laufzeit- und Standardbibliothek updaten,
* den DDP Sprachserver (DDPLS) updaten
  
Wenn das alles fehlerfrei geklappt hat solltest du das Terminal neustarten und mit `kddp version` die neuste Version sehen können.

## Optionen

`kddp update` nimmt mehrere optionale Kommandozeilen Optionen:

* `--wortreich`: sagt dir alles was während des Updates passiert
* `--vergleiche_version`: überprüft ob eine neue Version vorhanden ist, ohne diese zu installieren
* `--pre_release`: bezieht Github pre-releases in die Suche mit ein (wird unten erklärt)
* `--jetzt`: lädt eine neue Version, falls vorhanden, ohne zu fragen sofort runter und installiert sie

## Versionen

Eine DDP Version setzt sich aus 4 Teilen zusammen:

&gt;Major&lt;.&gt;Minor&lt;.&gt;Patch&lt;-&gt;Zusatz&lt;

Beispiel: 1.2.3-alpha

* Major: diese Zahl erhöht sich immer, wenn ein großes Feature herausgegeben wird, alter DDP-Code wird vielleicht nicht mehr funktionieren
* Minor: diese Zahl erhöht sich immer, wenn ein Feature herausgegeben wird, das mit altem Code kompatibel ist
* Patch: diese Zahl erhöht sich immer, wenn Bugfixes o.ä. herausgegeben werden
* Zusatz: einer von (-pre|-alpha|-beta). Signalisiert Test-Versionen

## Wie es funktioniert

`kddp update` schaut einfach in den [Github-Releases](https://github.com/DDP-Projekt/Kompilierer/releases) vom Kompilierer repo nach, ob es einen Release gibt, dessen Tag neuer als die aktuell installierte Version ist.
Sollte das der Fall sein lädt es die entsprechende Archiv-Datei herunter und updatet sich damit.

Sollte die `--pre_release` Option angegeben sein werden auch als "Pre-Release" markierte Releases in die Suche miteinbezogen.