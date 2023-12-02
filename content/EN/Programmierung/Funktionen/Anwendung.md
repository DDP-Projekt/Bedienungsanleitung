+++
title = "Function declarations"
weight = 1
type = "article"
+++

# Function declarations

Each function must be declared before it can be used. In DDP there are only global functions, i.e. those that are not within an instruction block.
Each function needs a unique **name**, a list of **parameters** and their **types**, a **return type**, and one or more **aliases**.

Here is the general form of a function declaration (optional parts are enclosed in parentheses):
```ddp
Die Funktion <name> (mit den Parametern <parameters> vom Typ <parameter types>,) gibt <return type> zurück, macht:
	<function body>
Und kann so benutzt werden:
	"Alias mit Parameter <x> <y> und <z>" (,
	"<z> Noch ein Alias mit allen Parametern <y> <x>" oder
	...)
```

## The name

of a function must be unique and should describe the function well, but is otherwise not important and is not used directly in the code. It becomes more important with external functions, which will be described later.

## The parameters

of a function are optional. There can be between 0 and infinitely many. Each parameter requires a type.
When writing the parameters, the number (singular, plural) and the enumeration rules must be followed.

If there is only one parameter, write
```ddp
... mit dem Parameter x vom Typ T, ...
```
If there are two, you write
```ddp
... mit den Parametern x und y vom Typ T1 und T2, ...
```
And at N parameters
```ddp
... mit den Parametern a, b, ... und z vom Typ T1, T2, ... und T26, ...
```

## The return type

a function must be present, but can be `nichts`.
`nichts` is not a type, but a placeholder for functions that do not generate a value and only have side effects.
The result of a function that returns `nichts` can't be used because, well, it doesn't exist.

If a function returns a value, there must be a return statement at the end of the function body.
A return statement looks like this:
```ddp
Gib <Wert vom Rückgabetyp> zurück.
```
A return statement can appear anywhere in the function body, and exits the function with the result of the returned value.

To exit functions that return `nichts` early, you can use `Verlasse die Funktion`.

## Aliases

are the way functions are called.<br>
Since the goal of DDP is to be as good German as possible, it is difficult to correctly invoke functions in every grammatical context. With aliases we found a way to make this possible.

An alias is a word, sentence element, sentence or similar in the form of a text literal that contains all the parameters of its function.
Aliases are interpreted by the lexer like normal DDP code, so any DDP keyword and any valid DDP name can be used in it.<br>

Each alias must be unique.
In this case, however, uniqueness is not to be seen so narrowly, because the parameters in aliases are type-sensitive.
That is, two functions that both take a parameter can have the same aliases if the parameters are of different types.<br>
The best example is the Duden. It defines 5 Write_X functions, one for each base type.
The aliases are all of the form `"Schreibe <p1>"`.
This is possible because p1 always has a different type, which is recognized when the function is called.

Each function needs at least one alias, but it is recommended to create more than one to respect grammatical context.<br>
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

Aliases can also be defined outside of the function declaration, with the form:
```ddp
Der Alias "<Alias>" steht für die Funktion <function name>.
```
This is useful for adding aliases to functions from libraries (e.g. the Duden) for grammatical contexts that the creator of the function couldn't think of.<br>

**But be careful!!**<br>
In such late alias declarations, any parameters must have the same name as in the function.<br>
For example, if you want to add an alias to a function from the dictionary, you have to look up the name of the function's parameters.


# Function calls

Functions are only called through their aliases.<br>
A function call is technically an expression, so it can be used anywhere an expression can go.

When passing parameters, there are a few simple rules to keep in mind:
- Parameters that are only 1 token long (e.g. literals like "hi", 22 or true) can simply stand in the place of the parameter
- Parameters longer than 1 token (e.g. another function call or a list literal) must be enclosed in parentheses