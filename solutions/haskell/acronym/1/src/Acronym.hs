module Acronym (abbreviate) where

import Data.Char (isLetter, isLower, isNumber, isSpace, isUpper, toUpper)
import Data.List (dropWhileEnd)
import Data.List.Split (wordsBy)

abbreviate :: String -> String
abbreviate = (concatMap getInitial) . (map trimNonLetterDigit) . (concatMap splitCamelCase) . splitString
  where
    splitString = wordsBy (\c -> isSpace c || c == '-')
    getInitial [] = []
    getInitial (c:_) = [toUpper c]
    trimNonLetterDigit = (dropWhileEnd nonLetterDigit) . (dropWhile nonLetterDigit)
    nonLetterDigit c = (not . isNumber) c && (not . isLetter) c
    splitCamelCase = foldr step []
      where
        step c [] = [[c]]
        step c acc@(str@(l:_):strs)
          | isLower c && isUpper l = [c] : acc
          | otherwise = (c:str) : strs