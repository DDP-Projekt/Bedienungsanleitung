+++
title = "Externe Funktionen"
weight = 3
type = "article"
+++

# Externe Funktionen

Eine externe Funktion ist eine Funktion, deren Körper nicht im DDP-Quellcode, sondern außerhalb (z.B. in anderen Sprachen oder statischen Bibliotheken) definiert ist.

Die Allgemeine Form sieht so aus:
```ddp
Die Funktion <Name> (mit den Parametern x, y und z vom Typ T1, T2 und T3,) gibt <Rückgabetyp> zurück,
ist in "Dateipfad" definiert
und kann so benutzt werden:
	"Alias mit Parameter <x> <y> und <z>" (,
	"<z> Noch ein Alias mit allen Parametern <y> <x>" oder
	...)
```

Man sieht, dass externe Funktionen identisch mit normalen Funktionen sind, allerdings ist ihr Körper kein DDP-Quellcode sondern ein Verweis auf eine Datei, in der der Funktionskörper dann definiert ist.
Mögliche Dateien sind .c, .o, .a und .lib Dateien.

Der Körper sollte in C implementiert sein.
Obwohl es möglich ist, ihn in anderen Sprachen zu definieren, wird das nicht empfohlen, da die DDP-Runtime in C geschrieben ist, und somit alles C-Calling-Conventions und C Datentypen benutzt und allgemein recht eng mit der Runtime verbunden ist.

Sollte die gegebene Datei eine .c Datei sein, so wird diese mit GCC zu einer .o Datei kompiliert und in die finale Ausgabedatei gelinkt.
.o, .a oder .lib Dateien werden direkt in die finale Ausgabedatei gelinkt.

## Anwendungsfälle

Purer DDP-Quellcode könnte nicht viel tun, nichtmal etwas auf die Konsole schreiben.
Deshalb sind Externe Funktionen enorm wichtig, da sie Schnittstellen zum Betriebssystem und vielen Bibliotheken bilden.
Ein gutes Beispiel dafür ist der Duden. Sehr viele Funktionen darin sind extern und in C geschrieben. Der Duden bildet einen guten Einstiegspunkt um das Nutzen von externen Funktionen zu lernen und sich Beispiele in C anzusehen wie sie mit der DDP-Runtime interagieren.

# Extern sichtbare Funktionen

Der DDP Kompilierer benutzt eine Technik namens [name mangling](https://en.wikipedia.org/wiki/Name_mangling). Das heißt, die Namen von Funktionen und Variablen
sind im Quellcode nicht dieselben wie in der entstehenden Binärdatei.

Wenn man also eine nicht-externe DDP Funktion aus C-Code benutzen will, muss man das name-mangling ausschalten:
```ddp
[test.ddp]
Die Funktion ddp_funktion gibt nichts zurück, macht:
    Schreibe "Ich bin eine DDP Funktion" auf eine Zeile.
Und kann so benutzt werden:
    "ddp_funktion"
```

```c
// test.c
extern void ddp_funktion(void);

int main(void) {
    ddp_funktion();
    return 0;
}
```