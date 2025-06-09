module ETL (transform) where

import Data.Char (toLower)
import Data.Map (Map)
import qualified Data.Map as M

transform :: Map a String -> Map Char a
transform = M.foldlWithKey' f M.empty
  where
    f a k v = let ins e = M.insert (toLower e) k
              in foldr ins a v