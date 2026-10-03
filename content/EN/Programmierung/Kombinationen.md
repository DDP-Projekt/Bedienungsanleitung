+++
title = "Combinations"
weight = 7
+++

# Combinations

Composite types, also called structures or records, are one of the most important data types in most programming languages.
They make it possible to define composite data types and thus also model more complex data.
In DDP we call them combinations (Kombinationen).

Similar to functions, readability is also a priority for combinations in DDP, although not quite as extreme, since they should also be usable.

## Declaration

A combination declaration generally looks like this:

```ddp
Wir nennen die (öffentliche) Kombination aus
    (der/dem (öffentlichen) <type name> <field name> mit Standardwert <expression>),
    ...
ein/eine/einen <combination name>, und erstellen sie so:
	"alias with field <x>" (,
	"another alias with the fields <x> and <y>" oder
	...)
```

Example:

```ddp
Wir nennen die öffentliche Kombination aus
	der Zahl x mit Standardwert 0, [private field]
	der öffentlichen Zahl y mit Standardwert 0, [public field]
einen Vektor2, und erstellen sie so:
	"Nullvektor2" oder
	"der Nullvektor2" oder [no parameter]
	"ein Vektor2 mit x gleich <x>" oder [1 parameter]
	"ein Vektor2 mit x gleich <x> und y gleich <y>" [all parameters]
```

Here a two-dimensional vector consisting of two numbers (x and y) is defined.
The combination itself and the field y are public, so they can also be used in other modules (more on this in the article [Modules](/en/Programmierung/Module)).

This notation may sound very mathematical, but it enables something very important.
The indefinite article (einer, eine or ein) indicates the grammatical gender (masculine, feminine or neuter) of the type name.
The Vektor2 combination defined above must therefore be treated as grammatically masculine:

```ddp
Der Vektor2 x ist der Nullvektor2. [correct]
Die Vektor2 y ist der Nullvektor2. [error: wrong article]

Die Vektor2 Liste vektoren ist eine leere Vektor2 Liste.
Für jeden [not jede or jedes!!] Vektor2 vek in vektoren, mache:
    ...
```

For practical reasons (especially with regard to generic types), the aliases and default values of combinations are optional.
That means that instead of the example above, you could also define a vector like this:

```ddp
Wir nennen die öffentliche Kombination aus
	der Zahl x,
	der öffentlichen Zahl y,
einen Vektor2.

Der Vektor2 v ist der Standardwert von einem Vektor2. [ x = 0, y = 0 ]
```

As you can see, in this case a combination can only be created with the help of the "Standardwert" operator.
Fields that have no explicit default value are also set to the default value of their type.

### Aliases

Combinations are created via aliases just like functions.
However, unlike with functions, a combination alias does not have to contain all fields.
The fields that are missing from a combination alias are simply set to their default value.

## Accessing fields

Let's assume we have declared the Vektor2 combination from above and created a Vektor2:

```ddp
Der Vektor2 vek ist der Nullvektor2.
```

The individual fields of vek are accessed with the `von` operator (`.` in other languages):

```ddp
Schreibe (x von vek). [0]
Schreibe (y von vek). [0]

Speichere 2 in x von vek. [vek.x is set to 2]

Schreibe (x von vek). [2]
Schreibe (y von vek). [0]
```

## Combination lists

Of course you can also have lists of combinations.
Assuming you have defined a combination `Vektor`, then a Vektor list looks like this:

```ddp
Die Vektor Liste vektoren ist eine leere Vektor Liste.

Der Vektor vek1 ist vektoren an der Stelle 1.
Die Zahl x ist x von (vektoren an der Stelle 1).
```

Unfortunately, for user-defined combinations it is not (yet) possible to decline the type name accordingly, so unlike with built-in types (Zahl -> Zahl*en* Liste vs. Vektor -> Vektor Liste) the name stays the same.

As you can see, the `von` operator also takes precedence over the `an der Stelle` operator, as specified in the [operator prioritization](/en/Programmierung/Operatoren/#operator-prioritization).

## Generic combinations

Just like there are [generic functions](/en/Programmierung/Funktionen/Generische-Funktionen), there are also generic combinations.
For these, the word `generische` is written before `Kombination`, and the fields can have type parameters like `T` as their type:

```ddp
Wir nennen die generische Kombination aus
	dem T erstes,
	dem R zweites,
ein Paar, und erstellen sie so:
	"ein Paar aus <erstes> und <zweites>"
```

Type parameters are grammatically neuter. That's why it says `dem T` here and not `der T`.

When you use a generic combination, you write the types for the type parameters with hyphens in front of the name.
The order is the order in which the type parameters first appear in the fields.
So for `Paar`, `T` comes first and then `R`:

```ddp
Das Zahl-Text-Paar p ist ein Paar aus 1 und "one".
Schreibe (erstes von p) auf eine Zeile. [1]
Schreibe (zweites von p) auf eine Zeile. [one]

Die Zahl-Text-Paar Liste liste ist eine leere Zahl-Text-Paar Liste.
```

If one of the types is a generic type itself, you put it in parentheses, for example `(Zahl-Paar)-Text-Paar`.

### Aliases of generic combinations

For an alias, the compiler must be able to find out all type parameters.
This works through the arguments of the alias or through the default values of the fields:

```ddp
Wir nennen die generische Kombination aus
	dem T x mit Standardwert 0,
	dem T y mit Standardwert 0,
einen Punkt, und erstellen sie so:
	"der Ursprung",
	"ein Punkt bei <x> und <y>"

Der Zahl-Punkt u ist der Ursprung. [T is a Zahl because of the default values]
Der Kommazahl-Punkt k ist ein Punkt bei 1,5 und 2,5. [T is a Kommazahl because of the arguments]
```

If the fields had no default values, the alias `"der Ursprung"` would not be allowed, because then it is not clear which type `T` is.
Without aliases, you can always create a generic combination with the `Standardwert` operator:

```ddp
Das Zahl-Text-Paar p ist der Standardwert von einem Zahl-Text-Paar.
```

### Generic combinations in functions

Generic functions can have generic combinations as parameters:

```ddp
Die generische Funktion Summe mit dem Parameter p vom Typ T-Punkt, gibt ein T zurück, macht:
	Gib x von p plus y von p zurück.
Und kann so benutzt werden:
	"die Summe von <p>"

Schreibe (die Summe von k) auf eine Zeile. [4]
```
