# Bedienungsanleitung
Die Dokumentation für [die Deutsche Programmiersprache](https://github.com/DDP-Projekt/Kompilierer)

Die einzelnen Artikel liegen im Ordner "[Artikel](Artikel)" und sind Markdown Dateien. Ein Javascript script wandelt dieses Markdown in HTML um, sodass es auf der Webseite angezeigt werden kann.

## Lokal Ausführen
### Vorraussetzungen
1. Go

### Starten
1. Git Repository clonen
2. `config.json` erstellen (Inhalt für Standardwerte: "{}")
3. `./generate_and_run.sh`
4. `localhost:<port>/Bedienungsanleitung` im Webbrowser öffnen

### Konfiguration
Die Standardkonfiguration sieht so aus:
```json
{
	"port": "8080",
	"useHTTPS": false,
	"certPath": "",
	"keyPath": ""
}
```
