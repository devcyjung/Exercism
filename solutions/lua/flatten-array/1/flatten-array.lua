local function is_array(tbl)
  if type(tbl) ~= 'table' then
    return false
  end
  local i = 0
  for _ in pairs(tbl) do
      i = i + 1
      if tbl[i] == nil then
        return false
      end
  end
  return true
end

local function flatten(input)
  if is_array(input) then
    local result = {}
    for _, value in ipairs(input) do
      for _, elem in ipairs(flatten(value)) do
        table.insert(result, elem)
      end
    end
    return result
  end
  if input then
    return {input}
  end
  return nil
end

return flatten