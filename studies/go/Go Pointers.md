# Go Pointers

#golang

- A pointer is a variable that holds the location of a memory.
- Many things that `c` pointers can be done,
  are _not_ available in `go` including _pointer arithematics_
- `&` is the address operator
  `pointerToX := &x`
- `*` is indicator operator
  `z := 5 + *pointerToX`
- program will panic if `nil` pointer is `dereferenced`
- make sure the pointer you're `dereferencig` is _not nil_
- built in fx `new` returns a pointer to [[Go DataTypes#The zero value|zv]] variable
- use a & before a [[Go DataTypes#Struct|Struct]] literal to create a pointer instance

## Pointers indicate a mutable parameter

- go idiom is if a fx has pointer param, it means it is mutable

## Pointers are last resort

- only use pointer param to modify var is when fx expects

## Pointer passing performance

- vs [[Go Funtions#Go is call by value]]
- using pointer to pass struct as param/return,
  is more performant when struct is very large (>10MB)

## Difference between slices and maps

- Slice and [[Go DataTypes#Maps|Maps]] pass the value,
  it's just that the value happens to be a **pointer**
- the _special case_ of [[Go DataTypes#Arrays and slices|Arrays and slices]] is that it stores,
  a struct of 3 values `(ptr,len int,cap int)`
  -- on passing it is copied, and when length is modified,
  it's modified in the called fx's new var not original.
  -- hence in slices we cannot use \*append\*, we can but won't reflect
  -- since the original vars len is never updated.

## Slices as buffers

- in case of a for loop you override a same value multiple times
  `for i := range(100){ i = i*4; fmt.prin...(i)}`
- in this case you'll _create 100s of allocations_,
  and burden the [[#Garbage collector|gc]]
- better to use a slice as a buffer,
  that will hold the diff data in same location
- eg. while reading a file

## Garbage collector

- try reducing the gc burden
- go provides users a couple of settings,
  to control the heap size

---

Origin: [[Learning Go]]
