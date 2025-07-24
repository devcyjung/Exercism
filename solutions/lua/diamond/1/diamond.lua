return function(which)
  local order <const> = string.byte(which) - string.byte('A')
  if order < 0 or 25 < order then
    error("input is out of valid range")
  end
  local side <const> = 2 * order + 1
  local buf <const> = {}
  for i = 1, side do
    buf[i] = ' '
  end
  local diamond <const> = {}
  local top = 1
  local bottom = side
  local left = order + 1
  local right = left
  for o = 0, order do
    local char <const> = string.char(o + string.byte('A'))
    buf[left], buf[right] = char, char
    diamond[top] = table.concat(buf)
    diamond[bottom] = diamond[top]
    buf[left], buf[right] = ' ', ' '
    left, right = left - 1, right + 1
    top, bottom = top + 1, bottom - 1
  end
  return table.concat(diamond, '\n') .. '\n'
end