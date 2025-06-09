module Isogram (isIsogram) where

import qualified Data.Set as Set
import Data.Char (isLetter, toLower)

isIsogram :: String -> Bool
isIsogram = go Set.empty
  where
    go _ "" = True
    go s (x:xs)
      | isLetter(x) =
        let c = toLower(x)
        in if Set.member c s
           then False
           else go (Set.insert c s) xs
      | otherwise =
        if x == ' ' || x == '-'
        then go s xs
        else False