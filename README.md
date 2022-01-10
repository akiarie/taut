# taut
A simple boolean algebra processor/calculator.

Application built in Go with simple interpreter for boolean algebra.

# Functions

1. Given a truth table, reduce it to a simplified statement (or sequence of statements) using
   standard operators. By _simplified_ we mean a statement composed of the only of the unary
   operators ID, NOT and the binary operators OR, AND and XOR; and a statement with a minimised
   number of symbols.
2. Given a statement, simplify it if possible.
3. Given a statement and constraining conditions, return output.

# Representation of Tables
Truth tables are represented by giving a series of output values, with each set of values being of
the same length 2^m.

Thus the unary operators are input as
```
τ: ID[01], NOT[10]
```
corresponding to the truth table

| _A_ | ID _A_ | _!A_ |
| --- | ------ | ---- |
| 0   |      0 |    1 |
| 1   |      1 |    0 |

The standard binary operators are given by
```
τ: OR[0111], AND[0001], XOR[0110]
```
corresponding to the truth table

| _A_ | _B_ | _A_ + _B_ | _A_ * _B_ | _A_ ^ _B_ |
| --- | --- | --------- | --------- | --------- |
| 0   | 0   |      0    |       0   |       0   |
| 0   | 1   |      1    |       0   |       1   |
| 1   | 0   |      1    |       0   |       1   |
| 1   | 1   |      1    |       1   |       0   |

The following are interchangeable (in taut):

| Operator | Symbol |
| -------- | ------ |
| AND	   |	\*  |
| OR	   |	+   |
| XOR	   |	^   |
| NOT	   |	!   |

In general, a truth table is a string of binary digits of length 2^m, with the highest letter being
the LSB in the "counting sequence" made by the values of the variables denoted by the letters.

For an explanation of the implementation, see [here](METHOD.md).
