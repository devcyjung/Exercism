local function tick(matrix)
  local rows = #matrix
  local cols = matrix[1] and #(matrix[1]) or 0
  local function alive(i, j)
    local total = 0
    for x = math.max(1, i-1), math.min(i+1, rows) do
      for y = math.max(1, j-1), math.min(j+1, cols) do
        if x ~= i or y ~= j then
          total = total + matrix[x][y]
        end
      end
    end
    return total
  end
  local to_one = {}
  local to_zero = {}
  for i, row in ipairs(matrix) do
    for j, cell in ipairs(row) do
      local n = alive(i, j)
      if n == 3 then
        table.insert(to_one, {i, j})
      elseif n ~= 2 then
        table.insert(to_zero, {i, j})
      end     
    end
  end
  for _, o in ipairs(to_one) do
    matrix[o[1]][o[2]] = 1
  end
  for _, z in ipairs(to_zero) do
    matrix[z[1]][z[2]] = 0
  end
  return matrix
end

return { tick = tick }