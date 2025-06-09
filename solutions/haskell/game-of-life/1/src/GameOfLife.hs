module GameOfLife (tick) where

tick :: [[Int]] -> [[Int]]
tick xs = [[isAlive i j | j <- [0..(ncols i)-1]] | i <- [0..nrows-1]]
  where
    nrows = length xs
    ncols i = length (xs !! i)
    vitality i j
      | i < 0 || i >= nrows = 0
      | j < 0 || j >= ncols i = 0
      | otherwise = (xs !! i) !! j
    acc i j = sum [vitality (i + di) (j + dj) | di <- [-1..1], dj <- [-1..1], (di, dj) /= (0, 0)]
    isAlive i j = case (vitality i j, acc i j) of
      (0, 3) ->  1
      (1, a) | a == 2 || a == 3 -> 1
      _ -> 0