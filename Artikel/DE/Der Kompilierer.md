# Der Kompilierer
<to-do></to-do>

In diesem Artikel wird der Kompilierer der Deutschen Programmiersprache (kurz kddp) und seine Nutzung erklärt.

## Übersicht

KDDP ist ein Konsolen Programm, das über eine Befehlszeile (wie Powershell, bash, etc.) benutzt wird.
Die Nutzung des kddp Befehls sieht allgemein so aus:

```terminal
$ kddp <Befehl> <Eingabedatei> <Optionen>
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

### kompiliere

`$ kddp kompiliere <Eingabedatei> <Optionen>` ist der wichtigste Befehl, und wird am meisten genutzt.

Die Eingabedatei muss eine .ddp Datei sein.

Mit der -o Option kann der Name der Ausgabedatei angegeben werden, typischerweise .exe auf Windows. 
Es können aber auch die Dateierweiterungen .ll, .o, .obj, .s und .asm.
Wenn die Erweiterung .ll angegeben wird, wird llvm-ir ausgegeben. Das könnte interessant sein für Menschen die einfach daran interessiert sind, um Bugs zu finden oder Ähnliches.
Bei den Erweiterungen .s oder .asm wird assembler-Sprache ausgegeben.
Bei den Erweiterungen .o oder .obj werden Objektdateien ausgegeben, falls man andere Programme zu diesen linken möchte.

Mit den Optionen --gcc_optionen und --externe_gcc_optionen können weitere Argumente angegeben werden, die beim Linken an GCC übergeben werden.

Die Argumente aus --gcc_optionen werden beim finalen Link-Schritt übergeben und werden Benutzt wenn man z.B. zu externen Bibliotheken linken möchte (zu einer Grafikbibliothek o.ä.).

Die Argumente aus --externe_gcc_optionen werden nur benutzt falls es [externe Funktionen](/Bedienungsanleitung/DE/Programmierung/Funktionen/Externe%20Funktionen) gibt, die in .c Dateien definiert werden. Sollte das der Fall sein, wird jede angegebene .c Datei seperat mit GCC kompiliert, und dabei werden die Argumente aus --externe_gcc_optionen übergeben.
Das ist nützlich, wenn man Include-Verzeichnisse (wie die DDP-Runtime) oder C-Präprozessor Direktiven angeben muss/möchte.