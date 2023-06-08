# Funktionen
## Quicksort
```
Eine Funktion, die die Zahlen Liste <liste> mit dem Quick-Sort Algorithmus sortiert und die sortierte Liste zurück gibt.
Zeitkomplexität: O(n * log(n))
```
* Parameter: `liste`
* Parameter Typ: `Zahlen Liste`
* Rückgabe Typ: `Zahlen Liste`

### Aliase
1. `"<liste> sortiert"`
2. `"<liste> mit quick-sort sortiert"`

### Implementation
```ddp
	Wenn die Länge von liste kleiner als, oder 1 ist, gib liste zurück. [End-Bedingung der Rekursion]
	
	Die Zahl Teiler ist 0. [für später]
	:
		Die Zahl pivot ist liste an der Stelle (die Länge von liste). [das pivot Element ist das letzte Element der Liste]
		[alle Elemente die größer als pivot sind kommen nach rechts, der Rest nach links, die Mitte wird in Teiler gespeichert]
		Die Zahl _links ist 1.
		Die Zahl _rechts ist die Länge von liste.
		Die Zahl i ist _links.
		Die Zahl j ist _rechts.

		Solange i kleiner als j ist, mache:
			Solange i kleiner als j ist und liste an der Stelle i kleiner als, oder pivot ist, erhöhe i um 1.
			Solange j größer als i ist und liste an der Stelle j größer als pivot ist, verringere j um 1.
			
			Wenn liste an der Stelle i größer als liste an der Stelle j ist, tausche (liste an der Stelle j) und (liste an der Stelle i).

		Wenn liste an der Stelle i größer als pivot ist, tausche (liste an der Stelle i) und (liste an der Stelle _rechts).
		Sonst speichere _rechts in i.
		Speichere i in Teiler.
	
	[sortiere die linke und rechte Hälfte rekursiv]
	Die Zahlen Liste linkeHälfte ist (liste von 1 bis (Teiler minus 1)) mit quick-sort sortiert.
	Die Zahlen Liste rechteHälfte ist (liste von Teiler bis (die Länge von liste)) mit quick-sort sortiert.

	Gib linkeHälfte verkettet mit rechteHälfte zurück.
```
