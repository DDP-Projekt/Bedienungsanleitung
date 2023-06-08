# Duden/Zeichen Funktionen
## IstDeutscherBuchstabe
```
Gibt wahr zurück wenn der Buchstabe b ein deutscher Buchstabe (a-Z, äöü, ÄÖÜ und ß) ist.
```
* Parameter: `b`
* Parameter Typ: `Buchstabe`
* Rückgabe Typ: `Boolean`

### Aliase
1. `"<b> ein deutscher Buchstabe ist"`

### Implementation
```ddp
	Die Zahl z ist b als Zahl.
	Gib 
		(z größer als, oder 65 ist und z kleiner als, oder 90 ist) [a-z] oder
		(z größer als, oder 97 ist und z kleiner als, oder 122 ist) [A-Z] oder
		(z gleich 196 ist oder z gleich 228 ist) [Ää] oder 
		(z gleich 214 ist oder z gleich 246 ist) [Öö] oder
		(z gleich 220 ist oder z gleich 252 ist) [Üü] oder
		(z gleich 223 ist) [ß]
		zurück.
```
## IstDeutscherBuchstabeOderZahl
```
Gibt wahr zurück wenn der Buchstabe b ein deutscher Buchstabe oder eine Zahl ist.
```
* Parameter: `b`
* Parameter Typ: `Buchstabe`
* Rückgabe Typ: `Boolean`

### Aliase
1. `"<b> ein deutscher Buchstabe oder eine Zahl ist"`

### Implementation
```ddp
	Gib b ein deutscher Buchstabe ist oder b eine Zahl ist zurück.
```
## Großgeschrieben
```
Gibt wahr zurück wenn der Buchstabe b ein großer deutscher Buchstabe ist.
```
* Parameter: `b`
* Parameter Typ: `Buchstabe`
* Rückgabe Typ: `Buchstabe`

### Aliase
1. `"<b> als großer Buchstabe"`

### Implementation
```ddp
	Wenn nicht b ein deutscher Buchstabe ist, dann:
		Gib b zurück.

	Gib (b als Zahl logisch und 223) als Buchstabe zurück.
```
## Kleingeschrieben
```
Gibt wahr zurück wenn der Buchstabe b ein kleiner deutscher Buchstabe ist.
```
* Parameter: `b`
* Parameter Typ: `Buchstabe`
* Rückgabe Typ: `Buchstabe`

### Aliase
1. `"<b> als kleiner Buchstabe"`

### Implementation
```ddp
	Wenn nicht b ein deutscher Buchstabe ist oder b ein kleiner Buchstabe ist, dann:
		Gib b zurück.

	Gib (b als Zahl plus 32) als Buchstabe zurück.
```
## IstLeer
```
Gibt wahr zurück wenn der Buchstabe b ein Leerzeichen (' '), eine neue Zeile ('\n'), ein Tabulator ('\t') oder ein Wagenrücklauf ('\r') ist.
```
* Parameter: `b`
* Parameter Typ: `Buchstabe`
* Rückgabe Typ: `Boolean`

### Aliase
1. `"<b> ein leeres Zeichen ist"`

### Implementation
```ddp
	Gib b gleich ' ' ist oder b gleich '\n' ist oder b gleich '\t' ist oder b gleich '\r' ist zurück.
```
## IstGroß
* Parameter: `b`
* Parameter Typ: `Buchstabe`
* Rückgabe Typ: `Boolean`

### Aliase
1. `"<b> ein großer Buchstabe ist"`

### Implementation
```ddp
	Die Zahl z ist b als Zahl.
	Der Boolean erg ist 
		(z größer als, oder 65 ist und z kleiner als, oder 90 ist) oder
		(z größer als, oder 192 ist und z kleiner als, oder 214 ist) oder
		(z größer als, oder 216 ist und z kleiner als, oder 222 ist).
		[... es gibt noch viel mehr ...]
	Gib erg zurück.
```
## IstKlein
* Parameter: `b`
* Parameter Typ: `Buchstabe`
* Rückgabe Typ: `Boolean`

### Aliase
1. `"<b> ein kleiner Buchstabe ist"`

### Implementation
```ddp
	Die Zahl z ist b als Zahl.
	Der Boolean erg ist 
		(z größer als, oder 96 ist und z kleiner als, oder 122 ist) oder
		(z größer als, oder 223 ist und z kleiner als, oder 246 ist) oder
		(z größer als, oder 248 ist und z kleiner als, oder 255 ist).
		[... es gibt noch viel mehr ...]
	Gib erg zurück.
```
## IstLeerzeichen
```
Gibt wahr zurück wenn der Buchstabe b ein Leerzeichen (' ') ist.
```
* Parameter: `b`
* Parameter Typ: `Buchstabe`
* Rückgabe Typ: `Boolean`

### Aliase
1. `"<b> ein Leerzeichen ist"`

### Implementation
```ddp
	Gib b gleich ' ' ist zurück.
```
## IstZahl
```
Gibt wahr zurück wenn der Buchstabe b eine Zahl (Code 48-57) ist.
```
* Parameter: `b`
* Parameter Typ: `Buchstabe`
* Rückgabe Typ: `Boolean`

### Aliase
1. `"<b> eine Zahl ist"`

### Implementation
```ddp
	Gib b als Zahl größer als, oder 48 ist und b als Zahl kleiner als, oder 57 ist zurück.
```
## IstKontroll
```
Gibt wahr zurück wenn der Buchstabe b ein Kontrollzeichen (Code 0-31) ist.
```
* Parameter: `b`
* Parameter Typ: `Buchstabe`
* Rückgabe Typ: `Boolean`

### Aliase
1. `"<b> ein Kontrollzeichen ist"`

### Implementation
```ddp
	Gib b als Zahl größer als, oder 0 ist und b als Zahl kleiner als, oder 31 ist zurück.
```

