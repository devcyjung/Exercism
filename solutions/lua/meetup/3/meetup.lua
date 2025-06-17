local WEEKDAYS = {Sunday = 1, Monday = 2, Tuesday = 3, Wednesday = 4, Thursday = 5, Friday = 6, Saturday = 7}
local START_DATES = {teenth = 13, first = 1, second = 8, third = 15, fourth = 22, last = 1}

return function(config)
  local date = {year = config.year, month = config.month, day = START_DATES[config.week]}
  if config.week == 'last' then
    date.month = date.month + 1
    date.day = date.day - 7
  end
  local meetup = os.date('*t', os.time(date))
  meetup.day = meetup.day + (WEEKDAYS[config.day] - meetup.wday + 7) % 7
  return meetup.day
end