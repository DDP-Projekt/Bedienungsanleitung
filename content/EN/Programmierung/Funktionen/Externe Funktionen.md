+++
title = "External functions"
weight = 3
+++

# External functions

An external function is a function whose body is not defined in the DDP source code but outside of it (e.g. in other languages or static libraries).

The general form looks like this:
```ddp
Die Funktion <name> (mit den Parametern x, y und z vom Typ T1, T2 und T3,) gibt <return type> zurück,
ist in "<file path>" definiert
und kann so benutzt werden:
	"alias with parameters <x> <y> and <z>" (,
	"<z> another alias with all parameters <y> <x>" oder
	...)
```

You can see that external functions are identical to normal functions, except that their body is not DDP source code but a reference to a file in which the function body is defined.
Possible files are .c, .o, .a and .lib files.

The body should be implemented in C.
Although it is possible to define it in other languages, this is not recommended, since the DDP runtime is written in C, and therefore everything uses C calling conventions and C data types and is generally quite tightly coupled to the runtime.

If the given file is a .c file, it is compiled to an .o file with GCC and linked into the final output file.
.o, .a or .lib files are linked directly into the final output file.

## Use cases

Pure DDP source code couldn't do much, not even write to the console.
That's why external functions are extremely important, since they form interfaces to the operating system and many libraries.
A good example of this is the Duden. Many functions in it are external and written in C. The Duden is a good starting point to learn how to use external functions and to look at examples in C of how they interact with the DDP runtime.

# Externally visible functions

The DDP compiler uses a technique called [name mangling](https://en.wikipedia.org/wiki/Name_mangling). This means that the names of functions and variables
in the source code are not the same as in the resulting binary.

So if you want to use a non-external DDP function from C code, you have to turn off name mangling:
```ddp
[test.ddp]
Die Funktion ddp_funktion gibt nichts zurück, ist extern sichtbar, macht:
    Schreibe "I am a DDP function" auf eine Zeile.
Und kann so benutzt werden:
    "ddp_funktion"
```

```c
// test.c
extern void ddp_funktion(void);

int main(void) {
    ddp_funktion();
    return 0;
}
```
