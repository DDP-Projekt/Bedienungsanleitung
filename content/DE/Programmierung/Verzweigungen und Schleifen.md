+++
title = "Verzweigungen"
weight = 5
+++

# Verzweigungen
Verzweigungen werden genutzt um Anweisungen basierend auf Bedingungen auszuführen.

## Einseitige Verzweigung
Bei einer einseitigen Verzweigung wird ein Anweisungs-Block nur dann ausgeführt, wenn die angegebene Bedingung `wahr` ergibt, ansonsten wird nichts getan und die nachfolgenden Anweisungen ausgeführt.

### Aufbau:
```ddp
Wenn <Bedingung>, dann:
	<Anweisungen>.
```

### Beispiel:
```ddp
Wenn 1 gleich 1 ist, dann:
	Schreibe den Text "Bedingung erfüllt!".
```
    
## Zweiseitige Verzweigung
Eine zweiseitige Verzweigung funktioniert wie die einseitige, mit dem Unterschied, dass ein `sonst` vorhanden ist.
Sollte die Bedingung `falsch` ergeben, wird der Code im `sonst`-Block ausgeführt.

### Aufbau:
```ddp
Wenn <Bedingung>, dann:
	<Anweisungen>.
Sonst:
	<Anweisungen>.
```

### Beispiel:
```ddp
Wenn 1 gleich 2 ist, dann:
	Schreibe den Text "Bedingung erfüllt!".
Sonst:
	Schreibe den Text "Bedingung nicht erfüllt!".
```

## Mehrseitige Verzweigung
Mehrseitige Verzweigungen erweitern die zweiseitige Verzweigung um beliebig viele Verzweigungen.
Nach dem `Wenn`-Block können beliebig viele `Wenn aber`-Blöcke angefügt werden.
Sollte die erste Bedingung im `Wenn`-Block `falsch` ergeben wird die Bedingung im ersten `Wenn aber`-Block überprüft, sollte diese `falsch` sein wird die nächste überprüft und so weiter.
Wenn alle Bedingungen falsch sein sollten wird, falls vorhanden, der `sonst`-Block ausgeführt.
Sobald ein Block ausgeführt wurde, werden alle folgenden Blöcke übersprungen.

### Aufbau:
```ddp
Wenn <Bedingung>, dann:
	<Anweisungen>
Wenn aber <2. Bedingung>, dann:
	<Anweisungen>.
Sonst:
	<Anweisungen>.
```

### Beispiel:
```ddp
Wenn 1 gleich 2 ist, dann:
	Schreibe den Text "Bedingung erfüllt!".
Wenn aber 2 gleich 2 ist, dann:
	Schreibe den Text "2. Bedingung erfüllt!".
Sonst:
	Schreibe den Text "Bedingung nicht erfüllt!".
```

# Schleifen
Schleifen werden genutzt um Code basierend auf Bedingungen mehrmals auszuführen.

## Kopfgesteuerte Schleife
Kopfgesteuerte Schleifen sind die einfachste Art von Schleifen.
Wenn die Bedingung 'wahr' ergibt, wird der Code-Block ausgeführt.
Das wiederholt sich so lange, wie die Bedingung 'wahr' ergibt.
```ddp
Solange <Bedingung>, mache:
	<Anweisungen>.
```

## Fußgesteuerte Schleife
Fußgesteuerte Schleifen ähneln stark den Kopfgesteuerten, mit dem einzigen Unterschied, dass der Code-Block mindestens einmal ausgeführt wird, und dann erst das Prüfen der Bedingung beginnt.
```ddp
Mache:
	<Anweisungen>.
Solange <Bedingung>.
```

## Wiederholung
Wiederholungen werden genutzt um einen Code-Block mehrere Male auszuführen.
Sie sind eine gekürzte Version von zählenden Schleifen, sparen Text und erhöhen die Lesbarkeit des Codes
wenn man die Zähler-Variable nicht benötigt.
```ddp
Wiederhole:
	<Anweisungen>.
<Anzahl> Mal.
```

## Zählende Schleife
Zählende Schleifen ermöglichen es ebenfalls Code mehrmals auszuführen, wobei gleichzeitig ein Zähler gegeben wird, der anderweitig genutzt werden kann.
Bei jeder Zählenden Schleife muss ein Zähler benannt werden (eine Variable vom Typ `Zahl`, `Byte` oder `Kommazahl`) und ein Start- und Endwert.
Optional kann auch eine Schrittgröße angegeben werden, mit der gezählt wird.

### Hochzählende Schleife
Zuerst wird der Zähler benannt und mit dem Startwert initialisiert.
Dann wird (wie bei jeder Iteration) überprüft, ob der Wert des Zählers kleiner oder gleich wie der des Endwerts ist.
Sollte diese Bedingung erfüllt sein, wird der Code-Block ausgeführt, und danach der Zähler um 1 erhöht.
Das wird solange wiederholt wie der Zähler den Endwert nicht überschreitet.
Im Code-Block kann der Zähler wie eine normale lokale Variable benutzt werden.

### Aufbau
```ddp
Für jede Zahl <Zähler> von <Startwert> bis <Endwert>, mache:
	<Anweisungen>.
```

### Beispiel
```ddp
Für jede Zahl i von 1 bis 100, mache:
	Schreibe die Zahl i.
```
### Runterzählende Schleife
Eine Runterzählende Schleife funktioniert wie die Hochzählende, bloß mit dem Unterschied, dass eine Schrittgröße von -1 (oder irgendeinem negativen Wert) spezifiziert wird. Dadurch wird der Zähler am Ende nicht um 1 sondern um die angegebene Schrittgröße erhöht (bzw. verringert).
Natürlich muss hierbei der Startwert größer als der Endwert sein, da es sonst eine Endlosschleife wird.

### Aufbau
```ddp
Für jede Zahl <Zähler> von <Startwert> bis <Endwert> mit Schrittgröße -1, mache:
	<Anweisungen>.
```

### Beispiel
```ddp
Für jede Zahl i von 100 bis 1 mit Schrittgröße -1, mache:
	Schreibe die Zahl i.
```
### Eigene Schrittgröße
Wie oben erwähnt, kann auch eine beliebige Schrittgröße n angegeben werden. Das sorgt dafür, das anstatt 1 der angegebene Wert zum Zähler addiert wird.

`mit Schrittgröße n` ist optional. Standardmäßig ist die Schrittgröße +1. 

### Aufbau
```ddp
Für jede Zahl <Zähler> von <Startwert> bis <Endwert> mit Schrittgröße <n>, mache:
	<Anweisungen>.
```

### Beispiel
```ddp
Für jede Zahl i von 1 bis 100 mit Schrittgröße 5, mache:
	<Anweisungen>.
```

## Iterierende Schleifen
Es ist auch möglich mit Schleifen durch jedes Element einer Liste durchzulaufen. 
```ddp
Die Zahlen Liste liste ist eine leere Zahlen Liste.

Für jede Zahl element in liste, mache:
	<Anweisungen>.
```

Bei solchen Schleifen kann man auch einen Index mit angeben.
```ddp
Die Text Liste liste ist eine Liste, die aus "hi", "Hallo", "Tschüss" besteht.

[Ausgabe:  hi 1 Hallo 2 Tschüss 3]
Für jeden Text element mit Index i in liste, mache:
	Schreibe (' ' verkettet mit element verkettet mit ' ').
	Schreibe index.
```

## Abbrechen/Fortführen von Schleifen

Schleifen können auch abgebrochen bzw. zur nächsten Iteration geführt werden (in anderen Sprachen entspricht das den Schlüsselwörtern `break` und `continue`).

### Abbrechen

```ddp
[Schreibt "1 2" in die Konsole]
Für jede Zahl i von 1 bis 5, mache:
	Wenn i gleich 3 ist, dann:
    	Verlasse die Schleife. [break]
    Schreibe i.
    Schreibe ' '.
```

### Fortführen

```ddp
[Schreibt "1 2 4 5" in die Konsole]
Für jede Zahl i von 1 bis 5, mache:
	Wenn i gleich 3 ist, dann:
    	Fahre mit der Schleife fort. [continue]
    Schreibe i.
    Schreibe ' '.
```

# Tipp
Fast jede der hier aufgelisteten Verzweigungen und Schleifen kann auch auf einer einzigen Zeile geschrieben werden,
falls nur eine Anweisung ausgeführt werden muss.

## Beispiele
```ddp
Wenn 1 gleich 1 ist, Schreibe den Text "Bedingung erfüllt!".

Wenn 1 gleich 2 ist, Schreibe den Text "Bedingung erfüllt!".
Sonst Schreibe den Text "Bedingung nicht erfüllt".

Wenn 1 gleich 2 ist, Schreibe den Text "Bedingung erfüllt!".
Wenn aber 1 kleiner als 2 ist, Schreibe den Text "Zweite Bedingung erfüllt!".
Sonst Schreibe den Text "Bedingung nicht erfüllt".


Solange i gleich 5 ist, Rufe eine Funktion auf, die i erhöht.

Schreibe den Text "Hi!" 5 Mal.


Für jede Zahl i von 1 bis 100, Schreibe die Zahl i.

Für jede Zahl i von 100 bis 1 mit Schrittgröße -1, Schreibe die Zahl i.
```