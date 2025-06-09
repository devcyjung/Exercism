module Pangram (isPangram) where

import Prelude (String, Bool (..), (==), (>=), (<=), (&&), (-), otherwise)
import Data.Char (ord)
import Data.Vector.Generic
import qualified Data.Vector as V

isPangram  :: String -> Bool
isPangram text  = all (== True) (listBools text) 

listBools  :: String -> V.Vector Bool
listBools []                = replicate 26 False
listBools (x:xs)
    | x >= 'a' && x <= 'z'  = (listBools xs) // [(ord x - ord 'a', True)]
    | x >= 'A' && x <= 'Z'  = (listBools xs) // [(ord x - ord 'A', True)]
    | otherwise             = listBools xs