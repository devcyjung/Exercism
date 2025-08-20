module Satellite (treeFromTraversals) where

import BinaryTree (BinaryTree (..))
import Data.List (elemIndices)

treeFromTraversals :: (Ord a) => [a] -> [a] -> Maybe (BinaryTree a)
treeFromTraversals [] [] = Nothing
treeFromTraversals preorder inorder = tree preorder inorder

tree :: (Ord a) => [a] -> [a] -> Maybe (BinaryTree a)
tree [] [] = Just Leaf
tree [] _ = Nothing
tree _ [] = Nothing
tree (root : preorder) inorder = case elemIndices root inorder of
  [] -> Nothing
  (index : []) -> 
    let (lhsInorder, rhsInorderPlusRoot) = splitAt index inorder
    in case rhsInorderPlusRoot of
      [] -> Nothing
      (_ : rhsInorder) -> 
        let (lhsPreorder, rhsPreorder) = splitAt (length lhsInorder) preorder
        in Branch <$> tree lhsPreorder lhsInorder <*> pure root <*> tree rhsPreorder rhsInorder
  _ -> Nothing
