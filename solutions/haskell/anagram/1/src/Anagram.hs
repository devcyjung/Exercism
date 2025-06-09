module Anagram (anagramsFor) where

import qualified Data.Map as Map
import Data.Char (toLower, isLetter)

anagramsFor :: String -> [String] -> [String]
anagramsFor word = filter isAnagram
  where
    counter :: String -> Map.Map Char Integer
    counter = foldr insertChar Map.empty
    insertChar c m
      | isLetter c = Map.insertWith (+) (toLower c) 1 m
      | otherwise = m
    baseCounter = counter baseWord
    baseWord = map toLower word
    isAnagram s = map toLower s /= baseWord && baseCounter == counter s