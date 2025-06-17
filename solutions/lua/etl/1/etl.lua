return {
  transform = function(dataset)
    local result = {}
    for key, values in pairs(dataset) do
      for _, value in ipairs(values) do
        result[value:lower()] = key
      end
    end
    return result
  end
}