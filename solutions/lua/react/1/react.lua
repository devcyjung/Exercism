local function Reactor()
  local reactor = {
    m_cells = {}
  }
  reactor.InputCell = function(value)
    local input = {
      m_value = value,
      m_id = #reactor.m_cells + 1,
    }
    input.get_value = function()
      return input.m_value
    end
    input.compute = input.get_value
    input.set_value = function(new_value)
      if new_value == input.m_value then
        return
      end
      input.m_value = new_value
      for i = input.m_id + 1, #reactor.m_cells do
        local cell = reactor.m_cells[i]
        local old_value = cell.m_value
        cell.m_value = cell.compute()
        if old_value ~= cell.m_value then
          for _, cb in ipairs(cell.m_callbacks) do
            cb(cell.m_value)
          end
        end
      end
    end
    table.insert(reactor.m_cells, input)
    return input
  end
  reactor.ComputeCell = function(...)
    local args = {...}
    local fn = table.remove(args)
    local computeFn = function()
      local values = {}
      for i, cell in ipairs(args) do
        values[i] = cell.get_value()
      end
      return fn(table.unpack(values))
    end
    local compute = {
      m_value = computeFn(),
      m_callbacks = {},
    }
    compute.get_value = function()
      return compute.m_value
    end
    compute.compute = computeFn
    compute.watch = function(callback)
      table.insert(compute.m_callbacks, callback)
    end
    compute.unwatch = function(callback)
      for i, cb in ipairs(compute.m_callbacks) do
        if cb == callback then
          table.remove(compute.m_callbacks, i)
          return
        end
      end
    end
    table.insert(reactor.m_cells, compute)
    return compute
  end
  return reactor
end

return { Reactor = Reactor }
