# Verzweigungen
Verzweigungen werden genutzt um Anweisungen basierend auf Bedingungen auszuführen.
Jede Verzweigung ist eine leicht variierte 'Wenn' Anweisung.

## Einseitige Verzweigung
Bei einer einseitigen Verzweigung wird ein Anweisungs-Block nur dann ausgeführt, wenn die angegebene Bedingung `wahr` ergibt, ansonsten wird nichts getan und die nachfolgenden Anweisungen ausgeführt.

### Aufbau:
```
Wenn <Bedingung> ist, dann:
	<Anweisung>.
```

### Bespiel:
```
Wenn 1 gleich 1 ist, dann:
	Schreibe den Text "Bedingung erfüllt!".
```
    
## Zweiseitige Verzweigung
Eine zweiseitige Verzweigung funktioniert wie die einseitige, mit dem Unterschied, dass ein `sonst` vorhanden ist.
Sollte die Bedingung `falsch` ergeben, wird der Code im `sonst`-Block ausgeführt.

### Aufbau:
```
Wenn <Bedingung>, dann:
	<Anweisung>.
Sonst:
	<Anweisung>.
```

### Bespiel:
```
Wenn 1 gleich 2 ist, dann:
Schreibe den Text "Bedingung erfüllt!" auf eine Zeile.
Sonst:
Schreibe den Text "Bedingung nicht erfüllt!" auf eine Zeile.
```

## Mehrseitige Verzweigung
Mehrseitige Verzweigungen erweitern die Zweiseitige Verzweigung um beliebig viele Verzweigungen.
Nach dem `Wenn`-Block können beliebig viele `Wenn aber`-Blöcke angefügt werden.
Sollte die erste Bedingung im `Wenn`-Block `falsch` ergeben wird die Bedingung im ersten `Wenn aber`-Block überprüft, sollte diese `falsch` sein wird die nächste überprüft und so weiter.
Wenn alle Bedingungen falsch sein sollten wird, falls vorhanden, der `sonst`-Block ausgeführt.
Sobald ein Block ausgeführt wurde werden alle folgenden Blöcke übersprungen.

### Aufbau:
```
Wenn <Bedingung>, dann:
	<Anweisung>
Wenn aber <2. Bedingung>, dann:
	<Anweisung>.
Sonst:
	<Anweisung>.
```

### Bespiel:
```
Wenn 1 gleich 2 ist, dann:
	Schreibe den Text "Bedingung erfüllt!" auf eine Zeile.
Wenn aber 2 gleich 2 ist, dann:
	Schreibe den Text "2. Bedingung erfüllt!" auf eine Zeile.
Sonst:
	Schreibe den Text "Bedingung nicht erfüllt!" auf eine Zeile.
```

# Schleifen
Schleifen werden genutzt um Code basierend auf Bedingungen mehrmals auszuführen.

## Kopfgesteuerte Schleife
Kopfgesteuerte Schleifen sind die einfachste Art von Schleifen.
Wenn die Bedingung 'wahr' ergibt wird der Code-Block ausgeführt.
Das wiederholt sich so lange, wie die Bedingung 'wahr' ergibt.
```
Solange i gleich 5 ist, mache:
	<Anweisung>.
```

## Fußgesteuerte Schleife
Fußgesteuerte Schleifen ähneln stark den Kopfgesteuerten, mit dem einzigen Unterschied, dass der Code-Block mindestens einmal ausgeführt wird, und dann erst das Prüfen der Bedingung beginnt.
```
Mache:
	<Anweisung>.
solange i gleich 5 ist.
```

## Zählende Schleife
Zählende Schleifen ermöglichen es ebenfalls Code mehrmals auszuführen, wobei gleichzeitig ein Zähler gegeben wird, der anderweitig genutzt werden kann.
Bei jeder Zählenden Schleife muss ein Zähler benannt werden (Eine Variable vom Typ 'Zahl') und ein Start- und Endwert.
Optional kann auch eine Schrittgröße angegeben werden, mit der gezählt wird.

### Hochzählende Schleife
Zuerst wird der Zähler benannt und mit dem Startwert initialisiert.
Dann wird (wie bei jeder Iteration) überprüft ob der Wert des Zählers kleiner oder gleich wie der des Endwerts ist.
Sollte diese Bedingung erfüllt sein, wird der Code-Block ausgeführt, und danach der Zähler um 1 erhöht.
Das wird solange wiederholt wie der Zähler den Endwert nicht überschreitet.
Im Code-Block kann der Zähler wie eine normale lokale Variable benutzt werden.

### Aufbau
```
Für jede Zahl <Zähler> von <Startwert> bis <Endwert>, mache:
	<Anweisung>.
```

### Beispiel
```
Für jede Zahl i von 1 bis 100, mache:
	Schreibe die Zahl i auf eine Zeile.
```
### Runterzählende Schleife
Eine Runterzählende Schleife funktioniert wie die Hochzählende, bloß mit dem Unterschied, dass eine Schrittgröße von -1 (oder irgendeinem negativen Wert) spezifiziert wird. Dadurch wird der Zähler am Ende nicht um 1 sondern um die angegebene Schrittgrößer erhöht (bzw. verringert).
Natürlich muss hierbei der Startwert größer als der Endwert sein, da es sonst eine Endlosschleife wird.

### Aufbau
```
Für jede Zahl <Zähler> von <Startwert> bis <Endwert> mit Schrittgröße -1, mache:
	<Anweisung>.
```

### Beispiel
```
Für jede Zahl i von 100 bis 1 mit Schrittgröße -1, mache:
	Schreibe die Zahl i auf eine Zeile.
```
### Eigene Schrittgröße
Wie oben erwähnt, kann auch eine Beliebige Schrittgröße n angegeben werden. Das sorgt dafür, das anstatt 1 der angegebene Wert zum Zähler addiert wird.

`mit Schrittgröße n` ist optional. Standartmäßig ist die Schrittgröße +1. 

### Aufbau
```
Für jede Zahl <Zähler> von <Startwert> bis <Endwert> mit Schrittgröße <n>, mache:
	<Anweisung>.
```

### Beispiel
```
Für jede Zahl i von 1 bis 100 mit Schrittgröße 5, mache:
	<Anweisung>.
```