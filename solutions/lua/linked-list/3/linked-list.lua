-- Here we have two versions of implementation: 1. Doubly Linked List 2. Doubly-Ended Queue

-- 1. Doubly Linked List
local List = {
  count_ = 0,
  head = nil,
  tail = nil,
}
List.__index = List

function List.new()
  local list = {}
  setmetatable(list, List)
  return list
end

function List:count()
  return self.count_
end

function List:pop()
  if self:count() == 0 then
    return
  end
  local result = self.tail.value
  if self.tail.prev then
    self.tail.prev.next = nil
  end
  self.tail = self.tail.prev
  self.count_ = self.count_ - 1
  return result
end

function List:push(value)
  local link = {
    value = value,
    prev = self.tail,
  }
  if self.tail then
    self.tail.next = link
  end
  self.tail = link
  if self:count() == 0 then
    self.head = link
  end
  self.count_ = self.count_ + 1
end

function List:shift()
  if self:count() == 0 then
    return
  end
  local result = self.head.value
  if self.head.next then
    self.head.next.prev = nil
  end
  self.head = self.head.next
  self.count_ = self.count_ - 1
  return result
end

function List:unshift(value)
  local link = {
    value = value,
    next = self.head,
  }
  if self.head then
    self.head.prev = link
  end
  self.head = link
  if self:count() == 0 then
    self.tail = link
  end
  self.count_ = self.count_ + 1
end

function List:delete(value)
  if self:count() == 0 then
    return
  end
  local link, prevlink, nextlink
  repeat
    if not self.head then
      break
    end
    link = nextlink or self.head
    prevlink = link.prev
    nextlink = link.next
    if link.value == value then
      link.prev = nil
      link.next = nil
      self.count_ = self.count_ - 1
      if prevlink then
        prevlink.next = nextlink
      end
      if nextlink then
        nextlink.prev = prevlink
      end
      if link == self.head then
        self.head = nextlink
      end
      if link == self.tail then
        self.tail = prevlink
      end
    end
  until link == self.tail
end

-- 2. Doubly-Ended Queue
-- when non-empty: both left and right index must have a non-nil value!
-- when empty: both left and right index must have a nil!
-- possible scenario
-- left = 0 & right = 0 <index 0 has a value>
-- left = 1 & right = 2 <index 1, 2 have a value>
-- left = 5 & right = 4 <empty>
local Deque = {
  left = 0,
  right = -1,
}
Deque.__index = Deque

function Deque.new()
  local dq = {}
  setmetatable(dq, Deque)
  return dq
end

function Deque:count()
  return self.right - self.left + 1
end

function Deque:pop()
  if self:count() == 0 then
    return
  end
  local result = self[self.right]
  self.right = self.right - 1
  return result
end

function Deque:push(value)
  self.right = self.right + 1
  self[self.right] = value
end

function Deque:shift()
  if self:count() == 0 then
    return
  end
  local result = self[self.left]
  self.left = self.left + 1
  return result
end

function Deque:unshift(value)
  self.left = self.left - 1
  self[self.left] = value
end

function Deque:delete(value)
  local temp
  for _ = 1, self:count() do
    temp = self:pop()
    if temp ~= value then
      self:unshift(temp)
    end
  end
end

return function()
  return List.new()
  -- return Deque.new()
end