# taut's methodology

Since [I am learning about compilers from _The Dragon Book_](https://github.com/akiarie/dragon-book)
it seemed appropriate to implement taut in the form of a simple interpreter.

We begin by specifying the grammar for inputting truth tables. `||` is the concatenation operator.
```
table -> op | table , op
op    -> id||[||bits||]
id    -> alpha | id||alpha
alpha -> A | B | ... | Z
bits  -> bit | bits||bit
bit   -> 0 | 1
```
We leave the question of the semantic correctness of the operators inserted, such as the bitstreams
having lengths of 2^m for some m ≥ 0.
