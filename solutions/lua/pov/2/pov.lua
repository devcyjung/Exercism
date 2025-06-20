local function clone_tree(tree)
  local clone = {}
  for k, v in pairs(tree) do
    if type(v) ~= 'table' then
      clone[k] = v
    else
      clone[k] = clone_tree(v)
    end
  end
  return clone
end

local function rearrange(from, trail)
  local cur = from
  while trail[cur] ~= nil do
    local parent = trail[cur]
    if #cur < 2 then
      cur[2] = {}
    end
    table.insert(cur[2], parent)
    for i, parent_child in ipairs(parent[2]) do
      if parent_child == cur then
        table.remove(parent[2], i)
        if #parent[2] == 0 then
          table.remove(parent, 2)
        end 
        break
      end
    end
    cur = parent
  end
  return from
end

local function pov_from(node_name)
  return {
    of = function(tree)
      local trail = {}
      local stack = { clone_tree(tree) }
      while #stack > 0 do
        local cur = table.remove(stack)
        if cur[1] == node_name then
          return rearrange(cur, trail)
        end
        if #cur >= 2 then
          for _, child in ipairs(cur[2]) do
            trail[child] = cur
            table.insert(stack, child)
          end
        end
      end
      error('node not found')
    end
  }
end

local function ancestry(tree, node_name)
  local trail = {}
  local stack = { tree }
  while #stack > 0 do
    local cur = table.remove(stack)
    if cur[1] == node_name then
      local ancestry_list = {}
      repeat
        table.insert(ancestry_list, cur[1])
        cur = trail[cur]
      until cur == nil
      return ancestry_list
    end
    if #cur >= 2 then
      for _, child in ipairs(cur[2]) do
        trail[child] = cur
        table.insert(stack, child)
      end
    end
  end
  error('node not found')
end

local function path_from(source)
  return {
    to = function(destination)
      return {
        of = function(tree)
          local from_src = ancestry(tree, source)
          local from_dst = ancestry(tree, destination)
          local last_removed = nil
          while from_src[#from_src] == from_dst[#from_dst] do
            last_removed = from_src[#from_src]
            table.remove(from_src)
            table.remove(from_dst)
          end
          if last_removed ~= nil then
            table.insert(from_src, last_removed)
          end
          for i = #from_dst, 1, -1 do
            table.insert(from_src, from_dst[i])
          end
          return from_src
        end
      }
    end
  }
end

return { pov_from = pov_from, path_from = path_from }