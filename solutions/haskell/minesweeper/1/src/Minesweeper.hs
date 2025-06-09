module Minesweeper (annotate) where

annotate :: [String] -> [String]
annotate board = [
    [annotation i j
      | j <- [0..((ncols i) - 1)]
    ]
    | i <- [0..(nrows - 1)]
  ]
  where
    nrows = length board
    ncols i = length (board !! i)
    minecount i j = sum[
        if i >= 0 && i < nrows && j >= 0 && j < ncols i && (board !! i) !! j == '*'
        then 1 else 0
        | i <- [i-1..i+1], j <- [j-1..j+1]
      ]
    annotation i j
      | (board !! i) !! j == '*' = '*'
      | otherwise = if minecount i j == 0
        then ' '
        else toEnum (fromEnum '0' + minecount i j)::Char