module Minesweeper (annotate) where

import Data.Char (intToDigit)
import Safe (atMay)

annotate :: [String] -> [String]
annotate board =
  [ [ annotateCell (i, j, v) | (j, v) <- [0..] `zip` row ]
  | (i, row) <- [0..] `zip` board
  ]
  where
    annotateCell (i, j, v) =
      case v of
        '*' -> '*'
        _ ->
          let count = length [ ()
                             | di <- [-1..1]
                             , dj <- [-1..1]
                             , (di, dj) /= (0, 0)
                             , getCell (i + di, j + dj) == Just '*'
                             ]
          in if count == 0
             then ' '
             else intToDigit count
    getCell (i, j) = do
      row <- board `atMay` i
      row `atMay` j