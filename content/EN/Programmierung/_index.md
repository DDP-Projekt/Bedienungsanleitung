+++
title = "Programming"
weight = 2
+++

# Programming

Many, if not all, of the features described in the following articles should already be familiar to anyone who has programmed before.
Nevertheless, it is worth reading everything carefully, if only for the syntax, since DDP often changes some details to do justice to the German language.

In order to get a rough overview of the language, it is probably sufficient to mainly look at the tables and code examples. However, if you want to know the details and don't want to be surprised by strange syntax, you should read all the articles completely.

The following articles describe the individual features and properties of the German programming language and how to use them.
Programming terminology is used in these articles, which is to be explained here.

## Expressions

An expression is a piece of source code that produces a value.
The mathematical expression `2 * pi * radius` produces the circumference of a circle.
In DDP this expression would look like this: `2 mal pi mal radius`.
Expressions can take multiple arguments of different types and produce a value of the same or different type.
The different data types are described in the article [data types](/en/Programmierung/Datentypen).

## Statements

A statement is a piece of source code that does something.
Statements usually end with a period (.).
For example, the statement `Schreibe den Text "Hallo".` outputs "Hallo" to the console.
Expressions followed by a period are also statements. Your result will be discarded.
For example, `1 plus 1.` is evaluated, but the result is never used.

A DDP program consists of any number of statements that are executed one after the other (in the source code from left to right and from top to bottom).

## Statement-block

A statement block is a sequence of statements that are executed one after the other.
It starts with a colon (:) and a newline, and any associated statements must be at least one indentation (one tab character or 4 spaces) deeper than the colon.
```ddp
:
	Schreibe den Text "I'm in the statement block!".
	Schreibe den Text "me too!".
Schreibe den Text "I'm not in the statement block.".
```
The statement block ends when the next character is indented the same or less than the colon.

statement blocks are often separated into [branches](/en/Programmierung/Verzweigungen-und-Schleifen#verzweigungen), [loops](/en/Programmierung/Verzweigungen-und-Schleifen#schleifen) or [functions](/en/Programmierung/Funktionen) used.
More on this in the following articles.

## Comments

Comments can be used to explain the source code of a DDP program and make it easier to understand.
Comments are written between square brackets ([]) and are only there to explain the source code, they don't change what it does.

```ddp
Schreibe den Text "test". [writes the text "test" to the console]
```

Comments can also appear in the middle of statements and expressions.
```ddp
Wenn x [wahr ist], Schreibe den Text "x ist wahr".
```

Also, comments can span multiple lines.
```ddp
[
	Task: This function writes "Hello world!" into the console.
	Parameters: no parameters.
	Return: no return.
]
Die Funktion hi gibt nichts zurück, macht:
	Schreibe den Text "Hallo Welt!" auf eine Zeile.
und kann so benutzt werden:
	"Schreibe Hallo Welt"
```

## literals

A literal is a notation in the source code that represents a value (constant in DDP).
A list of literals for the corresponding data types can be found in the [data types](/en/Programmierung/Datentypen) article.
