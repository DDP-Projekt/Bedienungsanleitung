# External functions

An external function is a function whose body is not defined in the DDP source code but outside (e.g. in other languages or static libraries).

The general form looks like this:
```ddp
Die Funktion <Name> (mit den Parametern x, y und z vom Typ T1, T2 und T3,) gibt <return type> zurück,
ist in "<file path>" definiert
und kann so benutzt werden:
	"Alias with parameters <x> <y> and <z>" (,
	"<z> another Alias with all parameters <y> <x>" oder
	...)
```

You can see that external functions are identical to normal functions, but their body is not DDP source code, but a reference to a file in which the function body is then defined.
Possible files are .c, .o, .a and .lib files.

The body should be implemented in C.
Although it is possible to define it in other languages, this is not recommended since the DDP runtime is written in C, and thus uses all C calling conventions and C data types, and is generally quite tightly coupled to the runtime.

If the given file is a .c file, it will be compiled to an .o file with GCC and linked into the final output file.
.o, .a or .lib files are linked directly into the final output file.

## Use cases

Pure DDP source code couldn't do much, not even write to the console.
Therefore, external functions are extremely important, since they form interfaces to the operating system and many libraries.
A good example of this is the Duden. Many functions in it are external and written in C. The Duden is a good entry point to learn the usefulness of external functions and look at examples in C of how they interact with the DDP runtime.