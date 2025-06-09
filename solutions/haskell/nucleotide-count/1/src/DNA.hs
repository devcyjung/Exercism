module DNA (nucleotideCounts, Nucleotide(..)) where

import Data.Map (Map, fromList, insertWith)

data Nucleotide = A | C | G | T deriving (Eq, Ord, Show)

nucleotideCounts :: String -> Either String (Map Nucleotide Int)
nucleotideCounts xs =
  foldr step initial xs
  where
    initial = Right $ fromList [(A, 0), (C, 0), (G, 0), (T, 0)]
    step ch acc = do
      r <- parse ch
      racc <- acc
      return $ insertWith (+) r 1 racc
    parse 'A' = Right A
    parse 'C' = Right C
    parse 'G' = Right G
    parse 'T' = Right T
    parse c = Left $ "invalid character: " ++ [c]