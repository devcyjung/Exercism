module SumOfMultiples (sumOfMultiples) where

import Data.Foldable (foldl')
import Data.Set (Set)
import qualified Data.Set as Set

sumOfMultiples :: [Integer] -> Integer -> Integer
sumOfMultiples factors limit = sum multiples
  where
    multiples = foldl' f Set.empty factors
    f set factor
      | factor <= 0 = set
      | otherwise = foldl' (\acc i -> Set.insert (i * factor) acc) set (takeWhile (\i -> i * factor < limit) [1..])