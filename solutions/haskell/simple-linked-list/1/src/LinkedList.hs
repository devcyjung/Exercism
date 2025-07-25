{-# LANGUAGE DeriveFoldable #-}

module LinkedList
    ( LinkedList
    , datum
    , fromList
    , isNil
    , new
    , next
    , nil
    , reverseLinkedList
    , toList
    ) where

import Data.Foldable (foldl')

data LinkedList a = Nil | Cons a (LinkedList a) deriving (Eq, Show, Foldable)

datum :: LinkedList a -> a
datum (Cons x _) = x

fromList :: [a] -> LinkedList a
fromList [] = Nil
fromList (x:xs) = Cons x (fromList xs)

isNil :: LinkedList a -> Bool
isNil Nil = True
isNil _ = False

new :: a -> LinkedList a -> LinkedList a
new x xs = Cons x xs

next :: LinkedList a -> LinkedList a
next Nil = Nil
next (Cons _ xs) = xs

nil :: LinkedList a
nil = Nil

reverseLinkedList :: LinkedList a -> LinkedList a
reverseLinkedList Nil = Nil
reverseLinkedList xs = foldl' (\acc elem -> Cons elem acc) Nil xs 

toList :: LinkedList a -> [a]
toList Nil = []
toList (Cons x xs) = x : toList xs