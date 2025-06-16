module Luhn (isValid) where

import Data.Bits ((.&.), shift)
import Data.Char (ord)

isValid :: String -> Bool
isValid xs
  | any (not . validChar) xs = False
  | length digits <= 1 = False
  | otherwise = total `mod` 10 == 0
  where
    validChar c = c `elem` ' ':['0'..'9']
    digits = map (\x -> ord x - ord '0') (filter (`elem` ['0'..'9']) xs)
    total = sum $ zipWith luhn [(0 :: Int)..] (reverse digits)
    luhn i x
      | i .&. 1 == 1  = let doubled = shift x 1
                        in if doubled > 9
                           then doubled - 9
                           else doubled
      | otherwise     = x