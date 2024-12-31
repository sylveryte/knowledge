# Tables

- Tables are ==hetrogenious==, can mix up anything
- Anything like a function inside table or function as key
- is ==1-indexed== 🤯
- Tables are maps and arrays

```lua
-- like an array
local list = { "first", 2, false, function() print("Fourth") end }
print("Yup, 1-indexed", list[1])
print("Fourth is 4.... ",list[4]())

-- like a map
local m = {
  literal_key = "a string",
  ["an expression"] = "also works",
  [function() end] = true --crazy 🤯
}
print("literal_key", m.literal_key) --a string
print("an expression", m["an expression"]) --also works
print("function() end", m[function() end]) --nil🤔coz diff fx
```

- Can get **length** with `#tableName`
