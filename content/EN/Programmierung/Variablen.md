+++
title = "Variables"
weight = 4
type = "article"
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
Der Wahrheitswert <variable name> ist <true or false>, wenn <expression>.
```
You can find a list of all data types in the article [data types](/en/Programmierung/Datentypen)

## Examples:

```ddp
Die Zahl a ist 10.
Die Kommazahl b ist 4,32.
Der Text c ist "Hallo!".
Der Wahrheitswert d ist wahr.
Der Wahrheitswert e ist falsch wenn 1 gleich 1 ist. 
```

# Assignment

Assignment is changing the value of a variable. There are several ways to change variables in DDP.

With the keyword `is` you can only assign variables to a literal:
```ddp
a ist 30.
```

To assign the result of an expression to a variable, `Speichere ... in` must be used:
```ddp
Speichere pi durch 2 in b.
```

# Special assignments

There are still some assignment operators that can be used to directly change variables,
without having to use them in an expression yourself.
In other languages, these are so-called "compound assignements", i.e. operators like `+=, -=, *=, etc.` .
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