+++
title = "Modules"
weight = 8
+++

# Modules

Larger and more complex programs often consist of several source files across which the program code is distributed.

DDP provides modules for this purpose.

## Principle

Every DDP file is a DDP module.
Every DDP module can be included by another one.
To do this, however, it must expose an externally visible (public) interface.
This is possible with the keyword "öffentliche" (public).
The public interface of a DDP module is the set of all variables, constants, functions, combinations, type aliases and type definitions declared as "öffentliche".

When a DDP module includes another one, it gets access to its public functions and variables and can use them itself.

In other languages, this feature is often implemented using the "import" keyword (or "#include" in C), and restricting the visibility of names is called encapsulation.

## Syntax

```ddp
Binde "<relative path>" ein.
Binde "Duden/<file from the standard library>" ein.
Binde a aus "<relative path>" ein.
Binde a, b und c aus "<relative path>" ein.
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

foo. [calls foo from A.ddp]
Speichere 2 in z. [sets z from A.ddp to 2]

[
	bar and zz are not recognized as a function or variable,
	since they are not declared as public in A.ddp.
]
bar. 
Speichere 2 in zz.

[
	Error:
	"Duden/Ausgabe" was included in A.ddp, but not in B.ddp,
	so the Schreibe function is not available.
]
Schreibe "Hello World" auf eine Zeile.
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
Speichere 2 in zz.
```

## Explanation

As you can see in the example, when including you either specify a relative path to another .ddp file (without the .ddp extension),
or a path to a file from the standard library. Files from the standard library are always included with `Duden/<file>`.

Either all public names or just some specific ones can be included.
This is useful to avoid including too many unnecessary function aliases that might cause trouble.

Includes are also not "inherited" between files.
So if B.ddp includes A.ddp, B.ddp does not take over A.ddp's includes (like "Duden/Ausgabe" in the example).

To navigate through directories, you can use unix file paths.

With this file structure:
- root
	- Folder1
		- A.ddp
	- Folder2
		- B.ddp

B.ddp would have to contain `Binde "../Folder1/A" ein.` to include A.ddp.


## Directory includes

For convenience, you can also include all modules in a directory.

- root
	- Folder1
		- A.ddp
        - B.ddp
	- Folder2
        - C.ddp
        - Folder3
            - D.ddp

```ddp
Binde alle Module aus "Folder1" ein. [includes A.ddp and B.ddp]
Binde rekursiv alle Module aus "Folder2" ein. [includes C.ddp and D.ddp]
```
