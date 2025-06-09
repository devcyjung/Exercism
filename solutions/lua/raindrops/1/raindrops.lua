local rules = {
  {when = function(n) return n % 3 == 0 end, then_ = "Pling"},
  {when = function(n) return n % 5 == 0 end, then_ = "Plang"},
  {when = function(n) return n % 7 == 0 end, then_ = "Plong"},
}

return function(n)
  local result = {}
  for _, rule in ipairs(rules) do
    if rule.when(n) then
      table.insert(result, rule.then_)
    end
  end
  return #result == 0 and tostring(n) or table.concat(result, "")
end