module Squares (difference, squareOfSum, sumOfSquares) where

difference :: Integral a => a -> a
difference n = squareOfSum n - sumOfSquares n

squareOfSum :: Integral a => a -> a
squareOfSum n = sum [1..n] ^ 2

sumOfSquares :: Integral a => a -> a
sumOfSquares n = sum [x ^ 2 | x <- [1..n]]