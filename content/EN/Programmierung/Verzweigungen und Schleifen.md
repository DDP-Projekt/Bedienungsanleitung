+++
title = "Branches and Loops"
weight = 5
+++

# Branches
Branches are used to execute statements based on conditions.

## If branch
In an if branch, a statement block is only executed if the given condition evaluates to `wahr`; otherwise nothing is done and the following statements are executed.

### Syntax:
```ddp
Wenn <condition>, dann:
	<statements>.
```

### Example:
```ddp
Wenn 1 gleich 1 ist, dann:
	Schreibe den Text "Condition met!".
```

## If-Else branch
An if-else branch works like the if branch, with the difference that a `sonst` is present.
If the condition evaluates to `falsch`, the code in the `sonst` block is executed.

### Syntax:
```ddp
Wenn <condition>, dann:
	<statements>.
Sonst:
	<statements>.
```

### Example:
```ddp
Wenn 1 gleich 2 ist, dann:
	Schreibe den Text "Condition met!".
Sonst:
	Schreibe den Text "Condition not met!".
```

## If-ElseIf-Else branch
If-ElseIf-Else branches extend the if-else branch with any number of branches.
Any number of `Wenn aber` blocks can be appended after the `Wenn` block.
If the first condition in the `Wenn` block evaluates to `falsch`, the condition in the first `Wenn aber` block is checked; if that one is `falsch` too, the next one is checked and so on.
If all conditions are false, the `sonst` block is executed, if present.
As soon as one block has been executed, all following blocks are skipped.

### Syntax:
```ddp
Wenn <condition>, dann:
	<statements>
Wenn aber <2nd condition>, dann:
	<statements>.
Sonst:
	<statements>.
```

### Example:
```ddp
Wenn 1 gleich 2 ist, dann:
	Schreibe den Text "Condition met!".
Wenn aber 2 gleich 2 ist, dann:
	Schreibe den Text "2nd condition met!".
Sonst:
	Schreibe den Text "Condition not met!".
```

# Loops
Loops are used to execute code multiple times based on conditions.

## While loop
While loops are the simplest kind of loop.
If the condition evaluates to 'wahr', the code block is executed.
This is repeated as long as the condition evaluates to 'wahr'.
```ddp
Solange <condition>, mache:
	<statements>.
```

## Do-While loop
Do-while loops are very similar to while loops, with the only difference that the code block is executed at least once, and only then the condition starts being checked.
```ddp
Mache:
	<statements>.
Solange <condition>.
```

## Repetition
Repetitions are used to execute a code block several times.
They are a shortened version of counting loops, save text and improve the readability of the code
if you don't need the counter variable.
```ddp
Wiederhole:
	<statements>.
<count> Mal.
```

## Counting loop
Counting loops also allow code to be executed multiple times, while at the same time providing a counter that can be used otherwise.
For every counting loop, a counter must be named (a variable of type `Zahl`, `Byte` or `Kommazahl`) together with a start and end value.
Optionally, a step size to count with can be specified as well.

### Countup loop
First, the counter is named and initialized with the start value.
Then (as in every iteration) it is checked whether the value of the counter is less than or equal to the end value.
If this condition is met, the code block is executed, and afterwards the counter is incremented by 1.
This is repeated as long as the counter does not exceed the end value.
Inside the code block, the counter can be used like a normal local variable.

### Syntax
```ddp
Für jede Zahl <counter> von <start value> bis <end value>, mache:
	<statements>.
```

### Example
```ddp
Für jede Zahl i von 1 bis 100, mache:
	Schreibe die Zahl i.
```
### Countdown loop
A countdown loop works like the countup loop, except that a step size of -1 (or any other negative value) is specified. Because of that, the counter is not incremented by 1 at the end, but by the given step size (i.e. decremented).
Of course, the start value must be greater than the end value, otherwise the loop body is not executed at all.

### Syntax
```ddp
Für jede Zahl <counter> von <start value> bis <end value> mit Schrittgröße -1, mache:
	<statements>.
```

### Example
```ddp
Für jede Zahl i von 100 bis 1 mit Schrittgröße -1, mache:
	Schreibe die Zahl i.
```
### Custom step size
As mentioned above, any step size n can be specified. This causes the given value, instead of 1, to be added to the counter.

`mit Schrittgröße n` is optional. By default, the step size is +1.

### Syntax
```ddp
Für jede Zahl <counter> von <start value> bis <end value> mit Schrittgröße <n>, mache:
	<statements>.
```

### Example
```ddp
Für jede Zahl i von 1 bis 100 mit Schrittgröße 5, mache:
	<statements>.
```

## Iterating loops
It is also possible to use loops to go through every element of a list.
```ddp
Die Zahlen Liste liste ist eine leere Zahlen Liste.

Für jede Zahl element in liste, mache:
	<statements>.
```

In such loops you can also specify an index.
```ddp
Die Text Liste liste ist eine Liste, die aus "hi", "hello", "bye" besteht.

[Output:  hi 1 hello 2 bye 3]
Für jeden Text element mit Index i in liste, mache:
	Schreibe (' ' verkettet mit element verkettet mit ' ').
	Schreibe i.
```

In the same way, you can use a loop to go through every letter of a text:
```ddp
[Output: Hello]
Für jeden Buchstaben b in "Hello", mache:
	Schreibe b.
```

## Breaking/continuing loops

Loops can also be exited or skipped to the next iteration (in other languages this corresponds to the keywords `break` and `continue`).

### Break

```ddp
[Writes "1 2" to the console]
Für jede Zahl i von 1 bis 5, mache:
	Wenn i gleich 3 ist, dann:
    	Verlasse die Schleife. [break]
    Schreibe i.
    Schreibe ' '.
```

### Continue

```ddp
[Writes "1 2 4 5" to the console]
Für jede Zahl i von 1 bis 5, mache:
	Wenn i gleich 3 ist, dann:
    	Fahre mit der Schleife fort. [continue]
    Schreibe i.
    Schreibe ' '.
```

# Tip
Almost every branch and loop listed here can also be written on a single line
if only one statement needs to be executed.

A single statement, for example a function call or an assignment, can also end with `<count> Mal`.
It is then repeated as often as the count says:
```ddp
Schreibe den Text "Hi!" 5 Mal.
Erhöhe x um 2 3 Mal. [x is increased by 6]
```

## Examples
```ddp
Wenn 1 gleich 1 ist, Schreibe den Text "Condition met!".

Wenn 1 gleich 2 ist, Schreibe den Text "Condition met!".
Sonst Schreibe den Text "Condition not met".

Wenn 1 gleich 2 ist, Schreibe den Text "Condition met!".
Wenn aber 1 kleiner als 2 ist, Schreibe den Text "Second condition met!".
Sonst Schreibe den Text "Condition not met".


Solange i gleich 5 ist, Rufe eine Funktion auf, die i erhöht.

Schreibe den Text "Hi!" 5 Mal.


Für jede Zahl i von 1 bis 100, Schreibe die Zahl i.

Für jede Zahl i von 100 bis 1 mit Schrittgröße -1, Schreibe die Zahl i.
```
