+++
title = "Function declarations"
weight = 1
+++

# Function declarations

Every function must be declared before it can be used. In DDP there are only global functions, i.e. those that are not inside a statement block.
Every function needs a unique **name**, a list of **parameters** and their **types**, a **return type**, and one or more **aliases**.

Here is the general form of a function declaration (optional parts are enclosed in parentheses):
```ddp
Die Funktion <name> (mit den Parametern x, y und z vom Typ T1, T2 und T3,) gibt <return type> zurück, macht:
	<function body>
Und kann so benutzt werden:
	"alias with parameters <x> <y> and <z>" (,
	"<z> another alias with all parameters <y> <x>" oder
	...)
```

## The name

of a function must be unique and should describe the function well, but is otherwise not important and is not used directly in the code. It becomes more important with external functions, which are described later.

## The parameters

of a function are optional. There can be between 0 and infinitely many. Every parameter requires a type.
When writing the parameters, the grammatical number (singular, plural) and the rules of enumeration must be followed.

With a single parameter you write
```ddp
... mit dem Parameter x vom Typ T, ...
```
With two you write
```ddp
... mit den Parametern x und y vom Typ T1 und T2, ...
```
And with N parameters
```ddp
... mit den Parametern a, b, ... und z vom Typ T1, T2, ... und T26, ...
```

## The return type

of a function must be present, but may be `nichts`.
`nichts` is not a type, but a placeholder for functions that do not produce a value and only have side effects.
The result of a function that returns `nichts` can't be used because, well, it doesn't exist.

If a function returns a value, there must be a return statement at the end of the function body.
A return statement looks like this:
```ddp
Gib <value of the return type> zurück.
```
A return statement can appear anywhere in the function body, and exits the function with the returned value as its result.

To exit functions that return `nichts` early, you can use `Verlasse die Funktion`.

## Aliases

are the way functions are called.<br>
Since the goal of DDP is to be as good German as possible, it is hard to call functions correctly in every grammatical context. With aliases we found a way to make that possible.

An alias is a word, phrase, sentence or similar in the form of a text literal that contains all the parameters of its function.
Aliases are interpreted by the lexer like normal DDP code, so any DDP keyword and any valid DDP name can be used in them.<br>

Every alias must be unique.
However, uniqueness is not to be taken too strictly here, because the parameters in aliases are type-sensitive.
That means two functions that both take one parameter can have the same aliases if the parameters are of different types.<br>
The best example is the Duden. It defines 5 Schreibe_X functions, one for each base type.
The aliases are all of the form `"Schreibe <p1>"`.
This is possible because p1 always has a different type, which is recognized when the function is called.

Every function needs at least one alias, but it is recommended to create several to do justice to different grammatical contexts.<br>
The general form of the alias declaration at the end of a function declaration looks like this:
```ddp
...
Und kann so benutzt werden:
	"Alias1",
	"Alias2" oder
	"Alias3"
```
The comma and 'oder' can be used interchangeably.
If a function has parameters, they must all be present in the alias in the form `<parameter-name>`.

### Alias negations

For functions that return a Wahrheitswert, there is special syntax to negate the function:
```ddp
Die öffentliche Funktion Ist_Text_Leer mit dem Parameter text vom Typ Text, gibt einen Wahrheitswert zurück, macht:
	Gib wahr, wenn die Länge von text gleich 0 ist, zurück.
Und kann so benutzt werden:
	"<text> <!nicht> leer ist"

Schreibe ("" leer ist). [wahr]
Schreibe ("" nicht leer ist). [falsch]

[ An alias negation corresponds to an additional "nicht" operator ]
Schreibe ( ("" nicht leer ist) gleich (nicht "" leer ist) ). [wahr]
```

If the token marked with the exclamation mark (`!`) in the alias is present, a `nicht` is inserted before the function call.

### Subsequent aliases

Aliases can also be defined outside of the function declaration, with the form:
```ddp
Der Alias "<alias>" steht für die Funktion <function name>.
```
This is useful to add aliases to functions from libraries (e.g. the Duden) for grammatical contexts that the creator of the function couldn't think of.<br>

**But be careful!!**<br>
In such late alias declarations, any parameters must have the same name as in the function.<br>
So if you want to add an alias to a function from the Duden, for example, you have to look up the names of the function's parameters.


## Forward declarations

A function can only be used after it was declared.
But sometimes two functions call each other. Then one of them must be used before it was declared.

For such cases there are forward declarations.
Instead of `macht:` and the function body, you write `wird später definiert` (will be defined later).
Later in the same module, the function is then defined with `Die Funktion <name> macht:`:

```ddp
Die Funktion Ist_Gerade mit dem Parameter n vom Typ Zahl, gibt einen Wahrheitswert zurück,
wird später definiert
und kann so benutzt werden:
	"<n> gerade ist"

Die Funktion Ist_Ungerade mit dem Parameter n vom Typ Zahl, gibt einen Wahrheitswert zurück, macht:
	Wenn n gleich 0 ist, gib falsch zurück.
	Gib (n minus 1) gerade ist zurück.
Und kann so benutzt werden:
	"<n> ungerade ist"

Die Funktion Ist_Gerade macht:
	Wenn n gleich 0 ist, gib wahr zurück.
	Gib (n minus 1) ungerade ist zurück.

Schreibe (7 ungerade ist) auf eine Zeile. [wahr]
```

In the definition, the parameters, the return type and the aliases are not repeated.
Every function that was declared with `wird später definiert` must also be defined.

## Public functions

With `öffentliche`, functions can also be used in other modules:

```ddp
Die öffentliche Funktion Hallo gibt nichts zurück, macht:
	Schreibe "Hello!" auf eine Zeile.
Und kann so benutzt werden:
	"Sag Hallo"
```

More about this in the article [Modules](/en/Programmierung/Module).

# Function calls

Functions are called exclusively through their aliases.<br>
A function call is technically an expression, so it can be used anywhere an expression can be.

When passing arguments, there are a few simple rules to keep in mind:
- Arguments that are only 1 token long (e.g. literals like "hi", 22 or wahr) can simply stand in the place of the parameter
- Arguments that are longer than 1 token (e.g. another function call or a list literal) must be enclosed in parentheses
- The `-` operator is an exception, since it looks nicer not to have to put `-22` in parentheses.

For functions that have the same alias with parameters of different types, the types of the passed arguments are recognized and the correct function is chosen.

## Examples

```ddp
Schreibe "Hi" auf eine Zeile.
Schreibe -22 auf eine Zeile.
Schreibe (einen Funktionsaufruf) auf eine Zeile.
```
