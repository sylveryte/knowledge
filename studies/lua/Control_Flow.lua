local fav = { "M4", 'm2', 'am', 'goth' }
-- for i= start,end do
for index = 1, #fav do
  print(index, fav[index])
end

for index,value in ipairs(fav) do
  print(index, value)
end
