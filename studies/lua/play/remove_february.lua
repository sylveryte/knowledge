-- 2025-02-February << input
-- 2025-02 << output


function getDateStringForMonth(s)
  local pattern = "(%d+)-(%d+)-"
  local y, m = s:match(pattern, 0)

  return y .. "-" .. m ..  "-1"
end

local s = "2025-02-February"
print("output is ", getDateStringForMonth(s))
