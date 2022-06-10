### Mit dem Syntax `binde <Text> ein` kannst du mehrere Dateien in einem Projekt benutzen.
#### Dabei ist das Verzeichnis indem sich die ausgeführte Datei befindet, das Stammverzeichnis.

## Beispiel:

#### Beispiel/test1.ddp
```ddp
die Funktion foo() macht:
  schreibeZeile("Hallo Welt!").

schreibeZeile("test").
```

#### Beispiel/test2.ddp
```ddp
binde "test1" ein.

foo().
schreibeZeile("Tschüss Welt!").
lese().
```

## `> test2 wird ausgeführt`

### Ausgabe:
```ddp
test
Hallo Welt!
Tschüss Welt!
```

## Man kann auch Dateien die sich in Über-/Unterverzeichnissen befinden einbinden:
`binde "../test" ein`\
`binde "Ordner/test" ein`