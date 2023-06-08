# Funktionen
## PolsterLinks
* Parameter: `text`, `zeichen`, `endlänge`
* Parameter Typen: `Text`, `Buchstabe`, `Zahl`
* Rückgabe Typ: `Text`

### Aliase
1. `"<text> mit <endlänge> <zeichen> links gepolstert"`

### Implementation
```ddp
	Die Zahl länge ist die Länge von text.
	Die Zahl gesuchteLänge ist endlänge minus länge.
	Wenn gesuchteLänge kleiner als, oder 0 ist, dann:
		Gib text zurück.
	
	Wiederhole:
		Stelle zeichen vor text.
	gesuchteLänge Mal.

	Gib text zurück.
```
## TrimEndeWert
* Parameter: `text`, `zeichen`
* Parameter Typen: `Text`, `Buchstabe`
* Rückgabe Typ: `Text`

### Aliase
1. `"<text> mit allen <zeichen> danach entfernt"`

### Implementation
```ddp
	Entferne alle zeichen nach text.
	Gib text zurück.
```
## Lösche_Text_Bereich
* Parameter: `text`, `start`, `end`
* Parameter Typen: `Text Referenz`, `Zahl`, `Zahl`
* Rückgabe Typ: `nichts`

### Aliase
1. `"Lösche alle Elemente von <start> bis <end> aus <text>"`

### Implementation
```ddp
	Wenn start gleich 1 ist, dann:
		Speichere text von end plus 1 bis (die Länge von text) in text.
	Sonst:
		Speichere text von 1 bis (start minus 1) verkettet mit text von (end plus 1) bis (die Länge von text) in text.
```
## Kleinschreiben
* Parameter: `text`
* Parameter Typ: `Text Referenz`
* Rückgabe Typ: `nichts`

### Aliase
1. `"Schreibe <text> klein"`

### Implementation
```ddp
	Speichere text klein geschrieben in text.
```
## EndetMitBuchstabe
* Parameter: `text`, `buchstabe`
* Parameter Typen: `Text`, `Buchstabe`
* Rückgabe Typ: `Boolean`

### Aliase
1. `"<buchstabe> am Ende von <text> steht"`

### Implementation
```ddp
	Wenn die Länge von text gleich 0 ist, gib falsch zurück.
	Gib (text an der Stelle (die Länge von text)) gleich buchstabe ist zurück.
```
## Text_an_Text_fügen
* Parameter: `text`, `elm`
* Parameter Typen: `Text Referenz`, `Text`
* Rückgabe Typ: `nichts`

### Aliase
1. `"Füge <elm> an <text> an"`

### Implementation
```ddp
	Speichere text verkettet mit elm in text.
```
## Buchstabe_an_Text_fügen
* Parameter: `text`, `elm`
* Parameter Typen: `Text Referenz`, `Buchstabe`
* Rückgabe Typ: `nichts`

### Aliase
1. `"Füge <elm> an <text> an"`

### Implementation
```ddp
	Speichere text verkettet mit elm in text.
```
## IndexVonText
* Parameter: `text`, `elm`
* Parameter Typen: `Text`, `Buchstabe`
* Rückgabe Typ: `Zahl`

### Aliase
1. `"der Index von <elm> in <text>"`

### Implementation
```ddp
	Wenn die Länge von text gleich 0 ist, gib -1 zurück.
	Für jede Zahl i von 1 bis (die Länge von text), Wenn text an der Stelle i gleich elm ist, gib i zurück.
	Gib -1 zurück.
```
## TrimEnde
* Parameter: `text`, `zeichen`
* Parameter Typen: `Text Referenz`, `Buchstabe`
* Rückgabe Typ: `nichts`

### Aliase
1. `"Entferne alle <zeichen> nach <text>"`

### Implementation
```ddp
	Die Zahl index ist die Länge von text.
	Wenn index gleich 0 ist, verlasse die Funktion.
	Solange (text an der Stelle index) gleich zeichen ist, verringere index um 1.
	Speichere text von 1 bis index in text.
```
## BeginntMitBuchstabe
* Parameter: `text`, `buchstabe`
* Parameter Typen: `Text`, `Buchstabe`
* Rückgabe Typ: `Boolean`

### Aliase
1. `"<buchstabe> am Anfang von <text> steht"`

### Implementation
```ddp
	Wenn die Länge von text gleich 0 ist, gib falsch zurück.
	Gib (text an der Stelle 1) gleich buchstabe ist zurück.
```
## EnthältText
* Parameter: `text`, `suchText`
* Parameter Typen: `Text`, `Text`
* Rückgabe Typ: `Boolean`

### Aliase
1. `"<text> <suchText> enthält"`

### Implementation
```ddp
	Die Zahl startIndex ist 0.
	Die Zahl endIndex ist die Länge von suchText.

	Solange endIndex kleiner als, oder die Länge von text ist, mache:
		Der Text subtext ist text von startIndex bis endIndex.

		Wenn subtext gleich suchText ist, gib wahr zurück.
		
		Speichere startIndex plus die Länge von suchText in endIndex.
		Erhöhe startIndex um 1.
	Gib falsch zurück.
```
## Lösche_Text
* Parameter: `text`, `index`
* Parameter Typen: `Text Referenz`, `Zahl`
* Rückgabe Typ: `nichts`

### Aliase
1. `"Lösche das Element an der Stelle <index> aus <text>"`

### Implementation
```ddp
	Speichere text von 1 bis (index minus 1) verkettet mit text von (index plus 1) bis (die Länge von text) in text.
```
## IstTextLeer
* Parameter: `text`
* Parameter Typ: `Text`
* Rückgabe Typ: `Boolean`

### Aliase
1. `"<text> leer ist"`

### Implementation
```ddp
	Gib [wahr wenn] die Länge von text gleich 0 ist zurück.
```
## GroßschreibenWert
* Parameter: `text`
* Parameter Typ: `Text`
* Rückgabe Typ: `Text`

### Aliase
1. `"<text> groß geschrieben"`

### Implementation
```ddp
	Der Text neuerText ist "".
	Für jeden Buchstaben b in text, mache:
		Füge (b als großer Buchstabe) an neuerText an.
	Gib neuerText zurück.
```
## TrimWert
* Parameter: `text`, `zeichen`
* Parameter Typen: `Text`, `Buchstabe`
* Rückgabe Typ: `Text`

### Aliase
1. `"<text> mit allen <zeichen> davor und danach entfernt"`

### Implementation
```ddp
	Entferne alle zeichen vor und nach text.
	Gib text zurück.
```
## EnthältBuchstabe
* Parameter: `text`, `zeichen`
* Parameter Typen: `Text`, `Buchstabe`
* Rückgabe Typ: `Boolean`

### Aliase
1. `"<text> <zeichen> enthält"`

### Implementation
```ddp
	Für jeden Buchstaben b in text, wenn b gleich zeichen ist, gib wahr zurück.
	Gib falsch zurück.
```
## BeginntMitText
* Parameter: `text`, `suchText`
* Parameter Typen: `Text`, `Text`
* Rückgabe Typ: `Boolean`

### Aliase
1. `"<suchText> am Anfang von <text> steht"`

### Implementation
```ddp
	Gib (text von 1 bis (die Länge von suchText)) gleich suchText ist zurück.
```
## Text_vor_Text_stellen
* Parameter: `text`, `elm`
* Parameter Typen: `Text Referenz`, `Text`
* Rückgabe Typ: `nichts`

### Aliase
1. `"Stelle <elm> vor <text>"`

### Implementation
```ddp
	Speichere elm verkettet mit text in text.
```
## KleinschreibenWert
* Parameter: `text`
* Parameter Typ: `Text`
* Rückgabe Typ: `Text`

### Aliase
1. `"<text> klein geschrieben"`

### Implementation
```ddp
	Der Text neuerText ist "".
	Für jeden Buchstaben b in text, mache:
		Füge (b als kleiner Buchstabe) an neuerText an.
	Gib neuerText zurück.
```
## Spalte
* Parameter: `text`, `zeichen`
* Parameter Typen: `Text`, `Buchstabe`
* Rückgabe Typ: `Text Liste`

### Aliase
1. `"<text> an <zeichen> gespalten"`

### Implementation
```ddp
	Die Text Liste endliste ist eine leere Text Liste.
	Die Zahl endIndex ist der Index von zeichen in text.
	Solange endIndex ungleich -1 ist und endIndex ungleich die Länge von text ist, mache:
		Wenn endIndex ungleich 1 ist, dann:
			Speichere endliste verkettet mit text von 1 bis (endIndex minus 1) in endliste.
			Speichere text von endIndex plus 1 bis (die Länge von text) in text.
		Sonst:
			Speichere endliste verkettet mit "" in endliste.
			Speichere text von 2 bis (die Länge von text) in text.
		Speichere der Index von zeichen in text in endIndex.
	
	Speichere endliste verkettet mit text in endliste.
	Gib endliste zurück.
```
## TrimAnfangWert
* Parameter: `text`, `zeichen`
* Parameter Typen: `Text`, `Buchstabe`
* Rückgabe Typ: `Text`

### Aliase
1. `"<text> mit allen <zeichen> davor entfernt"`

### Implementation
```ddp
	Entferne alle zeichen vor text.
	Gib text zurück.
```
## Fülle_Text
* Parameter: `text`, `elm`
* Parameter Typen: `Text Referenz`, `Buchstabe`
* Rückgabe Typ: `nichts`

### Aliase
1. `"Fülle <text> mit <elm>"`

### Implementation
```ddp
	Der Text neuerText ist "".
	Wiederhole:
		Speichere neuerText verkettet mit elm in neuerText.
	Die Länge von text Mal.

	Speichere neuerText in text.
```
## Text_leeren
* Parameter: `text`
* Parameter Typ: `Text Referenz`
* Rückgabe Typ: `nichts`

### Aliase
1. `"Leere <text>"`

### Implementation
```ddp
	Speichere "" in text.
```
## Text_in_Text_einfügen
* Parameter: `text`, `index`, `elm`
* Parameter Typen: `Text Referenz`, `Zahl`, `Text`
* Rückgabe Typ: `nichts`

### Aliase
1. `"Setze <elm> an die Stelle <index> von <text>"`

### Implementation
```ddp
	Speichere text von 1 bis (index minus 1) verkettet mit elm verkettet mit text von index bis (die Länge von text) in text.
```
## Buchstabe_vor_Text_stellen
* Parameter: `text`, `elm`
* Parameter Typen: `Text Referenz`, `Buchstabe`
* Rückgabe Typ: `nichts`

### Aliase
1. `"Stelle <elm> vor <text>"`

### Implementation
```ddp
	Speichere elm verkettet mit text in text.
```
## Ist_Zahl
```
ob ein Text in eine Zahl umgewandelt werden kann
```
* Parameter: `t`
* Parameter Typ: `Text`
* Rückgabe Typ: `Boolean`

### Aliase
1. `"<t> eine Zahl ist"`

### Implementation
```ddp
	Die Zahl l ist die Länge von t.
	Wenn l kleiner als 1 ist, gib falsch zurück.

	Der Buchstabe Vorzeichen ist t an der Stelle 1.
	Wenn Vorzeichen eine Zahl ist, gib wahr zurück.
	Wenn Vorzeichen ungleich '+' ist und Vorzeichen ungleich '-' ist, gib falsch zurück.
	Wenn l kleiner als 2 ist oder nicht (t an der Stelle 2) eine Zahl ist, gib falsch zurück.
	Gib wahr zurück.
```
## Trim
* Parameter: `text`, `zeichen`
* Parameter Typen: `Text Referenz`, `Buchstabe`
* Rückgabe Typ: `nichts`

### Aliase
1. `"Entferne alle <zeichen> vor und nach <text>"`

### Implementation
```ddp
	Die Zahl startIndex ist 1.
	Die Zahl stopIndex ist die Länge von text.
	Solange (text an der Stelle startIndex) gleich zeichen ist, erhöhe startIndex um 1.
	Solange (text an der Stelle stopIndex) gleich zeichen ist, verringere stopIndex um 1.
	Speichere text von startIndex bis stopIndex in text.
```
## EndetMitText
* Parameter: `text`, `suchText`
* Parameter Typen: `Text`, `Text`
* Rückgabe Typ: `Boolean`

### Aliase
1. `"<suchText> am Ende von <text> steht"`

### Implementation
```ddp
	Gib (text von die Länge von text minus die Länge von suchText plus 1 bis (die Länge von text)) gleich suchText ist zurück.
```
## Großschreiben
* Parameter: `text`
* Parameter Typ: `Text Referenz`
* Rückgabe Typ: `nichts`

### Aliase
1. `"Schreibe <text> groß"`

### Implementation
```ddp
	Speichere text groß geschrieben in text.
```
## PolsterRechts
* Parameter: `text`, `zeichen`, `endlänge`
* Parameter Typen: `Text`, `Buchstabe`, `Zahl`
* Rückgabe Typ: `Text`

### Aliase
1. `"<text> mit <endlänge> <zeichen> rechts gepolstert"`

### Implementation
```ddp
	Die Zahl länge ist die Länge von text.
	Die Zahl gesuchteLänge ist endlänge minus länge.
	Wenn gesuchteLänge kleiner als, oder 0 ist, dann:
		Gib text zurück.
	
	Wiederhole:
		Füge zeichen an text an.
	gesuchteLänge Mal.
	
	Gib text zurück.
```
## TrimAnfang
* Parameter: `text`, `zeichen`
* Parameter Typen: `Text Referenz`, `Buchstabe`
* Rückgabe Typ: `nichts`

### Aliase
1. `"Entferne alle <zeichen> vor <text>"`

### Implementation
```ddp
	Wenn die Länge von text gleich 0 ist, verlasse die Funktion.
	Die Zahl index ist 1.
	Solange (text an der Stelle index) gleich zeichen ist, erhöhe index um 1.
	Speichere text von index bis (die Länge von text) in text.
```
## Buchstabe_in_Text_einfügen
* Parameter: `text`, `index`, `elm`
* Parameter Typen: `Text Referenz`, `Zahl`, `Buchstabe`
* Rückgabe Typ: `nichts`

### Aliase
1. `"Setze <elm> an die Stelle <index> von <text>"`

### Implementation
```ddp
	Speichere text von 1 bis (index minus 1) verkettet mit elm verkettet mit text von index bis (die Länge von text) in text.
```

