+++
title = "Variables"
weight = 4
+++

# Variables

Variables in DDP, as in other languages, are "containers" with names in which values are stored.

# Declaration

Variables can be created (or declared) as follows:

```ddp
<data type with article> <variable name> ist <expression>.
```

There is a special feature for variables with the data type "Wahrheitswert".\
If you want to declare such variables with an expression, you should use this syntax instead:
```ddp
Der Wahrheitswert <variable name> ist <wahr or falsch>, wenn <expression>.
```

This syntax also works with return statements in [functions](/en/Programmierung/Funktionen):
```ddp
Die öffentliche Funktion Ist_Leer_Text mit dem Parameter liste vom Typ Text Liste, gibt einen Wahrheitswert zurück, macht:
	Gib wahr, wenn die Länge von liste gleich 0 ist, zurück.
Und kann so benutzt werden:
	"<liste> leer ist"
```

You can find a list of all data types in the article [data types](/en/Programmierung/Datentypen)

## Examples:

```ddp
Die Zahl a ist 10.
Die Kommazahl b ist 4,32.
Der Text c ist "Hello!".
Der Wahrheitswert d ist wahr.
Der Wahrheitswert e ist falsch, wenn 1 gleich 1 ist. 
```

# Assignment

Assignment is changing the value of a variable. There are several ways to change variables in DDP.

With the keyword `ist` you can only assign literals to variables:
```ddp
a ist 30.
```

To assign the result of an expression to a variable, `Speichere ... in` must be used:
```ddp
Speichere pi durch 2 in b.
```

It is also possible to assign a value of a different numeric type to a numeric variable. The value is then [converted](/en/Programmierung/Typkonvertierung) accordingly. For example, here a Zahl is assigned to a Kommazahl:
```ddp
Die Kommazahl k ist 5.
```

# Externally visible variables

The DDP compiler uses a technique called [name mangling](https://en.wikipedia.org/wiki/Name_mangling). This means that the names of functions and variables
in the source code are not the same as in the resulting binary.

So if you want to share a variable between DDP and C source code, you have to turn off name mangling by marking the variable as "extern sichtbar" (externally visible):
```ddp
[test.ddp]
Die extern sichtbare Zahl z ist 22.
```

```c
// test.c
#include "DDP/ddptypes.h"
#include <stdio.h>

extern ddpint z;

int main(void) {
    printf("%ld", z);
    return 0;
}
```

# Constants

In addition to variables, you can also declare constants.
Constants cannot be changed and really only serve as an alias name for a value.

```ddp
Die Konstante pi ist 3,1415.
```

Constants can only be literals (e.g. `1`, `""` or `1,5` but NOT `1 plus 1`) and cannot be changed.
The type of a constant is derived from its value.

# Special assignments

There are some more assignment operators that can be used to change variables directly,
without having to use them in an expression yourself.
In other languages, these are so-called "compound assignments", i.e. operators like `+=, -=, *=, etc.` .
These operators can be used to make code more readable.

## Addition

```ddp
Erhöhe <variable> um <a>
```  
equivalent to  
```ddp
Speichere <variable> plus <a> in <variable>
```

## Subtraction

```ddp
Verringere <variable> um <a>
```  
equivalent to  
```ddp
Speichere <variable> minus <a> in <variable>
```

## Multiplication

```ddp
Vervielfache <variable> um <a>.
```
equivalent to  
```ddp
Speichere <variable> mal <a> in <variable>
```

## Division

```ddp
Teile <variable> durch <a>.
```
equivalent to  
```ddp
Speichere <variable> durch <a> in <variable>
```

## Negation

```ddp
Negiere <variable>.
```
equivalent to  
```ddp
Speichere -<variable> in <variable>
```
or  
```ddp
Speichere nicht <variable> in <variable>
```

## Bitshift
```ddp
Verschiebe <variable> um <a> Bit nach Links
```
```ddp
Verschiebe <variable> um <a> Bit nach Rechts
```
equivalent to  
```ddp
Speichere <variable> um <a> Bit nach Links verschoben in <variable>
```
```ddp
Speichere <variable> um <a> Bit nach Rechts verschoben in <variable>
```
