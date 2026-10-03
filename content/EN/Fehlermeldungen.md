+++
title = "Errorcodes"
weight = 5
+++

# Errorcodes

When the compiler finds an error in the source code, it prints an error message.
An error message looks like this, for example:

```terminal
Typ Fehler (3001) in test.ddp (Z: 1, S: 16)

1 |  Die Zahl z ist "hallo".
  |                 ^^^^^^^

Ein Wert vom Typ Text kann keiner Variable vom Typ Zahl zugewiesen werden.
```

The first line contains:
- the kind of error (here `Typ Fehler`, a type error),
- the error code in parentheses (here `3001`),
- the file in which the error was found,
- the line (`Z` for Zeile) and column (`S` for Spalte) of the error.

Below that, the faulty line is shown. The `^` characters mark the exact position.
At the end there is a description of the error.

Some messages are only warnings (Warnung). With a warning, the program is still compiled.

The first digit of the error code tells you the kind of error:

| Error codes | Kind                                 | Meaning                                                         |
| ----------- | ------------------------------------ | --------------------------------------------------------------- |
| 0000        | Sonstiger Fehler (other error)       | Errors that don't fit into any other category                   |
| 1000 - 1999 | Syntax Fehler (syntax error)         | The source code is not structured correctly                     |
| 2000 - 2999 | Semantischer Fehler (semantic error) | The source code is structured correctly, but makes no sense     |
| 3000 - 3999 | Typ Fehler (type error)              | A value has the wrong type                                      |

## List of Errorcodes

### Other errors

| Code | Description |
| ---- | ----------- |
| 0000 | A module could not be included, for example because the file does not exist or two modules include each other |

### Syntax errors

| Code | Description |
| ---- | ----------- |
| 1000 | An unexpected word or character was found. Often a period or a comma is missing |
| 1001 | A literal was expected, but an expression was found |
| 1002 | A name was expected, for example a parameter name |
| 1003 | A type name was expected |
| 1004 | A capital letter was expected |
| 1005 | A literal is malformed, for example because of an invalid escape sequence |
| 1006 | The path of an include is malformed |
| 1007 | An alias is malformed |
| 1008 | The text is not valid UTF-8 |
| 1009 | The article or pronoun does not match the grammatical gender of the type, for example `Der Zahl z ist 1.` |
| 1010 | The name of an overloaded operator is invalid |

### Semantic errors

| Code | Description |
| ---- | ----------- |
| 2000 | The name is already in use |
| 2001 | The name has not been declared yet |
| 2002 | The number of parameter names and parameter types of a function is different |
| 2003 | No path to a .c, .lib, .a or .o file was given for an external function |
| 2004 | A function was not declared in the global scope |
| 2005 | A function is missing its return statement |
| 2006 | An alias is faulty, for example because parameters are missing |
| 2007 | The alias is already used by another function or combination |
| 2008 | An alias from an included module already exists |
| 2009 | A subsequent alias was not declared in the global scope |
| 2010 | The arguments of a subsequent alias don't match the function |
| 2011 | A return statement is outside of a function |
| 2012 | A function name was used instead of a variable name or vice versa |
| 2013 | A non-global declaration was marked as public |
| 2014 | A type was not declared in the global scope |
| 2015 | A field does not exist in this combination |
| 2016 | The word öffentlich (public) is missing or in the wrong place |
| 2017 | `Verlasse die Schleife` or `Fahre mit der Schleife fort` is outside of a loop |
| 2018 | A type is unknown or was not included |
| 2019 | An external function was unnecessarily also marked as externally visible |
| 2020 | The number of parameters does not match the overloaded operator |
| 2021 | The operator is already overloaded for these parameter types |
| 2022 | (Warning) A `...` placeholder was found |
| 2023 | A type cannot be defined as `Variable` |
| 2024 | A function was declared with `wird später definiert`, but never defined |
| 2025 | A function from another module was defined |
| 2026 | A function was already defined |
| 2027 | A generic function has no type parameter |
| 2028 | A generic combination has no type parameter |
| 2029 | A generic function was marked as externally visible |
| 2030 | A generic function was declared without a definition |
| 2031 | There were errors when creating a generic function for specific types |
| 2032 | A non-generic type was used with type parameters |
| 2033 | For an alias of a generic combination, not all type parameters could be found out |
| 2034 | A value cannot be assigned to a constant |

### Type errors

| Code | Description |
| ---- | ----------- |
| 3000 | An operator or a function got a value of the wrong type |
| 3001 | A value of the wrong type was assigned to a variable |
| 3002 | An indexing is faulty, for example because the index is not a number |
| 3003 | The values in a list literal don't all have the same type |
| 3004 | A type conversion with `als` is not possible |
| 3005 | A reference parameter did not get a variable |
| 3006 | A value cannot be passed as a reference, for example a letter in a text |
| 3007 | The condition is not a Wahrheitswert |
| 3008 | A value in a counting loop has the wrong type |
| 3009 | The returned value does not match the return type of the function |
| 3010 | The `von` operator was used with a value that is not a combination |
| 3011 | A private field was used from another module |
| 3012 | An overloaded operator returns `nichts` |
| 3013 | A type parameter was given as a reference |
| 3014 | The type of a generic field could not be found out |
| 3015 | A generic type could not be created with the given types |
| 3016 | A generic parameter of an external function is not a list or reference |
