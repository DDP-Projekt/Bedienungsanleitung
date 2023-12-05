+++
title = "References"
weight = 2
type = "article"
+++

# References

Normally, function arguments are passed by value.
This means that if you call e.g. `Schreibe den Text t.`, a copy of the variable t is made and passed to the function. So the function cannot change `t`.

In DDP, functions can also specify so-called *references* as parameter types.
It looks like this:

```ddp
Die Funktion foo mit dem Parameter param vom Typ Text Referenz, gibt nichts zurück, macht:
    Speichere "new!" in param.
Und kann so benutzt werden:
    "Verändere <t>"
```

And you can call the function like this:

```ddp
Der Text t ist "old".
Schreibe t. [writes "old" to the Console]
Verändere t.
Schreibe t. [writes "new!" to the Console]
```

As you can see, reference parameters take variables as arguments.
If you passed them a value, e.g. a literal, there would be an error:

```ddp
Verändere "Fehler". [compile error: foo expected a variable]
```

Any type name can be made a reference type by appending "Referenz" to the type name.
So `Zahlen Liste` becomes `Zahlen Listen Referenz`, `Wahrheitswert` becomes `Wahrheitswert Referenz`, etc.

Normal variables cannot be reference types, only function parameters.