module PerfectNumbers (classify, Classification(..)) where

data Classification = Deficient | Perfect | Abundant deriving (Eq, Show)

classify :: Int -> Maybe Classification
classify i
  | i <= 0 = Nothing
  | otherwise = Just (if aliquot == i then Perfect else if aliquot > i then Abundant else Deficient)
    where aliquot = sum [x | x <- [1..(i - 1)], i `mod` x == 0]