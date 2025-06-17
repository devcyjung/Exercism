local wdays = {Sunday = 1, Monday = 2, Tuesday = 3, Wednesday = 4, Thursday = 5, Friday = 6, Saturday = 7}
local startDates = {teenth = 13, first = 1, second = 8, third = 15, fourth = 22, last = 1}

return function(config)
  local wday = wdays[config.day]
  local date = {year = config.year, month = config.month, day = startDates[config.week]}
  if config.week == 'last' then
    date.month = date.month + 1
    date.day = date.day - 7
  end
  local date_table = os.date('*t', os.time(date))
  return date_table.day + (wday - date_table.wday + 7) % 7
end