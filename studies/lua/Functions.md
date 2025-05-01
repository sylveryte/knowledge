# Functions

- Color calling (:) sets first arg as `self` while (.) doesn't
- Functions are ==first class== in lua
- Can be made ==as values==

```lua
local function hello(name)
 print("Hello!",name)
end
hello('cholo')

local greet = function (name)
  -- .. is concatenating
  print("Greeings, ".. name.."!")
end
greet("cholo")
```

- Higher order functions

```lua
local function adder(value)
  return function (v)
    return v + value
  end
end
local addTwo = adder(2)
print(addTwo(3))
```
