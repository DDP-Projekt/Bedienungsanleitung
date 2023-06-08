# Funktionen
## SchreibeZeile_Zahl
```
Die Funktion SchreibeZeile_Zahl schreibt eine gegebene Zahl (p1) gefolgt von einer neuen Zeile in den Standart Output Stream.
```
* Parameter: `p1`
* Parameter Typ: `Zahl`
* Rückgabe Typ: `nichts`

### Aliase
1. `"Schreibe die Zahl <p1> auf eine Zeile"`
2. `"Schreibe <p1> auf eine Zeile"`

### Implementation
```ddp
	Schreibe p1.
	Schreibe '\n'.
```
## Schreibe_Kommazahlen_Liste
```
Die Funktion Schreibe_Kommazahlen_Liste schreibt alle Elemente einer Kommazahlen Liste getrennt mit einem Komma in den Standart Output Stream.
```
* Parameter: `p1`
* Parameter Typ: `Kommazahlen Liste`
* Rückgabe Typ: `nichts`

### Aliase
1. `"Schreibe die Kommazahlen Liste <p1>"`
2. `"Schreibe <p1>"`

### Implementation
```ddp
	Schreibe p1 mit dem Seperator ", ".
```
## Schreibe_Boolean_Liste
```
Die Funktion Schreibe_Boolean_Liste schreibt alle Elemente einer Boolean Liste getrennt mit einem Komma in den Standart Output Stream.
```
* Parameter: `p1`
* Parameter Typ: `Boolean Liste`
* Rückgabe Typ: `nichts`

### Aliase
1. `"Schreibe die Boolean Liste <p1>"`
2. `"Schreibe <p1>"`

### Implementation
```ddp
	Schreibe p1 mit dem Seperator ", ".
```
## Schreibe_Buchstabe
```
Die Funktion Schreibe_Buchstabe schreibt einen gegebenen Buchstaben (p1) in den Standart Output Stream.
```
* Parameter: `p1`
* Parameter Typ: `Buchstabe`
* Rückgabe Typ: `nichts`

### Aliase
1. `"Schreibe den Buchstaben <p1>"`
2. `"Schreibe <p1>"`

### Implementation
Implementiert in "libddpstdlib.a"
## Schreibe_Text_Liste_Getrennt
```
Die Funktion Schreibe_Text_Liste_Getrennt schreibt alle Elemente einer gegebenen Text Liste (liste) getrennt mit einem Text (seperator) in den Standart Output Stream.
```
* Parameter: `liste`, `seperator`
* Parameter Typen: `Text Liste`, `Text`
* Rückgabe Typ: `nichts`

### Aliase
1. `"Schreibe die Text Liste <liste> mit dem Seperator <seperator>"`
2. `"Schreibe <liste> mit dem Seperator <seperator>"`

### Implementation
```ddp
	Wenn die Länge von liste größer als 0 ist, dann:
		Wenn die Länge von liste größer als 1 ist, dann:
			Für jede Zahl i von 1 bis die Länge von liste minus 1, mache:
				Schreibe (liste an der Stelle i).
				Schreibe seperator.
		Schreibe (liste an der Stelle (die Länge von liste)).
```
## Schreibe_Buchstaben_Liste
```
Die Funktion Schreibe_Buchstaben_Liste schreibt alle Elemente einer Buchstaben Liste getrennt mit einem Komma in den Standart Output Stream.
```
* Parameter: `p1`
* Parameter Typ: `Buchstaben Liste`
* Rückgabe Typ: `nichts`

### Aliase
1. `"Schreibe die Buchstaben Liste <p1>"`
2. `"Schreibe <p1>"`

### Implementation
```ddp
	Schreibe p1 mit dem Seperator ", ".
```
## Schreibe_Zahlen_Liste_Zeile
```
Die Funktion Schreibe_Zahlen_Liste schreibt alle Elemente einer Zahlen Liste getrennt mit einem Komma und gefolgt von einer neuen Zeile in den Standart Output Stream.
```
* Parameter: `p1`
* Parameter Typ: `Zahlen Liste`
* Rückgabe Typ: `nichts`

### Aliase
1. `"Schreibe die Zahlen Liste <p1> auf eine Zeile"`
2. `"Schreibe <p1> auf eine Zeile"`

### Implementation
```ddp
	Schreibe p1 mit dem Seperator ", ".
	Schreibe '\n'.
```
## Schreibe_Buchstaben_Liste_Getrennt
```
Die Funktion Schreibe_Buchstaben_Liste_Getrennt schreibt alle Elemente einer gegebenen Buchstaben Liste (liste) getrennt mit einem Text (seperator) in den Standart Output Stream.
```
* Parameter: `liste`, `seperator`
* Parameter Typen: `Buchstaben Liste`, `Text`
* Rückgabe Typ: `nichts`

### Aliase
1. `"Schreibe die Buchstaben Liste <liste> mit dem Seperator <seperator>"`
2. `"Schreibe <liste> mit dem Seperator <seperator>"`

### Implementation
```ddp
	Wenn die Länge von liste größer als 0 ist, dann:
		Wenn die Länge von liste größer als 1 ist, dann:
			Für jede Zahl i von 1 bis die Länge von liste minus 1, mache:
				Schreibe (liste an der Stelle i).
				Schreibe seperator.
		Schreibe (liste an der Stelle (die Länge von liste)).
```
## Schreibe_Kommazahlen_Liste_Zeile
```
Die Funktion Schreibe_Kommazahlen_Liste schreibt alle Elemente einer Kommazahlen Liste getrennt mit einem Komma und gefolgt von einer neuen Zeile in den Standart Output Stream.
```
* Parameter: `p1`
* Parameter Typ: `Kommazahlen Liste`
* Rückgabe Typ: `nichts`

### Aliase
1. `"Schreibe die Kommazahlen Liste <p1> auf eine Zeile"`
2. `"Schreibe <p1> auf eine Zeile"`

### Implementation
```ddp
	Schreibe p1 mit dem Seperator ", ".
	Schreibe '\n'.
```
## Schreibe_Boolean
```
Die Funktion Schreibe_Boolean schreibt einen gegebenen Boolean (p1) in den Standart Output Stream.
```
* Parameter: `p1`
* Parameter Typ: `Boolean`
* Rückgabe Typ: `nichts`

### Aliase
1. `"Schreibe den Boolean <p1>"`
2. `"Schreibe <p1>"`

### Implementation
Implementiert in "libddpstdlib.a"
## SchreibeZeile_Buchstabe
```
Die Funktion SchreibeZeile_Buchstabe schreibt einen gegebenen Buchstaben (p1) gefolgt von einer neuen Zeile in den Standart Output Stream.
```
* Parameter: `p1`
* Parameter Typ: `Buchstabe`
* Rückgabe Typ: `nichts`

### Aliase
1. `"Schreibe den Buchstaben <p1> auf eine Zeile"`
2. `"Schreibe <p1> auf eine Zeile"`

### Implementation
```ddp
	Schreibe p1.
	Schreibe '\n'.
```
## SchreibeZeile_Text
```
Die Funktion SchreibeZeile_Text schreibt einen gegebenen Text (p1) gefolgt von einer neuen Zeile in den Standart Output Stream.
```
* Parameter: `p1`
* Parameter Typ: `Text`
* Rückgabe Typ: `nichts`

### Aliase
1. `"Schreibe den Text <p1> auf eine Zeile"`
2. `"Schreibe <p1> auf eine Zeile"`

### Implementation
```ddp
	Schreibe p1.
	Schreibe '\n'.
```
## Schreibe_Zahlen_Liste_Getrennt
```
Die Funktion Schreibe_Zahlen_Liste_Getrennt schreibt alle Elemente einer gegebenen Zahlen Liste (liste) getrennt mit einem Text (seperator) in den Standart Output Stream.
```
* Parameter: `liste`, `seperator`
* Parameter Typen: `Zahlen Liste`, `Text`
* Rückgabe Typ: `nichts`

### Aliase
1. `"Schreibe die Zahlen Liste <liste> mit dem Seperator <seperator>"`
2. `"Schreibe <liste> mit dem Seperator <seperator>"`

### Implementation
```ddp
	Wenn die Länge von liste größer als 0 ist, dann:
		Wenn die Länge von liste größer als 1 ist, dann:
			Für jede Zahl i von 1 bis die Länge von liste minus 1, mache:
				Schreibe (liste an der Stelle i).
				Schreibe seperator.
		Schreibe (liste an der Stelle (die Länge von liste)).
```
## Schreibe_Kommazahlen_Liste_Getrennt
```
Die Funktion Schreibe_Kommazahlen_Liste_Getrennt schreibt alle Elemente einer gegebenen Kommazahlen Liste (liste) getrennt mit einem Text (seperator) in den Standart Output Stream.
```
* Parameter: `liste`, `seperator`
* Parameter Typen: `Kommazahlen Liste`, `Text`
* Rückgabe Typ: `nichts`

### Aliase
1. `"Schreibe die Kommazahlen Liste <liste> mit dem Seperator <seperator>"`
2. `"Schreibe <liste> mit dem Seperator <seperator>"`

### Implementation
```ddp
	Wenn die Länge von liste größer als 0 ist, dann:
		Wenn die Länge von liste größer als 1 ist, dann:
			Für jede Zahl i von 1 bis die Länge von liste minus 1, mache:
				Schreibe (liste an der Stelle i).
				Schreibe seperator.
		Schreibe (liste an der Stelle (die Länge von liste)).
```
## Schreibe_Boolean_Liste_Getrennt
```
Die Funktion Schreibe_Boolean_Liste_Getrennt schreibt alle Elemente einer gegebenen Boolean Liste (liste) getrennt mit einem Text (seperator) in den Standart Output Stream.
```
* Parameter: `liste`, `seperator`
* Parameter Typen: `Boolean Liste`, `Text`
* Rückgabe Typ: `nichts`

### Aliase
1. `"Schreibe die Boolean Liste <liste> mit dem Seperator <seperator>"`
2. `"Schreibe <liste> mit dem Seperator <seperator>"`

### Implementation
```ddp
	Wenn die Länge von liste größer als 0 ist, dann:
		Wenn die Länge von liste größer als 1 ist, dann:
			Für jede Zahl i von 1 bis die Länge von liste minus 1, mache:
				Schreibe (liste an der Stelle i).
				Schreibe seperator.
		Schreibe (liste an der Stelle (die Länge von liste)).
```
## Schreibe_Text_Liste_Zeile
```
Die Funktion Schreibe_Text_Liste schreibt alle Elemente einer Text Liste getrennt mit einem Komma und gefolgt von einer neuen Zeile in den Standart Output Stream.
```
* Parameter: `p1`
* Parameter Typ: `Text Liste`
* Rückgabe Typ: `nichts`

### Aliase
1. `"Schreibe die Text Liste <p1> auf eine Zeile"`
2. `"Schreibe <p1> auf eine Zeile"`

### Implementation
```ddp
	Schreibe p1 mit dem Seperator ", ".
	Schreibe '\n'.
```
## Schreibe_Zahlen_Liste
```
Die Funktion Schreibe_Zahlen_Liste schreibt alle Elemente einer Zahlen Liste getrennt mit einem Komma in den Standart Output Stream.
```
* Parameter: `p1`
* Parameter Typ: `Zahlen Liste`
* Rückgabe Typ: `nichts`

### Aliase
1. `"Schreibe die Zahlen Liste <p1>"`
2. `"Schreibe <p1>"`

### Implementation
```ddp
	Schreibe p1 mit dem Seperator ", ".
```
## Schreibe_Text_Liste
```
Die Funktion Schreibe_Text_Liste schreibt alle Elemente einer Text Liste getrennt mit einem Komma in den Standart Output Stream.
```
* Parameter: `p1`
* Parameter Typ: `Text Liste`
* Rückgabe Typ: `nichts`

### Aliase
1. `"Schreibe die Text Liste <p1>"`
2. `"Schreibe <p1>"`

### Implementation
```ddp
	Schreibe p1 mit dem Seperator ", ".
```
## Schreibe_Zahl
```
Die Funktion Schreibe_Zahl schreibt eine gegebene Zahl (p1) in den Standart Output Stream.
```
* Parameter: `p1`
* Parameter Typ: `Zahl`
* Rückgabe Typ: `nichts`

### Aliase
1. `"Schreibe die Zahl <p1>"`
2. `"Schreibe <p1>"`

### Implementation
Implementiert in "libddpstdlib.a"
## Schreibe_Kommazahl
```
Die Funktion Schreibe_Kommazahl schreibt eine gegebene Kommazahl (p1) in den Standart Output Stream.
```
* Parameter: `p1`
* Parameter Typ: `Kommazahl`
* Rückgabe Typ: `nichts`

### Aliase
1. `"Schreibe die Kommazahl <p1>"`
2. `"Schreibe <p1>"`

### Implementation
Implementiert in "libddpstdlib.a"
## Schreibe_Text
```
Die Funktion Schreibe_Text schreibt einen gegebenen Text (p1) in den Standart Output Stream.
```
* Parameter: `p1`
* Parameter Typ: `Text`
* Rückgabe Typ: `nichts`

### Aliase
1. `"Schreibe den Text <p1>"`
2. `"Schreibe <p1>"`

### Implementation
Implementiert in "libddpstdlib.a"
## Schreibe_Fehler
```
Die Funktion Schreibe_Text schreibt einen gegebenen Text (fehler) in den Standart Error Stream.
```
* Parameter: `fehler`
* Parameter Typ: `Text`
* Rückgabe Typ: `nichts`

### Aliase
1. `"Schreibe den Fehler <fehler>"`

### Implementation
Implementiert in "libddpstdlib.a"
## SchreibeZeile_Kommazahl
```
Die Funktion SchreibeZeile_Kommazahl schreibt eine gegebene Kommazahl (p1) gefolgt von einer neuen Zeile in den Standart Output Stream.
```
* Parameter: `p1`
* Parameter Typ: `Kommazahl`
* Rückgabe Typ: `nichts`

### Aliase
1. `"Schreibe die Kommazahl <p1> auf eine Zeile"`
2. `"Schreibe <p1> auf eine Zeile"`

### Implementation
```ddp
	Schreibe p1.
	Schreibe '\n'.
```
## SchreibeZeile_Boolean
```
Die Funktion SchreibeZeile_Boolean schreibt einen gegebenen Boolean (p1) gefolgt von einer neuen Zeile in den Standart Output Stream.
```
* Parameter: `p1`
* Parameter Typ: `Boolean`
* Rückgabe Typ: `nichts`

### Aliase
1. `"Schreibe den Boolean <p1> auf eine Zeile"`
2. `"Schreibe <p1> auf eine Zeile"`

### Implementation
```ddp
	Schreibe p1.
	Schreibe '\n'.
```
## Schreibe_Boolean_Liste_Zeile
```
Die Funktion Schreibe_Boolean_Liste schreibt alle Elemente einer Boolean Liste getrennt mit einem Komma und gefolgt von einer neuen Zeile in den Standart Output Stream.
```
* Parameter: `p1`
* Parameter Typ: `Boolean Liste`
* Rückgabe Typ: `nichts`

### Aliase
1. `"Schreibe die Boolean Liste <p1> auf eine Zeile"`
2. `"Schreibe <p1> auf eine Zeile"`

### Implementation
```ddp
	Schreibe p1 mit dem Seperator ", ".
	Schreibe '\n'.
```
## Schreibe_Buchstaben_Liste_Zeile
```
Die Funktion Schreibe_Buchstaben_Liste schreibt alle Elemente einer Buchstaben Liste getrennt mit einem Komma und gefolgt von einer neuen Zeile in den Standart Output Stream.
```
* Parameter: `p1`
* Parameter Typ: `Buchstaben Liste`
* Rückgabe Typ: `nichts`

### Aliase
1. `"Schreibe die Buchstaben Liste <p1> auf eine Zeile"`
2. `"Schreibe <p1> auf eine Zeile"`

### Implementation
```ddp
	Schreibe p1 mit dem Seperator ", ".
	Schreibe '\n'.
```

