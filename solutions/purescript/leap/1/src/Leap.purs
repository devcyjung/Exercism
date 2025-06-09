module Leap
  ( isLeapYear
  ) where

import Prelude

isLeapYear :: Int -> Boolean
isLeapYear n = (mod n 400 == 0 || (mod n 4 == 0 && not (mod n 100 == 0)))