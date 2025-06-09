module CryptoSquare (encode) where

import Data.Char (isAlphaNum, toLower)
import Data.List (intercalate)

encode :: String -> String
encode xs = intercalate " " [[charAt (ri * c + ci) | ri <- [0..r-1]] | ci <- [0..c-1]]
  where
    sanitized = [toLower x | x <- xs, isAlphaNum x]
    size = length sanitized
    (r, c) = let root = (floor . sqrt . fromIntegral) size
             in if root * root == size
                then (root, root)
                else if root * (root + 1) >= size
                     then (root, root + 1)
                     else (root + 1, root + 1)
    charAt i = if i < size then sanitized !! i else ' '