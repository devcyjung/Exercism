{-# LANGUAGE OverloadedStrings #-}
module Bob (responseFor) where

import qualified Data.Text as T
import qualified Data.Char as C

responseFor :: T.Text -> T.Text
responseFor xs
  | isSilent = "Fine. Be that way!"
  | isYelling && isAsking = "Calm down, I know what I'm doing!"
  | isYelling = "Whoa, chill out!"
  | isAsking = "Sure."
  | otherwise = "Whatever."
  where
    trimmed = T.stripEnd xs
    uppered = T.toUpper xs
    isSilent = trimmed == T.empty
    isYelling = T.any C.isLetter xs && xs == uppered
    isAsking = T.isSuffixOf "?" trimmed