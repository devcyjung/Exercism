module DNA (toRNA) where

toRNA :: String -> Either Char String
toRNA "" = Right ""
toRNA (x:xs) =
  parse x >>= \r ->
  toRNA xs >>= \rs ->
  Right (r:rs)
  where
    parse ch = case ch of
      'G' -> Right 'C'
      'C' -> Right 'G'
      'T' -> Right 'A'
      'A' -> Right 'U'
      ch -> Left ch