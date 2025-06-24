module Meetup (Weekday(..), Schedule(..), meetupDay) where

import Data.Time.Calendar (Year, MonthOfYear, Day, firstDayOfWeekOnAfter, fromGregorian, gregorianMonthLength)

data Weekday = Monday | Tuesday | Wednesday | Thursday | Friday | Saturday | Sunday deriving Enum

data Schedule = First | Second | Third | Fourth | Last | Teenth deriving Enum

meetupDay :: Schedule -> Weekday -> Year -> MonthOfYear -> Day
meetupDay sch wd y m = firstDayOfWeekOnAfter (toEnum $ fromEnum wd + 1) (fromGregorian y m (case sch of
      Last -> gregorianMonthLength y m - 6
      Teenth -> 13
      _ -> 1 + 7 * fromEnum sch))