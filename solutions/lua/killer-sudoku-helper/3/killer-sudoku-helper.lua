local function find_combinations(target, length, start, exclude_set, current, results)
    if target < 0 then return end
  
    if #current == length then
        if target == 0 then
            table.insert(results, { table.unpack(current) })
        end
        return
    end

    for i = start, target do
        if not exclude_set[i] then
            table.insert(current, i)
            find_combinations(target - i, length, i + 1, exclude_set, current, results)
            table.remove(current)
        end
    end
end

local function combinations(sum, size, exclude)
    local results = {}
    local exclude_set = {}
    for _, excl in ipairs(exclude or {}) do
        exclude_set[excl] = true
    end
    find_combinations(sum, size, 1, exclude_set, {}, results)
    return results
end

return { combinations = combinations }
