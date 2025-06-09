module Pangram (isPangram) where

import qualified Data.Set as Set
import Data.Char (isAscii, isLetter, toLower)

isPangram :: String -> Bool
isPangram text = Set.size letterSet == 26
  where
    letterSet = Set.fromList [toLower(c) | c <- text, isAscii c, isLetter c]