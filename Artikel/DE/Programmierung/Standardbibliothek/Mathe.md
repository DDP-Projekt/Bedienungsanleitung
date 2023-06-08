# Duden/Mathe Funktionen
## tan
* Parameter: `v`
* Parameter Typ: `Kommazahl`
* Rückgabe Typ: `Kommazahl`

### Aliase
1. `"der Tangens von <v>"`
2. `"den Tangens von <v>"`

### Implementation
Implementiert in "libddpstdlib.a"
## asin
* Parameter: `v`
* Parameter Typ: `Kommazahl`
* Rückgabe Typ: `Kommazahl`

### Aliase
1. `"der Arkussinus von <v>"`
2. `"den Arkussinus von <v>"`

### Implementation
Implementiert in "libddpstdlib.a"
## atan
* Parameter: `v`
* Parameter Typ: `Kommazahl`
* Rückgabe Typ: `Kommazahl`

### Aliase
1. `"der Arkustangens von <v>"`
2. `"den Arkustangens von <v>"`

### Implementation
Implementiert in "libddpstdlib.a"
## sinh
* Parameter: `v`
* Parameter Typ: `Kommazahl`
* Rückgabe Typ: `Kommazahl`

### Aliase
1. `"der Hyperbelsinus von <v>"`
2. `"den Hyperbelsinus von <v>"`

### Implementation
Implementiert in "libddpstdlib.a"
## PHI

* Rückgabe Typ: `Kommazahl`

### Aliase
1. `"PHI"`

### Implementation
```ddp
	Gib 1,618033989 zurück.
```
## max
* Parameter: `a`, `b`
* Parameter Typen: `Zahl`, `Zahl`
* Rückgabe Typ: `Zahl`

### Aliase
1. `"die größere Zahl von <a> und <b>"`

### Implementation
```ddp
	Wenn a größer als, oder b ist, gib a zurück.
	Gib b zurück.
```
## min
* Parameter: `a`, `b`
* Parameter Typen: `Zahl`, `Zahl`
* Rückgabe Typ: `Zahl`

### Aliase
1. `"die kleinere Zahl von <a> und <b>"`

### Implementation
```ddp
	Wenn a kleiner als, oder b ist, gib a zurück.
	Gib b zurück.
```
## clamp
* Parameter: `wert`, `max`, `min`
* Parameter Typen: `Zahl`, `Zahl`, `Zahl`
* Rückgabe Typ: `Zahl`

### Aliase
1. `"<wert> zwischen <min> und <max>"`

### Implementation
```ddp
	Wenn wert größer als max ist, gib max zurück.
	Wenn wert kleiner als min ist, gib min zurück.
	Gib wert zurück.
```
## groesster_gemeinsamer_Teiler
```
Eine Funktion, die den größten gemeinsamen Teiler zweier Zahlen, <a> und <b>, als Zahl zurück gibt.
Zeitkomplexität: O(n)
```
* Parameter: `a`, `b`
* Parameter Typen: `Zahl`, `Zahl`
* Rückgabe Typ: `Zahl`

### Aliase
1. `"der größte gemeinsame Teiler von <a> und <b>"`

### Implementation
```ddp
	Die Zahl t ist 0.
	Solange b ungleich 0 ist, mache:
		Speichere b in t.
		Speichere (a modulo t) in b.
		Speichere t in a.
	Gib a zurück.
```
## kleinster_gemeinsamer_Teiler
```
Eine Funktion, die den kleinsten gemeinsamen Teiler zweier Zahlen, <a> und <b>, als Zahl zurück gibt.
Zeitkomplexität: O(n)
```
* Parameter: `a`, `b`
* Parameter Typen: `Zahl`, `Zahl`
* Rückgabe Typ: `Zahl`

### Aliase
1. `"das kleinste gemeinsame Vielfache von <a> und <b>"`

### Implementation
```ddp
	Gib (der Betrag von (a mal b) durch (der größte gemeinsame Teiler von a und b)) als Zahl zurück.
```
## cos
* Parameter: `v`
* Parameter Typ: `Kommazahl`
* Rückgabe Typ: `Kommazahl`

### Aliase
1. `"der Kosinus von <v>"`
2. `"den Kosinus von <v>"`

### Implementation
Implementiert in "libddpstdlib.a"
## cosh
* Parameter: `v`
* Parameter Typ: `Kommazahl`
* Rückgabe Typ: `Kommazahl`

### Aliase
1. `"der Hyperbelkosinus von <v>"`
2. `"den Hyperbelkosinus von <v>"`

### Implementation
Implementiert in "libddpstdlib.a"
## E

* Rückgabe Typ: `Kommazahl`

### Aliase
1. `"E"`

### Implementation
```ddp
	Gib 2,718281828 zurück.
```
## TAU

* Rückgabe Typ: `Kommazahl`

### Aliase
1. `"TAU"`

### Implementation
```ddp
	Gib 6,283185307 zurück.
```
## tanh
* Parameter: `v`
* Parameter Typ: `Kommazahl`
* Rückgabe Typ: `Kommazahl`

### Aliase
1. `"der Hyperbeltangens von <v>"`
2. `"den Hyperbeltangens von <v>"`

### Implementation
Implementiert in "libddpstdlib.a"
## Primfaktorzerlegung
```
Eine Funktion, die eine Zahlen Liste von allen Primfaktoren der Zahl <z> gibt.  
Zeitkomplexität: O(sqrt(n))
```
* Parameter: `z`
* Parameter Typ: `Zahl`
* Rückgabe Typ: `Zahlen Liste`

### Aliase
1. `"die Primfaktoren von <z>"`
2. `"alle Primfaktoren von <z>"`

### Implementation
```ddp
	Die Zahlen Liste faktoren ist eine leere Zahlen Liste.

	Solange z modulo 2 gleich 0 ist, mache:
		Speichere faktoren verkettet mit 2 in faktoren.
		Speichere (z durch 2) als Zahl in z.

	Die Zahl i ist 3.
	Solange i kleiner als, oder die 2. Wurzel von z als Zahl ist, mache:
		Solange z modulo i gleich 0 ist, mache:
			Speichere faktoren verkettet mit i in faktoren.
			Speichere (z durch i) als Zahl in z.
		Erhöhe i um 2.

	Wenn z größer als 2 ist, dann:
		Speichere faktoren verkettet mit z in faktoren.

	Gib faktoren zurück.
```
## trunc
* Parameter: `wert`
* Parameter Typ: `Kommazahl`
* Rückgabe Typ: `Kommazahl`

### Aliase
1. `"<wert> trunkiert"`

### Implementation
```ddp
	Gib (wert als Zahl) als Kommazahl zurück.
```
## sin
```
Trigonometrische Funktionen
```
* Parameter: `v`
* Parameter Typ: `Kommazahl`
* Rückgabe Typ: `Kommazahl`

### Aliase
1. `"der Sinus von <v>"`
2. `"den Sinus von <v>"`

### Implementation
Implementiert in "libddpstdlib.a"
## acos
* Parameter: `v`
* Parameter Typ: `Kommazahl`
* Rückgabe Typ: `Kommazahl`

### Aliase
1. `"der Arkuskosinus von <v>"`
2. `"den Arkuskosinus von <v>"`

### Implementation
Implementiert in "libddpstdlib.a"
## Teilerzerlegung
```
Gibt eine Zahlen Liste von alle Zahlen, die durch <z> geteilt werden können.
Zeitkomplexität: O(n)
```
* Parameter: `z`
* Parameter Typ: `Zahl`
* Rückgabe Typ: `Zahlen Liste`

### Aliase
1. `"alle Teiler von <z>"`

### Implementation
```ddp
	Die Zahlen Liste teiler ist eine leere Zahlen Liste.

	Für jede Zahl i von z bis 1 mit Schrittgröße -1, mache:
		Wenn z modulo i gleich 0 ist, speichere teiler verkettet mit i in teiler.	

	Gib teiler zurück.
```
## PI

* Rückgabe Typ: `Kommazahl`

### Aliase
1. `"PI"`

### Implementation
```ddp
	Gib 3,141592654 zurück.
```
## sign
* Parameter: `wert`
* Parameter Typ: `Zahl`
* Rückgabe Typ: `Zahl`

### Aliase
1. `"das Vorzeichen von <wert>"`

### Implementation
```ddp
	Wenn wert kleiner als 0 ist, gib -1 zurück.
	Wenn aber wert größer als 0 ist, gib 1 zurück.
	Gib 0 zurück.
```
## floor
* Parameter: `wert`
* Parameter Typ: `Kommazahl`
* Rückgabe Typ: `Kommazahl`

### Aliase
1. `"<wert> nach unten gerundet"`

### Implementation
```ddp
	Gib wert minus (wert minus wert als Zahl) zurück.
```
## ceil
* Parameter: `wert`
* Parameter Typ: `Kommazahl`
* Rückgabe Typ: `Kommazahl`

### Aliase
1. `"<wert> nach oben gerundet"`

### Implementation
```ddp
	Gib wert plus (1 minus (wert minus wert als Zahl)) zurück.
```

