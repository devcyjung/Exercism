module Bob (responseFor) where

import qualified Data.Text as T
import           Data.Text (Text)
import           Data.Char (isAlpha)

responseFor :: Text -> Text
responseFor xs
  | isSilent xs = T.pack "Fine. Be that way!"
  | isYelling xs && isAsking xs = T.pack "Calm down, I know what I'm doing!"
  | isYelling xs = T.pack "Whoa, chill out!"
  | isAsking xs = T.pack "Sure."
  | otherwise = T.pack "Whatever."

isSilent :: Text -> Bool
isSilent xs = T.stripStart xs == T.empty

isYelling :: Text -> Bool
isYelling xs = T.any isAlpha xs && xs == T.toUpper xs 

isAsking :: Text -> Bool
isAsking xs = T.isSuffixOf (T.pack "?") (T.stripEnd xs)