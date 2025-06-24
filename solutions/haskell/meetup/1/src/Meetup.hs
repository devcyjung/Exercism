module Meetup (Weekday(..), Schedule(..), meetupDay) where

import Data.Time.Calendar

data Weekday = Monday
             | Tuesday
             | Wednesday
             | Thursday
             | Friday
             | Saturday
             | Sunday
             deriving Enum

data Schedule = First
              | Second
              | Third
              | Fourth
              | Last
              | Teenth

meetupDay :: Schedule -> Weekday -> Integer -> Int -> Day
meetupDay s wd y m = firstDayOfWeekOnAfter (toEnum $ fromEnum wd + 1) (YearMonthDay y m d)
  where
    d = case s of
      First -> 1
      Second -> 8
      Third -> 15
      Fourth -> 22
      Last -> gregorianMonthLength y m - 6
      Teenth -> 13