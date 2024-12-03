# Bedienungsanleitung 
[![Deploy Hugo site to Pages](https://github.com/DDP-Projekt/Bedienungsanleitung/actions/workflows/hugo.yaml/badge.svg)](https://github.com/DDP-Projekt/Bedienungsanleitung/actions/workflows/hugo.yaml)

Die Dokumentation für [die Deutsche Programmiersprache](https://github.com/DDP-Projekt/Kompilierer)

Die einzelnen Artikel liegen im Ordner "[content](content)" und sind Markdown Dateien.
Hugo nutzt die Markdown Dateien um die Website zu generieren.

## Lokal Ausführen
### Voraussetzungen
1. [Hugo](https://gohugo.io/installation/)
2. [Go](https://go.dev/dl/)

### Starten
1. Git Repository klonen
2. `go run ./gen` - Artikel für die Standardbibliothek generieren
3. `hugo serve` - Webserver starten

Mit dem Befehl `hugo` generiert man alle html Dateien, welche im Ordner `/public` gespeichert werden.

## Mitwirken
Über Issues oder Pull-Requests kannst du uns helfen Fehler in der Dokumentation zu beheben.

### Artikel erstellen
Um Artikel zu erstellen benutzt man den `hugo new content <pfad>` Befehl.

Alle Pfade gehen von `/content/DE/` aus, also würde `hugo new content hallo.md` die Datei `/content/DE/hallo.md` erstellen.

Jeder Artikel beginnt mit einem front-matter der so aussieht:
```
+++
title = ""
weight = 1
type = "article"
+++
```

Das Feld `title` gibt den Text an, der in der Seitenleiste für den Artikel angezeigt wird. `weight` bestimmt die Reihenfolge. `type` muss auf "article" bleiben. Diese Felder müssen immer ausgefüllt werden.

Der Pfad der Datei bestimmt die URL des Artikels.

### Sections erstellen
Sections sind Verzeichnisse die mehrere Artikel enthalten. In der Sidebar werden Sections als Dropdown-Menü angezeigt.

Jede Section hat einen speziellen Artikel mit dem Namen: "`_index.md`".
Diese Datei beginnt mit einem front-matter, welches so aussieht:
```
+++
title = ""
weight = 1
type = "article"
layout = "single"
+++
```

### Die Website selbst
Das HTML für die Website außerhalb der Artikel befindet sich im [`/layouts`](/layouts/) Verzeichnis.

[`/_default/baseof.html`](/layouts/_default/baseof.html) ist das Kern der Website und bindet alle partials ein.

[`/article/single.html`](/layouts/article/single.html) enthält das Layout für Artikel.

[`/partials`](/layouts/partials/) enthält kleinere Teile der Website:
- [`artikel.html`](/layouts/partials/artikel.html)
- [`head.html`](/layouts/partials/head.html)
- [`header.html`](/layouts/partials/header.html)
- [`sidebar.html`](/layouts/partials/sidebar.html)

[`/layouts/404.html`](/layouts/404.html) definiert die Fehlermeldung falls eine Seite nicht gefunden wurde.

Bilder, Fonts, CSS und JS Dateien befinden sich in [`/assets`](/assets/) ihren jeweiligen Ordnern.

Das favicon liegt in [`/static/favicon`](/static/favicon/).

Übersetzungsschlüssel für jede Sprache findet man in [`/i18n`](/i18n/).
