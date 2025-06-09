module ReverseString (reverseString) where

import Data.Foldable (foldl')

reverseString :: String -> String
reverseString = foldl' (flip (:)) []