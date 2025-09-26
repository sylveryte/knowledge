# Go Funtions

#golang

## Declare

- basic syntax
  ```go
  func div(num int, denom int) int {
  if denom == 0 {
  return 0
  }
  return num / denom
  }
  ```
- for consecutive input params can put type in the end
  `func div(num, denom, remainder int)`

## Simulating named and optional params

- make a person [[Go DataTypes#Struct|Struct]]
- use optional as well
  ```go
  fullname(Person{firstname: "ram",lastname: "swar"})
  fullname(Person{firstname: "tiwari"})
  ```

## Variadic input params and slices

- _spread operator_ same concept
  ```go
  func main() {
  sl := []int{4,5,2,3}
  fmt.Println(sum(sl...)) // 14
  fmt.Println(sum(2,5,6)) //13
  }
  func sum( ak ...int) int{
  var total int
  for \_,v := range ak{
  total += v;
  }
  return total;
  }
  ```

## Multiple Return Values

- must return/catch all by ,
- can ignore return value with `_` as usual
  ```go
  func main() {
  a,b,c := div(15, 2)
  fmt.Println(a,b,c)
  }
  func div(num int, denom int) (int, int, int) {
  if denom == 0 {
  return 0,0,0
  }
  return num / denom, denom, num % denom
  }
  ```

## Named return values

- predeclaring variables to use with the function
- _warning_ you don't have to return them,
  can return diff things
- notice the types
  ```go
  func main() {
  a,b,c := div(15, 2)
  fmt.Println(a,b,c)
  }
  func div(num int, denom int) (result int, remainder int, err error) {
  if denom == 0 {
  err = errors.New("divide by 0")
  return result, remainder, err
  }
  result = num /denom
  remainder = num % denom
  return result, remainder, err
  }
  ```
- On invalid input it returns the [[Go DataTypes#The zero value|zv]] values immediately

## Blank return do not use

    🚧
    - It confuses wether it returned

## Function Type

- `type fxTypeName func(int int) int`

## Annonymous Funtions

- use `func` without the name
- they can be used to assign
- they can be passed into a func as variable

### For eg sorting

- inplace sorting remember [[#Go is call by value]]

```go
sort.Slice(p, func(i,j int)bool{return p[i].name<p[j].name})
```

## Closure

- a function with a function that retains the outer scoped variable,
  even when outer function is done executing

## Defer

- `defer` keyword makes sure that,
  the derered statement runs at the end of function
- can use defer with multiple statement, runs in LIFO
- tip can use intermediately called function to group actions
- best use of [[#Named return values]] can be used with defer,
  to provide context to defered actions (eg: error)
  ```go
  //simple cat uses defer to close file
  func main() {
  if len(os.Args) < 2 {
  log.Fatal("no file specified")
  }
  f,err := os.Open(os.Args[1])
  if err != nil {
  log.Fatal(err)
  }
  defer f.Close()
  data := make([]byte, 2048)
  for {
  count, err := f.Read(data)
  os.Stdout.Write(data[:count])
  if err != nil {
  if err != io.EOF{
  log.Fatal(err)
  }
  }
  break
  }
  }
  ```
- _pro tip_ you can use `defer` to rollback or commit based on error,
  with [[#Named return values]] or declare value beforehand, no issue
- another _pro tip_ can make helper functions for eg,
  readFile that returns closer fx, file and error,
  then you defer the closer function right away

## Go is call by value

- always passes a _copy_ of variable but 🍉
- paases `ref` for [[Go DataTypes#Maps]]]]
- can modify values, but cannot grow size of [[DataTypes#Arrays and slices]]
- example
  ```go
  p := Person{}
  p.name = "helo"
  i := 4
  s := []int{23, 5}
  modify(p,i,s)
  //printed in modify {bello 0} 89 [98765 5]
  // printed again in main {helo 0} 4 [98765 5]
  // doesn't change in main for struct and int but for slice
  ```

---

Origin: [[Learning Go]]
