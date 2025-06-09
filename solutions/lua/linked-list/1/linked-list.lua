-- when non-empty: both left and right index must have a non-nil value!
-- when empty: both left and right index must have a nil!
-- possible scenario
-- left = 0 & right = 0 <index 0 has a value>
-- left = 1 & right = 2 <index 1, 2 have a value>
-- left = 5 & right = 4 <empty>
local List = {
  left = 0,
  right = -1,
}
List.__index = List

function List:count()
  return self.right - self.left + 1
end

function List:pop()
  if self:count() == 0 then
    return
  end
  local result = self[self.right]
  self.right = self.right - 1
  return result
end

function List:push(value)
  self.right = self.right + 1
  self[self.right] = value
end

function List:shift()
  if self:count() == 0 then
    return
  end
  local result = self[self.left]
  self.left = self.left + 1
  return result
end

function List:unshift(value)
  self.left = self.left - 1
  self[self.left] = value
end

function List:delete(value)
  local temp
  for _ = 1, self:count() do
    temp = self:pop()
    if temp ~= value then
      self:unshift(temp)
    end
  end
end

return function()
  local list = {}
  setmetatable(list, List)
  return list
end