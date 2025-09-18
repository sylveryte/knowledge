# Basics of Go / golang

this gives basic knowledge

## Data types

### The zero value

short **zv**

- Like many langs, go assigns a zero value to variables that are declared but not assigned

### Types

#### Literal

is code literal value

- number defaults to int
- Literals are untyped
  - go waits for use to determine the type, (ofcourse only compatible)

#### Boolean

`var flag bool` [zv] = false

#### Numeric

`int[8-64] & uint[8-64]` [zv] = 0

- int8 -128 127 aka byte
- int16 -32768 32767
- int32 - 2147483648 2147483647
- int64 -9223372036854775808 9223372036854775807
- uint8 0 255
- uint16 0 65535
- uint32 0 4294967295
- uint64 0 18446744073709551615
- _int/uint_ is int32/unint on 32machine and 64 on...

#### Float

[zv] is like int 0

- float32 has only /six or seven decimal/
- float64
- Use lib for currency

#### Strings and Runes

- strings are immutable
- for comp same == != > etc..
- runes are 'a' single char type it's actual type is int32/rune `var k rune = 'k'`

#### Explicit type conversion

- In go only explict type conversion is allowed,
- `if(inttype)` _won't_ work

## Variables

### Declarations

```go
var x int = 10;
var x = 10;
var x int;
var x, y int = 10, 20;
var x, y int;
var x, y, 10, "hello";
var (
x int
y = 10
z int = 30
d, e = 10, "hello"
f, g string
)

var x = 10;
x := 10
```

### vs const

- const are very limited in go, can only be used to store const literals.
  Not even _const object or array ref_

### vs :=

```go
var x = 10;
x := 10

x, y := 40, "hello"
```

- one special thing := does is,
  it can assign value to existing values, that var cannot do
- := is _not legal_ outside a function
- use var in func when you want to assign a [zv], `var x int`

### Arrays and slices

- `[3]int` makes a _array_ vs `[]int` makes a _slice_
- don't use array unless you know the size,
- check in terms of pointers [[Go Pointers#Difference between slices and maps]]
- use slice it's array but with blank `[]int`
  ```go
  var k = [3]int{23, 4, 5} // [3]int type array
  var j = [4]int{23, 4} // [4]int type array different than above [3]int type
  var l = []int{1:23, 4:4} // is a slice, looks like array but empty []
  ```
- indices can be given like in above l
- multi dim slices
  ```go
  var m = [][]int{{3, 4, 5}, {4, 5, 8: 8}} //[[3,4,5],[4,5,0,0....8]]
  const fr = m[1][0] //4
  ```
- no negative index

#### Slice operations

##### len append

- can use _len_ or _append_ builtins
  ```go
  l = append(l, 78)
  fmt.Printf("\nWare %d \n", len(l))
  ```

##### copy

- go makes a _copy of value passed_ even for slices
  -- in above example append makes a copy and returns the new appended slice

##### capacity

- slices has _capacity_, auto increased

##### make

- _make_ is used to make _empty_ slices with specified len and capacity
  ```go
  x := make([]int, 5) //len 5 capacity 5
  y := make([]int,5,10) //len 5 capacity 10
  ```

##### emptying a slice

- we can use builtin `clear`
  ```go
  s := []string{"firt","secon","thurd"}
  fmt.Println(s, len(s)) // ["first",....] 3
  clear(s)
  fmt.Println(s, len(s)) // [] 3
  ```
- retains the len

##### compare

standard lib provides two functions
~ `slices.Equal` compares len same, elements Equal
~ `slices.EqualFunc` same as above with a compare function

##### declare

- might stay nil `var data []int`
- empty slice literal `var x = []int{}`
  - this creates slice with 0 len 0 capacity
  - it is different than `nil` slice
  - favor `nil` slices
  - 0 len slice is useful when converting to json

##### convert

- when converting _slice_ to an _array_ new memory is created
- can also use type conversion to convert a slice into a _*pointer*_

##### Sort

todo should be #For sorting

- [[Go Funtions#Declare]]

#### Subslices

- Sublices are `k:=[2,3,5,6]` then `k[0:2]` will be `[2,3]`
- Similarily `k[1:]` will be `[3,5,6]`
- Subslices are ref to original slice and _not_ a copy
- Avoid overriding by not appending in sub slices
- using {\*\*\*\* copy} on small slice will copy as many,
  limited by whichever slice is smaller

### Strings bytes and runes

- a single rune or byte canbe converted to string
- making int into string
  ```go
  x := 'a'
  a := string(x)
  fmt.Println(x, a) // 97 a
  fmt.Println(string(12)) //error cannot change int to string
  ```

## Maps

- declare `var nilMap map[string]int`
- with size `ages := make(map[int][]string,10)`
- grow automatically
- check in terms of pointers {:$/studies/go/Pointers:\*\* Difference between slices and maps}[Difference between slices and maps]
- [zv] is nil

### assign

```go
tw := map[string]int{}
tw["red"] = 1
tw["lp"] = 2
fmt.Println(tw["red"]) //1
fmt.Println(tw["newknight"]) //0 [zv] of int
tw["newknight"]++
fmt.Println(tw["newknight"]) //1
tw["lp"] = 3
fmt.Println(tw["lp"]) //3
```

- IMPORTANT go stores value copies for struct like types
  - If you assign a struct to map and then modify, it will
    not stick, because while assignment it makes a copy and assigns it
  - Same with getting value out is also a copy
- in context of [[DataTypes#comma ok idiom|comma ok]] tell the diff between a key,
  thats associated with 0 or key not in the map
  ```go
  v, ok := tw["panda"]
  fmt.Println("panda", v, ", okay", ok) // panda 0 , okay false
  tw["lp"] = 0
  lv, lok := tw["lp"]
  fmt.Println("lp v", lv, ", lokay", lok) // lp v 0 , lokay true
  ```

### deleting

```go
m := map[string]int{
"hi": 6,
"hello": 10,
}
fmt.Println(m) //map[hello:10, hi:6]
delete(m, "hello")
fmt.Println(m) //map[hi:6]
```

### emptying

`clear(m)`

### compare

- `maps.Equal`
- `maps.EqualFunc`

### maps as sets

- keep value as boolean and make it true on add 😆

## Struct

### declare assign

```go
type person struct {
name string
age int
pet string
}
var fred person
// with key value
beth := person{
name: "Beth",
age: 40,
pet: "sam",
}
// with value in seq
julia := person{
"Julia",
25,
"zoro",
}

```

- [zv] is 0 `{0}`
- no diff between `var fred person` and `var fred person = {}`
- no diff between assiging an empty struct vs no value / [zv]

### nested structs

- nested structs cannot be initialized in one line like

```go
  type ErrorObject struct {
  Error struct {
  Message string `json:"msg"`
  Code string `json:"code"`
  } `json:"error"`
  }

// ❌
err ErrorObject{
Error {
Message:"something" //doesnt work
}
}

// ✅
// 1. define Error separate
// 2. do make it like this
err := ErrorObject{}
err.Error.Message = "Something"
```

### anonymous struct

- can declare a var with struct implmentation,
  without giving struct type a name
  ```go
  var person struct {
  name string
  age int
  }
  person.name = "beth"
  person.age = 23
  fmt.Println(person)
  fmt.Println(person.age)
  ```

### comparison & conversion

- struct with entirely comparable fields are comparable
- _no_ magic = to override and get `a == b` working
- go allows struct type conversion if fields, order and type are same

  ```go
  type firstPerson struct {
  name string
  age int
  }

  ```

- firstPerson instance can be converted to secondPerson
- but _cannot_ be compared, since those two are diff types
  ```go
  type secondPerson struct {
  name string
  age int
  }
  ```
- firstPerson _cannot_ be converted to thirdPerson,
  since diff order
  ```go
  type thirdPerson struct {
  age int
  name string
  }
  ```
- firstPerson _cannot_ be converted to fourthPerson,
  since field names don't match
  ```go
  type fourthPerson struct{
  firstName string
  age int
  }
  ```
- firstPerson _cannot_ be converted to fifthPerson,
  since there's additional field
  ```go
  type fifthPerson struct {
  name string
  age int
  favoriteColor string
  }
  ```

## any type

- any type is simply `interface{}` which is an empty interface
- empty interface can store any value that has 0 or more methods
- to use {:$/studies/go/Types Methods and Interfaces:\*\* Type assertion}[Type assertion]
- to use {:$/studies/go/Types Methods and Interfaces:\*\* Type Switch}[Type Switch]

## Channels

- Used for concurrency {:$/studies/go/Goroutines:}[Channels]

## comma ok idiom

- `v, ok := tw["panda"]`

## Conversions

### byte to string

- `string(thatbyte)`
- `fmt.Sprint(p.ageint)`

### string to int

---

[[Learning Go]]
