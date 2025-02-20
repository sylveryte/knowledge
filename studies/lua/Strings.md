# Strings

- Match / Date

```lua

local pattern = "(%d+)-(%d+)-(%d+) (%d+):(%d+):(%d+)"
local year, month, day, hour, minute, second = dateString:match(pattern)
local dateString = "2023-09-21 15:00:00"

local timestamp = os.time({
  year = year,
  month = month,
  day = day,
  hour = hour,
  min = minute,
  sec = second
})

-- Convert timestamp back to a human-readable format
local formattedDate = os.date("%Y-%m-%d %H:%M:%S", timestamp)
print(formattedDate)  -- Output: 2023-09-21 15:00:00
```

- GMatch
