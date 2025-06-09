local COLORCODES = {
  black = 0,
  brown = 1,
  red = 2,
  orange = 3,
  yellow = 4,
  green = 5,
  blue = 6,
  violet = 7,
  grey = 8,
  white = 9,
}

local function is_array(t)
  if type(t) ~= "table" then
    return false
  end
  local i = 1
  for _ in pairs(t) do
    if t[i] == nil then
      return false
    end
    i = i + 1
  end
  return true
end

return {
  value = function(colors)
    if not is_array(colors) then
      error(string.format([[          
          value() expects type array (type: table) of input.
          received type: %s
          if the displayed type is "table", it is a non-array table.
          ]], type(colors)))
    end
    if #colors < 2 then
      error(string.format([[          
          value() expects an array of length 2 or longer.
          received length %d
          ]], #colors))
    end
    for _, v in ipairs(colors) do
      if COLORCODES[v] == nil then
        error(string.format([[            
            color_code() could not find the following color.
            color: %s
        ]], v))
      end
    end
    return tonumber(string.format("%d%d", COLORCODES[colors[1]], COLORCODES[colors[2]]))
  end
}
