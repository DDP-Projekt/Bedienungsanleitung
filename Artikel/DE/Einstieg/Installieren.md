# Installieren

<to-do></to-do>

In diesem Artikel wird beschrieben wie man den Kompilierer (kddp) installiert.

## 1. Download

Gehe zum [Github-Repo des DDP Kompilierers](https://github.com/DDP-Projekt/Kompilierer) und lade den neusten [Release](https://github.com/DDP-Projekt/Kompilierer/releases) für dein System herunter.
Für Windows ist das DDP-&lt;Version&gt;-&lt;Betriebssystem&gt;-&lt;CPU-Architektur&gt;.zip und für Linux DDP-&lt;Version&gt;-&lt;Betriebssystem&gt;-&lt;CPU-Architektur&gt;.tar.gz.
Für Windows Benutzer, die gcc bereits installiert haben gibt es auch -no-mingw Varianten.

Beispielhafte Dateinamen:
- DDP-0.0.1-windows-amd64.zip
- DDP-0.0.1-windows-amd64-no-mingw.zip
- DDP-0.0.1-linux-amd64.tar.gz

## 2. Entpacken

Entpacke das heruntergeladene Archiv an die Stelle an die du DDP installieren möchtest.
Diese Stelle sollte nach der Installation nicht mehr verändert werden!

Auf Windows einfach auf das Archiv rechts-klicken und "Alle Extrahieren" wählen.
Auf Linux den `$tar` Befehl verwenden.

Nach dem extrahieren solltest du den enstprechend entpackten Ordner sehen. Es kann sein, dass dieser geschachtelt ist (also DDP-0.0.1-windows-amd64/DDP-0.0.1-windows-amd64). Falls das passiert ist verschiebe den Ordner VOR DER INSTALLATION an die richtige Stelle.

## 3. Installieren

Gehe in das DDP Verzeichnis das gerade erstellt wurde und führe die Datei `ddp-setup.exe` (bzw. `ddp-setup` auf Linux) aus. 
Es sollte sich ein Konsolen Fenster öffnen, das dich durch die Installation führt.

Im Moment ist das Setup nur auf Englisch um es allen zugänglich zu machen, also nicht wundern.

Das Setup wird dir Fragen stellen (blau) und dir Warnungen (gelb) und Fehler (rot) melden.
Die Fragen sind ja/nein fragen, die mit y/n (yes/no) benatwortet werden können.
Warnungen sollten gelesen werden, können aber unter Umständen ignoriert werden.
Sollte ein Fehler auftreten ist es sehr wahrscheinlich, dass die gesamte Installation fehlgeschlagen ist.

Das Setup ist so ausgelegt, dass du Standardmäßig jede Frage mit ja (y) beantworten solltest um sicherzustellen, dass alles möglichst gut funktioniert.

Das Setup wird
* auf Windows falls nötig Mingw64 installieren
* die kddp Version mit der (installierten) GCC Version abgleichen und
* falls nötig die DDP-runtime/stdlib neu kompilieren
* die Umgebungsvariable DDPPATH setzen (diese MUSS gesetzt sein, sonst funktioniert der Kompilierer nicht)
* kddp deinem PATH hinzufügen
* vscode-ddp (die DDP-Erweiterung für Visual Studio Code) installieren, falls VSCode installiert ist
* am Ende nicht mehr benötigte Dateien löschen um Platz zu sparen

Eine fertige DDP Installation kann zwischen 20mb und etwa 1,1Gb groß sein, je nach Betriebssystem und je nachdem ob mingw64 extra installiert werden musste oder nicht.

## 4. Testen

Um die Installation zu testen, siehe [Erstes Programm](/Bedienungsanleitung/DE/Einstieg/Erstes%20Programm)