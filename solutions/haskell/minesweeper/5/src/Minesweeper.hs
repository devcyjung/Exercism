module Minesweeper (annotate) where

import Data.Bool (bool)
import Data.Char (intToDigit)
import Safe (atMay)

annotate :: [String] -> [String]
annotate board =
  [[annotateCell (i, j, v) | (j, v) <- zip [0..] row] | (i, row) <- zip [0..] board]
  where
    annotateCell (i, j, v) =
      case v of
        '*' -> '*'
        _ -> let count = length [() | di <- [-1..1], dj <- [-1..1], getCell (i + di, j + dj) == Just '*']
             in bool (intToDigit count) ' ' (count == 0)
    getCell (i, j) = do
      row <- board `atMay` i
      row `atMay` j