local grains = {}

function grains.square(n)
  return 2 ^ (n - 1)
end

function grains.total()
  return 2 ^ 64 - 1
end

return grains