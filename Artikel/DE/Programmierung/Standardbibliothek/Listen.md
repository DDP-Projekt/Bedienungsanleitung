# Duden/Listen Funktionen
## Füllen_Z
* Parameter: `liste`, `elm`
* Parameter Typen: `Zahlen Listen Referenz`, `Zahl`
* Rückgabe Typ: `nichts`

### Aliase
1. `"Fülle <liste> mit <elm>"`

### Implementation
```ddp
	Die Zahlen Liste neueListe ist die Länge von liste Mal elm.
	Speichere neueListe in liste.
```
## IndexVon_B
* Parameter: `liste`, `elm`
* Parameter Typen: `Boolean Liste`, `Boolean`
* Rückgabe Typ: `Zahl`

### Aliase
1. `"der Index von <elm> in <liste>"`

### Implementation
```ddp
	Für jede Zahl i von 1 bis (die Länge von liste), Wenn liste an der Stelle i gleich elm ist, Gib i zurück.
	Gib -1 zurück.
```
## Voranstellen_T
* Parameter: `liste`, `elm`
* Parameter Typen: `Text Listen Referenz`, `Text`
* Rückgabe Typ: `nichts`

### Aliase
1. `"Stelle <elm> vor <liste>"`

### Implementation
```ddp
	Speichere elm verkettet mit liste in liste.
```
## IstLeer_B
* Parameter: `liste`
* Parameter Typ: `Boolean Liste`
* Rückgabe Typ: `Boolean`

### Aliase
1. `"<liste> leer ist"`

### Implementation
```ddp
	Gib [wahr wenn] die Länge von liste gleich 0 ist zurück.
```
## EinfügenT
* Parameter: `liste`, `index`, `elm`
* Parameter Typen: `Text Listen Referenz`, `Zahl`, `Text`
* Rückgabe Typ: `nichts`

### Aliase
1. `"Setze <elm> an die Stelle <index> von <liste>"`

### Implementation
```ddp
	Speichere liste von 1 bis (index minus 1) verkettet mit elm verkettet mit liste von index  bis (die Länge von liste) in liste.
```
## Hinzufüge_K
* Parameter: `liste`, `elm`
* Parameter Typen: `Kommazahlen Listen Referenz`, `Kommazahl`
* Rückgabe Typ: `nichts`

### Aliase
1. `"Füge <elm> an <liste> an"`
2. `"füge <elm> an <liste> an"`

### Implementation
```ddp
	Speichere liste verkettet mit elm in liste.
```
## Voranstellen_K
* Parameter: `liste`, `elm`
* Parameter Typen: `Kommazahlen Listen Referenz`, `Kommazahl`
* Rückgabe Typ: `nichts`

### Aliase
1. `"Stelle <elm> vor <liste>"`

### Implementation
```ddp
	Speichere elm verkettet mit liste in liste.
```
## Lösche_B
* Parameter: `liste`, `index`
* Parameter Typen: `Boolean Listen Referenz`, `Zahl`
* Rückgabe Typ: `nichts`

### Aliase
1. `"Lösche das Element an der Stelle <index> aus <liste>"`

### Implementation
```ddp
	Speichere liste von 1 bis (index minus 1) verkettet mit liste von (index plus 1) bis (die Länge von liste) in liste.
```
## Voranstellen_C
* Parameter: `liste`, `elm`
* Parameter Typen: `Buchstaben Listen Referenz`, `Buchstabe`
* Rückgabe Typ: `nichts`

### Aliase
1. `"Stelle <elm> vor <liste>"`

### Implementation
```ddp
	Speichere elm verkettet mit liste in liste.
```
## IndexVon_C
* Parameter: `liste`, `elm`
* Parameter Typen: `Buchstaben Liste`, `Buchstabe`
* Rückgabe Typ: `Zahl`

### Aliase
1. `"der Index von <elm> in <liste>"`

### Implementation
```ddp
	Für jede Zahl i von 1 bis (die Länge von liste), Wenn liste an der Stelle i gleich elm ist, Gib i zurück.
	Gib -1 zurück.
```
## Aneinandergehängt_C
* Parameter: `liste`
* Parameter Typ: `Buchstaben Liste`
* Rückgabe Typ: `Text`

### Aliase
1. `"<liste> aneinandergehängt"`

### Implementation
```ddp
	Der Text t ist "".
	Für jeden Buchstabe b in liste, mache:
		Speichere t verkettet mit b in t.
	Gib t zurück.
```
## IndexVon_Z
* Parameter: `liste`, `elm`
* Parameter Typen: `Zahlen Liste`, `Zahl`
* Rückgabe Typ: `Zahl`

### Aliase
1. `"der Index von <elm> in <liste>"`

### Implementation
```ddp
	Für jede Zahl i von 1 bis (die Länge von liste), wenn liste an der Stelle i gleich elm ist, gib i zurück.
	Gib -1 zurück.
```
## Füllen_K
* Parameter: `liste`, `elm`
* Parameter Typen: `Kommazahlen Listen Referenz`, `Kommazahl`
* Rückgabe Typ: `nichts`

### Aliase
1. `"Fülle <liste> mit <elm>"`

### Implementation
```ddp
	Die Kommazahlen Liste neueListe ist die Länge von liste Mal elm.
	Speichere neueListe in liste.
```
## Leere_C
```
Buchstaben Listen Funktionen
```
* Parameter: `liste`
* Parameter Typ: `Buchstaben Listen Referenz`
* Rückgabe Typ: `nichts`

### Aliase
1. `"Leere <liste>"`

### Implementation
```ddp
	Speichere eine leere Buchstaben Liste in liste.
```
## Lösche_Bereich_C
* Parameter: `liste`, `start`, `end`
* Parameter Typen: `Buchstaben Listen Referenz`, `Zahl`, `Zahl`
* Rückgabe Typ: `nichts`

### Aliase
1. `"Lösche alle Elemente von <start> bis <end> aus <liste>"`

### Implementation
```ddp
	Speichere liste von 1 bis (start minus 1) verkettet mit liste von (end plus 1) bis (die Länge von liste) in liste.
```
## Leere_Z
```
Zahlen Listen Funktionen
```
* Parameter: `liste`
* Parameter Typ: `Zahlen Listen Referenz`
* Rückgabe Typ: `nichts`

### Aliase
1. `"Leere <liste>"`

### Implementation
```ddp
	Speichere eine leere Zahlen Liste in liste.
```
## Voranstellen_B
* Parameter: `liste`, `elm`
* Parameter Typen: `Boolean Listen Referenz`, `Boolean`
* Rückgabe Typ: `nichts`

### Aliase
1. `"Stelle <elm> vor <liste>"`

### Implementation
```ddp
	Speichere elm verkettet mit liste in liste.
```
## Einfügen_Bereich_B
* Parameter: `liste`, `index`, `range`
* Parameter Typen: `Boolean Listen Referenz`, `Zahl`, `Boolean Liste`
* Rückgabe Typ: `nichts`

### Aliase
1. `"Setze die Elemente in <range> an die Stelle <index> von <liste>"`

### Implementation
```ddp
	Speichere liste von 1 bis (index minus 1) verkettet mit range verkettet mit liste von index  bis (die Länge von liste) in liste.
```
## Leere_T
```
Text Listen Funktionen
```
* Parameter: `liste`
* Parameter Typ: `Text Listen Referenz`
* Rückgabe Typ: `nichts`

### Aliase
1. `"Leere <liste>"`

### Implementation
```ddp
	Speichere eine leere Text Liste in liste.
```
## Enthält_T
* Parameter: `liste`, `elm`
* Parameter Typen: `Text Liste`, `Text`
* Rückgabe Typ: `Boolean`

### Aliase
1. `"<liste> <elm> enthält"`

### Implementation
```ddp
	Für jeden Text t in liste, wenn t gleich elm ist, gib wahr zurück.
	Gib falsch zurück.
```
## Einfügen_Bereich_K
* Parameter: `liste`, `index`, `range`
* Parameter Typen: `Kommazahlen Listen Referenz`, `Zahl`, `Kommazahlen Liste`
* Rückgabe Typ: `nichts`

### Aliase
1. `"Setze die Elemente in <range> an die Stelle <index> von <liste>"`

### Implementation
```ddp
	Speichere liste von 1 bis (index minus 1) verkettet mit range verkettet mit liste von index  bis (die Länge von liste) in liste.
```
## Lösche_K
* Parameter: `liste`, `index`
* Parameter Typen: `Kommazahlen Listen Referenz`, `Zahl`
* Rückgabe Typ: `nichts`

### Aliase
1. `"Lösche das Element an der Stelle <index> aus <liste>"`

### Implementation
```ddp
	Speichere liste von 1 bis (index minus 1) verkettet mit liste von (index plus 1) bis (die Länge von liste) in liste.
```
## Lösche_Bereich_K
* Parameter: `liste`, `start`, `end`
* Parameter Typen: `Kommazahlen Listen Referenz`, `Zahl`, `Zahl`
* Rückgabe Typ: `nichts`

### Aliase
1. `"Lösche alle Elemente von <start> bis <end> aus <liste>"`

### Implementation
```ddp
	Speichere liste von 1 bis (start minus 1) verkettet mit liste von (end plus 1) bis (die Länge von liste) in liste.
```
## Hinzufüge_B
* Parameter: `liste`, `elm`
* Parameter Typen: `Boolean Listen Referenz`, `Boolean`
* Rückgabe Typ: `nichts`

### Aliase
1. `"Füge <elm> an <liste> an"`
2. `"füge <elm> an <liste> an"`

### Implementation
```ddp
	Speichere liste verkettet mit elm in liste.
```
## Lösche_Bereich_B
* Parameter: `liste`, `start`, `end`
* Parameter Typen: `Boolean Listen Referenz`, `Zahl`, `Zahl`
* Rückgabe Typ: `nichts`

### Aliase
1. `"Lösche alle Elemente von <start> bis <end> aus <liste>"`

### Implementation
```ddp
	Speichere liste von 1 bis (start minus 1) verkettet mit liste von (end plus 1) bis (die Länge von liste) in liste.
```
## Füllen_B
* Parameter: `liste`, `elm`
* Parameter Typen: `Boolean Listen Referenz`, `Boolean`
* Rückgabe Typ: `nichts`

### Aliase
1. `"Fülle <liste> mit <elm>"`

### Implementation
```ddp
	Die Boolean Liste neueListe ist die Länge von liste Mal elm.
	Speichere neueListe in liste.
```
## EinfügenZ
* Parameter: `liste`, `index`, `elm`
* Parameter Typen: `Zahlen Listen Referenz`, `Zahl`, `Zahl`
* Rückgabe Typ: `nichts`

### Aliase
1. `"Setze <elm> an die Stelle <index> von <liste>"`

### Implementation
```ddp
	Speichere liste von 1 bis (index minus 1) verkettet mit elm verkettet mit liste von index  bis (die Länge von liste) in liste.
```
## Voranstellen_Z
* Parameter: `liste`, `elm`
* Parameter Typen: `Zahlen Listen Referenz`, `Zahl`
* Rückgabe Typ: `nichts`

### Aliase
1. `"Stelle <elm> vor <liste>"`

### Implementation
```ddp
	Speichere elm verkettet mit liste in liste.
```
## EinfügenK
* Parameter: `liste`, `index`, `elm`
* Parameter Typen: `Kommazahlen Listen Referenz`, `Zahl`, `Kommazahl`
* Rückgabe Typ: `nichts`

### Aliase
1. `"Setze <elm> an die Stelle <index> von <liste>"`

### Implementation
```ddp
	Speichere liste von 1 bis (index minus 1) verkettet mit elm verkettet mit liste von index  bis (die Länge von liste) in liste.
```
## IndexVon_K
* Parameter: `liste`, `elm`
* Parameter Typen: `Kommazahlen Liste`, `Kommazahl`
* Rückgabe Typ: `Zahl`

### Aliase
1. `"der Index von <elm> in <liste>"`

### Implementation
```ddp
	Für jede Zahl i von 1 bis (die Länge von liste), Wenn liste an der Stelle i gleich elm ist, Gib i zurück.
	Gib -1 zurück.
```
## Füllen_C
* Parameter: `liste`, `elm`
* Parameter Typen: `Buchstaben Listen Referenz`, `Buchstabe`
* Rückgabe Typ: `nichts`

### Aliase
1. `"Fülle <liste> mit <elm>"`

### Implementation
```ddp
	Die Buchstaben Liste neueListe ist die Länge von liste Mal elm.
	Speichere neueListe in liste.
```
## Hinzufüge_T
* Parameter: `liste`, `elm`
* Parameter Typen: `Text Listen Referenz`, `Text`
* Rückgabe Typ: `nichts`

### Aliase
1. `"Füge <elm> an <liste> an"`
2. `"füge <elm> an <liste> an"`

### Implementation
```ddp
	Speichere liste verkettet mit elm in liste.
```
## Füllen_T
* Parameter: `liste`, `elm`
* Parameter Typen: `Text Listen Referenz`, `Text`
* Rückgabe Typ: `nichts`

### Aliase
1. `"Fülle <liste> mit <elm>"`

### Implementation
```ddp
	Die Text Liste neueListe ist die Länge von liste Mal elm.
	Speichere neueListe in liste.
```
## Lösche_Z
* Parameter: `liste`, `index`
* Parameter Typen: `Zahlen Listen Referenz`, `Zahl`
* Rückgabe Typ: `nichts`

### Aliase
1. `"Lösche das Element an der Stelle <index> aus <liste>"`

### Implementation
```ddp
	Speichere liste von 1 bis (index minus 1) verkettet mit liste von (index plus 1) bis (die Länge von liste) in liste.
```
## Enthält_K
* Parameter: `liste`, `elm`
* Parameter Typen: `Kommazahlen Liste`, `Kommazahl`
* Rückgabe Typ: `Boolean`

### Aliase
1. `"<liste> <elm> enthält"`

### Implementation
```ddp
	Für jede Kommazahl z in liste, wenn z gleich elm ist, gib wahr zurück.
	Gib falsch zurück.
```
## Enthält_B
* Parameter: `liste`, `elm`
* Parameter Typen: `Boolean Liste`, `Boolean`
* Rückgabe Typ: `Boolean`

### Aliase
1. `"<liste> <elm> enthält"`

### Implementation
```ddp
	Für jeden Boolean b in liste, wenn b gleich elm ist, gib wahr zurück.
	Gib falsch zurück.
```
## Einfügen_Bereich_C
* Parameter: `liste`, `index`, `range`
* Parameter Typen: `Buchstaben Listen Referenz`, `Zahl`, `Buchstaben Liste`
* Rückgabe Typ: `nichts`

### Aliase
1. `"Setze die Elemente in <range> an die Stelle <index> von <liste>"`

### Implementation
```ddp
	Speichere liste von 1 bis (index minus 1) verkettet mit range verkettet mit liste von index  bis (die Länge von liste) in liste.
```
## IstLeer_C
* Parameter: `liste`
* Parameter Typ: `Buchstaben Liste`
* Rückgabe Typ: `Boolean`

### Aliase
1. `"<liste> leer ist"`

### Implementation
```ddp
	Gib [wahr wenn] die Länge von liste gleich 0 ist zurück.
```
## Lösche_T
* Parameter: `liste`, `index`
* Parameter Typen: `Text Listen Referenz`, `Zahl`
* Rückgabe Typ: `nichts`

### Aliase
1. `"Lösche das Element an der Stelle <index> aus <liste>"`

### Implementation
```ddp
	Speichere liste von 1 bis (index minus 1) verkettet mit liste von (index plus 1) bis (die Länge von liste) in liste.
```
## Lösche_Bereich_Z
* Parameter: `liste`, `start`, `end`
* Parameter Typen: `Zahlen Listen Referenz`, `Zahl`, `Zahl`
* Rückgabe Typ: `nichts`

### Aliase
1. `"Lösche alle Elemente von <start> bis <end> aus <liste>"`

### Implementation
```ddp
	Speichere liste von 1 bis (start minus 1) verkettet mit liste von (end plus 1) bis (die Länge von liste) in liste.
```
## Enthält_Z
* Parameter: `liste`, `elm`
* Parameter Typen: `Zahlen Liste`, `Zahl`
* Rückgabe Typ: `Boolean`

### Aliase
1. `"<liste> <elm> enthält"`

### Implementation
```ddp
	Für jede Zahl z in liste, wenn z gleich elm ist, gib wahr zurück.
	Gib falsch zurück.
```
## Enthält_C
* Parameter: `liste`, `elm`
* Parameter Typen: `Buchstaben Liste`, `Buchstabe`
* Rückgabe Typ: `Boolean`

### Aliase
1. `"<liste> <elm> enthält"`

### Implementation
```ddp
	Für jeden Buchstaben b in liste, wenn b gleich elm ist, gib wahr zurück.
	Gib falsch zurück.
```
## EinfügenB
* Parameter: `liste`, `index`, `elm`
* Parameter Typen: `Boolean Listen Referenz`, `Zahl`, `Boolean`
* Rückgabe Typ: `nichts`

### Aliase
1. `"Setze <elm> an die Stelle <index> von <liste>"`

### Implementation
```ddp
	Speichere liste von 1 bis (index minus 1) verkettet mit elm verkettet mit liste von index  bis (die Länge von liste) in liste.
```
## Lösche_Bereich_T
* Parameter: `liste`, `start`, `end`
* Parameter Typen: `Text Listen Referenz`, `Zahl`, `Zahl`
* Rückgabe Typ: `nichts`

### Aliase
1. `"Lösche alle Elemente von <start> bis <end> aus <liste>"`

### Implementation
```ddp
	Speichere liste von 1 bis (start minus 1) verkettet mit liste von (end plus 1) bis (die Länge von liste) in liste.
```
## Einfügen_Bereich_T
* Parameter: `liste`, `index`, `range`
* Parameter Typen: `Text Listen Referenz`, `Zahl`, `Text Liste`
* Rückgabe Typ: `nichts`

### Aliase
1. `"Setze die Elemente in <range> an die Stelle <index> von <liste>"`

### Implementation
```ddp
	Speichere liste von 1 bis (index minus 1) verkettet mit range verkettet mit liste von index  bis (die Länge von liste) in liste.
```
## IndexVon_T
* Parameter: `liste`, `elm`
* Parameter Typen: `Text Liste`, `Text`
* Rückgabe Typ: `Zahl`

### Aliase
1. `"der Index von <elm> in <liste>"`

### Implementation
```ddp
	Für jede Zahl i von 1 bis (die Länge von liste), Wenn liste an der Stelle i gleich elm ist, Gib i zurück.
	Gib -1 zurück.
```
## Einfügen_Bereich_Z
* Parameter: `liste`, `index`, `range`
* Parameter Typen: `Zahlen Listen Referenz`, `Zahl`, `Zahlen Liste`
* Rückgabe Typ: `nichts`

### Aliase
1. `"Setze die Elemente in <range> an die Stelle <index> von <liste>"`

### Implementation
```ddp
	Speichere liste von 1 bis (index minus 1) verkettet mit range verkettet mit liste von index  bis (die Länge von liste) in liste.
```
## IstLeer_Z
* Parameter: `liste`
* Parameter Typ: `Zahlen Liste`
* Rückgabe Typ: `Boolean`

### Aliase
1. `"<liste> leer ist"`

### Implementation
```ddp
	Gib [wahr wenn] die Länge von liste gleich 0 ist zurück.
```
## Leere_K
```
Kommazahlen Listen Funktionen
```
* Parameter: `liste`
* Parameter Typ: `Kommazahlen Listen Referenz`
* Rückgabe Typ: `nichts`

### Aliase
1. `"Leere <liste>"`

### Implementation
```ddp
	Speichere eine leere Kommazahlen Liste in liste.
```
## IstLeer_K
* Parameter: `liste`
* Parameter Typ: `Kommazahlen Liste`
* Rückgabe Typ: `Boolean`

### Aliase
1. `"<liste> leer ist"`

### Implementation
```ddp
	Gib [wahr wenn] die Länge von liste gleich 0 ist zurück.
```
## Leere_B
```
Boolean Listen Funktionen
```
* Parameter: `liste`
* Parameter Typ: `Boolean Listen Referenz`
* Rückgabe Typ: `nichts`

### Aliase
1. `"Leere <liste>"`

### Implementation
```ddp
	Speichere eine leere Boolean Liste in liste.
```
## Hinzufüge_C
* Parameter: `liste`, `elm`
* Parameter Typen: `Buchstaben Listen Referenz`, `Buchstabe`
* Rückgabe Typ: `nichts`

### Aliase
1. `"Füge <elm> an <liste> an"`
2. `"füge <elm> an <liste> an"`

### Implementation
```ddp
	Speichere liste verkettet mit elm in liste.
```
## IstLeer_T
* Parameter: `liste`
* Parameter Typ: `Text Liste`
* Rückgabe Typ: `Boolean`

### Aliase
1. `"<liste> leer ist"`

### Implementation
```ddp
	Gib [wahr wenn] die Länge von liste gleich 0 ist zurück.
```
## Hinzufüge_Z
* Parameter: `liste`, `elm`
* Parameter Typen: `Zahlen Listen Referenz`, `Zahl`
* Rückgabe Typ: `nichts`

### Aliase
1. `"Füge <elm> an <liste> an"`
2. `"füge <elm> an <liste> an"`

### Implementation
```ddp
	Speichere liste verkettet mit elm in liste.
```
## EinfügenC
* Parameter: `liste`, `index`, `elm`
* Parameter Typen: `Buchstaben Listen Referenz`, `Zahl`, `Buchstabe`
* Rückgabe Typ: `nichts`

### Aliase
1. `"Setze <elm> an die Stelle <index> von <liste>"`

### Implementation
```ddp
	Speichere liste von 1 bis (index minus 1) verkettet mit elm verkettet mit liste von index  bis (die Länge von liste) in liste.
```
## Lösche_C
* Parameter: `liste`, `index`
* Parameter Typen: `Buchstaben Listen Referenz`, `Zahl`
* Rückgabe Typ: `nichts`

### Aliase
1. `"Lösche das Element an der Stelle <index> aus <liste>"`

### Implementation
```ddp
	Speichere liste von 1 bis (index minus 1) verkettet mit liste von (index plus 1) bis (die Länge von liste) in liste.
```

