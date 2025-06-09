module Knapsack (maximumValue) where

maximumValue :: Int -> [(Int, Int)] -> Int
maximumValue _ [] = 0
maximumValue mw ((w, v):xs)
  | w > mw = excluded
  | otherwise = max included excluded
  where
    excluded = maximumValue mw xs
    included = v + maximumValue (mw - w) xs