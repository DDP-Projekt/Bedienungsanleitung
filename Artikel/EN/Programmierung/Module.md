# Modules

Larger and more complex programs often consist of several source files to which the program code is distributed.

DDP provides modules for this purpose.

## Principle

Each DDP file represents a DDP module.
Each DDP module can be linked by another one.
To do this, however, it must reveal an externally visible (public) interface.
This is possible with the keyword "public".
The public interface of a DDP module is the set of all variables and functions declared as "public".

When a DDP module integrates another, it gets access to its public functions and variables and can use them yourself.

In other languages, this feature is often implemented using the "import" keyword (or "#include" in C), and restricting the visibility of names is called encapsulation.

## syntax

```ddp
Binde "<relativer Pfad>" ein.
Binde "Duden/<Datei aus der Standardbibliothek>" ein.
Binde a aus "<relativer Pfad>" ein.
Binde a, b und c aus "<relativer Pfad>" ein.
```

### Concrete example
A.ddp:
```ddp
Binde "Duden/Ausgabe" ein.

Die öffentliche Zahl z ist 1.

Die öffentliche Funktion foo gibt nichts zurück, macht:
	Schreibe "foo" auf eine Zeile.
Und kann so benutzt werden:
	"foo"

Die Zahl zz ist 1.

Die Funktion bar gibt nichts zurück, macht:
	Schreibe "bar" auf eine Zeile.
Und kann so benutzt werden:
	"bar"
```

B.ddp:
```ddp
Binde "A" ein.

foo. [calls foo form A.ddp]
Speichere 2 in z. [sets z from A.ddp to 2]

[
	bar and zz are not recognized as a function or variable, since they are not declared as public in A.ddp.
]
bar. 
Speichere 2 in zz. [sets z from A.ddp to 2]

[
	Error:
	"Duden/Output" was included in A.ddp, but not in B.ddp, so the 'Schreibe' function is not available
]
Schreibe "Hallo Welt" auf eine Zeile.
```

C.ddp:
```ddp
Binde foo aus "A" ein.

foo. [calls foo from A.ddp]
[
	Error:
	Only foo was included, but not z.
]
Speichere 2 in z.

[
	Error:
	bar and zz are not recognized as a function or variable,
	since they are not declared as public in A.ddp.
]
bar. 
Speichere 2 in zz. [setzt z aus A.ddp auf 2]
```

## Explanation

As you can see in the example, you either enter a relative path to another .DDP file (without the .ddp ending) when embedding, or a path to a file from the standard library. The files from the standard library are always included with `Duden/<File>`.

Either all public names or just a few special ones can be included.
This is useful for not including too many unnecessary function aliases that might cause trouble.

Also, includes are not "inherited" between files.
So if B.ddp includes A.ddp, B.ddp does not accept A.ddp's inclusions (as in the "Duden/Output" example).

To navigate through directories one can use unix file paths.

With this file structure:
- root
	- Ordner1
		- A.ddp
	- Ordner2
		- B.ddp

B.ddp would have to contain `Binde "../Ordner1/A" ein.` to include A.ddp.