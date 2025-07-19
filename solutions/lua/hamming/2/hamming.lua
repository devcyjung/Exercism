local function compute(a, b)
  if string.len(a) ~= string.len(b) then
    error("strands must be of equal length")
  end
  local acc = 0
  for i = 1, string.len(a) do
    if a:sub(i, i) ~= b:sub(i, i) then
      acc = acc + 1
    end
  end
  return acc
end

return {compute = compute}