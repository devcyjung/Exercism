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
  color_code = function(color)
    if type(color) ~= "string" then
      error(string.format([[          
          color_code() expects a string input.
          received: %s
          ]], type(color)))
    end
    if COLORCODES[color] == nil then
      error(string.format([[          
          color_code() could not find the following color.
          color: %s
          ]], color))
    end
    return COLORCODES[color]
  end
}