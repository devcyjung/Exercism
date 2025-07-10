module PerfectNumbers (classify, Classification(..)) where

data Classification = Deficient | Perfect | Abundant deriving (Eq, Show)

classify :: Int -> Maybe Classification
classify i
  | i <= 0 = Nothing
  | otherwise = Just classification
    where
      aliquot = sum [x | x <- [1..(i `quot` 2)], i `mod` x == 0]
      classification = case compare aliquot i of
        EQ -> Perfect
        GT -> Abundant
        LT -> Deficient