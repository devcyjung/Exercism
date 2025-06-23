module WordCount (wordCount) where

import Data.Char (isLetter, isNumber, toLower)
import Data.List (dropWhileEnd)
import qualified Data.Map as M

wordCount :: String -> [(String, Int)]
wordCount = M.toAscList . toCounter . toTuple . trimApos . splitWords
  where
    isApos x = x == '\''
    translateToSpace x
      | isApos x || isLetter x || isNumber x = toLower x
      | otherwise = ' '
    splitWords = words . map translateToSpace
    trimAposPre = dropWhile isApos
    trimAposSuf = dropWhileEnd isApos
    trimApos = map $ trimAposSuf . trimAposPre
    toTuple = flip zip $ repeat 1
    toCounter = M.fromListWith (+)