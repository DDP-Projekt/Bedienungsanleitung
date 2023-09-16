# Bedienungsanleitung
Die Dokumentation für die Deutsche Programmiersprache


Die einzelnen Artikel liegen im Ordner "Artikel" und sind Markdown Dateien. Ein Javascript script wandelt dieses Markdown in HTML um, sodass es auf der Webseite angezeigt werden kann.

## Lokal Ausführen
### Vorraussetzungen
1. Go

### Starten
1. Git Repository clonen
2. `config.json` erstellen (Inhalt für Standardwerte: "{}")
3. `go run ./server`

### Konfiguration
Die Standardkonfiguration sieht so aus:
````json
{
	"port": "8080",
	"useHTTPS": false,
	"certPath": "",
	"keyPath": ""
}
``