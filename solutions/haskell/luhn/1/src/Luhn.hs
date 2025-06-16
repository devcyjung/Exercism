module Luhn (isValid) where

import Data.Bits ((.&.), shift)
import Data.Char (ord)

isValid :: String -> Bool
isValid xs
  | any (not . (`elem` ' ':['0'..'9'])) xs = False
  | length trimmed <= 1 = False
  | otherwise = (sum (map f (zip (reverse trimmed) [(0::Int)..]))) `mod` 10 == 0
  where
    trimmed = map (\x -> ord x - ord '0') (filter (`elem` ['0'..'9']) xs)
    f (x, i)
      | i .&. 1 == 1 = let double = shift x 1
                       in if double > 9 then double - 9 else double
      | otherwise = x