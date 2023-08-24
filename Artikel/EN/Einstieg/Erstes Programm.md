# First program
Now that we have successfully installed the programming language, we can start writing some programs in it.

## Hello World
With the command `Schreibe den Text` from the standard library followed by a text you can display it in the console.<br>
This is what a typical "Hello World" program looks like:
```ddp
Binde "Duden/Ausgabe" ein.

Schreibe den Text "Hallo Welt!".
```
<h6 style="margin-top: 0; margin-left: 10px">File: HelloWorld.ddp </h6>


This must now be saved as a `.ddp` file.

Finally, to run the program, use these commands:
```terminal
$ kddp kompiliere HalloWelt.ddp
$ ./HalloWelt
```

Alternatively, you can use the `kddp starte` command:
```terminal
$ kddp starte HalloWelt.ddp
```