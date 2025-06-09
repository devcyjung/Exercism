module CollatzConjecture (collatz) where

collatz :: Integer -> Maybe Integer
collatz n
  | n < 1         = Nothing
  | n == 1        = Just 0
  | mod n 2 == 0  = fmap (1 +) (collatz (div n 2))
  | otherwise     = fmap (1 +) (collatz (3 * n + 1))