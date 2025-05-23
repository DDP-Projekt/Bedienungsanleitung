+++
title = "Beispielprogramme"
weight = 4
+++

# Beispielprogramme
## Fizzbuzz
```ddp
Binde "Duden/Ausgabe" ein.

Für jede Zahl i von 1 bis 100, mache:
	Wenn i modulo 3 gleich 0 ist und i modulo 5 gleich 0 ist, Schreibe den Text "FizzBuzz" auf eine Zeile.
	Sonst:
		Wenn i modulo 3 gleich 0 ist, Schreibe den Text "Fizz" auf eine Zeile.
		Wenn aber i modulo 5 gleich 0 ist, Schreibe den Text "Buzz" auf eine Zeile.
		Sonst Schreibe die Zahl i auf eine Zeile.
```

## Fibonacci
```ddp
Binde "Duden/Ausgabe" ein.

Die Funktion fib mit dem Parameter n vom Typ Zahl, gibt eine Zahl zurück, macht:
	Wenn n kleiner als, oder 1 ist, gib n zurück.

    Gib die (n minus 2). Fibonacci Nummer plus die (n minus 1). Fibonacci Nummer zurück.
Und kann so benutzt werden:
	"die <n>. Fibonacci Nummer"

Schreibe (die 6. Fibonacci Nummer).
```

## Tic Tac Toe
```ddp
Binde "Duden/Eingabe" ein.
Binde "Duden/Ausgabe" ein.
Binde "Duden/Laufzeit" ein.

Der Wahrheitswert ist_Spieler_1_dran ist wahr.
Die Zahlen Liste Spielfeld ist 9 Mal 0.

Die Funktion spieler1dran gibt einen Wahrheitswert zurück, macht:
    Gib ist_Spieler_1_dran zurück.
Und kann so benutzt werden:
    "Spieler 1 dran ist"

Die Funktion wechsel_spieler gibt nichts zurück, macht:
    Negiere ist_Spieler_1_dran.
Und kann so benutzt werden:
    "Wechsel den Spieler"

Die Funktion momentane_Spieler gibt eine Zahl zurück, macht:
    Wenn Spieler 1 dran ist, gib 1 zurück.
    [Ansonsten] Gib 2 zurück.
Und kann so benutzt werden:
    "die Zahl für den aktuellen Spieler"

Die Funktion spiel_nicht_gewonnen gibt einen Wahrheitswert zurück, macht:
    Die Zahl Spieler_zahl ist 0.
    Wenn Spieler 1 dran ist, speichere 2 in Spieler_zahl.
    Sonst speichere 1 in Spieler_zahl.

    Der Wahrheitswert Gewonnen ist wahr, wenn
        Spielfeld an der Stelle 1 gleich Spielfeld an der Stelle 2 ist und Spielfeld an der Stelle 1 gleich Spielfeld an der Stelle 3 ist und Spielfeld an der Stelle 1 gleich Spieler_zahl ist oder
        Spielfeld an der Stelle 4 gleich Spielfeld an der Stelle 5 ist und Spielfeld an der Stelle 4 gleich Spielfeld an der Stelle 6 ist und Spielfeld an der Stelle 4 gleich Spieler_zahl ist oder
        Spielfeld an der Stelle 7 gleich Spielfeld an der Stelle 8 ist und Spielfeld an der Stelle 7 gleich Spielfeld an der Stelle 9 ist und Spielfeld an der Stelle 7 gleich Spieler_zahl ist oder
        Spielfeld an der Stelle 1 gleich Spielfeld an der Stelle 4 ist und Spielfeld an der Stelle 1 gleich Spielfeld an der Stelle 7 ist und Spielfeld an der Stelle 1 gleich Spieler_zahl ist oder
        Spielfeld an der Stelle 2 gleich Spielfeld an der Stelle 5 ist und Spielfeld an der Stelle 2 gleich Spielfeld an der Stelle 8 ist und Spielfeld an der Stelle 2 gleich Spieler_zahl ist oder
        Spielfeld an der Stelle 3 gleich Spielfeld an der Stelle 6 ist und Spielfeld an der Stelle 3 gleich Spielfeld an der Stelle 9 ist und Spielfeld an der Stelle 3 gleich Spieler_zahl ist oder
        Spielfeld an der Stelle 1 gleich Spielfeld an der Stelle 5 ist und Spielfeld an der Stelle 1 gleich Spielfeld an der Stelle 9 ist und Spielfeld an der Stelle 1 gleich Spieler_zahl ist oder
        Spielfeld an der Stelle 3 gleich Spielfeld an der Stelle 5 ist und Spielfeld an der Stelle 3 gleich Spielfeld an der Stelle 7 ist und Spielfeld an der Stelle 3 gleich Spieler_zahl ist.
    
    Gib nicht Gewonnen zurück.
Und kann so benutzt werden:
    "das Spiel nicht gewonnen ist"

Die Funktion spiel_unentschieden gibt einen Wahrheitswert zurück, macht:
    Für jede Zahl feld in Spielfeld, wenn feld gleich 0 ist, gib falsch zurück.
    [ Falls alle felder belegt sind ] Gib wahr zurück.
Und kann so benutzt werden:
    "das Spiel unentschieden ist"

Die Funktion schreibe_spielfeld gibt nichts zurück, macht:
    Für jede Zahl i von 0 bis 2, mache:
        Schreibe den Text "+-+-+-+" auf eine Zeile.
        Für jede Zahl j von 0 bis 2, mache:
            Schreibe den Buchstaben '|'.
            Die Zahl Feld ist Spielfeld an der Stelle (i mal 3 plus j plus 1).

            Wenn Feld gleich 0 ist, Schreibe den Buchstaben ' '.
            Wenn aber Feld gleich 1 ist, Schreibe den Buchstaben 'X'.
            Wenn aber Feld gleich 2 ist, Schreibe den Buchstaben 'O'.
        Schreibe den Buchstaben '|' auf eine Zeile.
    Schreibe den Text "+-+-+-+" auf eine Zeile.
Und kann so benutzt werden:
    "Zeige das Spielfeld"

[ Programm anfang ]
Solange das Spiel nicht gewonnen ist, mache:
    Wenn das Spiel unentschieden ist, dann:
        Zeige das Spielfeld.
        Schreibe den Text "Das Spiel ist unentschieden" auf eine Zeile.
        Beende das Programm.

    Schreibe den Text "Bitte nehme einen Zug, Spieler ".
    Schreibe (die Zahl für den aktuellen Spieler) auf eine Zeile.
    Schreibe den Text "Das Spielfeld sieht so aus:" auf eine Zeile.
    Zeige das Spielfeld.
    Schreibe den Text "Schreibe die Position, wo du dein Zeichen setzen willst. [1-9]" auf eine Zeile.

    [Nutzereingabe validieren]
    Die Zahl Eingabe ist die nächste Zahl.
    Solange Eingabe kleiner als 1 ist oder Eingabe größer als 9 ist, mache:
        Schreibe den Text "Ungültige Eingabe! Bitte schreibe nur Zahlen von 1 bis 9!" auf eine Zeile.
        Speichere die nächste  Zahl in Eingabe.

    [Spielfeld aktualisieren und Spieler welchseln]
    Wenn Spielfeld an der Stelle Eingabe gleich 0 ist, dann:
        Speichere (die Zahl für den aktuellen Spieler) in Spielfeld an der Stelle Eingabe.
        Wechsel den Spieler.

Zeige das Spielfeld.

Wechsel den Spieler.
Schreibe den Text "Spieler ".
Schreibe (die Zahl für den aktuellen Spieler).
Schreibe den Text " hat gewonnen!" auf eine Zeile.
```