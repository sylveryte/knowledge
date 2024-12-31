# Controls

## Variable shadowing

- blocks shadow variable with same name from outer var
- Since {## Universe Block} exists, shadowing is possible for packages & other important entities as well

## Universe Block

- Special block that encompasses all
- Go is a small lang with only 25keywords hence,
  built-in types (int, string) and consts like(true false),
  function like (make & close) defines in universe block
- be careful to not override/redefine any identifiers in universe block

## If

- no parenthesis around the condition
- **special** ability to declare var for if only scope

  ```go
  if n := rand.Intn(23); n%2 == 0 {
  fmt.Println("accessible here ", n)
  } else {
  fmt.Println("accessible here too ", n)
  }

  ```

- Technically you can put any statement before comp

## For

- same above if block, declare and scope apply here

### A complete C-style for

```go
for i := 0 ; i<10;i++ {
  fmt.Println("accessible here ", i)
}
```

### A condition-only for

- notice ; before the condition

```go
i := 0
for ; i < 10; i++ {
fmt.Println("accessible here ", i)
}
```

### An infinite for

- can utilise break to break it

```go
i := 0
for {
  i++
  fmt.Println("accessible here ", i)
  if i > 10 {
    fmt.Println("breaking")
    break
  }
}
```

### for-range

- value is a #copy#, will not modify compound types
- can use {### break and continue}
- can range slices,arrays,maps, ints too
- note {:$/studies/go/DataTypes:# Maps 🗺️}[maps] iteration order varies

  ```go
  // slice/arrays
  evnV := []int{4, 6, 8, 10, 4}
  for i, v := range evnV {
  fmt.Println("i index ", i, "v is value", v)
  }
  // maps
  uniqu := map[string]bool{"fred":true,"monty":false}
  for k, v := range uniqu {
  fmt.Println("k is key", k, "v is value", v)
  }
  // int
  for i := range 10 {
  fmt.Println("now ",i)
  }
  ```

- **should** use for {:$/studies/go/DataTypes:### Strings and Runes}[Strings] type,
  it give you multi byte {:$/studies/go/DataTypes:### Strings and Runes}[rune] as one
- it iterates over runes instead of bytes
- there's also for-range for {:$/studies/go/Goroutines:### Using for-range}[Channels]

### break and continue

- can _label_ the loops

```go
uniqu := map[string]bool{"fred": true, "monty": false}
outer:
for k, v := range uniqu {
  fmt.Println("k is key", k, "v is value", v)
  for i := range 20 {
    i++
    if i > 10 {
      break outer
    }
  }
}
```

## Switch

- Unlike other langs switch in go is #useful#
- **Won't fall down** instead only matched will work,
  no need to break;
- Var declared in a block is limited to that var only
- empty case means nothing will happen
  `case 6;`
- has same statement expression as in [[Controls#If|If]] && [[Controls#For|For]]

  ```go
  switch uniqu := map[string]bool{"fred": true, "monty": false}; len(uniqu) {
  case 1,2,3: {
  fmt.Println("len is within 4")
  }
  case 4: {
  r:= 43
  fmt.Println("len is exaclty 4",r)
  }
  case 5,6: {
  fmt.Println("len greater then 4, r not accessible here")
  }
  default:{
  fmt.Println("this is defualt")
  }
  }
  ```

- go has `fallthrough` key for default switch behavior like other langs,
  but you shouldnt be using it
- if switch is inside a {## For}[for] then put a label for loop,
  go assumes a {### break and continue}[break] is for breaking the case

### Blank switch

- can have a blank switch no need for a value to compare with
- is very #powerful#
- for above swith you can put condition inside the case condi
  ```go
  j := []int{3, 4, 5}
  switch l := len(j); {
  case l < 4:
  {
  fmt.Println("len is less than 4")
  }
  }
  // truly blank
  k := len(j)
  switch {
  case k < 4:
  {
  fmt.Println("len is less than 4")
  }
  }
  ```

```

```

### Type Switch

{:$/studies/go/Types Methods and Interfaces:## Type Switch}[Type Switch]

## Goto

- yes the black sheep `goto`
- you can't jump anywhere go forbids jumps
  that skip over a var declaration or
  into an inner or parallel block

```go
k := 10
begin:
k++
fmt.Println("Ware here k",k)
if k < 20 {
 goto begin
}
```

---

[[Learning Go]]
[[Controls]]
