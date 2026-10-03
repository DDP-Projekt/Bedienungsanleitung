+++
title = "Operator overloading" 
weight = 4
type = "article"
+++

# Operator overloading

The built-in operators like `plus` or `als` only work with certain types.
With operator overloading you can use them for other types too, for example for your own [combinations](/en/Programmierung/Kombinationen).
To do this, you write a function that is called when the operator is used with the matching types.

## Syntax

A function that overloads an operator has no aliases.
Instead, it ends with `Und überlädt den "<operator>" Operator.`:

```ddp
Wir nennen die Kombination aus
	der Zahl x mit Standardwert 0,
	der Zahl y mit Standardwert 0,
einen Vektor, und erstellen sie so:
	"ein Vektor aus <x> und <y>"

Die Funktion Vektor_Plus mit den Parametern a und b vom Typ Vektor und Vektor, gibt einen Vektor zurück, macht:
	Gib ein Vektor aus (x von a plus x von b) und (y von a plus y von b) zurück.
Und überlädt den "plus" Operator.

Der Vektor a ist ein Vektor aus 1 und 2.
Der Vektor b ist ein Vektor aus 2 und 2.
Der Vektor c ist a plus b. [x = 3, y = 4]
```

Now, when two vectors are added with `plus`, `Vektor_Plus` is called.
The [operator prioritization](/en/Programmierung/Operatoren#operator-prioritization) stays the same.

You can also overload operators for built-in types.
An overload is always used before the built-in operator.

```ddp
Die Funktion Text_Plus_Zahl mit den Parametern t und z vom Typ Text und Zahl, gibt einen Text zurück, macht:
	Gib t verkettet mit z als Text zurück.
Und überlädt den "plus" Operator.

Schreibe ("Number " plus 5) auf eine Zeile. [Number 5]
```

## Rules

- The number of parameters must match the operator. A unary operator like `Betrag` needs one parameter, a binary operator like `plus` needs two and a ternary operator like `zwischen` needs three.
- The function must return a value. So the return type must not be `nichts`.
- An operator can only be overloaded once for the same parameter types.
- The `als` operator is an exception. It can be overloaded several times for the same parameter type, as long as the return type is different each time.
- Parameters may also be [references](/en/Programmierung/Funktionen/Referenzen).
- Inside the function itself, using the operator does not call the overload, but the normal operator.
- Public overloads also apply in other modules that include the module.

## The als operator

By overloading the `als` operator you can write your own type conversions.
The return type of the function is the type that is converted to:

```ddp
Die Funktion Vektor_Als_Text mit dem Parameter v vom Typ Vektor, gibt einen Text zurück, macht:
	Gib "(" verkettet mit (x von v) als Text verkettet mit ", " verkettet mit (y von v) als Text verkettet mit ")" zurück.
Und überlädt den "als" Operator.

Schreibe (ein Vektor aus 3 und 4 als Text) auf eine Zeile. [(3, 4)]
```

## Generic overloads

[Generic functions](/en/Programmierung/Funktionen/Generische-Funktionen) can also overload operators.
The overload is then used for all types the function body works with.

## List of operators

These names can be written between the quotation marks:

| Kind    | Operators |
| ------- | --------- |
| unary   | `"Betrag"`, `"Länge"`, `"unäres minus"`, `"nicht"`, `"logisch nicht"` |
| binary  | `"plus"`, `"minus"`, `"mal"`, `"durch"`, `"modulo"`, `"hoch"`, `"logarithmus"`, `"verkettet mit"`, `"an der Stelle"`, `"ab dem"`, `"bis zum"`, `"und"`, `"oder"`, `"entweder ... oder"`, `"logisch und"`, `"logisch oder"`, `"logisch kontra"`, `"links verschiebung"`, `"rechts verschiebung"`, `"gleich"`, `"ungleich"`, `"kleiner als"`, `"größer als"`, `"kleiner als, oder"`, `"größer als, oder"` |
| ternary | `"von bis"` (for `im Bereich von ... bis`), `"zwischen"` |
| type conversion | `"als"` |
