+++
title = "Combinations"
weight = 7
+++

# Combinations

Composite types, also called structures or records, are one of the most important data types in most programming languages
They make it possible to define composite data types and thus also model more complex data.\
In DDP we call them combinations.

Similar to functions, readability is also a priority when it comes to combinations in DDP, although not quite as extreme as they should also be usable

## Declaration

A combination declaration generally looks like this:

```ddp
Wir nennen die (öffentliche) Kombination aus
    (der/dem (öffentlichen) <Typ-Name> <Feld-Name> mit Standardwert <Ausdruck>),
    ...
ein/eine/einen <Kombinations-Name>, und erstellen sie so:
	"Alias mit Feld <x>" (,
	"Noch ein Alias mit den Feldern <x> und <y>" oder
	...)
```

Beispiel:

```ddp
Wir nennen die öffentliche Kombination aus
	der Zahl x mit Standardwert 0, [privates Feld]
	der öffentlichen Zahl y mit Standardwert 0, [öffentliches Feld]
einen Vektor2, und erstellen sie so:
	"Nullvektor2" oder
	"der Nullvektor2" oder [kein Parameter]
	"ein Vektor2 mit x gleich <x>" oder [1 Parameter]
	"ein Vektor2 mit x gleich <x> und y gleich <y>" [alle Parameter]
```

Here a two-dimensional vector consisting of two numbers (x and y) is defined.
The combination itself and the field y are public, so they can also be used in other modules (more on this in the article [Module](/en/Programmierung/Module/)).

This notation sounds very mathematical, but it enables something very important.
The indefinite article (einer, eine or ein) indicates the grammatical gender (masculine, feminine or neuter) of the type name.
The Vektor2 combination defined above must therefore be treated as grammatically masculine:

```ddp
Der Vektor2 x ist der Nullvektor2. [korrekt]
Die Vektor2 y ist der Nullvektor2. [Fehler: falscher Artikel]

Die Vektor2 Liste vektoren ist eine leere Vektor2 Liste.
Für jeden [nicht jede oder jedes!!] Vektor2 vek in vektoren, mache:
    ...
```

### Aliases

Combinations are created via aliases just like functions.
However, unlike functions, a combination salias does not have to contain all fields.\
The fields that are missing from a combination alias are simply set to the default value.

## Access to fields

Let's assume we declared the above Vektor2 combination and created a Vektor2:

```ddp
Der Vektor2 vek ist der Nullvektor.
```

The individual fields of vek are accessed using the `von` operator (`.` in other languages):

```ddp
Schreibe (x von vek). [0]
Schreibe (y von vek). [0]

Speichere 2 in x von vek. [vek.x wird auf 2 gesetzt]

Schreibe (x von vek). [2]
Schreibe (y von vek). [0]
```

## Combination lists

Of course you can also have lists of combinations.
Assuming you have defined a combination `Vektor`, then a Vektor list looks like this:

```ddp
Die Vektor Liste vektoren ist eine leere Vektoren Liste.

Der Vektor vek1 ist vektoren an der Stelle 1.
Die Zahl x ist x von (vektoren an der Stelle 1).
```
Unfortunately, with user-defined combinations it is not (yet) possible to decline the type name accordingly, so it is different than with built-in types (Zahl -> Zahl*en* Liste vs. Vektor -> Vektor Liste).

As you can see, the `von` operator also has priority over the `an der Stelle` operator, as specified in the [prioritization of operators](/en/Programmierung/Operatoren/#operator-prioritization).

## Generic combinations
<to-do></to-do>