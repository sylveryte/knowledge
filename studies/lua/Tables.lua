local list = { "first", 2, false, function() print("Fourth") end }
print("Yup, 1-indexed", list[1])
print("Fourth is 4.... ", list[4]())

local m = {
  literal_key = "a string",
  ["an expression"] = "also works",
  [function() end] = true --crazy 🤯
}
print("literal_key", m.literal_key) --a string
print("an expression", m["an expression"]) --also works
print("function() end", m[function() end]) --nil 🤔

-- experiments
print("list length",#list)
