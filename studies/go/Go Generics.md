# Generics

#golang

## Types

- [example](./action/generic_stack/main.go)

```go
type Stack[T any] struct{
  vals []T;
}
j := Stack[int]{
  vals: []int{},
}
```

- cannot use `any` with ==
- with [[Go DataTypes#any type]] you can only store and retreive value
- go types can be compared with ==
- hencec there is a built-in [[Go Types Methods and Interfaces]] in the universe
  block `comparable`
- To make a type that holds any two value of same type

```go
type Pair[T any] struct{
  val1 T
  val2 T
}
```

## Functions

- To abstract away functions
- Generic is specified after the name in `[]`
  eg: `func Map[t1 , t2 any](s T1[]...`
- To use `k:=[]int{};s := Map(k,...`
- example of `Map, Filter, Reduce`
  [ex](./action/generic_functions_map_filter_reduce/main.go)

```go
func Filter[t any](s []t, f func(a t) bool) []t {
  r := []t{}
  for _, v := range s {
    if f(v) {
      r = append(r, v)
    }
  }
  return r
}
```

## Interfaces

- interface with type parameters

```go
type Differ[T any] interface {
  fmt.Stringer
  Diff(T) float64
}
```

## Ultimate example

[example](./action/generic_pairs/main.go)

## Type Elements

- Type element is composed of one or more _type terms_
  within an interface
- _type terms_ can also be
  `slices maps arrays channels structs even functions`

```go
type Integer interface{
  int | int8 | int16 | int32 | int64 | uint...
}
```

- 🍉 Be aware the interfaces with type elements are
  valid only as type constraints
- By default /type terms/ match exactly
- We can specify /type terms/ to match with underlying type
  as well with `~int`

```go
type Integer interface{
  ~int | ~int8...
}
```

- above will match with `type MyInt int`, `MyInt` as well
- ther are interfaces in `cmp` package like `Ordered`
  interface, which support the `==, !=, <, >, <=, and >=`
  operators
- It is legal to have both `type term` and method elements
  in an interface used for type parameters

```go
type PrintableInt interface{
  ~int,
  String() string
}
```

---

[[Learning Go]]
