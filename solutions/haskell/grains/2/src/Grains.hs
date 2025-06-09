module Grains (square, total) where

import Data.Bits (complement, shift)
import Data.Word (Word64)

square :: Integer -> Maybe Integer
square n
  | 1 <= n && n <= 64 = Just (fromIntegral (1 :: Word64) `shift` (fromInteger n - 1))
  | otherwise = Nothing

total :: Integer
total = (fromIntegral . complement) (0 :: Word64)