local function append(xs, ys)
  local result = {}
  for _, x in ipairs(xs) do
    table.insert(result, x)
  end
  for _, y in ipairs(ys) do
    table.insert(result, y)
  end
  return result
end

local function concat(...)
  local xss = {...}
  local result = {}
  for _, xs in ipairs(xss) do
    for _, x in ipairs(xs) do
      table.insert(result, x)
    end
  end
  return result
end

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

local function foldl(xs, init, f)
  local acc = init
  for _, x in ipairs(xs) do
    acc = f(acc, x)
  end
  return acc
end

local function foldr(xs, init, f)
  local acc = init
  for i = #xs, 1, -1 do
    acc = f(acc, xs[i])
  end
  return acc
end

local function reverse(xs)
  local result = {}
  for i = #xs, 1, -1 do
    table.insert(result, xs[i])
  end
  return result
end

local function length(xs)
  return #xs
end

return { append = append, concat = concat, filter = filter, foldl = foldl, foldr = foldr, length = length, map = map, reduce = reduce, reverse = reverse}