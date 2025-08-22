local function combination(list, count)
    local result = {}
    if count == 0 then return result end
    if count == 1 then
        for _, elem in ipairs(list) do
            table.insert(result, {elem})
        end
        return result
    end
    for i, elem in ipairs(list) do
        for _, comb in ipairs(combination({table.unpack(list, i + 1)}, count - 1)) do
            table.insert(comb, 1, elem)
            table.insert(result, comb)
        end
    end
    return result
end

local function combinations(sum, size, exclude)
    local result = {}
    for _, comb in ipairs(combination({1, 2, 3, 4, 5, 6, 7, 8, 9}, size)) do
        local acc = 0
        local hasExclude = false
        for _, elem in ipairs(comb) do
            acc = acc + elem
            for _, excl in ipairs(exclude or {}) do
                if excl == elem then hasExclude = true end
            end
        end
        if acc == sum and not hasExclude then table.insert(result, comb) end
    end
    return result
end

return { combinations = combinations }
