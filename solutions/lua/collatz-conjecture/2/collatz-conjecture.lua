return function(n)
  assert(n > 0, 'Only positive numbers are allowed')
  local acc = 0
  while n ~= 1 do
    while n & 1 == 0 do
      acc = acc + 1
      n = n >> 1
    end
    if n ~= 1 then
      acc = acc + 1
      n = n + 1 + (n << 1)
    end
  end
  return acc
end