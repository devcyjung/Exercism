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
    splitWords = words . fmap toLower . fmap (bool ' ' <*> isApos <||> isLetter <||> isNumber)
    trimApos = fmap $ dropWhile isApos . dropWhileEnd isApos
    toCounter = fromListWith (+) . flip zip (repeat 1)