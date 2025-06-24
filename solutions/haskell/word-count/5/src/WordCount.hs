{-# LANGUAGE TupleSections #-}

module WordCount (wordCount) where

import Data.Bool (bool)
import Data.Char (isLetter, isNumber, toLower)
import Data.List (dropWhileEnd)
import Data.Map (fromListWith, toAscList)

wordCount :: String -> [(String, Int)]
wordCount = toAscList . toCounter . trimApos . splitWords
  where
    (<||>) = (<*>) . ((||) <$>)
    isApos = (== '\'')
    mapNonWordToSpace = bool ' ' <*> isApos <||> isLetter <||> isNumber
    splitWords = words . fmap toLower . fmap mapNonWordToSpace
    trimApos = fmap $ dropWhile isApos . dropWhileEnd isApos
    toCounter = fromListWith (+) . fmap (,1)