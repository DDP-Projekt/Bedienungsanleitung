# Der Kompilierer
<to-do></to-do>

In diesem Artikel wird der Kompilierer der Deutschen Programmiersprache (kurz kddp) und seine Nutzung erklärt.

## Übersicht

KDDP ist ein Konsolen Programm, das über eine Befehlszeile (wie Powershell, bash, etc.) benutzt wird.
Die Nutzung des kddp Befehls sieht allgemein so aus:

```
$kddp <Befehl> <Eingabedatei> <Optionen>
```

Nicht alle Befehle brauchen eine Eingabedatei, und die möglichen Optionen sind immer optional.

Hier ist eine Liste aller Befehle und Optionen mit einer kurzen Erklärung.
Details folgen unten.

| Befehlsname | Befehlssyntax | Befehlsbeschreibung | Befehlsoptionen | Optionsbeschreibungen |
|-------------|---------------|---------------------|-----------------|-----------------------|
| hilfe| `hilfe <Befehl>`| Zeigt Nutzungsinformationen über den Befehl| - | - |
| kompiliere  | `kompiliere <Eingabedatei> <Optionen>` | Kompiliert die gegebene .ddp Datei zu einer ausführbaren Datei | `-o <Ausgabepfad>`<hr>`--wortreich`<hr>`--nichts_loeschen`<hr>`--gcc_optionen`<hr>`--externe_gcc_optionen` | Optionaler Pfad der Ausgabedatei<hr>Gibt wortreiche Informationen während des Befehls<hr>Temporäre Dateien werden nicht gelöscht<hr>Benutzerdefinierte Optionen, die gcc übergeben werden<hr>Benutzerdefinierte Optionen, die gcc für jede externe .c Datei übergeben werden |
| parse       | `parse <Eingabedatei> <Optionen>`      | Parse die Eingabedatei zu einem Abstrakten Syntaxbaum          | `-o <Ausgabepfad>` | Optionaler Pfad der Ausgabedatei |
| version     | `version <Optionen>` | Zeige informationen zu dieser DDP Version | `--wortreich`<hr>`--go_build_info` | Zeige wortreiche Informationen<hr>Zeige Go build Informationen |
| starte | `starte <Eingabedatei> <Optionen>` | Kompiliert und führt die gegebene .ddp Datei aus | `--wortreich`<hr>`--gcc_optionen`<hr>`--externe_gcc_optionen` | Gibt wortreiche Informationen während des Befehls<hr>Benutzerdefinierte Optionen, die gcc übergeben werden<hr>Benutzerdefinierte Optionen, die gcc für jede externe .c Datei übergeben werden |

### hilfe

Der `$kddp hilfe <Befehl>` Befehl zeigt einfach die Informationen aus der Tabelle über den gegebenen Befehl, oder die gesamte Tabelle, falls kein Befehl angegeben wurde.

### kompiliere

`$kddp kompiliere <Eingabedatei> <Optionen>` ist der wichtigste Befehl, und wird am meisten genutzt.

Die Eingabedatei muss eine .ddp Datei sein.
