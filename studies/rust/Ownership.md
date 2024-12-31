# Ownership

Set of rust rules that manages memory and makes it memory safe.

## The stack and the heap

### Stack

- quick add quick remove
- lifo

### Heap

- slower than stack
- pointer system
- data structures that are on heap, implment [trait] [drop] eg: String

## Ownership rules

- Each value has an owner
- Only one owner at a time
- When owner goes out of scope, the value will be dropped

## Variable scope

- Starts when declared
- Follows { } braces

## The String Type

- 'string literal' is _not_ String type
- string literal is static _not_ String
- data is on heap, is dynamic length
- eg `let mut s = String::from("hello")`
