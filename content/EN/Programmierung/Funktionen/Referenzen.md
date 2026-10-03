+++
title = "References"
weight = 2
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
Verändere "error". [compile error: foo expected a variable]
```

Any type name can be made a reference type by appending "Referenz" to the type name.
So `Zahlen Liste` becomes `Zahlen Listen Referenz`, `Wahrheitswert` becomes `Wahrheitswert Referenz`, etc.

Normal variables cannot be reference types, only function parameters.

## Type conversions as references

With [type aliases and type definitions](/en/Programmierung/Typ-Aliase-und-Typ-Definitionen), you can convert a variable with `als` and still pass it as a reference.
This only works if the type actually stays the same, for example with a type definition and its base type:

```ddp
Wir definieren eine Hausnummer als eine Zahl.

Die Funktion Erhöhe_Zahl mit dem Parameter z vom Typ Zahlen Referenz, gibt nichts zurück, macht:
	Erhöhe z um 1.
Und kann so benutzt werden:
	"Zähle <z> hoch"

Die Hausnummer h ist 41 als Hausnummer.
Zähle (h als Zahl) hoch.
Schreibe (h als Zahl) auf eine Zeile. [42]
```

Other conversions, for example from a Zahl to a Kommazahl, cannot be passed as a reference.