local function reduce(xs, value, f)
  local acc = value
  for _, x in ipairs(xs) do
    acc = f(x, acc)
  end
  return acc
end

local function map(xs, f)
  local result = {}
  for _, x in ipairs(xs) do
    table.insert(result, f(x))
  end
  return result
end

local function filter(xs, pred)
  local result = {}
  for _, x in ipairs(xs) do
    if pred(x) then
      table.insert(result, x)
    end
  end
  return result
end

return { map = map, reduce = reduce, filter = filter }