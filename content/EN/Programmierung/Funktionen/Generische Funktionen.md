+++
title = "Generic Functions"
weight = 5
+++

# Generic functions

Sometimes you want to write the same function for different types.
A function that returns the bigger of two values, for example, works with Zahlen just like with Kommazahlen.
So that you don't have to write the function for every type separately, there are generic functions.

## Declaration

A generic function is declared like a normal function.
The only addition is the word `generische`.
You can then use any new name as the type of a parameter, for example `T`.
This name is a so-called type parameter. It stands for a type that is only known when the function is called.

```ddp
Die generische Funktion Größtes mit den Parametern a und b vom Typ T und T, gibt ein T zurück, macht:
	Gib a, falls a größer als b ist, ansonsten b zurück.
Und kann so benutzt werden:
	"das Größere von <a> und <b>"

Schreibe (das Größere von 3 und 7) auf eine Zeile. [7]
Schreibe (das Größere von 2,5 und 1,5) auf eine Zeile. [2,5]
```

In the first call `T` is a Zahl, in the second call it is a Kommazahl.
The compiler recognizes the type from the arguments and creates a separate version of the function for every type.

Type parameters are grammatically neuter. So you write `ein T` and `das T`.

A generic function needs at least one type parameter.
But there can also be several, for example `T` and `R`.

## Type parameters in other types

Type parameters can also be used in lists and references:

```ddp
Die generische Funktion Erstes mit dem Parameter liste vom Typ T Liste, gibt ein T zurück, macht:
	Gib liste an der Stelle 1 zurück.
Und kann so benutzt werden:
	"das erste Element von <liste>"

Schreibe (das erste Element von (eine Liste, die aus 4, 5, 6 besteht)) auf eine Zeile. [4]
Schreibe (das erste Element von (eine Liste, die aus "a", "b" besteht)) auf eine Zeile. [a]
```

```ddp
Die generische Funktion Tausche mit den Parametern a und b vom Typ T Referenz und T Referenz, gibt nichts zurück, macht:
	Das T temp ist a.
	Speichere b in a.
	Speichere temp in b.
Und kann so benutzt werden:
	"Tausche <a> und <b>"

Der Text x ist "left".
Der Text y ist "right".
Tausche x und y.
Schreibe x auf eine Zeile. [right]
```

Inside the function body you can use `T` like any other type, for example to declare variables.

## Errors in generic functions

The function body is only checked when the function is called with a specific type.
If you call `Größtes` with two Wahrheitswerte, for example, there is an error, because the `größer als` operator does not work with Wahrheitswerte.
The error is then reported at the call.

## Restrictions

- A generic function cannot be [externally visible](/en/Programmierung/Funktionen/Externe-Funktionen#externally-visible-functions).
- A generic function must be defined immediately. A [forward declaration](/en/Programmierung/Funktionen/Anwendung#forward-declarations) is not possible.
- Generic functions can also be defined [externally](/en/Programmierung/Funktionen/Externe-Funktionen#generic-external-functions). But then there are a few additional rules.

Generic functions can also overload operators. More about this in the article [Operator overloading](/en/Programmierung/Funktionen/Operatoren-Ueberladung).
How to create generic combinations is described in the article [Combinations](/en/Programmierung/Kombinationen#generic-combinations).
