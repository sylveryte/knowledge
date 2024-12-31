local function hello(name)
 print("Hello!",name)
end
hello('Cholo')

local greet = function (name)
  -- .. is concatenating
  print("Greetings, ".. name.."!")
end
greet("Cholo")

local function adder(value)
  return function (v)
    return v + value
  end
end
local addTwo = adder(2)
print(addTwo(3))
