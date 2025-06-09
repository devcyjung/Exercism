local function flatten(input)
  if type(input) ~= 'table' then
    return {input}
  end
  local result = {}
  for _, value in ipairs(input) do
    for _, elem in ipairs(flatten(value)) do
      table.insert(result, elem)
    end
  end
  return result
end

return flatten