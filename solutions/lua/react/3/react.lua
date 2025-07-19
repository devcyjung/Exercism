local function Reactor()
  local reactor = {}
  local reactor_cells = {}
  
  local function add_cell(cell)
    table.insert(reactor_cells, cell)
    return cell
  end

  local function update_after(index)
    for i = index + 1, #reactor_cells do
      reactor_cells[i]._update()
    end
  end
  
  reactor.InputCell = function(initial_value)
    local input_cell = {}
    local value = initial_value
    local index = #reactor_cells + 1
    
    input_cell._update = function()
    end
    
    input_cell.get_value = function()
      return value
    end
    
    input_cell.set_value = function(new_value)
      if value == new_value then return end
      value = new_value
      update_after(index)
    end
    
    return add_cell(input_cell)
  end
  
  reactor.ComputeCell = function(...)
    local args = {...}
    local fn = table.remove(args)
    local compute_cell = {}
    local arg_vals = {}
    
    local function compute()
      for i, cell in ipairs(args) do
        arg_vals[i] = cell.get_value()
      end
      return fn(table.unpack(arg_vals))
    end

    local value = compute()
    local callbacks = {}

    compute_cell._update = function()
      local new_value = compute()
      if value == new_value then return end
      value = new_value
      for _, cb in ipairs(callbacks) do
        cb(value)
      end
    end

    compute_cell.get_value = function()
      return value
    end

    compute_cell.watch = function(callback)
      table.insert(callbacks, callback)
    end

    compute_cell.unwatch = function(callback)
      for i, cb in ipairs(callbacks) do
        if cb == callback then return table.remove(callbacks, i) end
      end
    end
    
    return add_cell(compute_cell)
  end
  
  return reactor
end

return { Reactor = Reactor }
