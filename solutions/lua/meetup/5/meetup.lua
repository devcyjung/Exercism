local WEEKDAYS = {Sunday = 1, Monday = 2, Tuesday = 3, Wednesday = 4, Thursday = 5, Friday = 6, Saturday = 7}
local START_DATES = {teenth = 13, first = 1, second = 8, third = 15, fourth = 22, last = -6}

return function(config)
  local meetup = os.date('*t', os.time({
    year = config.year,
    month = config.week ~= 'last' and config.month or config.month + 1,
    day = START_DATES[config.week]
  }))
  meetup.day = meetup.day + (WEEKDAYS[config.day] - meetup.wday + 7) % 7
  return meetup.day
end