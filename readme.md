# SCRAM

Scram is a scheme dialect, inspired in large part by [lispy](http://norvig.com/lispy.html).

The syntax for scram is a mashup of common lisp and scheme, with some of its own idosyncracies.

## Primitives

There are 8 primitive syntactic forms in scram.

1. quote
2. if
3. cond
4. set
5. def
6. lambda
7. begin
8. go

## Builtins

There are 16 builtin functions in scram.

### List manipulation

1. cons
2. car
3. cdr
4. list

### Concurrency

1. chan
2. ->
3. <-

### Arithmetic

1. +
2. -
3. *
4. /

### Comparison

1. eq
2. lt
3. lte
4. gt
5. gte

## Status

Early development.

## Why

Educational value in interpreter / programming language design.
