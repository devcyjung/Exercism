module ETL (transform) where

import Data.Char (toLower)
import Data.Foldable (foldl')
import Data.Map (Map)
import qualified Data.Map as M

transform :: Map a String -> Map Char a
transform = M.foldlWithKey' (flip go) M.empty
  where
    go k = let ins e = M.insert (toLower e) k
           in foldr ins