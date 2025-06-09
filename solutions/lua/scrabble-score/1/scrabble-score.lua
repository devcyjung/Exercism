local SCORES = {
  A = 1, E = 1, I = 1, O = 1, U = 1, L = 1, N = 1, R = 1, S = 1, T = 1,
  D = 2, G = 2,
  B = 3, C = 3, M = 3, P = 3,
  F = 4, H = 4, V = 4, W = 4, Y = 4,
  K = 5,
  J = 8, X = 8,
  Q = 10, Z = 10,
}

return {
  score = function(word)
    if word == nil then
      return 0
    end
    if type(word) ~= "string" then
      error(string.format([[          
          score() expects string type input
          received type: %s
          received value: %s
      ]], type(word), tostring(word)))
    end
    local total = 0
    for _, cp in utf8.codes(word) do
      local char = string.upper(utf8.char(cp))
      local num = SCORES[char]
      total = total + (num or 0)
    end
    return total
  end
}