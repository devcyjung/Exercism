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

return {
  value = function(colors)
    if #colors < 2 then
      error(string.format("value() expects an array of length 2 or longer, received length %d", #colors))
    end
    return tonumber(string.format("%d%d", COLORCODES[colors[1]], COLORCODES[colors[2]]))
  end
}
