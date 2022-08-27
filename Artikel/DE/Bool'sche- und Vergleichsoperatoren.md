# Bool'sche Operatoren

Mithilfe von Bool'schen Operatoren können komplexe Bedingungen ausgedrückt und zusammengefasst werden und z.B. für Schleifen verwendet werden.

| Operator | Beschreibung | Beispiel | Ergebnis |
| -------- | ------------ | -------- | -------- |
| und | Wahr, wenn beide Argumente wahr sind. | `wahr und wahr`<br>`wahr und falsch`<br>`falsch und wahr`<br>`falsch und falsch` | `wahr`<br>`falsch`<br>`falsch`<br>`falsch` |
| oder | Wahr, wenn eines der Beiden Argumente wahr ist. | `wahr oder wahr`<br>`wahr oder falsch`<br>`falsch oder wahr`<br>`falsch oder falsch` | `wahr`<br>`wahr`<br>`wahr`<br>`falsch` |
| nicht | Der Wert des Arguments wird umgekehrt. | `nicht wahr` <br>`nicht falsch` | `falsch`<br>`wahr` |

# Vergleichsoperatoren

| Operator | Beschreibung | Beispiel | Ergebnis |
| -------- | ------------ | -------- | -------- |
| gleich | Wahr, wenn beide Argumente den gleichen Wert haben. | `1 gleich 1`<br>`1 gleich 2` | `wahr`<br>`falsch` |
| ungleich | Wahr, wenn die Beiden Argumente verschiedene Werte haben. | `1 ungleich 1`<br>`1 ungleich 2` | `wahr`<br>`falsch` |
| kleiner als | Wahr, wenn das Linke Argument einen kleineren Wert als das Rechte hat. | `5 kleiner als 10`<br>`30 kleiner als 15` | `wahr`<br>`falsch` |
| größer als | Wahr, wenn das Linke Argument einen größeren Wert als das Rechte hat. | `7 größer als 3`<br>`5 größer als 8` | `wahr`<br>`falsch` |
| kleiner als, oder | Wahr, wenn das Linke Argument einen kleineren oder denselben Wert wie das Rechte hat. | `5 kleiner als, oder 10`<br>`30 kleiner als, oder 15`<br>`5 kleiner als, oder 5` | `wahr`<br>`falsch`<br>`wahr` |
| größer als, oder | Wahr, wenn das Linke Argument einen größeren oder denselben Wert wie das Rechte hat. | `7 größer als, oder 3`<br>`5 größer als, oder 8`<br>`5 größer als, oder 5` | `wahr`<br>`falsch`<br>`wahr` |