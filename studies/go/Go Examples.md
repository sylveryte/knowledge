# Examples

## example 3 type inheritance

```go
type car struct {
wheels int
color  string
}

func (c *car) horn() {
fmt.Println("honk honk, car with " + fmt.Sprint(c.wheels) + " wheels and " + c.color + " color!")
}

type truck car

func (t *truck) sprint() {
fmt.Println("broom broom, truck with " + fmt.Sprint(t.wheels) + " wheels and " + t.color + " color!")
}

func main() {
beetl := car{
wheels: 4,
color:  "Yellow",
}
beetl.horn()
f1 := beetl.horn
f1()  //still attached to beetl
// beetl.sprint(); // error

prime := truck{
wheels: 8,
color:  "Red",
}
// prime.horn(); // error
prime.sprint()
}
```

## example 2 binary tree

- used
  - [[Go DataTypes#Struct|Struct]]
  * [[Go Types Methods and Interfaces#Methods|Methods]]
  * [[Go Types Methods and Interfaces#User defined types|User defined types]]

```go
type IntTree struct {
val   int
left  *IntTree
right *IntTree
}

func (it *IntTree) insert(val int) *IntTree {
if it == nil {
return &IntTree{ val: val }
}
if it.val > val {
it.left = it.left.insert(val)
} else if it.val < val {
it.right = it.right.insert(val)
}
return it
}

func (it *IntTree) contains(val int) bool {
switch {
case it == nil:
return false
case it.val > val:
return it.left.contains(val)
case it.val < val:
return it.right.contains(val)
default:
return true
}
}

func (it *IntTree) print() {
if it == nil {
fmt.Println("empty node")
} else {
fmt.Println("val => ", it.val)
it.left.print()
it.right.print()
}
}

func main() {
var t *IntTree
t = t.insert(23)
t = t.insert(43)
t = t.insert(87)
t = t.insert(2)
t = t.insert(73)
t = t.insert(99)
t = t.insert(10)
fmt.Println(t.contains(2))   // true
fmt.Println(t.contains(8))   // false
fmt.Println(t.contains(73))  // true
fmt.Println(t.contains(100)) // false
t.print()
}
```

## example 1 function ops from map

- used
- [[Go Funtions#Function Type|Function Type]]
- [[Go DataTypes#Maps|Maps]]
- [[Go DataTypes#Conversions|Conversions]] string to int

```go
type op func(int, int) (int, error)

func add(a, b int) (int, error) {
return a + b, nil
}
func div(a, b int) (int, error) {
if b == 0 {
return 0, errors.New("divide by 0")
}
return a / b, nil
}
func sub(a, b int) (int, error) {
return a - b, nil
}

func main() {
m := map[string]op{
"+": add,
"-": sub,
"/": div,
}

exprs := []string{
"14 + 56",
"4 / 2",
"16 - 4",
"50 / 0",
"7 + 9",
}

for _,e := range exprs {
seps := strings.Split(e, " ")
if len(seps) != 3 {
fmt.Println("expression not proper",e, len(seps),seps)
continue
}
op := m[string(seps[1])]
ai,_ := strconv.Atoi(seps[0])
bi,_ := strconv.Atoi(seps[2])
v,err := op(ai,bi)
if err != nil{
fmt.Println(e, " = error", err)
continue
}
fmt.Println(e," = ",v)
}
}
```

---

[[Learning Go]]
