# Funktionen
## War_EOF
```
Gibt wahr zurück wenn ein EOF Zeichen gelesen wurde.
```

* Rückgabe Typ: `Boolean`

### Aliase
1. `"die Benutzereingabe zu Ende gewesen ist"`

### Implementation
```ddp
	Gib war_eof zurück.
```
## Ist_Nicht_EOF
```
Gibt wahr zurück wenn noch kein EOF Zeichen gelesen wurde.
```

* Rückgabe Typ: `Boolean`

### Aliase
1. `"die Benutzereingabe nicht vorbei ist"`

### Implementation
```ddp
	Gib nicht war_eof zurück.
```
## Benutzereingabe_Zuruecksetzten
```
Setzt die interne war_eof variable auf falsch.
```

* Rückgabe Typ: `nichts`

### Aliase
1. `"setzte die Benutzereingabe zurück"`
2. `"Setzte die Benutzereingabe zurück"`

### Implementation
```ddp
	war_eof ist falsch.
```
## Lies_Token_in_Puffer
```
Liest eine einzelne Eingabe in Eingabe_Puffer.

Dabei werden vorrangehende Leerzeichen ignoriert
und ein einzelnes darrauffolgendes Leerzeichen in
Gepufferter_Buchstabe gespeichert
```

* Rückgabe Typ: `Text`

### Aliase
1. `"die nächste Eingabe"`
2. `"die naechste Eingabe"`

### Implementation
```ddp
	[vorangehende Leerstellen überspringen]
	Speichere das nächste Zeichen aus der Standardeingabe in Gepufferter_Buchstabe.
	Solange Gepufferter_Buchstabe ein Leerzeichen ist oder Gepufferter_Buchstabe ein Kontrollzeichen ist, mache:
		Wenn die Benutzereingabe zu Ende gewesen ist, gib "" zurück.
		Speichere das nächste Zeichen aus der Standardeingabe in Gepufferter_Buchstabe.

	[Text bis Leerstelle Lesen]
	Der Text token ist "".
	Solange nicht (Gepufferter_Buchstabe ein Leerzeichen ist oder Gepufferter_Buchstabe ein Kontrollzeichen ist), mache:
		Speichere token verkettet mit Gepufferter_Buchstabe in token.
		Speichere das nächste Zeichen aus der Standardeingabe in Gepufferter_Buchstabe.

	Gib token zurück.
```

