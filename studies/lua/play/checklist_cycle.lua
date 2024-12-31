print("hola")
local s = " - [>] Admin should not be deactivateable "

function reverseTable(t)
  local r = {}
  for key, value in pairs(t) do
    r[value] = key
  end
  return r
end

local baseBucket = {
  ["[ ]"] = "[/]",
  ["[/]"] = "[x]",
  ["[x]"] = "[>]",
  ["[>]"] = "[?]",
  ["[?]"] = "[!]",
  ["[!]"] = "[ ]",
}

local reverseBaseBucket = reverseTable(baseBucket)
local specialBucket = {
  ["[ ]"] = "[/]",
  ["[/]"] = "[x]",
  ["[x]"] = "[>]",
  ["[>]"] = "[?]",
  ["[?]"] = "[!]",
  ["[!]"] = "[*]",
  ["[*]"] = "[\"]",
  ["[\"]"] = "[l]",
  ["[l]"] = "[b]",
  ["[b]"] = "[i]",
  ["[i]"] = "[S]",
  ["[S]"] = "[I]",
  ["[I]"] = "[p]",
  ["[p]"] = "[c]",
  ["[c]"] = "[f]",
  ["[f]"] = "[k]",
  ["[k]"] = "[w]",
  ["[w]"] = "[u]",
  ["[u]"] = "[d]",
  ["[d]"] = "[ ]",
}

function cycle(s,bucket)
  local state_pattern = "%[.-%]"

  return s:gsub(state_pattern, function(state)
    return bucket[state] or state
  end)
end

print("c:", cycle(s,reverseBaseBucket))
