# Funktionen
## Programm_Beenden
```
Beendet das Programm mit dem gegebenen Code.
```
* Parameter: `Code`
* Parameter Typ: `Zahl`
* Rückgabe Typ: `nichts`

### Aliase
1. `"Beende das Programm mit Code <Code>"`
2. `"beende das Programm mit Code <Code>"`

### Implementation
Implementiert in "libddpstdlib.a"
## Programm_Beenden_Standard
```
Beendet das Programm mit Code 0.
```

* Rückgabe Typ: `nichts`

### Aliase
1. `"Beende das Programm"`
2. `"beende das Programm"`

### Implementation
```ddp
	Beende das Programm mit Code 0.
```
## Laufzeitfehler
```
Wirft einen Laufzeitfehler mit einer Nachricht und einem Code.
```
* Parameter: `Nachricht`, `Code`
* Parameter Typen: `Text`, `Zahl`
* Rückgabe Typ: `nichts`

### Aliase
1. `"Löse einen Laufzeitfehler mit der Nachricht <Nachricht> und dem Code <Code> aus"`
2. `"löse einen Laufzeitfehler mit der Nachricht <Nachricht> und dem Code <Code> aus"`

### Implementation
Implementiert in "libddpstdlib.a"
## Befehlszeilenargumente
```
Gibt eine Text Liste zurück die alle übergebenen Befehlszeilenargumente enthält.
```

* Rückgabe Typ: `Text Liste`

### Aliase
1. `"die Befehlszeilenargumente"`
2. `"den Befehlszeilenargumenten"`

### Implementation
Implementiert in "libddpruntime.a"
## Betriebssystem
```
Gibt je nach Betriebssystem entweder "Windows" oder "Linux" zurück.
```

* Rückgabe Typ: `Text`

### Aliase
1. `"das Betriebssystem"`

### Implementation
Implementiert in "libddpruntime.a"
## Ist_Befehlszeile
```
Entspricht C's `isatty` funktion.
```

* Rückgabe Typ: `Boolean`

### Aliase
1. `"die Benutzereingabe eine Befehlszeile ist"`

### Implementation
Implementiert in "libddpstdlib.a"

