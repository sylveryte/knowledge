# Rust basics

# Cargo

- Cargo is for rust, like npm for javascript
- Create new project
  `cargo new project_name`
- Cargo can build
  `cargo build`
- Cargo can get docs of all dependencies
  `cargo doc --open`

# Variables and Mutability

- Variables are immutable by default
- Will infer type from assignment

## Let

- `let x=5` immutable, cannot do `x=6`
- `let mut x=5` mutable, can do `x=6`

## Constants

- Always immutable
- Naming convention `THIS_IS_CONTANT`

## Shadowing

- Can use name again with immutable also
- Have to use with let keyword
- Different from mut coz no accidental reassigment
- Can change type while shadowing

```rust
let x = 5;
println!("this is x {x}"); // 5
let x = "again";
println!("this is x {x}"); // again
```

# Data Types

## Scalar Types

- Represents single value

## Integer Type

- i8 to i128 to isize signed
- u8 to u128 to usize unsigned
- isize & usize depends on architecture of app running computer
- Can be decimal, hex, octal, binary, bytes

## Integer overflow

- Overflow to none `wrapping_*` = add,substract etc
- return Overflow to wrap `overflowing_*` methods
- Indication if overflow

### Floating Point Types

- f32 f64 use _f64_ as negligible perf diff for modern cpus

### Numeric Operations

- sum `+`
- diff `-`

* prod `_`
* div `/`
* remainder `%`

## Boolean type

- `bool`

## Character Type

- With single quatation `'c'`

## Compound Types

- Group of multiple values in one type

### The Tuple Type

- Different type group together
- Have fixed length, once done _no_ grow or shrink
- `let tup: (i32, f64) = (500, 6.5)`
- Access using .0 .1 `tup.0 //500`
- Tuple without any value is a [[#Unit]]

  ## Unit

- Is an empty tuple `()`
- [[#Statements and Expressions|Expressions]] implicitly return [[#Unit]] if no return

## Array

- Every element is _same type_
- Have fixed length
- `let a = [1,2,3,4]`
- A [[#Vector]] is a dynamic sized [[#Array]]
- `let a:[i32:5] = [1,2,3,4,5]`
- `let a = [3; 5] //[3,3,3,3,3]`
- Access using `let first = a[0]`
- Invalid access crashes the program

### Vector

- Dynmic sized [[#Array]]

### The String Type

- [[Ownership#The String Type]]

## Functions

- Uses snake_case for [[#Functions]] and [[#Variables and Mutability]]
- Doesnt matter where it is defined in a scope of caller

### Parameters

- Must declare all parameters types
- Functions have parameters, you send arguments to function call

## Statements and Expressions

- _Statements_ are instructions that _do not return any value_
- Expressions evaluate to a _resultant value_
- Expressions `x=6` doesn't returns anything, so `x=y=6` will give error
- Calling a [[#Functions]] is an expression
- Calling a [marco] is an expression
- A new scope block created with curly brackets is an expression

```rust
let y = {
  let x = 3;
  x+1
};
println!("This value of y is: {y}") // 4
let onep = plus_one(6);
println!("This value of onep is: {onep}"); // 7
}
fn plus_one(x: i32) -> i32 {
  x + 1
}
```

- Putting `;` at `x+1;` will give error, since it will be a statement

## Comments

- uses `//this is comment`
- uses `//this is multicomment`
- There is also a document comments for libs

## Control Flow

### if - Only bool expression, unlike ruby, js

```rust
if number < 5 {

    } else if number < 3 {

    } else {

    }
```

## Loops - loop

```rust
    loop {
      println("again")
    }
```

- Returning values from loops `break counter _ 2`
- Loops labels with`'outerloop`can use with`break 'outerloop` `continue 'outerloop`
  - Must begin with singlequote
- Conditional`while`loop

```rust
while number !=0 {
  number -= 1;
}
```

- Collection loop with`for`loop
- Excute some code with each item of collection
  ```rust
  let a = [1,3,4,5,6]
  for el in a {
    println!("el is {el}")
  }
  ```
- Arbitary range loop with`for`

```rust
  for number in (1..4).rev() {
    println!("number is {el}")
  }
```

- `.rev` to reverse the range

  ***

#rust
