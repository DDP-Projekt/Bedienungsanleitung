+++
title = "Branches and Loops"
weight = 5
+++

# Branches
Branches are used to execute statements based on conditions.

## If branch
With a If branch, a block of statements is executed only if the specified condition evaluates to `wahr`, otherwise nothing is done and subsequent statements are executed.

### Syntax:
```ddp
Wenn <condition>, dann:
	<instructions>.
```

### Example:
```ddp
Wenn 1 gleich 1 ist, dann:
	Schreibe den Text "condition true!".
```
    
## If-Else branch
A If-Else branch works like the one-way branch, except that there is an `Sonst`.
If the condition evaluates to `falsch`, the code in the `Sonst` block is executed.

### Syntax:
```ddp
Wenn <condition>, dann:
	<instructions>.
Sonst:
	<instructions>.
```

### Example:
```ddp
Wenn 1 gleich 2 ist, dann:
	Schreibe den Text "condition true!".
Sonst:
	Schreibe den Text "condition false!".
```

## If-ElseIf-Else branch
If-ElseIf-Else branches extend the If-Else branch by any number of branches.
Any number of 'Wenn aber' blocks can be added after the 'Wenn' block.
If the first condition in the `Wenn` block is `falsch` then the condition in the first `Wenn aber` block is checked, if it is `falsch` the next one is checked and so on.
If all conditions are false, the `Sonst` block, if any, is executed.
Once a block has been executed, all following blocks are skipped.

### Syntax:
```ddp
Wenn <condition>, dann:
	<instructions>
Wenn aber <2nd condition>, dann:
	<instructions>.
Sonst:
	<instructions>.
```

### Example:
```ddp
Wenn 1 gleich 2 ist, dann:
	Schreibe den Text "condition true!".
Wenn aber 2 gleich 2 ist, dann:
	Schreibe den Text "2nd condition true!".
Sonst:
	Schreibe den Text "condition false!".
```

# Loops
Loops are used to run code multiple times based on conditions.

## While loop
While loops are the simplest type of loop.
If the condition evaluates to true, the code block is executed.
This is repeated as long as the condition is 'true'.
```ddp
Solange <condition>, mache:
	<instructions>.
```

## Do-While loop
Do-While loops are very similar to While loops, with the only difference being that the block of code is executed at least once before the condition testing begins.
```ddp
Mache:
	<instructions>.
Solange <condition>.
```

## Repetition
Repetition are used to run a block of code multiple times.
They are an abbreviated version of counting loops, saving text and increasing code readability if you don't need the counter variable.
```ddp
Wiederhole:
	<instructions>.
<number> Mal.
```

## Counting loop
Counting loops also allow code to be run multiple times while giving a counter that can be used elsewhere.
For each counting loop, a counter must be named (a variable of type 'Zahl') and a start and end value.
Optionally, a step size can also be specified, which is used for counting.

### Countup loop
First the counter is named and initialized with the start value.
Then (as with every iteration) it is checked whether the value of the counter is less than or equal to the final value.
If this condition is met, the code block is executed and then the counter is increased by 1.
This is repeated as long as the counter does not exceed the final value.
In the code block, the counter can be used like a normal local variable.

### Syntax
```ddp
Für jede Zahl <counter> von <start value> bis <end value>, mache:
	<instructions>.
```

### Example
```ddp
Für jede Zahl i von 1 bis 100, mache:
	Schreibe die Zahl i.
```
### Countdown loop
A countdown loop works like the countup loop, except that a step size of -1 (or some negative value) is specified. This means that the counter is not increased (or decreased) by 1 at the end, but by the specified increment.
Of course, the start value must be greater than the end value, otherwise it becomes an endless loop.

### Syntax
```ddp
Für jede Zahl <counter> von <start value> bis <end value> mit Schrittgröße -1, mache:
	<instructions>.
```

### Example
```ddp
Für jede Zahl i von 100 bis 1 mit Schrittgröße -1, mache:
	Schreibe die Zahl i.
```
### Custom step size
As mentioned above, any step size n can also be specified. This ensures that the specified value is added to the counter instead of 1.

`mit Schrittgröße n` is optional. By default, the step size is +1.

### Syntax
```ddp
Für jede Zahl <counter> von <start value> bis <end value> mit Schrittgröße <n>, mache:
	<instructions>.
```

### Example
```ddp
Für jede Zahl i von 1 bis 100 mit Schrittgröße 5, mache:
	<instructions>.
```

## Iterating loops
It is also possible with loops to iterate through each element of a list.
```ddp
Die Zahlen Liste liste ist eine leere Zahlen Liste.

Für jede Zahl element in liste, mache:
	<instructions>.
```

# Tip
Almost any of the branches and loops listed here can also be written on a single line, if only one statement needs to be executed.

## Examples
```ddp
Wenn 1 gleich 1 ist, Schreibe den Text "Bedingung erfüllt!".

Wenn 1 gleich 2 ist, Schreibe den Text "Bedingung erfüllt!".
Sonst Schreibe den Text "Bedingung nicht erfüllt".

Wenn 1 gleich 2 ist, Schreibe den Text "Bedingung erfüllt!".
Wenn aber 1 kleiner als 2 ist, Schreibe den Text "Zweite Bedingung erfüllt!".
Sonst Schreibe den Text "Bedingung nicht erfüllt".


Solange i gleich 5 ist, Rufe eine Funktion auf, die i erhöht.

Schreibe den Text "Hi!" 5 Mal.


Für jede Zahl i von 1 bis 100, Schreibe die Zahl i.

Für jede Zahl i von 100 bis 1 mit Schrittgröße -1, Schreibe die Zahl i.
```