module Clock (addDelta, fromHourMin, toString) where

import Text.Printf (printf)

data Clock = Clock
  { hours  :: Int
  , mins   :: Int
  }
  deriving Eq

fromHourMin :: Int -> Int -> Clock
fromHourMin h m = Clock
  { hours = minutes `div` minPerHour
  , mins = minutes `mod` minPerHour
  }
  where
    minPerHour = 60
    minPerDay = 24 * minPerHour
    minutes = (h * minPerHour + m) `mod` (minPerDay)

toString :: Clock -> String
toString Clock{ hours = h, mins = m } = printf "%02d:%02d" h m

addDelta :: Int -> Int -> Clock -> Clock
addDelta h m Clock{ hours = ch, mins = cm } = fromHourMin (ch + h) (cm + m)