module Bob (responseFor) where

import qualified Data.Text as T
import           Data.Char (isAlpha)

responseFor :: T.Text -> T.Text
responseFor xs
  | isSilent = T.pack "Fine. Be that way!"
  | isYelling && isAsking = T.pack "Calm down, I know what I'm doing!"
  | isYelling = T.pack "Whoa, chill out!"
  | isAsking = T.pack "Sure."
  | otherwise = T.pack "Whatever."
  where
    trimmed = T.stripEnd xs
    isSilent = trimmed == T.empty
    isYelling = T.any isAlpha xs && xs == T.toUpper xs
    isAsking = T.isSuffixOf (T.pack "?") trimmed