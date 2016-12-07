# taut
A simple boolean algebra processor/calculator.

Application built in Go with simple interpreter for boolean algebra.

# Functions

1. Given a truth table, reduce it to (unsimplified) statement (or sequence of statements).
2. Given statement, simplify if possible.
3. Given a statement and condition, return output.

# Representation of Tables
Statements are represented using truth tables with one output, i.e.:

    A, B, ... -> Y[ijkl...], X[mnop...]
where *A,B,...* are input boolean values; *Y,X...* are the output.
Note that the length of each of the values must be equal.
All space characters will be ignored.
Don't care will be implemented in a later version, most probably as _.

Below are the truth tables for the logical **AND** and **OR** operations
and the associated representation in taut syntax.

    Table 1: AND/OR
    A	B	A+B	A*B
    0	0	 0	 0
    0	1	 1	 0
    1	0	 1	 0
    1	1	 1	 1

*A+B* would be represented like

    A, B -> OR[0111]
while A\*B would be

    A, B -> AND[0001]
    
# Representation of Statements
Sequences of uppercase letters will be used to represent booleans.

The following logical operators are permitted and understood:
    
    Table 2: Operators
    Operator	Representation
    AND			*
    OR			+
    NOT			!

For version 2, operators will be customizable and definable.

We make use of the fact that statements can be mapped to truth tables.
We distinguish between two types of statements, molecular and atomic.
Atomic statements are single-variables, and can be written
in the form *A* without any operators.

Molecular statments are composed of other statements (molecular and
atomic) using operators, and can be written in the form *A[O]B* where
*[O]* represents an operator.

Here are some examples to that effect:

    Atomic: A,B,C...
    Molecular: A+B, A*B, (A*B)+C
