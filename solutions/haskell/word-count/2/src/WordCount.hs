module WordCount (wordCount) where

import Data.Bool (bool)
import Data.Char (isLetter, isNumber, toLower)
import Data.List (dropWhileEnd)
import Data.Map (fromListWith, toAscList)
import GHC.Utils.Misc ((<||>))

wordCount :: String -> [(String, Int)]
wordCount = toAscList . toCounter . toTuple . trimApos . splitWords
  where
    isApos = (== '\'')
    isNotPunct = isApos <||> isLetter <||> isNumber
    translatePunctToSpace = bool ' ' <*> isNotPunct
    normalize = toLower . translatePunctToSpace
    splitWords = words . map normalize
    trimAposPre = dropWhile isApos
    trimAposSuf = dropWhileEnd isApos
    trimApos = map $ trimAposSuf . trimAposPre
    toTuple = flip zip $ repeat 1
    toCounter = fromListWith (+)