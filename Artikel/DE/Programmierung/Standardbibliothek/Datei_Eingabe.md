# Duden/Datei_Eingabe Funktionen
## Lies_Text_Datei
```
Die Funktion Lies_Text_Datei speichert den Inhalt der Datei, die an dem gegebenen Pfad liegt, in ref und gibt die Anzahl der Bytes der gelesenen Datei zurück.
Wenn ein Fehler auftreten sollte, ist der zurückgegebene Wert negativ und die Fehler meldung in ref geschrieben
```
* Parameter: `Pfad`, `ref`
* Parameter Typen: `Text`, `Text Referenz`
* Rückgabe Typ: `Zahl`

### Aliase
1. `"Lies den Text in <Pfad> und speichere ihn in <ref>"`
2. `"die Anzahl der Bytes, die aus <Pfad> gelesen und in <ref> gespeichert wurden"`

### Implementation
Implementiert in "libddpstdlib.a"

