+++
title = "Type aliases and type definitions"
weight = 8
type = "article"
+++

# Type aliases and type definitions

Type aliases and type definitions allow you to create new names for existing types or to
define new types.
This feature makes it possible to make DDP programs considerably more readable with as little effort as possible.

## Type aliases

A type alias is simply another name for an existing type.
They are created like this:

```ddp
Wir nennen eine Zahl auch eine Hausnummer.
Wir nennen eine Zahl öffentlich auch eine Postleitzahl. 

Die Hausnummer h ist 22.
```

As you can see, type aliases can be public or private, just like combinations.
Type aliases just create a new name for the same type.
That means all operations are kept and the type can be passed to functions that expect the base type:

```ddp
Binde "Duden/Ausgabe" ein.

Wir nennen eine Zahl auch eine Hausnummer.
Die Hausnummer h ist 22.

[ Works even though the parameter of Schreibe_Zahl was declared as Zahl ]
Schreibe die Zahl h auf eine Zeile.

[ Works without the 'als' operator ]
Die Zahl z ist h.
```

## Type definitions

Type definitions look similar to type aliases, but have one crucial difference: they create *new* types.
That means operators and functions that expect the base type do not work with the new type:

```ddp
Wir definieren einen Zeiger als eine Zahl.
Wir definieren einen String öffentlich als einen Text.

[ Error: Ein Wert vom Typ Zahl kann keiner Variable vom Typ Zeiger zugewiesen werden (3001) ]
Der Zeiger zeiger ist 2.
[ Error: Der plus Operator erwartet einen Ausdruck vom Typ 'Zahl', oder 'Kommazahl' aber hat 'Zeiger' bekommen (3000) ]
Der Zeiger zeiger2 ist (2 als Zeiger) plus (2 als Zeiger).

[ OK ]
Der Zeiger z ist 2 als Zeiger.
[ OK ]
Der Zeiger z2 ist (z als Zahl plus z als Zahl) als Zeiger.
```

As you can see, you always have to convert type definitions correctly to be able to work with them.

Type definitions can be converted into their base type and back. However, this does not work recursively:

```ddp
Wir definieren einen Zeiger als eine Zahl.
Wir definieren einen Index als einen Zeiger.

[ OK ]
Der Zeiger z ist 0 als Zeiger.
[ OK ]
Der Index i ist z als Index.
[ Error: Ein Ausdruck vom Typ Index kann nicht in den Typ Zahl umgewandelt werden (3004) ]
Die Zahl z2 ist 1 als Zeiger als Index als Zahl.
```
