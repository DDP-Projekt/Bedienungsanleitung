# Bedienungsanleitung
Die Dokumentation für [die Deutsche Programmiersprache](https://github.com/DDP-Projekt/Kompilierer)

Die einzelnen Artikel liegen im Ordner "[content](content)" und sind Markdown Dateien.
Hugo nutzt die Markdown Dateien um die Website zu generieren.

## Lokal Ausführen
### Vorraussetzungen
1. Hugo
2. Go

### Starten
1. Git Repository Klonen
2. `cd gen` 
3. `go run .` - Artikel für die Standardbibliothek generieren
4. `cd ..`
5. `hugo serve` - Webserver starten

Mit dem Befehl `hugo` generiert man alle html Dateien welche im Ordner `/public` gespeichert werden.

## Mitwirken
Über Issues oder Pull-Requests kannst du uns helfen Fehler in der Dokumenation zu beheben.

### Artikel erstellen
Um Artikel zu erstellen benutzt man den `hugo new content <pfad>` Befehl.

Alle Pfade gehen von /content/DE/ aus, also würde `hugo new content hallo.md` die Datei `/content/DE/hallo.md` erstellen.

Jeder Artikel hat einen header der so aussieht:
```toml
+++
title = ""
weight = 1
+++
```

Das Feld "title" gibt den Text an, der in der Seitenleiste für den Artikel angezeigt wird. "weight" bestimmt die Reihenfolge. Diese Felder müssen immer ausgefüllt werden.

Der Pfad der Datei bestimmt die URL des Artikels.

## Die Website selbst
Das HTML für die Website außerhalb der Artikel befindet sich im [`/layouts`](/layouts/) Verzeichnis.

[`/layouts/_default`](/layouts/_default/) enthält das Layout für jeden Seiten Typ. [`list.html`](/layouts/_default/list.html) und [`single.html`](/layouts/_default/single.html) sind identisch und zeigen nur den Artikel an.
[`baseof.html`](/layouts/_default/baseof.html) ist das Kern der Website und bindet alle partials ein.

[`/layouts/partials`](/layouts/partials/) enthält kleinere Teile der Website:
- [`artikel.html`](/layouts/partials/artikel.html)
- [`head.html`](/layouts/partials/head.html)
- [`header.html`](/layouts/partials/header.html)
- [`sidebar.html`](/layouts/partials/sidebar.html)

[`/layouts/404.html`](/layouts/404.html) definiert die Fehlermeldung falls eine Seite nicht gefunden wurde.