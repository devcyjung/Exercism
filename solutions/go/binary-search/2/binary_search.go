package binarysearch

import "slices"

func SearchInts(list []int, key int) int {
	pos, ok := slices.BinarySearch(list, key)
    if !ok {
        return -1
    }
    return pos
}