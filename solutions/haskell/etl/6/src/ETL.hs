module ETL (transform) where

import Data.Char (toLower)
import Data.Map (Map, fromList, toList)

transform :: Map a String -> Map Char a
transform = fromList . concatMap (\(k, vs) -> [(toLower v, k) | v <- vs]) . toList