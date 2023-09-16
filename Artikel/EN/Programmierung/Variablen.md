# Variables

Variables in DDP, as in other languages, are "containers" with names in which values are stored.

# Declaration

Variables can be created (or declared) as follows:

```ddp
<data type with article> <variable name> ist <expression>.
```

There is a special feature for variables with the data type "Boolean".\
If you want to declare such variables with an expression, you should use this syntax instead:
```ddp
Der Boolean <variable name> ist <true or false>, wenn <expression>.
```
You can find a list of all data types in the article [data types](/Bedienungsanleitung/EN/Programmierung/Datentypen)

## Examples:

```dpp
Die Zahl a ist 10.
Die Kommazahl b ist 4,32.
Der Text c ist "Hallo!".
Der Boolean d ist wahr.
Der Boolean e ist falsch wenn 1 gleich 1 ist. 
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

`Erhöhe <variable> um <a>`  
equivalent to  
`Speichere <variable> plus <a> in <variable>`

## Subtraction

`Verringere <variable> um <a>`  
equivalent to  
`Speichere <variable> minus <a> in <variable>`

## Multiplication

`Vervielfache <variable> um <a>.`  
equivalent to  
`Speichere <variable> mal <a> in <variable>`

## Division

`Teile <variable> durch <a>.`  
equivalent to  
`Speichere <variable> durch <a> in <variable>`

## Negation

`Negiere <variable>.`  
equivalent to  
`Speichere -<variable> in <variable>` resp.  
`Speichere nicht <variable> in <variable>`

## Bitshift

`Verschiebe <variable> um <a> Bit nach Links`  
`Verschiebe <variable> um <a> Bit nach Rechts`  
equivalent to  
`Speichere <variable> um <a> Bit nach Links verschoben in <variable>`  
`Speichere <variable> um <a> Bit nach Rechts verschoben in <variable>`