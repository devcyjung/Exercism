module Zipper
 ( BinTree(BT)
 , fromTree
 , left
 , right
 , setLeft
 , setRight
 , setValue
 , toTree
 , up
 , value
 ) where

import Data.Foldable (foldl')

data BinTree a = BT { btValue :: a
                    , btLeft  :: Maybe (BinTree a)
                    , btRight :: Maybe (BinTree a)
                    } deriving (Eq, Show)

type Zipper a = (BinTree a, [Trace a])

data Trace a = LeftTrace a (Maybe (BinTree a)) | RightTrace a (Maybe (BinTree a)) deriving (Eq, Show)

fromTree :: BinTree a -> Zipper a
fromTree tree = (tree, [])

toTree :: Zipper a -> BinTree a
toTree (tree, traces) = foldl' f tree traces
  where
    f t trace = case trace of
      LeftTrace parent maybeRightTree -> BT { btValue = parent
                                            , btLeft = Just t
                                            , btRight = maybeRightTree
                                            }
      RightTrace parent maybeLeftTree -> BT { btValue = parent
                                            , btLeft = maybeLeftTree
                                            , btRight = Just t
                                            }

value :: Zipper a -> a
value (BT { btValue = btVal }, _) = btVal

left :: Zipper a -> Maybe (Zipper a)
left (BT { btValue = btVal
         , btLeft = maybeLeftTree
         , btRight = maybeRightTree
         }
     , traces) = case maybeLeftTree of
  Nothing -> Nothing
  Just leftTree -> Just (leftTree, LeftTrace btVal maybeRightTree : traces) 
              
right :: Zipper a -> Maybe (Zipper a)
right (BT { btValue = btVal
          , btLeft = maybeLeftTree
          , btRight = maybeRightTree
          }
      , traces) = case maybeRightTree of
  Nothing -> Nothing
  Just rightTree -> Just (rightTree, RightTrace btVal maybeLeftTree : traces)

up :: Zipper a -> Maybe (Zipper a)
up (tree, ts) = case ts of
  [] -> Nothing
  (LeftTrace btVal maybeRightTree : traces) -> Just (BT { btValue = btVal
                                                        , btLeft = Just tree
                                                        , btRight = maybeRightTree
                                                        }, traces)
  (RightTrace btVal maybeLeftTree : traces) -> Just (BT { btValue = btVal
                                                        , btLeft = maybeLeftTree
                                                        , btRight = Just tree
                                                        }, traces)

setValue :: a -> Zipper a -> Zipper a
setValue x (tree, traces) = (tree { btValue = x }, traces)

setLeft :: Maybe (BinTree a) -> Zipper a -> Zipper a
setLeft maybeTree (tree, traces) = (tree { btLeft = maybeTree }, traces) 

setRight :: Maybe (BinTree a) -> Zipper a -> Zipper a
setRight maybeTree (tree, traces) = (tree { btRight = maybeTree }, traces)