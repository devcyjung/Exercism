local function digits(n)
  local len = 1
  while n > 10 do
    n = n / 10
    len = len + 1
  end
  return len
end

local function arm_sum(n)
  local len <const> = digits(n)
  local sum = 0
  while n > 0 do
    local rem = math.fmod(n, 10)
    sum = sum + rem ^ len
    n = (n - rem) / 10
  end
  return sum
end

return {
  is_armstrong_number = function(n)
    return n == arm_sum(n)
  end
}