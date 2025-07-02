return {
  score = function(x, y)
    local radius = math.sqrt(x ^ 2 + y ^ 2)
    if radius > 10 then
      return 0
    elseif radius > 5 then
      return 1
    elseif radius > 1 then
      return 5
    else 
      return 10
    end
  end
}