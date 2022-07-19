# Beispielprogramme
## Fizzbuzz
```ddp
Binde "Duden/Ausgabe" ein.

Für jede Zahl i von 1 bis 100, mache:
	Wenn i modulo 3 gleich 0 und i modulo 5 gleich 0 ist, dann:
		Schreibe den Text "FizzBuzz" auf eine Zeile.
	Sonst:
		Wenn i modulo 3 gleich 0 ist, dann:
			Schreibe den Text "Fizz" auf eine Zeile.
		Wenn aber i modulo 5 gleich 0 ist, dann:
			Schreibe den Text "Buzz" auf eine Zeile.
		Sonst:
			Schreibe die Zahl i auf eine Zeile.
```

## Fibonacci
```ddp
Binde "Duden/Ausgabe" ein.

Die Funktion fib mit den Parametern n vom Typ Zahl, gibt eine Zahl zurück, macht:
	Wenn n kleiner als, oder 1 ist, dann:
        Gib n zurück.

    Gib die (n minus 2). Fibonacci nummer plus die (n minus 1). Fibonacci nummer zurück.
und kann so benutzt werden:
	"die *1. Fibonacci nummer"

Die Zahl num ist die 5. Fibonacci nummer.
Schreibe die Zahl num.
```

## Tic Tac Toe
```ddp
Binde "Duden/Ausgabe" ein.

Die Buchstaben brett sind 9 mal ' '.
Der Boolean Spieler1 ist wahr.

Die Funktion zeichneBrett gibt nichts zurück, macht:
	Für jede Zahl i von 1 bis 9, mache:
		Schreibe den Buchstaben brett an der Stelle i.
		Schreibe den Buchstaben ' '.
		Wenn i modulo 3 gleich 0 ist, dann überspringe eine Zeile.
und kann so benutzt werden:
	"Zeichne das Brett"

Die Funktion überprüfe mit den Parametern spieler vom Typ Buchstabe, gibt einen Boolean zurück, macht:
	Wenn brett an der Stelle 1 gleich spieler und brett an der Stelle 2 gleich spieler und brett an der Stelle 3 gleich spieler ist, dann gib wahr zurück.
	Wenn brett an der Stelle 4 gleich spieler und brett an der Stelle 5 gleich spieler und brett an der Stelle 6 gleich spieler ist, dann gib wahr zurück.
	Wenn brett an der Stelle 7 gleich spieler und brett an der Stelle 8 gleich spieler und brett an der Stelle 9 gleich spieler ist, dann gib wahr zurück.

	Wenn brett an der Stelle 1 gleich spieler und brett an der Stelle 4 gleich spieler und brett an der Stelle 7 gleich spieler ist, dann gib wahr zurück.
	Wenn brett an der Stelle 2 gleich spieler und brett an der Stelle 5 gleich spieler und brett an der Stelle 8 gleich spieler ist, dann gib wahr zurück.
	Wenn brett an der Stelle 3 gleich spieler und brett an der Stelle 6 gleich spieler und brett an der Stelle 9 gleich spieler ist, dann gib wahr zurück.

	Wenn brett an der Stelle 1 gleich spieler und brett an der Stelle 5 gleich spieler und brett an der Stelle 9 gleich spieler ist, dann gib wahr zurück.
	Wenn brett an der Stelle 3 gleich spieler und brett an der Stelle 5 gleich spieler und brett an der Stelle 7 gleich spieler ist, dann gib wahr zurück.

	Gib falsch zurück.
und kann so benutzt werden:
	"Spieler *1 gewonnen hat"


Der Buchstabe eingabe ist die nächste Taste.
solange eingabe ungleich 'v' ist, mache:
	Die Zahl feld ist eingabe als Zahl.
	Wenn brett an der Stelle feld gleich ' ' ist, dann:
		Wenn Spieler1 gleich wahr ist, dann:
			brett an der Stelle feld ist 'X'.
		Sonst:
			brett an der Stelle feld ist 'O'.
	Sonst:
		Solange brett an der Stelle feld ungleich ' ' ist, mache:
			Schreibe den Text "Dieses Feld ist bereits belegt, gib ein anderes ein: ".
			eingabe ist die nächste Benutzereingabe.
			feld ist eingabe als Zahl.
			Überspringe eine Zeile.

		Wenn Spieler1 gleich wahr ist, dann:
			brett an der Stelle feld ist 'X'.
		Sonst:
			brett an der Stelle feld ist 'O'.
	
	Wenn Spieler 'X' gewonnen hat, dann:
		Schreibe den Text "Spieler 1 hat gewonnen".
		Warte auf die nächste Benutzereingabe.
		Beende das Programm.
	Wenn aber Spieler 'O' gewonnen hat, dann:
		Schreibe den Text "Spieler 2 hat gewonnen".
		Warte auf die nächste Benutzereingabe.
		Beende das Programm.
	Wenn aber nicht brett ' ' enthält, dann:
		Schreibe den Text "Unentschieden".
		Warte auf die nächste Benutzereingabe.
		Beende das Programm.

	Zeichne das Brett.
	Negiere Spieler1.
	eingabe ist die nächste Taste.
```