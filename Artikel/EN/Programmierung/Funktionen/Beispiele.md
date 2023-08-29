# Beispiele

## no parameters, without return value
```ddp
Die Funktion mache_nichts gibt nichts zurück, macht:
	[ nichts ]
Und kann so benutzt werden:
	"Mache nichts"
```

## no parameters, with return value
```ddp
Die Funktion PiDurchZwei gibt eine Kommazahl zurück, macht:
	Gib pi durch 2 zurück.
Und kann so benutzt werden:
	"pi durch zwei"
```

## one parameter, without return value
```ddp
Die Funktion Warten mit dem Parameter sekunden vom Typ Kommazahl, gibt nichts zurück, macht:
	Die Zahl start_zeit ist die Zeit seit Programmstart.
	Solange (die Zeit seit Programmstart) kleiner als start_zeit plus (sekunden mal 1000) ist, mache:
		[ nichts ]
Und kann so benutzt werden:
	"Warte <sekunden> Sekunden"
```

## one parameter, with return value
```ddp
Die Funktion Kehrwert mit dem Parameter wert vom Typ Kommazahl, gibt eine Kommazahl zurück, macht:
	Gib 1 durch wert zurück.
Und kann so benutzt werden:
	"Der Kehrwert von <wert>"
```

## two parameters, with return value
```ddp
Die Funktion durchschnitt mit den Parametern a und b vom Typ Kommazahl und Kommazahl, gibt eine Kommazahl zurück, macht:
	Gib (a plus b) durch 2 zurück.
Und kann so benutzt werden:
	"der Durchschnitt von <a> und <b>"
```

## multiple parameters, without return value
```ddp
Die Funktion Tausche_Zahl mit den Parametern a und b vom Typ Zahlen Referenz und Zahlen Referenz, gibt nichts zurück, macht:
	Die Zahl temp ist a.
	Speichere b in a.
	Speichere temp in b.
Und kann so benutzt werden:
	"tausche <a> und <b>"
```

## multiple parameters, with return value
```ddp
Die Funktion durchschnitt mit den Parametern a, b und c vom Typ Kommazahl, Kommazahl und Kommazahl, gibt eine Kommazahl zurück, macht:
	Gib (a plus b plus c) durch 3 zurück.
Und kann so benutzt werden:
	"der Durchschnitt von <a>, <b> und <c>"
```

## Extern functions

### DDP:
```ddp
Die Funktion Arkushyperbelsinus mit dem Parameter wert vom Typ Kommazahl, gibt eine Kommazahl zurück,
ist in "datei.c" definiert
und kann so benutzt werden:
	"Arkushyperbelsinus von <wert>"
```
### C:
```c
#include "DDP-Runtime/include/ddptypes.h"
#include <math.h>

ddpfloat Arkushyperbelsinus(ddpfloat wert) {
    return asinh(wert);
}
```
