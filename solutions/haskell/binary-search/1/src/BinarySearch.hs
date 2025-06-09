module BinarySearch (find) where

import Data.Array

find :: Ord a => Array Int a -> a -> Maybe Int
find array x = binarySearch l r
  where
    (l, r) = bounds array
    binarySearch li ri
      | li > ri = Nothing
      | midVal == x = Just midIdx
      | midVal < x = binarySearch (midIdx + 1) ri
      | midVal > x = binarySearch li (midIdx - 1)
      where
        midIdx = (li + ri) `div` 2
        midVal = array ! midIdx