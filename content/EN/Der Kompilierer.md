+++
title = "The Compiler"
weight = 4
+++

# The Compiler

This article explains the compiler of the German programming language (kddp for short) and how to use it.

## Overview

KDDP is a console program that is used from a command line (like Powershell, bash, etc.).
In general, using the kddp command looks like this:

```terminal
$ kddp <command> [options] [arguments]
```

where [arguments] is usually an input file that is processed by the command.
Not every command needs an input file, and the options are always optional.

## Help

To see all available commands and options, kddp itself can be used:

```terminal
$ kddp
Der Kompilierer der deutschen Programmiersprache (DDP)

Nutzung:
  kddp <Befehl> [Optionen] [Argumente]

Verfügbare Befehle:
  formatiere     Formatiert eine .ddp Datei
  hilfe          Zeigt Informationen zu einem Befehl
  kompiliere     Kompiliert eine .ddp Datei
  parse          Parst eine .ddp Datei in einen Abstrakten Syntaxbaum
  starte         Kompiliert und führt die angegebene .ddp Datei aus
  update         Aktualisiert kddp
  version        Zeigt Versionsinformationen des Kompilierers

Optionen:
  -h, --hilfe        Zeigt Informationen zum Befehl
  -v, --version      Zeigt die Version des Kompilierers
  -w, --wortreich    Gibt wortreiche Informationen aus
  -t, --zeitmessen   Gibt in Kombination mit --wortreich Zeitmessungen aus

Probiere "kddp hilfe <Befehl>" oder "kddp <Befehl> [-h | --hilfe]" für mehr Informationen zu einem Befehl.
```

Example help for the kompiliere command:

```terminal
$ kddp hilfe kompiliere
Kompiliert eine .ddp Datei in eine ausführbare, llvm oder objekt Datei.

Nutzung:
  kddp kompiliere [-o Ausgabe-Datei [--main main.o] [--gcc-flags GCC-Flags] [--extern-gcc-flags Externe-GCC-Flags] [--nodeletes] [--verbose] [--link-modules] [--link-list-defs] [--gcc-executable Pfad-zu-GCC>] <Datei>

Optionen:
  -o, --ausgabe string                Optionaler Pfad der Ausgabedatei (.exe, .ll, .o, .obj, .s, .asm).
      --externe-gcc-optionen string   Benutzerdefinierte Optionen, die gcc für jede externe .c Datei übergeben werden
      --gcc-executable string         Pfad zur gcc executable, die genutzt werden soll (default "gcc")
      --gcc-optionen string           Benutzerdefinierte Optionen, die gcc übergeben werden
  -h, --hilfe                         Zeigt Informationen zum Befehl
      --list-defs-linken              Ob die eingebauten Listen Definitionen in das Hauptmodul gelinkt werden sollen (default true)
      --main string                   Optionaler Pfad zur main.o Datei
      --module-linken                 Ob alle Module in das Hauptmodul gelinkt werden sollen (default true)
      --nichts-loeschen               Keine temporären Dateien löschen
  -O, --optimierungs-stufe uint       Menge und Art der Optimierungen, die angewandt werden (default 1)

Globale Optionen:
  -w, --wortreich    Gibt wortreiche Informationen aus
  -t, --zeitmessen   Gibt in Kombination mit --wortreich Zeitmessungen aus
```

## Commands and options

Here is a list of all commands and options with a short explanation.
Details follow below.

| Command name | Command syntax                         | Description                                          | Options | Option descriptions |
| ------------ | -------------------------------------- | ---------------------------------------------------- | ------- | ------------------- |
| hilfe        | `hilfe <command>`                       | Shows usage information about the command            | - | - |
| kompiliere   | `kompiliere <input file> <options>` | Compiles the given .ddp file into an executable      | `-o, --ausgabe <output path>`<hr>`-O, --optimierungs-stufe <level>`<hr>`--nichts-loeschen`<hr>`--gcc-optionen`<hr>`--externe-gcc-optionen`<hr>`--gcc-executable <path>`<hr>`--main <path>`<hr>`--module-linken`<hr>`--list-defs-linken` | Optional path of the output file<hr>Amount and kind of optimizations that are applied (default: 1)<hr>Temporary files are not deleted<hr>Custom options that are passed to gcc<hr>Custom options that are passed to gcc for every external .c file<hr>Path to the gcc executable that should be used<hr>Optional path to the main.o file<hr>Whether all modules should be linked into the main module (default: true)<hr>Whether the built-in list definitions should be linked into the main module (default: true) |
| starte       | `starte <input file> <options>`     | Compiles and runs the given .ddp file                | `--gcc-optionen`<hr>`--externe-gcc-optionen` | Custom options that are passed to gcc<hr>Custom options that are passed to gcc for every external .c file |
| formatiere   | `formatiere <input file> <options>` | Formats the given .ddp file                          | `--leerzeichen` | Use spaces instead of tabs for indentation |
| parse        | `parse <input file> <options>`      | Parses the input file into an abstract syntax tree   | `-o, --ausgabe <output path>` | Optional path of the output file |
| update       | `update <options>`                    | Updates kddp (see [Updates](/en/Einstieg/Updates))   | `--jetzt`<hr>`--vergleiche-version`<hr>`--pre-release` | Updates immediately without asking<hr>Compares the new version with the installed one<hr>Updates to a pre-release version |
| version      | `version <options>`                   | Shows information about this DDP version             | `--go-build-info` | Shows Go build information |

The global options `-w, --wortreich` (prints verbose information during the command) and `-t, --zeitmessen` (prints time measurements together with `--wortreich`) can be passed to every command.

### kompiliere

`$ kddp kompiliere <input file> <options>` is the most important and most used command.

The input file must be a .ddp file.

The -o option specifies the name of the output file, typically .exe on Windows.
The file extensions .ll, .o, .obj, .s and .asm can be used as well.
If the extension .ll is given, llvm-ir is output. This might be interesting for people who are simply curious, or to find bugs and the like.
With the extensions .s or .asm, assembly is output.
With the extensions .o or .obj, object files are output, in case you want to link other programs against them.

The options --gcc-optionen and --externe-gcc-optionen can be used to pass further arguments to GCC.

The arguments from --gcc-optionen are passed in the final link step and are used, for example, if you want to link against external libraries (a graphics library or similar).

The arguments from --externe-gcc-optionen are only used if there are [external functions](/en/Programmierung/Funktionen/Externe-Funktionen/) that are defined in .c files. If that is the case, every given .c file is compiled separately with GCC, and the arguments from --externe-gcc-optionen are passed along.
This is useful if you need/want to specify include directories (like the DDP runtime) or C preprocessor directives.

The option `-O` (or `--optimierungs-stufe`) sets how much the program is optimized:
- `0`: no optimizations
- `1`: only the LLVM optimizations (default)
- `2`: all optimizations

With `--nichts-loeschen`, the temporary files that are created during compilation are not deleted.
This is mostly useful for debugging.

### starte

`$ kddp starte <input file> <options>` compiles the input file into a temporary directory and then runs the program right away.
The executable is deleted afterwards.

All arguments after the input file are passed to the program:

```terminal
$ kddp starte Program.ddp hello 42
```

Here the program gets the command line arguments `hello` and `42`.
How to use them in the program is described in the article about the Duden module [Befehlszeile](/en/Programmierung/Standardbibliothek/Befehlszeile).

### formatiere

`$ kddp formatiere <input file>` formats a .ddp file, for example its indentation.
The file is overwritten directly.

By default, tabs are used for indentation. With the option `--leerzeichen`, spaces are used instead.

### parse

`$ kddp parse <input file>` reads the input file and prints the abstract syntax tree.
This is mostly useful for developing the compiler.
With the option `-o`, the syntax tree is written into a file instead of being shown in the console.

### update

`$ kddp update` updates kddp to the newest version.
More about this in the article [Updates](/en/Einstieg/Updates).

### version

`$ kddp version` shows the version of kddp.
With `--wortreich`, the versions of GCC, LLVM and Go are shown as well.
The option `--go-build-info` shows even more information about how kddp was built.

### Global options

The option `--wortreich` (or `-w`) can be used with every command.
kddp then prints what it is currently doing.
Together with `--zeitmessen` (or `-t`), it also shows how long the individual steps took.
