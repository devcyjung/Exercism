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

local QUANTIFIERS = {
  [0] = "ohms",
  [3] = "kiloohms",
  [6] = "megaohms",
  [9] = "gigaohms",
}

return {
  decode = function(c1, c2, c3)
    for _, c in ipairs({c1, c2, c3}) do
      if type(c) ~= "string" then
        error(string.format([[            
            decode() expects 3 string type inputs.
            received types: %s, %s, %s
            received values: %s, %s, %s
        ]], type(c1), type(c2), type(c3), c1, c2, c3))
      end
      if COLORCODES[c] == nil then
        error(string.format([[            
            decode() could not find the following color.
            color: %s
        ]], c))
      end
    end
    local digits = (10 * COLORCODES[c1] + COLORCODES[c2]) * math.pow(10, COLORCODES[c3] % 3)
    local exponent = COLORCODES[c3] // 3
    while digits >= 1000 do
      digits = digits // 1000
      exponent = exponent + 1
    end
    if QUANTIFIERS[exponent * 3] == nil then
      error(string.format([[          
          input exceeds the valid numeric range.
          exponent of ten: %d
          inputs: %d, %d, %d          
      ]], exponent * 3, COLORCODES[c1], COLORCODES[c2], COLORCODES[c3]))
    end
    return digits, QUANTIFIERS[exponent * 3]
  end
}