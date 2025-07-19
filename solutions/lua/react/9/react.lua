local function reactor()
	local update_fns <const> = {}

	local function update_after(index)
		for i = index + 1, #update_fns do
			update_fns[i]()
		end
	end

	local function input_cell(initial_value)
		local value = initial_value
		local index <const> = #update_fns

		local function get_value()
			return value
		end

		local function set_value(new_value)
			if value == new_value then
				return
			end
			value = new_value
			update_after(index)
		end

		return { get_value = get_value, set_value = set_value }
	end

	local function compute_cell(...)
		local args <const> = { ... }
		local fn <const> = table.remove(args)
		local arg_vals <const> = {}

		local function compute()
			for i, cell in ipairs(args) do
				arg_vals[i] = cell.get_value()
			end
			return fn(table.unpack(arg_vals))
		end

		local value = compute()
		local callbacks <const> = {}

		table.insert(update_fns, function()
			local new_value <const> = compute()
			if value == new_value then
				return
			end
			value = new_value
			for cb, _ in pairs(callbacks) do
				cb(value)
			end
		end)

		local function get_value()
			return value
		end

		local function watch(callback)
			callbacks[callback] = true
		end

		local function unwatch(callback)
			callbacks[callback] = nil
		end

		return { get_value = get_value, watch = watch, unwatch = unwatch }
	end

	return { InputCell = input_cell, ComputeCell = compute_cell }
end

return { Reactor = reactor }
