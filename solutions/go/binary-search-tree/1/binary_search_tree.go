package binarysearchtree

type BinarySearchTree struct {
    data, size	int
	left, right *BinarySearchTree
}

func NewBst(i int) *BinarySearchTree {
	return &BinarySearchTree{ i, 1, nil, nil }
}

func (bst *BinarySearchTree) Insert(i int) {
    bst.size++
	switch {
    case bst.data < i:
        if bst.right == nil {
            bst.right = NewBst(i)
        } else {
            bst.right.Insert(i)
        }
    case i < bst.data:
        if bst.left == nil {
            bst.left = NewBst(i)
        } else {
            bst.left.Insert(i)
        }
    default:
        if bst.left == nil {
            bst.left = NewBst(i)
        } else if bst.right == nil {
            bst.right = NewBst(i)
        } else {
            bst.left.Insert(i)
        }
    }
}

func (bst *BinarySearchTree) SortedData() []int {
    if bst == nil {
        return nil
    }
	sorted := make([]int, 0, bst.size)
    sorted = append(sorted, bst.left.SortedData()...)
    sorted = append(sorted, bst.data)
    sorted = append(sorted, bst.right.SortedData()...)
    return sorted
}