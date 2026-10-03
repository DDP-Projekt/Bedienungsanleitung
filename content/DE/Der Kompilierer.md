+++
title = "Der Kompilierer"
weight = 4
+++

# Der Kompilierer
<to-do></to-do>

In diesem Artikel wird der Kompilierer der Deutschen Programmiersprache (kurz kddp) und seine Nutzung erklärt.

## Übersicht

KDDP ist ein Konsolen Programm, das über eine Befehlszeile (wie Powershell, bash, etc.) benutzt wird.
Die Nutzung des kddp Befehls sieht allgemein so aus:

```terminal
$ kddp <Befehl> [Optionen] [Argumente]
```

wobei [Argumente] meistens eine Eingabedatei ist, die vom Befehl verarbeitet wird.
Nicht alle Befehle brauchen eine Eingabedatei, und die möglichen Optionen sind immer optional.

## Hilfe

Um alle möglichen Befehle und Optionen zu sehen, kann kddp selbst benutzt werden:

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

Beispiel hilfe für den kompiliere Befehl:

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

## Detaillierte Erklärung zu einigen Optionen

<to-do></to-do>

Hier ist eine Liste aller Befehle und Optionen mit einer kurzen Erklärung.
Details folgen unten.

| Befehlsname | Befehlssyntax                          | Befehlsbeschreibung                                            | Befehlsoptionen | Optionsbeschreibungen |
| ----------- | -------------------------------------- | -------------------------------------------------------------- | --------------- | --------------------- |
| hilfe       | `hilfe <Befehl>`                       | Zeigt Nutzungsinformationen über den Befehl                    | - | - |
| kompiliere  | `kompiliere <Eingabedatei> <Optionen>` | Kompiliert die gegebene .ddp Datei zu einer ausführbaren Datei | `-o, --ausgabe <Ausgabepfad>`<hr>`-O, --optimierungs-stufe <Stufe>`<hr>`--nichts-loeschen`<hr>`--gcc-optionen`<hr>`--externe-gcc-optionen`<hr>`--gcc-executable <Pfad>`<hr>`--main <Pfad>`<hr>`--module-linken`<hr>`--list-defs-linken` | Optionaler Pfad der Ausgabedatei<hr>Menge und Art der Optimierungen, die angewandt werden (Standard: 1)<hr>Temporäre Dateien werden nicht gelöscht<hr>Benutzerdefinierte Optionen, die gcc übergeben werden<hr>Benutzerdefinierte Optionen, die gcc für jede externe .c Datei übergeben werden<hr>Pfad zur gcc executable, die genutzt werden soll<hr>Optionaler Pfad zur main.o Datei<hr>Ob alle Module in das Hauptmodul gelinkt werden sollen (Standard: wahr)<hr>Ob die eingebauten Listen Definitionen in das Hauptmodul gelinkt werden sollen (Standard: wahr) |
| starte      | `starte <Eingabedatei> <Optionen>`     | Kompiliert und führt die gegebene .ddp Datei aus               | `--gcc-optionen`<hr>`--externe-gcc-optionen` | Benutzerdefinierte Optionen, die gcc übergeben werden<hr>Benutzerdefinierte Optionen, die gcc für jede externe .c Datei übergeben werden |
| formatiere  | `formatiere <Eingabedatei> <Optionen>` | Formatiert die gegebene .ddp Datei                             | `--leerzeichen` | Leerzeichen statt Tabs für die Einrückung benutzen |
| parse       | `parse <Eingabedatei> <Optionen>`      | Parse die Eingabedatei zu einem Abstrakten Syntaxbaum          | `-o, --ausgabe <Ausgabepfad>` | Optionaler Pfad der Ausgabedatei |
| update      | `update <Optionen>`                    | Aktualisiert kddp (siehe [Updates](/Einstieg/Updates))          | `--jetzt`<hr>`--vergleiche-version`<hr>`--pre-release` | Aktualisiert sofort ohne zu fragen<hr>Vergleicht neue Version mit der installierten<hr>Aktualisiert auf eine Vorabversion |
| version     | `version <Optionen>`                   | Zeige Informationen zu dieser DDP Version                      | `--go-build-info` | Zeige Go build Informationen |

Die globalen Optionen `-w, --wortreich` (gibt wortreiche Informationen während des Befehls aus) und `-t, --zeitmessen` (gibt zusammen mit `--wortreich` Zeitmessungen aus) können bei jedem Befehl angegeben werden.

### kompiliere

`$ kddp kompiliere <Eingabedatei> <Optionen>` ist der wichtigste Befehl, und wird am meisten genutzt.

Die Eingabedatei muss eine .ddp Datei sein.

Mit der -o Option kann der Name der Ausgabedatei angegeben werden, typischerweise .exe auf Windows. 
Es können aber auch die Dateierweiterungen .ll, .o, .obj, .s und .asm.
Wenn die Erweiterung .ll angegeben wird, wird llvm-ir ausgegeben. Das könnte interessant sein für Menschen die einfach daran interessiert sind, um Bugs zu finden oder Ähnliches.
Bei den Erweiterungen .s oder .asm wird assembler-Sprache ausgegeben.
Bei den Erweiterungen .o oder .obj werden Objektdateien ausgegeben, falls man andere Programme zu diesen linken möchte.

Mit den Optionen --gcc-optionen und --externe-gcc-optionen können weitere Argumente angegeben werden, die beim Linken an GCC übergeben werden.

Die Argumente aus --gcc-optionen werden beim finalen Link-Schritt übergeben und werden Benutzt wenn man z.B. zu externen Bibliotheken linken möchte (zu einer Grafikbibliothek o.ä.).

Die Argumente aus --externe-gcc-optionen werden nur benutzt falls es [externe Funktionen](/Programmierung/Funktionen/Externe-Funktionen/) gibt, die in .c Dateien definiert werden. Sollte das der Fall sein, wird jede angegebene .c Datei separat mit GCC kompiliert, und dabei werden die Argumente aus --externe-gcc-optionen übergeben.
Das ist nützlich, wenn man Include-Verzeichnisse (wie die DDP-Runtime) oder C-Präprozessor Direktiven angeben muss/möchte.
